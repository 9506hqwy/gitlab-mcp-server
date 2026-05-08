package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type PostContainerRegistryEventEventsRequest struct {
}

func registerPostContainerRegistryEventEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostContainerRegistryEventEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_container_registry_event_events",
		mcp.WithDescription("This feature was introduced in GitLab 12.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postContainerRegistryEventEventsHandler))
}

func postContainerRegistryEventEventsHandler(ctx context.Context, request mcp.CallToolRequest, req PostContainerRegistryEventEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ContainerRegistryEventEvents(ctx, authorizationHeader))
}
