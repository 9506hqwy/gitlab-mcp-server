package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetWebCommitsPublicKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_web_commits_public_key",
		mcp.WithDescription("This feature was introduced in GitLab 17.4."),
	)

	s.AddTool(tool, getWebCommitsPublicKeyHandler)
}

func getWebCommitsPublicKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4WebCommitsPublicKey(ctx, authorizationHeader))
}
