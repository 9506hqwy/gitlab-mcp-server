package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetFeatureFlagsUnleashProjectIdRequest struct {
	ProjectId string                                             `json:"project_id" jsonschema:"description=The ID of a project"`
	Params    *client.GetApiV4FeatureFlagsUnleashProjectIdParams `json:"params,omitempty"`
}

func registerGetFeatureFlagsUnleashProjectId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetFeatureFlagsUnleashProjectIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_feature_flags_unleash_project_id",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getFeatureFlagsUnleashProjectIdHandler))
}

func getFeatureFlagsUnleashProjectIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetFeatureFlagsUnleashProjectIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4FeatureFlagsUnleashProjectId(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type GetFeatureFlagsUnleashProjectIdFeaturesRequest struct {
	ProjectId string                                                     `json:"project_id" jsonschema:"description=The ID of a project"`
	Params    *client.GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams `json:"params,omitempty"`
}

func registerGetFeatureFlagsUnleashProjectIdFeatures(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetFeatureFlagsUnleashProjectIdFeaturesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_feature_flags_unleash_project_id_features",
		mcp.WithDescription("Get a list of features (deprecated, v2 client support)"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getFeatureFlagsUnleashProjectIdFeaturesHandler))
}

func getFeatureFlagsUnleashProjectIdFeaturesHandler(ctx context.Context, request mcp.CallToolRequest, req GetFeatureFlagsUnleashProjectIdFeaturesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type GetFeatureFlagsUnleashProjectIdClientFeaturesRequest struct {
	ProjectId string                                                           `json:"project_id" jsonschema:"description=The ID of a project"`
	Params    *client.GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams `json:"params,omitempty"`
}

func registerGetFeatureFlagsUnleashProjectIdClientFeatures(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetFeatureFlagsUnleashProjectIdClientFeaturesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_feature_flags_unleash_project_id_client_features",
		mcp.WithDescription("Get a list of features"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getFeatureFlagsUnleashProjectIdClientFeaturesHandler))
}

func getFeatureFlagsUnleashProjectIdClientFeaturesHandler(ctx context.Context, request mcp.CallToolRequest, req GetFeatureFlagsUnleashProjectIdClientFeaturesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type PostFeatureFlagsUnleashProjectIdClientRegisterRequest struct {
	ProjectId string `json:"project_id" jsonschema:"description=The ID of a project"`

	Body client.PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody `json:"body"`
}

func registerPostFeatureFlagsUnleashProjectIdClientRegister(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostFeatureFlagsUnleashProjectIdClientRegisterRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_feature_flags_unleash_project_id_client_register",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postFeatureFlagsUnleashProjectIdClientRegisterHandler))
}

func postFeatureFlagsUnleashProjectIdClientRegisterHandler(ctx context.Context, request mcp.CallToolRequest, req PostFeatureFlagsUnleashProjectIdClientRegisterRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4FeatureFlagsUnleashProjectIdClientRegister(ctx, req.ProjectId, req.Body, authorizationHeader))
}

type PostFeatureFlagsUnleashProjectIdClientMetricsRequest struct {
	ProjectId string `json:"project_id" jsonschema:"description=The ID of a project"`

	Body client.PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody `json:"body"`
}

func registerPostFeatureFlagsUnleashProjectIdClientMetrics(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostFeatureFlagsUnleashProjectIdClientMetricsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_feature_flags_unleash_project_id_client_metrics",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postFeatureFlagsUnleashProjectIdClientMetricsHandler))
}

func postFeatureFlagsUnleashProjectIdClientMetricsHandler(ctx context.Context, request mcp.CallToolRequest, req PostFeatureFlagsUnleashProjectIdClientMetricsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4FeatureFlagsUnleashProjectIdClientMetrics(ctx, req.ProjectId, req.Body, authorizationHeader))
}
