package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPostIntegrationsSlackInteractions(s *server.MCPServer) {
	tool := mcp.NewTool("post_integrations_slack_interactions",
		mcp.WithDescription("null"),
	)

	s.AddTool(tool, postIntegrationsSlackInteractionsHandler)
}

func postIntegrationsSlackInteractionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4IntegrationsSlackInteractions(ctx, authorizationHeader))
}

func registerPostIntegrationsSlackOptions(s *server.MCPServer) {
	tool := mcp.NewTool("post_integrations_slack_options",
		mcp.WithDescription("null"),
	)

	s.AddTool(tool, postIntegrationsSlackOptionsHandler)
}

func postIntegrationsSlackOptionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4IntegrationsSlackOptions(ctx, authorizationHeader))
}
