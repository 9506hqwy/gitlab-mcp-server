package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PutApplicationPlanLimitsRequest struct {
	Body client.PutApiV4ApplicationPlanLimitsJSONRequestBody `json:"body"`
}

func registerPutApplicationPlanLimits(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutApplicationPlanLimitsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_application_plan_limits",
		mcp.WithDescription("Modify the limits of a plan on the GitLab instance."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putApplicationPlanLimitsHandler))
}

func putApplicationPlanLimitsHandler(ctx context.Context, request mcp.CallToolRequest, req PutApplicationPlanLimitsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ApplicationPlanLimits(ctx, req.Body, authorizationHeader))
}

type GetApplicationPlanLimitsRequest struct {
	Params *client.GetApiV4ApplicationPlanLimitsParams `json:"params,omitempty"`
}

func registerGetApplicationPlanLimits(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetApplicationPlanLimitsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_application_plan_limits",
		mcp.WithDescription("List the current limits of a plan on the GitLab instance."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getApplicationPlanLimitsHandler))
}

func getApplicationPlanLimitsHandler(ctx context.Context, request mcp.CallToolRequest, req GetApplicationPlanLimitsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ApplicationPlanLimits(ctx, req.Params, authorizationHeader))
}

type GetApplicationAppearanceRequest struct {
}

func registerGetApplicationAppearance(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetApplicationAppearanceRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_application_appearance",
		mcp.WithDescription("Get the current appearance"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getApplicationAppearanceHandler))
}

func getApplicationAppearanceHandler(ctx context.Context, request mcp.CallToolRequest, req GetApplicationAppearanceRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ApplicationAppearance(ctx, authorizationHeader))
}

type GetApplicationStatisticsRequest struct {
}

func registerGetApplicationStatistics(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetApplicationStatisticsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_application_statistics",
		mcp.WithDescription("Get the current application statistics"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getApplicationStatisticsHandler))
}

func getApplicationStatisticsHandler(ctx context.Context, request mcp.CallToolRequest, req GetApplicationStatisticsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ApplicationStatistics(ctx, authorizationHeader))
}
