package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPostContainerRegistryEventEvents(s *server.MCPServer) {
	tool := mcp.NewTool("post_container_registry_event_events",
		mcp.WithDescription("This feature was introduced in GitLab 12.10"),
	)

	s.AddTool(tool, postContainerRegistryEventEventsHandler)
}

func postContainerRegistryEventEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ContainerRegistryEventEvents(ctx, authorizationHeader))
}
