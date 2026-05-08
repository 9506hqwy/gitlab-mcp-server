package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetFeaturesRequest struct {
}

func registerGetFeatures(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetFeaturesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_features",
		mcp.WithDescription("Get a list of all persisted features, with its gate values."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getFeaturesHandler))
}

func getFeaturesHandler(ctx context.Context, request mcp.CallToolRequest, req GetFeaturesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Features(ctx, authorizationHeader))
}

type GetFeaturesDefinitionsRequest struct {
}

func registerGetFeaturesDefinitions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetFeaturesDefinitionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_features_definitions",
		mcp.WithDescription("Get a list of all feature definitions."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getFeaturesDefinitionsHandler))
}

func getFeaturesDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetFeaturesDefinitionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4FeaturesDefinitions(ctx, authorizationHeader))
}

type DeleteFeaturesNameRequest struct {
	Name int32 `json:"name" jsonschema:"description=null"`
}

func registerDeleteFeaturesName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteFeaturesNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_features_name",
		mcp.WithDescription("Removes a feature gate. Response is equal when the gate exists, or doesn't."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteFeaturesNameHandler))
}

func deleteFeaturesNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteFeaturesNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4FeaturesName(ctx, req.Name, authorizationHeader))
}

type PostFeaturesNameRequest struct {
	Name int32 `json:"name" jsonschema:"description=null"`

	Body client.PostApiV4FeaturesNameJSONRequestBody `json:"body"`
}

func registerPostFeaturesName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostFeaturesNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_features_name",
		mcp.WithDescription("Set a feature's gate value. If a feature with the given name doesn't exist yet, it's created. The value can be a boolean, or an integer to indicate percentage of time."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postFeaturesNameHandler))
}

func postFeaturesNameHandler(ctx context.Context, request mcp.CallToolRequest, req PostFeaturesNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4FeaturesName(ctx, req.Name, req.Body, authorizationHeader))
}
