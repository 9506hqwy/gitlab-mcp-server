package gitlab

import (
	"context"
	"math"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func registerGetPersonalAccessTokensSelf(s *server.MCPServer) {
	tool := mcp.NewTool("personal_access_tokens_self",
		mcp.WithDescription("Get the details of a personal access token by passing it to the API in a header"),
	)

	s.AddTool(tool, getPersonalAccessTokensSelfHandler)
}

func getPersonalAccessTokensSelfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PersonalAccessTokensSelf(ctx, authorizationHeader))
}

func registerGetPersonalAccessTokensSelfAssociations(s *server.MCPServer) {
	tool := mcp.NewTool("personal_access_tokens_self_associations",
		mcp.WithDescription("Get groups and projects this personal access token can access by passing it to the API in a header"),
		mcp.WithNumber("min_access_level",
			mcp.Description("Limit by minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getPersonalAccessTokensSelfAssociationsHandler)
}

func getPersonalAccessTokensSelfAssociationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetPersonalAccessTokensSelfAssociations(request)
	return toResult(c.GetApiV4PersonalAccessTokensSelfAssociations(ctx, &params, authorizationHeader))
}

func parseGetPersonalAccessTokensSelfAssociations(request mcp.CallToolRequest) client.GetApiV4PersonalAccessTokensSelfAssociationsParams {
	params := client.GetApiV4PersonalAccessTokensSelfAssociationsParams{}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	return params
}

func registerGetPersonalAccessTokens(s *server.MCPServer) {
	tool := mcp.NewTool("personal_access_tokens_",
		mcp.WithDescription("Get all personal access tokens the authenticated user has access to."),
		mcp.WithNumber("user_id",
			mcp.Description("Filter PATs by User ID (example: 2)"),
		),
		mcp.WithBoolean("revoked",
			mcp.Description("Filter tokens where revoked state matches parameter"),
		),
		mcp.WithString("state",
			mcp.Description("Filter tokens which are either active or not (example: active)"),

			mcp.Enum("active", "inactive"),
		),
		mcp.WithString("created_before",
			mcp.Description("Filter tokens which were created before given datetime (example: 2022-01-01)"),
		),
		mcp.WithString("created_after",
			mcp.Description("Filter tokens which were created after given datetime (example: 2021-01-01)"),
		),
		mcp.WithString("last_used_before",
			mcp.Description("Filter tokens which were used before given datetime (example: 2021-01-01)"),
		),
		mcp.WithString("last_used_after",
			mcp.Description("Filter tokens which were used after given datetime (example: 2022-01-01)"),
		),
		mcp.WithString("expires_before",
			mcp.Description("Filter tokens which expire before given datetime (example: 2022-01-01)"),
		),
		mcp.WithString("expires_after",
			mcp.Description("Filter tokens which expire after given datetime (example: 2021-01-01)"),
		),
		mcp.WithString("search",
			mcp.Description("Filters tokens by name (example: token)"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort tokens (example: created_at_desc)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getPersonalAccessTokensHandler)
}

func getPersonalAccessTokensHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetPersonalAccessTokens(request)
	return toResult(c.GetApiV4PersonalAccessTokens(ctx, &params, authorizationHeader))
}

func parseGetPersonalAccessTokens(request mcp.CallToolRequest) client.GetApiV4PersonalAccessTokensParams {
	params := client.GetApiV4PersonalAccessTokensParams{}

	user_id := request.GetInt("user_id", math.MinInt)
	if user_id != math.MinInt {
		user_id := int32(user_id)
		params.UserId = &user_id
	}

	revoked := request.GetBool("revoked", false)
	params.Revoked = &revoked

	state := request.GetString("state", "")
	if state != "" {

		params.State = &state
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
	}

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	last_used_before := request.GetString("last_used_before", "")
	if last_used_before != "" {
		last_used_before, _ := time.Parse(time.RFC3339, last_used_before)
		params.LastUsedBefore = &last_used_before
	}

	last_used_after := request.GetString("last_used_after", "")
	if last_used_after != "" {
		last_used_after, _ := time.Parse(time.RFC3339, last_used_after)
		params.LastUsedAfter = &last_used_after
	}

	expires_before := request.GetString("expires_before", "")
	if expires_before != "" {
		expires_before, _ := time.Parse(time.DateOnly, expires_before)
		params.ExpiresBefore = &openapi_types.Date{Time: expires_before}
	}

	expires_after := request.GetString("expires_after", "")
	if expires_after != "" {
		expires_after, _ := time.Parse(time.DateOnly, expires_after)
		params.ExpiresAfter = &openapi_types.Date{Time: expires_after}
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	return params
}

func registerGetPersonalAccessTokensId(s *server.MCPServer) {
	tool := mcp.NewTool("personal_access_tokens_id",
		mcp.WithDescription("Get a personal access token by using the ID of the personal access token."),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPersonalAccessTokensIdHandler)
}

func getPersonalAccessTokensIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4PersonalAccessTokensId(ctx, id, authorizationHeader))
}
