package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetKeysIdRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of an SSH key"`
}

func registerGetKeysId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetKeysIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_keys_id",
		mcp.WithDescription("Get SSH key with user by ID of an SSH key. Note only administrators can lookup SSH key with user by ID\\ of an SSH key"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getKeysIdHandler))
}

func getKeysIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetKeysIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4KeysId(ctx, req.Id, authorizationHeader))
}

type GetKeysRequest struct {
	Params *client.GetApiV4KeysParams `json:"params,omitempty"`
}

func registerGetKeys(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetKeysRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_keys",
		mcp.WithDescription("You can search for a user that owns a specific SSH key. Note only administrators can lookup SSH key\\ with the fingerprint of an SSH key"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getKeysHandler))
}

func getKeysHandler(ctx context.Context, request mcp.CallToolRequest, req GetKeysRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Keys(ctx, req.Params, authorizationHeader))
}
