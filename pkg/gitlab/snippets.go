package gitlab

import (
	"context"
	"math"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetSnippets(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithString("created_after",
			mcp.Description("Return snippets created after the specified time"),
		),
		mcp.WithString("created_before",
			mcp.Description("Return snippets created before the specified time"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getSnippetsHandler)
}

func getSnippetsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetSnippets(request)
	return toResult(c.GetApiV4Snippets(ctx, &params, authorizationHeader))
}

func parseGetSnippets(request mcp.CallToolRequest) client.GetApiV4SnippetsParams {
	params := client.GetApiV4SnippetsParams{}

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
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

func registerGetSnippetsPublic(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets_public",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithString("created_after",
			mcp.Description("Return snippets created after the specified time"),
		),
		mcp.WithString("created_before",
			mcp.Description("Return snippets created before the specified time"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getSnippetsPublicHandler)
}

func getSnippetsPublicHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetSnippetsPublic(request)
	return toResult(c.GetApiV4SnippetsPublic(ctx, &params, authorizationHeader))
}

func parseGetSnippetsPublic(request mcp.CallToolRequest) client.GetApiV4SnippetsPublicParams {
	params := client.GetApiV4SnippetsPublicParams{}

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
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

func registerGetSnippetsAll(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets_all",
		mcp.WithDescription("This feature was introduced in GitLab 16.3."),
		mcp.WithString("created_after",
			mcp.Description("Return snippets created after the specified time"),
		),
		mcp.WithString("created_before",
			mcp.Description("Return snippets created before the specified time"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("repository_storage",
			mcp.Description("Filter by repository storage used by the snippet"),
		),
	)

	s.AddTool(tool, getSnippetsAllHandler)
}

func getSnippetsAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetSnippetsAll(request)
	return toResult(c.GetApiV4SnippetsAll(ctx, &params, authorizationHeader))
}

func parseGetSnippetsAll(request mcp.CallToolRequest) client.GetApiV4SnippetsAllParams {
	params := client.GetApiV4SnippetsAllParams{}

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
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

	repository_storage := request.GetString("repository_storage", "")
	if repository_storage != "" {

		params.RepositoryStorage = &repository_storage
	}

	return params
}

func registerGetSnippetsId(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithNumber("id",
			mcp.Description("The ID of a snippet"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getSnippetsIdHandler)
}

func getSnippetsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4SnippetsId(ctx, id, authorizationHeader))
}

func registerGetSnippetsIdRaw(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets_id_raw",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithNumber("id",
			mcp.Description("The ID of a snippet"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getSnippetsIdRawHandler)
}

func getSnippetsIdRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4SnippetsIdRaw(ctx, id, authorizationHeader))
}

func registerGetSnippetsIdFilesRefFilePathRaw(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets_id_files_ref_file_path_raw",
		mcp.WithDescription("Get raw snippet file contents from the repository"),
		mcp.WithString("file_path",
			mcp.Description("The url encoded path to the file, e.g. lib%2Fclass%2Erb"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("The name of branch, tag or commit"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getSnippetsIdFilesRefFilePathRawHandler)
}

func getSnippetsIdFilesRefFilePathRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	ref := request.GetString("ref", "")
	file_path := request.GetString("file_path", "")

	return toResult(c.GetApiV4SnippetsIdFilesRefFilePathRaw(ctx, id, ref, file_path, authorizationHeader))
}

func registerGetSnippetsIdUserAgentDetail(s *server.MCPServer) {
	tool := mcp.NewTool("get_snippets_id_user_agent_detail",
		mcp.WithDescription("Get the user agent details for a snippet"),
		mcp.WithNumber("id",
			mcp.Description("The ID of a snippet"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getSnippetsIdUserAgentDetailHandler)
}

func getSnippetsIdUserAgentDetailHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4SnippetsIdUserAgentDetail(ctx, id, authorizationHeader))
}
