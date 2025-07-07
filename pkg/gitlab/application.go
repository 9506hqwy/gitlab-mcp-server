package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetApplicationPlanLimits(s *server.MCPServer) {
	tool := mcp.NewTool("get_application_plan_limits",
		mcp.WithDescription("List the current limits of a plan on the GitLab instance."),
		mcp.WithString("plan_name",
			mcp.Description("Name of the plan to get the limits from. Default: default. (default: default)"),

			mcp.Enum("default", "free", "bronze", "silver", "premium", "gold", "ultimate", "ultimate_trial", "ultimate_trial_paid_customer", "premium_trial", "opensource"),
		),
	)

	s.AddTool(tool, getApplicationPlanLimitsHandler)
}

func getApplicationPlanLimitsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetApplicationPlanLimits(request)
	return toResult(c.GetApiV4ApplicationPlanLimits(ctx, &params, authorizationHeader))
}

func parseGetApplicationPlanLimits(request mcp.CallToolRequest) client.GetApiV4ApplicationPlanLimitsParams {
	params := client.GetApiV4ApplicationPlanLimitsParams{}

	plan_name := request.GetString("plan_name", "")
	if plan_name != "" {

		params.PlanName = &plan_name
	}

	return params
}

func registerGetApplicationAppearance(s *server.MCPServer) {
	tool := mcp.NewTool("get_application_appearance",
		mcp.WithDescription("Get the current appearance"),
	)

	s.AddTool(tool, getApplicationAppearanceHandler)
}

func getApplicationAppearanceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ApplicationAppearance(ctx, authorizationHeader))
}

func registerGetApplicationStatistics(s *server.MCPServer) {
	tool := mcp.NewTool("get_application_statistics",
		mcp.WithDescription("Get the current application statistics"),
	)

	s.AddTool(tool, getApplicationStatisticsHandler)
}

func getApplicationStatisticsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ApplicationStatistics(ctx, authorizationHeader))
}
