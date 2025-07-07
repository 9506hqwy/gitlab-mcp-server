package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetMetadata(s *server.MCPServer) {
	tool := mcp.NewTool("get_metadata",
		mcp.WithDescription("This feature was introduced in GitLab 15.2."),
	)

	s.AddTool(tool, getMetadataHandler)
}

func getMetadataHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Metadata(ctx, authorizationHeader))
}
