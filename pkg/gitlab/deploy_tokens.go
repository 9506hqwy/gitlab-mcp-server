package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetDeployTokensRequest struct {
	Params *client.GetApiV4DeployTokensParams `json:"params,omitempty"`
}

func registerGetDeployTokens(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetDeployTokensRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_deploy_tokens",
		mcp.WithDescription("Get a list of all deploy tokens across the GitLab instance. This endpoint requires administrator access. This feature was introduced in GitLab 12.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getDeployTokensHandler))
}

func getDeployTokensHandler(ctx context.Context, request mcp.CallToolRequest, req GetDeployTokensRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4DeployTokens(ctx, req.Params, authorizationHeader))
}
