package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetUsageDataMetricDefinitions(s *server.MCPServer) {
	tool := mcp.NewTool("get_usage_data_metric_definitions",
		mcp.WithDescription("This feature was introduced in GitLab 13.11."),
		mcp.WithBoolean("include_paths",
			mcp.Description("Include file paths in the metric definitions (example: true) (default: false)"),
		),
	)

	s.AddTool(tool, getUsageDataMetricDefinitionsHandler)
}

func getUsageDataMetricDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetUsageDataMetricDefinitions(request)
	return toResult(c.GetApiV4UsageDataMetricDefinitions(ctx, &params, authorizationHeader))
}

func parseGetUsageDataMetricDefinitions(request mcp.CallToolRequest) client.GetApiV4UsageDataMetricDefinitionsParams {
	params := client.GetApiV4UsageDataMetricDefinitionsParams{}

	include_paths := request.GetBool("include_paths", false)
	params.IncludePaths = &include_paths

	return params
}

func registerGetUsageDataServicePing(s *server.MCPServer) {
	tool := mcp.NewTool("get_usage_data_service_ping",
		mcp.WithDescription("Introduces in Gitlab 16.9. Requires personal access token with read_service_ping scope."),
	)

	s.AddTool(tool, getUsageDataServicePingHandler)
}

func getUsageDataServicePingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataServicePing(ctx, authorizationHeader))
}

func registerGetUsageDataNonSqlMetrics(s *server.MCPServer) {
	tool := mcp.NewTool("get_usage_data_non_sql_metrics",
		mcp.WithDescription("This feature was introduced in GitLab 13.11."),
	)

	s.AddTool(tool, getUsageDataNonSqlMetricsHandler)
}

func getUsageDataNonSqlMetricsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataNonSqlMetrics(ctx, authorizationHeader))
}

func registerGetUsageDataQueries(s *server.MCPServer) {
	tool := mcp.NewTool("get_usage_data_queries",
		mcp.WithDescription("This feature was introduced in GitLab 13.11."),
	)

	s.AddTool(tool, getUsageDataQueriesHandler)
}

func getUsageDataQueriesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataQueries(ctx, authorizationHeader))
}
