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
    TOO_NAME=${TOO_NAME//\(/} # '(' -> ''
    TOO_NAME=${TOO_NAME//\)/} # ')' -> ''
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
    VALUE=${VALUE%%_}
    echo "${VALUE}" | \
        sed -E \
            -e 's/projects/pjs/' \
            -e 's/groups/grps/' \
            -e 's/packages/pkgs/' \
            -e 's/pipelines/pls/' \
            -e 's/repository/repo/' \
            -e 's/merge_requests/mrs/' \
            -e 's/_+/_/g'
}

function write-preceding() {
    OP_PATH="$1"

    FILE_PATH=$(filename "${OP_PATH}")
    cat > "${FILE_PATH}" <<EOF
package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)
EOF
}

function write-request-struct(){
    local METHOD="$1"
    local OP_PATH="$2"
    local PATH_INFO="$3"
    local REQ_BODY="${4:-}"

    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")
    NUM_PARAMETER=$(yq -r '[.parameters[] | select(.in == "query")] | length' <<<"${PATH_INFO}")

    QUERY_SMTM_STR=""
    if [[ ${NUM_PARAMETER} -ne 0 ]]; then
        QUERY_SMTM_STR="Params *client.${METHOD}ApiV4${API_NAME}Params \`json:\"params,omitempty\"\`"
    fi

    # parameters
    PATH_STMT=()
    PATH_ARG=()
    while read -r NAME
    do
        NAME="${NAME//{/}" # '{' -> ''
        NAME="${NAME//\}/}" # '}' -> ''

        PARAMETER=$(yq ".parameters[] | select(.name == \"${NAME}\")" <<<"${PATH_INFO}")

        VAR_NAME=${NAME//-/_} # '-' -> '_'
        VAR_NAME=${VAR_NAME//./}
        VAR_NAME=${VAR_NAME//[/_} # '[' -> '_'
        VAR_NAME=${VAR_NAME//]/_} # ']' -> '_'
        if [[ "${VAR_NAME}" == 'type' ]]; then
            VAR_NAME='_type'
        elif [[ "${VAR_NAME}" == 'params' ]]; then
            VAR_NAME='_params'
        fi

        PARAM_NAME=$(capitalize "${VAR_NAME}")

        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        DESCRIPTION=$(yq -r '.description' <<<"${PARAMETER}")

        DESCRIPTION=$(normalize "${DESCRIPTION}")
        DESCRIPTION="${DESCRIPTION//\\/\\\\}" # '\' -> '\\'
        DESCRIPTION="${DESCRIPTION//\"/\\\"}"
        DESCRIPTION="${DESCRIPTION//\`/\'}" # '`' -> '\''

        if [[ "${TY}" == "integer" ]]; then
            FORMAT=$(yq -r '.schema.format' <<<"${PARAMETER}")
            if [[ $FORMAT == "int32" ]]; then
                PATH_STMT+=("${PARAM_NAME} int32 \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
            else
                PATH_STMT+=("${PARAM_NAME} int \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
            fi
        elif [[ "${TY}" == "boolean" ]]; then
            PATH_STMT+=("${PARAM_NAME} bool \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
        else
            PATH_STMT+=("${PARAM_NAME} string \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
        fi
    done < <(echo "${OP_PATH}" | grep -oP '{[^}]+}')

    PATH_STMT_STR="$(IFS=$'\n'; echo "${PATH_STMT[*]}")"

    # body
    BODY_SMTM_STR=""
    if [[ -n "${REQ_BODY}" && "${REQ_BODY}" != 'null' ]]; then
        BODY_SMTM_STR="Body client.${METHOD}ApiV4${API_NAME}JSONRequestBody \`json:\"body\"\`"
    fi

    FILE_PATH=$(filename "${OP_PATH}")
    cat >> "${FILE_PATH}" <<EOF
type ${METHOD}${API_NAME}Request struct {
    ${PATH_STMT_STR}
    ${QUERY_SMTM_STR}
    ${BODY_SMTM_STR}
}
EOF

}

function write-register-func() {
    local METHOD="$1"
    local OP_PATH="$2"
    local PATH_INFO="$3"

    DESCRIPTION=$(yq -r ".description" <<<"${PATH_INFO}")
    DESCRIPTION=$(normalize "${DESCRIPTION}")
    DESCRIPTION="${DESCRIPTION//\\/\\\\}" # '\' -> '\\'
    DESCRIPTION="${DESCRIPTION//\"/\\\"}"
    DESCRIPTION="${DESCRIPTION//\`/\'}" # '`' -> '\''

    TOOL_NAME=$(toolname "${OP_PATH}")
    TOOL_SNAME=$(shortname "${TOOL_NAME}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    echo "| ${METHOD,,}_${TOOL_SNAME} | ${DESCRIPTION} |"

    FILE_PATH=$(filename "${OP_PATH}")
    # signature
    cat >> "${FILE_PATH}" <<EOF

func register${METHOD}${API_NAME}(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&${METHOD}${API_NAME}Request{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("${METHOD,,}_${TOOL_SNAME}",
		mcp.WithDescription("${DESCRIPTION}"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(${METHOD,}${API_NAME}Handler))
}
EOF
}

function write-handler-func() {
    local METHOD="$1"
    local OP_PATH="$2"
    local PATH_INFO="$3"
    local REQ_BODY="${4:-}"

    TOOL_NAME=$(toolname "${OP_PATH}")
    API_NAME=$(capitalize "${TOOL_NAME}")
    NUM_PARAMETER=$(yq -r '[.parameters[] | select(.in == "query")] | length' <<<"${PATH_INFO}")

    QUERY_SMTM_STR=""
    if [[ ${NUM_PARAMETER} -ne 0 ]]; then
        QUERY_SMTM_STR=", req.Params"
    fi

    # parameters
    PATH_ARG=()
    while read -r NAME
    do
        NAME="${NAME//{/}" # '{' -> ''
        NAME="${NAME//\}/}" # '}' -> ''

        PARAMETER=$(yq ".parameters[] | select(.name == \"${NAME}\")" <<<"${PATH_INFO}")

        VAR_NAME=${NAME//-/_} # '-' -> '_'
        VAR_NAME=${VAR_NAME//./}
        VAR_NAME=${VAR_NAME//[/_} # '[' -> '_'
        VAR_NAME=${VAR_NAME//]/_} # ']' -> '_'
        if [[ "${VAR_NAME}" == 'type' ]]; then
            VAR_NAME='_type'
        elif [[ "${VAR_NAME}" == 'params' ]]; then
            VAR_NAME='_params'
        fi

        PARAM_NAME=$(capitalize "${VAR_NAME}")

        PATH_ARG+=("req.${PARAM_NAME}")
    done < <(echo "${OP_PATH}" | grep -oP '{[^}]+}')

    PATH_ARG_STR="$(IFS=,; echo "${PATH_ARG[*]}")"
    if [[ -n "${PATH_ARG_STR}" ]]; then
        PATH_ARG_STR=", ${PATH_ARG_STR}"
    fi

    BODY_ARG_STR=""
    if [[ -n "${REQ_BODY}" && "${REQ_BODY}" != 'null' ]]; then
        BODY_ARG_STR=", req.Body"
    fi

    FILE_PATH=$(filename "${OP_PATH}")
    cat >> "${FILE_PATH}" <<EOF

func ${METHOD,}${API_NAME}Handler(ctx context.Context, request mcp.CallToolRequest, req ${METHOD}${API_NAME}Request) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.${METHOD}ApiV4${API_NAME}(ctx ${PATH_ARG_STR} ${QUERY_SMTM_STR} ${BODY_ARG_STR}, authorizationHeader))
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

yq '.paths | to_entries | .[] | . style="flow"' "${TMP_DIR}/openapi.yml" | while read -r PATH_INFO
do
    OP_PATH=$(yq -r '.key' <<<"${PATH_INFO}")

    OP_DELETE=$(yq '.value.delete | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_DELETE" != 'null' ]]; then
        write-request-struct "Delete" "${OP_PATH}" "${OP_DELETE}"
        write-register-func "Delete" "${OP_PATH}" "${OP_DELETE}"
        write-handler-func "Delete" "${OP_PATH}" "${OP_DELETE}"
    fi

    OP_POST=$(yq '.value.post | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.requestBody | . style="flow"' <<<"${OP_POST}")
    IS_JSON=$(yq -r '.content | ."application/json"' <<<"${OP_REQ_BODY}")
    if [[ "$OP_POST" != 'null' && "${IS_JSON}" != "null" ]]; then
        write-request-struct "Post" "${OP_PATH}" "${OP_POST}" "${OP_REQ_BODY}"
        write-register-func "Post" "${OP_PATH}" "${OP_POST}"
        write-handler-func "Post" "${OP_PATH}" "${OP_POST}" "${OP_REQ_BODY}"
    fi

    OP_PUT=$(yq '.value.put | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.requestBody | . style="flow"' <<<"${OP_PUT}")
    IS_JSON=$(yq -r '.content | ."application/json"' <<<"${OP_REQ_BODY}")
    if [[ "$OP_PUT" != 'null' && "${IS_JSON}" != "null" ]]; then
        write-request-struct "Put" "${OP_PATH}" "${OP_PUT}" "${OP_REQ_BODY}"
        write-register-func "Put" "${OP_PATH}" "${OP_PUT}"
        write-handler-func "Put" "${OP_PATH}" "${OP_PUT}" "${OP_REQ_BODY}"
    fi

    OP_GET=$(yq '.value.get | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_GET" != 'null' ]]; then
        write-request-struct "Get" "${OP_PATH}" "${OP_GET}"
        write-register-func "Get" "${OP_PATH}" "${OP_GET}"
        write-handler-func "Get" "${OP_PATH}" "${OP_GET}"
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

yq '.paths | to_entries | .[] | . style="flow"' "${TMP_DIR}/openapi.yml" | while read -r PATH_INFO
do
    OP_PATH=$(yq -r '.key' <<<"${PATH_INFO}")

    TOOL_NAME=$(toolname "${OP_PATH}")
    TOOL_SNAME=$(shortname "${TOOL_NAME}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    # If the tool name is too long, we need to use a shorter name
    # to avoid exceeding the maximum length of 46 characters.
    # See https://github.com/microsoft/vscode/blob/1.101.2/src/vs/workbench/contrib/mcp/common/mcpTypes.ts#L710-L714

    OP_DELETE=$(yq '.value.delete' <<<"${PATH_INFO}")
    if [[ "$OP_DELETE" != 'null' ]]; then
        METHOD="Delete"
        if [[ ${#TOOL_SNAME} -gt 39 ]]; then # delete_XXX
            echo "//if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        else
            echo "if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        fi
    fi

    OP_POST=$(yq '.value.post' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.requestBody | . style="flow"' <<<"${OP_POST}")
    IS_JSON=$(yq -r '.content | ."application/json"' <<<"${OP_REQ_BODY}")
    if [[ "$OP_POST" != 'null' ]]; then
        METHOD="Post"
        if [[ ${#TOOL_SNAME} -gt 41 ]]; then # post_XXX
            echo "//if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        elif [[ "${IS_JSON}" == "null" ]]; then
            echo "//if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        else
            echo "if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        fi
    fi

    OP_PUT=$(yq '.value.put' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.requestBody | . style="flow"' <<<"${OP_PUT}")
    IS_JSON=$(yq -r '.content | ."application/json"' <<<"${OP_REQ_BODY}")
    if [[ "$OP_PUT" != 'null' ]]; then
        METHOD="Put"
        if [[ ${#TOOL_SNAME} -gt 42 ]]; then # put_XXX
            echo "//if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        elif [[ "${IS_JSON}" == "null" ]]; then
            echo "//if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        else
            echo "if !readonly { register${METHOD}${API_NAME}(s) }" >> "${TOOLS_PATH}"
        fi
    fi

    OP_GET=$(yq '.value.get' <<<"${PATH_INFO}")
    if [[ "$OP_GET" != 'null' ]]; then
        METHOD="Get"
        if [[ ${#TOOL_SNAME} -gt 42 ]]; then # get_XXX
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
