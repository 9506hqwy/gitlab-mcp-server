package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetVersion(s *server.MCPServer) {
	tool := mcp.NewTool("version_",
		mcp.WithDescription("This feature was introduced in GitLab 8.13 and deprecated in 15.5. We recommend you instead use the Metadata API."),
	)

	s.AddTool(tool, getVersionHandler)
}

func getVersionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Version(ctx, authorizationHeader))
}
