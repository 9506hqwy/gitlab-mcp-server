package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetFeatureFlagsUnleashProjectId(s *server.MCPServer) {
	tool := mcp.NewTool("feature_flags_unleash_project_id",
		mcp.WithDescription("null"),
		mcp.WithString("project_id",
			mcp.Description("The ID of a project"),
			mcp.Required(),
		),
		mcp.WithString("instance_id",
			mcp.Description("The instance ID of Unleash Client"),
		),
		mcp.WithString("app_name",
			mcp.Description("The application name of Unleash Client"),
		),
	)

	s.AddTool(tool, getFeatureFlagsUnleashProjectIdHandler)
}

func getFeatureFlagsUnleashProjectIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseGetFeatureFlagsUnleashProjectId(request)
	return toResult(c.GetApiV4FeatureFlagsUnleashProjectId(ctx, project_id, &params, authorizationHeader))
}

func parseGetFeatureFlagsUnleashProjectId(request mcp.CallToolRequest) client.GetApiV4FeatureFlagsUnleashProjectIdParams {
	params := client.GetApiV4FeatureFlagsUnleashProjectIdParams{}

	instance_id := request.GetString("instance_id", "")
	if instance_id != "" {

		params.InstanceId = &instance_id
	}

	app_name := request.GetString("app_name", "")
	if app_name != "" {

		params.AppName = &app_name
	}

	return params
}

func registerGetFeatureFlagsUnleashProjectIdFeatures(s *server.MCPServer) {
	tool := mcp.NewTool("feature_flags_unleash_project_id_features",
		mcp.WithDescription("Get a list of features (deprecated, v2 client support)"),
		mcp.WithString("project_id",
			mcp.Description("The ID of a project"),
			mcp.Required(),
		),
		mcp.WithString("instance_id",
			mcp.Description("The instance ID of Unleash Client"),
		),
		mcp.WithString("app_name",
			mcp.Description("The application name of Unleash Client"),
		),
	)

	s.AddTool(tool, getFeatureFlagsUnleashProjectIdFeaturesHandler)
}

func getFeatureFlagsUnleashProjectIdFeaturesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseGetFeatureFlagsUnleashProjectIdFeatures(request)
	return toResult(c.GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx, project_id, &params, authorizationHeader))
}

func parseGetFeatureFlagsUnleashProjectIdFeatures(request mcp.CallToolRequest) client.GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams {
	params := client.GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams{}

	instance_id := request.GetString("instance_id", "")
	if instance_id != "" {

		params.InstanceId = &instance_id
	}

	app_name := request.GetString("app_name", "")
	if app_name != "" {

		params.AppName = &app_name
	}

	return params
}

func registerGetFeatureFlagsUnleashProjectIdClientFeatures(s *server.MCPServer) {
	tool := mcp.NewTool("feature_flags_unleash_project_id_client_features",
		mcp.WithDescription("Get a list of features"),
		mcp.WithString("project_id",
			mcp.Description("The ID of a project"),
			mcp.Required(),
		),
		mcp.WithString("instance_id",
			mcp.Description("The instance ID of Unleash Client"),
		),
		mcp.WithString("app_name",
			mcp.Description("The application name of Unleash Client"),
		),
	)

	s.AddTool(tool, getFeatureFlagsUnleashProjectIdClientFeaturesHandler)
}

func getFeatureFlagsUnleashProjectIdClientFeaturesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseGetFeatureFlagsUnleashProjectIdClientFeatures(request)
	return toResult(c.GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx, project_id, &params, authorizationHeader))
}

func parseGetFeatureFlagsUnleashProjectIdClientFeatures(request mcp.CallToolRequest) client.GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams {
	params := client.GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams{}

	instance_id := request.GetString("instance_id", "")
	if instance_id != "" {

		params.InstanceId = &instance_id
	}

	app_name := request.GetString("app_name", "")
	if app_name != "" {

		params.AppName = &app_name
	}

	return params
}
