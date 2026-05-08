package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostDeployKeysRequest struct {
	Body client.PostApiV4DeployKeysJSONRequestBody `json:"body"`
}

func registerPostDeployKeys(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostDeployKeysRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_deploy_keys",
		mcp.WithDescription("Create a deploy key for the GitLab instance. This endpoint requires administrator access."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postDeployKeysHandler))
}

func postDeployKeysHandler(ctx context.Context, request mcp.CallToolRequest, req PostDeployKeysRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4DeployKeys(ctx, req.Body, authorizationHeader))
}

type GetDeployKeysRequest struct {
	Params *client.GetApiV4DeployKeysParams `json:"params,omitempty"`
}

func registerGetDeployKeys(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetDeployKeysRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_deploy_keys",
		mcp.WithDescription("Get a list of all deploy keys across all projects of the GitLab instance. This endpoint requires administrator access and is not available on GitLab.com."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getDeployKeysHandler))
}

func getDeployKeysHandler(ctx context.Context, request mcp.CallToolRequest, req GetDeployKeysRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4DeployKeys(ctx, req.Params, authorizationHeader))
}
