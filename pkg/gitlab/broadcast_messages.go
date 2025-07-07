package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetBroadcastMessages(s *server.MCPServer) {
	tool := mcp.NewTool("get_broadcast_messages",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getBroadcastMessagesHandler)
}

func getBroadcastMessagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetBroadcastMessages(request)
	return toResult(c.GetApiV4BroadcastMessages(ctx, &params, authorizationHeader))
}

func parseGetBroadcastMessages(request mcp.CallToolRequest) client.GetApiV4BroadcastMessagesParams {
	params := client.GetApiV4BroadcastMessagesParams{}

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

func registerGetBroadcastMessagesId(s *server.MCPServer) {
	tool := mcp.NewTool("get_broadcast_messages_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithNumber("id",
			mcp.Description("Broadcast message ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getBroadcastMessagesIdHandler)
}

func getBroadcastMessagesIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4BroadcastMessagesId(ctx, id, authorizationHeader))
}
