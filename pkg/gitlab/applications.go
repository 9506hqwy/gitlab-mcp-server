package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetApplications(s *server.MCPServer) {
	tool := mcp.NewTool("applications_",
		mcp.WithDescription("List all registered applications"),
	)

	s.AddTool(tool, getApplicationsHandler)
}

func getApplicationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Applications(ctx, authorizationHeader))
}
