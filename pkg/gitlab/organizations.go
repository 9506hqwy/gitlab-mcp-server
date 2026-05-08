package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostOrganizationsRequest struct {
	Body client.PostApiV4OrganizationsJSONRequestBody `json:"body"`
}

func registerPostOrganizations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostOrganizationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_organizations",
		mcp.WithDescription("This feature was introduced in GitLab 17.5. \\ This feature is currently in an experimental state. \\ This feature is behind the 'allow_organization_creation' feature flag."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postOrganizationsHandler))
}

func postOrganizationsHandler(ctx context.Context, request mcp.CallToolRequest, req PostOrganizationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4Organizations(ctx, req.Body, authorizationHeader))
}
