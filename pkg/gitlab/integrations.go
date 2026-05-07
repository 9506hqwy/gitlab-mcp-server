package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type PostIntegrationsSlackInteractionsRequest struct {
}

func registerPostIntegrationsSlackInteractions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostIntegrationsSlackInteractionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_integrations_slack_interactions",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postIntegrationsSlackInteractionsHandler))
}

func postIntegrationsSlackInteractionsHandler(ctx context.Context, request mcp.CallToolRequest, req PostIntegrationsSlackInteractionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4IntegrationsSlackInteractions(ctx, authorizationHeader))
}

type PostIntegrationsSlackOptionsRequest struct {
}

func registerPostIntegrationsSlackOptions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostIntegrationsSlackOptionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_integrations_slack_options",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postIntegrationsSlackOptionsHandler))
}

func postIntegrationsSlackOptionsHandler(ctx context.Context, request mcp.CallToolRequest, req PostIntegrationsSlackOptionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4IntegrationsSlackOptions(ctx, authorizationHeader))
}
