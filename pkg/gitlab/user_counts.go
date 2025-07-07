package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetUserCounts(s *server.MCPServer) {
	tool := mcp.NewTool("get_user_counts",
		mcp.WithDescription("Assigned open issues, assigned MRs and pending todos count"),
	)

	s.AddTool(tool, getUserCountsHandler)
}

func getUserCountsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UserCounts(ctx, authorizationHeader))
}
