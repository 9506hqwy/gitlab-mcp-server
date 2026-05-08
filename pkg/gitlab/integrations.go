package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostIntegrationsSlackEventsRequest struct {
	Body client.PostApiV4IntegrationsSlackEventsJSONRequestBody `json:"body"`
}

func registerPostIntegrationsSlackEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostIntegrationsSlackEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_integrations_slack_events",
		mcp.WithDescription("Receive Slack events"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postIntegrationsSlackEventsHandler))
}

func postIntegrationsSlackEventsHandler(ctx context.Context, request mcp.CallToolRequest, req PostIntegrationsSlackEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4IntegrationsSlackEvents(ctx, req.Body, authorizationHeader))
}

type PostIntegrationsJiraConnectSubscriptionsRequest struct {
	Body client.PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody `json:"body"`
}

func registerPostIntegrationsJiraConnectSubscriptions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostIntegrationsJiraConnectSubscriptionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_integrations_jira_connect_subscriptions",
		mcp.WithDescription("Subscribes the namespace to the JiraConnectInstallation"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postIntegrationsJiraConnectSubscriptionsHandler))
}

func postIntegrationsJiraConnectSubscriptionsHandler(ctx context.Context, request mcp.CallToolRequest, req PostIntegrationsJiraConnectSubscriptionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4IntegrationsJiraConnectSubscriptions(ctx, req.Body, authorizationHeader))
}
