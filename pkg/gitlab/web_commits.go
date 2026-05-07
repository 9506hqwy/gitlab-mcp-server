package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type GetWebCommitsPublicKeyRequest struct {
}

func registerGetWebCommitsPublicKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetWebCommitsPublicKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_web_commits_public_key",
		mcp.WithDescription("This feature was introduced in GitLab 17.4."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getWebCommitsPublicKeyHandler))
}

func getWebCommitsPublicKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetWebCommitsPublicKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4WebCommitsPublicKey(ctx, authorizationHeader))
}
