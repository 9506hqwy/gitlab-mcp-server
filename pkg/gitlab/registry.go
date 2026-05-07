package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetRegistryRepositoriesIdRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID of the repository"`
	Params *client.GetApiV4RegistryRepositoriesIdParams `json:"params,omitempty"`
}

func registerGetRegistryRepositoriesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetRegistryRepositoriesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_registry_repositories_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getRegistryRepositoriesIdHandler))
}

func getRegistryRepositoriesIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetRegistryRepositoriesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4RegistryRepositoriesId(ctx, req.Id, req.Params, authorizationHeader))
}
