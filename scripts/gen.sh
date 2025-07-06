#!/bin/bash
set -euo pipefail

BASE_DIR=$(dirname "$(dirname "$0")")
TMP_DIR=/tmp
BASE_DIR="${BASE_DIR}/pkg/gitlab"

function capitalize() {
    VALUE="$1"
    # shellcheck disable=SC2206
    IFS=_ ARR=(${VALUE})

    CAP_ARR=()
    for V in "${ARR[@]}"
    do
        CAP_ARR+=("${V^}")
    done

    IFS='' CAP="${CAP_ARR[*]}"

    echo "${CAP}"
}

function groupname() {
    OP_PATH="$1"
    echo "${OP_PATH}" | cut -d '/' -f 4
}

function apiname() {
    OP_PATH="$1"
    echo "${OP_PATH}" | cut -d '/' -f 5-
}

function filename() {
    OP_PATH="$1"
    OP_GROUP=$(groupname "${OP_PATH}")
    echo "${BASE_DIR}/${OP_GROUP}.go"
}

function toolname() {
    OP_PATH="$1"
    OP_GROUP=$(groupname "${OP_PATH}")
    OP_API=$(apiname "${OP_PATH}")
    TOO_NAME="${OP_GROUP,,}_${OP_API,,}"
    TOO_NAME=${TOO_NAME//./_} # '.' -> '_'
    TOO_NAME=${TOO_NAME//-/_} # '-' -> '_'
    TOO_NAME=${TOO_NAME////_} # '/' -> '_'
    TOO_NAME=${TOO_NAME//$/_} # '$' -> '_'
    TOO_NAME=${TOO_NAME//{/} # '{' -> ''
    TOO_NAME=${TOO_NAME//\}/} # '}' -> ''
    echo "${TOO_NAME}"
}

function normalize() {
    VALUE="$1"
    # shellcheck disable=SC2206
    ARR=(${VALUE})
    echo "${ARR[*]}"
}

function shortname() {
    VALUE="$1"
    echo "${VALUE}" | \
        sed -e 's/projects/pjs/' \
            -e 's/groups/grps/' \
            -e 's/packages/pkgs/' \
            -e 's/pipelines/pls/' \
            -e 's/repository/repo/' \
            -e 's/merge_requests/mrs/'
}

function write-preceding() {
    OP_PATH="$1"

    FILE_PATH=$(filename "${OP_PATH}")
    cat > "${FILE_PATH}" <<EOF
package gitlab

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
	openapi_types "github.com/oapi-codegen/runtime/types"
)
EOF
}

function write-register-func() {
    METHOD="$1"
    OP_PATH="$2"
    QUERY="$3"

    DESCRIPTION=$(yq -r "${QUERY}.description" "${TMP_DIR}/openapi.yml")
    DESCRIPTION=$(normalize "${DESCRIPTION}")
    DESCRIPTION="${DESCRIPTION//\"/\\\"}"

    TOOL_NAME=$(toolname "${OP_PATH}")
    TOOL_SNAME=$(shortname "${TOOL_NAME}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    echo "| ${TOOL_SNAME} | ${DESCRIPTION} |"

    FILE_PATH=$(filename "${OP_PATH}")
    # signature
    cat >> "${FILE_PATH}" <<EOF

func register${METHOD}${API_NAME}(s *server.MCPServer) {
	tool := mcp.NewTool("${TOOL_SNAME}",
		mcp.WithDescription("${DESCRIPTION}"),
EOF

    # parameters
    yq "${QUERY}.parameters[] | . style=\"flow\"" "${TMP_DIR}/openapi.yml" | while read -r PARAMETER
    do
        NAME=$(yq -r '.name' <<<"${PARAMETER}")
        DESCRIPTION=$(yq -r '.description' <<<"${PARAMETER}")
        DESCRIPTION=$(normalize "${DESCRIPTION}")
        DESCRIPTION="${DESCRIPTION//\"/\\\"}"

        EXAMPLE=$(yq -r '.example' <<<"${PARAMETER}")
        if [[ $EXAMPLE != "null" && $EXAMPLE != "{}" ]]; then
            DESCRIPTION="${DESCRIPTION} (example: ${EXAMPLE})"
        fi

        REQUIRED=""
        if [[ $(yq -r '.required' <<<"${PARAMETER}") == "true" ]]; then
            REQUIRED="mcp.Required(),"
        fi

        ENUM=""
        ENUM_VALUES=$(yq -r '[.schema.enum[] | "\"\(.)\""] | join(",")' <<<"${PARAMETER}")
        if [[ -n "${ENUM_VALUES}" && "${ENUM_VALUES}" != "\"\"" ]]; then
            ENUM="mcp.Enum(${ENUM_VALUES}),"
        fi

        DEFAULT=""
        DEFAULT_VALUE=$(yq -r '.schema.default' <<<"${PARAMETER}")

        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        if [[ "${TY}" == "integer" ]]; then
            if [[ "$DEFAULT_VALUE" != "null" && "$DEFAULT_VALUE" != "{}" ]]; then
                # https://github.com/microsoft/vscode/issues/250599
                # https://github.com/microsoft/vscode-copilot-release/issues/13588
                #DEFAULT="mcp.DefaultNumber(float64(${DEFAULT_VALUE})),"
                DESCRIPTION="${DESCRIPTION} (default: ${DEFAULT_VALUE})"
            fi

            cat >> "${FILE_PATH}" <<EOF
		mcp.WithNumber("${NAME}",
			mcp.Description("${DESCRIPTION}"),
			${REQUIRED}
			${ENUM}
			${DEFAULT}
		),
EOF
        elif [[ "${TY}" == "boolean" ]]; then
            if [[ "$DEFAULT_VALUE" != "null" && "$DEFAULT_VALUE" != "{}" ]]; then
                #DEFAULT="mcp.DefaultBool(${DEFAULT_VALUE}),"
                DESCRIPTION="${DESCRIPTION} (default: ${DEFAULT_VALUE})"
            fi

            cat >> "${FILE_PATH}" <<EOF
		mcp.WithBoolean("${NAME}",
			mcp.Description("${DESCRIPTION}"),
			${REQUIRED}
			${ENUM}
			${DEFAULT}
		),
EOF
        else
            if [[ "$DEFAULT_VALUE" != "null" && "$DEFAULT_VALUE" != "{}" ]]; then
                #DEFAULT="mcp.DefaultString(\"${DEFAULT_VALUE}\"),"
                DESCRIPTION="${DESCRIPTION} (default: ${DEFAULT_VALUE})"
            fi

            cat >> "${FILE_PATH}" <<EOF
		mcp.WithString("${NAME}",
			mcp.Description("${DESCRIPTION}"),
			${REQUIRED}
			${ENUM}
			${DEFAULT}
		),
EOF
        fi
    done

    # complete
    cat >> "${FILE_PATH}" <<EOF
	)

	s.AddTool(tool, ${METHOD,}${API_NAME}Handler)
}
EOF
}

function write-handler-func() {
    METHOD="$1"
    OP_PATH="$2"
    QUERY="$3"

    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")
    NUM_PARAMETER=$(yq -r "[${QUERY}.parameters[] | select(.in == \"query\")] | length" "${TMP_DIR}/openapi.yml")

    # parameters
    PATH_STMT=()
    PATH_ARG=()
    while read -r NAME
    do
        NAME="${NAME//{/}" # '{' -> ''
        NAME="${NAME//\}/}" # '}' -> ''

        PARAMETER=$(yq "${QUERY}.parameters[] | select(.name == \"${NAME}\")" "${TMP_DIR}/openapi.yml")

        VAR_NAME=${NAME//-/_} # '-' -> '_'
        VAR_NAME=${VAR_NAME//./}
        VAR_NAME=${VAR_NAME//[/_} # '[' -> '_'
        VAR_NAME=${VAR_NAME//]/_} # ']' -> '_'
        if [[ "${VAR_NAME}" == 'type' ]]; then
            VAR_NAME='_type'
        elif [[ "${VAR_NAME}" == 'params' ]]; then
            VAR_NAME='_params'
        fi

        PATH_ARG+=("${VAR_NAME}")

        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        FORMAT=$(yq -r '.schema.format' <<<"${PARAMETER}")
        if [[ "${TY}" == "integer" ]]; then
            CONV_STMT="request.GetInt(\"${NAME}\", math.MinInt)"
            if [[ $FORMAT == "int32" ]]; then
                CONV_STMT="int32(${CONV_STMT})"
            fi

            PATH_STMT+=("${VAR_NAME} := ${CONV_STMT}")
        elif [[ "${TY}" == "boolean" ]]; then
            PATH_STMT+=("${VAR_NAME} := request.GetBool(\"${NAME}\", false)")
        else
            PATH_STMT+=("${VAR_NAME} := request.GetString(\"${NAME}\", \"\")")
        fi
    done < <(echo "${OP_PATH}" | grep -oP '{[^}]+}')

    PATH_STMT_STR="$(IFS=$'\n'; echo "${PATH_STMT[*]}")"
    PATH_ARG_STR="$(IFS=,; echo "${PATH_ARG[*]}")"
    if [[ -n "${PATH_ARG_STR}" ]]; then
        PATH_ARG_STR=", ${PATH_ARG_STR}"
    fi

    PARSE_STMT="params := parse${METHOD}${API_NAME}(request)"
    PARAM_ARG=",&params"
    if [[ ${NUM_PARAMETER} -eq 0 ]]; then
        PARSE_STMT=""
        PARAM_ARG=""
    fi

    FILE_PATH=$(filename "${OP_PATH}")
    cat >> "${FILE_PATH}" <<EOF

func ${METHOD,}${API_NAME}Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	${PATH_STMT_STR}
	${PARSE_STMT}
	return toResult(c.${METHOD}ApiV4${API_NAME}(ctx ${PATH_ARG_STR} ${PARAM_ARG}, authorizationHeader))
}
EOF
}

function write-parser-func() {
    METHOD="$1"
    OP_PATH="$2"
    QUERY="$3"

    NUM_PARAMETER=$(yq -r "[${QUERY}.parameters[] | select(.in == \"query\")] | length" "${TMP_DIR}/openapi.yml")
    if [[ ${NUM_PARAMETER} -eq 0 ]]; then
        return
    fi

    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    FILE_PATH=$(filename "${OP_PATH}")
    # signature
    cat >> "${FILE_PATH}" <<EOF

func parse${METHOD}${API_NAME}(request mcp.CallToolRequest) client.${METHOD}ApiV4${API_NAME}Params {
	params := client.${METHOD}ApiV4${API_NAME}Params{}
EOF

    # parameters
    yq "${QUERY}.parameters[] | select(.in == \"query\") | . style=\"flow\"" "${TMP_DIR}/openapi.yml" | while read -r PARAMETER
    do
        NAME=$(yq -r '.name' <<<"${PARAMETER}")

        VAR_NAME=${NAME//-/_} # '-' -> '_'
        VAR_NAME=${VAR_NAME//./}
        VAR_NAME=${VAR_NAME//[/_} # '[' -> '_'
        VAR_NAME=${VAR_NAME//]/_} # ']' -> '_'
        if [[ "${VAR_NAME}" == 'type' ]]; then
            VAR_NAME='_type'
        elif [[ "${VAR_NAME}" == 'params' ]]; then
            VAR_NAME='_params'
        fi

        REF_NAME="&${VAR_NAME}"
        if [[ $(yq -r '.required' <<<"${PARAMETER}") == "true" ]]; then
            REF_NAME="${VAR_NAME}"
        fi

        PARAM_NAME=$(capitalize "${VAR_NAME}")

        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        FORMAT=$(yq -r '.schema.format' <<<"${PARAMETER}")
        if [[ "${TY}" == "integer" ]]; then
            DEFAULT_VALUE=$(yq -r '.schema.default' <<<"${PARAMETER}")
            if [[ "$DEFAULT_VALUE" == "null" || "$DEFAULT_VALUE" == "{}" ]]; then
                unset DEFAULT_VALUE
            fi

            CONV_STMT=""
            if [[ $FORMAT == "int32" ]]; then
                CONV_STMT="${VAR_NAME} := int32(${VAR_NAME})"
            fi

            cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetInt("${NAME}", ${DEFAULT_VALUE:-math.MinInt})
	if ${VAR_NAME} != math.MinInt {
        ${CONV_STMT}
		params.${PARAM_NAME} = ${REF_NAME}
	}
EOF
        elif [[ "${TY}" == "boolean" ]]; then
            DEFAULT_VALUE=$(yq -r '.schema.default' <<<"${PARAMETER}")
            if [[ "$DEFAULT_VALUE" == "null" || "$DEFAULT_VALUE" == "{}" ]]; then
                unset DEFAULT_VALUE
            fi

            cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetBool("${NAME}", ${DEFAULT_VALUE:-false})
	params.${PARAM_NAME} = ${REF_NAME}
EOF
        elif [[ "${TY}" == "array" ]]; then
            FORMAT=$(yq -r '.schema.items.format' <<<"${PARAMETER}")
            ITEM_TY=$(yq -r '.schema.items.type' <<<"${PARAMETER}")
            if [[ "${ITEM_TY}" == "integer" ]]; then
                # TODO: error handling
                TY_NAME="int"
                CONV_STMT="intValue"
                if [[ $FORMAT == "int32" ]]; then
                    TY_NAME="int32"
                    CONV_STMT="int32(${CONV_STMT})"
                fi
                cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetString("${NAME}", "")
	if ${VAR_NAME} != "" {
        ${VAR_NAME} := strings.Split(${VAR_NAME}, ",")
        var intSlice []${TY_NAME}
        for _, v := range ${VAR_NAME} {
            intValue, _ := strconv.Atoi(v)
            intSlice = append(intSlice, ${CONV_STMT})
        }
		params.${PARAM_NAME} = &intSlice
	}
EOF
            else
                cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetString("${NAME}", "")
	if ${VAR_NAME} != "" {
        ${VAR_NAME} := strings.Split(${VAR_NAME}, ",")
		params.${PARAM_NAME} = ${REF_NAME}
	}
EOF
            fi
        else
            CONV_STMT=""
            if [[ "${FORMAT}" == "date-time" ]]; then
                # TODO: error handling
                CONV_STMT="${VAR_NAME}, _ := time.Parse(time.RFC3339, ${VAR_NAME})"
            elif [[ "${FORMAT}" == "date" ]]; then
                # TODO: error handling
                CONV_STMT="${VAR_NAME}, _ := time.Parse(time.DateOnly, ${VAR_NAME})"
                REF_NAME="&openapi_types.Date{Time: ${VAR_NAME}}"
            fi
            cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetString("${NAME}", "")
	if ${VAR_NAME} != "" {
        ${CONV_STMT}
		params.${PARAM_NAME} = ${REF_NAME}
	}
EOF
        fi

    done

    # complete
    cat >> "${FILE_PATH}" <<EOF

	return params
}
EOF
}

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    FILE_PATH=$(filename "${OP_PATH}")
    rm -rf "${FILE_PATH}"
done

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    write-preceding "${OP_PATH}"
done

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    OD_QUERY=".paths.\"${OP_PATH}\".get"
    OP_GET=$(yq "${OD_QUERY}" "${TMP_DIR}/openapi.yml")
    if [[ "$OP_GET" != 'null' ]]; then
        write-register-func "Get" "${OP_PATH}" "${OD_QUERY}"
        write-handler-func "Get" "${OP_PATH}" "${OD_QUERY}"
        write-parser-func "Get" "${OP_PATH}" "${OD_QUERY}"
    fi
done

# write tools.go
TOOLS_PATH="${BASE_DIR}/tools.go"
cat > "${TOOLS_PATH}" <<EOF
package gitlab

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
EOF

yq -r '.paths | keys | .[]' "${TMP_DIR}/openapi.yml" | while read -r OP_PATH
do
    TOOL_NAME=$(toolname "${OP_PATH}")
    TOOL_SNAME=$(shortname "${TOOL_NAME}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    OP_GET=$(yq ".paths.\"${OP_PATH}\".get" "${TMP_DIR}/openapi.yml")
    if [[ "$OP_GET" != 'null' ]]; then
        METHOD="Get"
        if [[ ${#TOOL_SNAME} -gt 46 ]]; then
            # If the tool name is too long, we need to use a shorter name
            # to avoid exceeding the maximum length of 46 characters.
            # See https://github.com/microsoft/vscode/blob/1.101.2/src/vs/workbench/contrib/mcp/common/mcpTypes.ts#L710-L714
            echo "//register${METHOD}${API_NAME}(s)" >> "${TOOLS_PATH}"
        else
            echo "register${METHOD}${API_NAME}(s)" >> "${TOOLS_PATH}"
        fi
    fi
done

cat >> "${TOOLS_PATH}" <<EOF
}
EOF

# format
go fmt ./...
