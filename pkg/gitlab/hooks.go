package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetHooks(s *server.MCPServer) {
	tool := mcp.NewTool("get_hooks",
		mcp.WithDescription("Get a list of all system hooks"),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getHooksHandler)
}

func getHooksHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetHooks(request)
	return toResult(c.GetApiV4Hooks(ctx, &params, authorizationHeader))
}

func parseGetHooks(request mcp.CallToolRequest) client.GetApiV4HooksParams {
	params := client.GetApiV4HooksParams{}

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

func registerGetHooksHookId(s *server.MCPServer) {
	tool := mcp.NewTool("get_hooks_hook_id",
		mcp.WithDescription("Get a system hook by its ID. Introduced in GitLab 14.9."),
		mcp.WithNumber("hook_id",
			mcp.Description("The ID of the system hook"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getHooksHookIdHandler)
}

func getHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	hook_id := int32(request.GetInt("hook_id", math.MinInt))

	return toResult(c.GetApiV4HooksHookId(ctx, hook_id, authorizationHeader))
}
