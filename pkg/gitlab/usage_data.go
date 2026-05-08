package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostUsageDataIncrementCounterRequest struct {
	Body client.PostApiV4UsageDataIncrementCounterJSONRequestBody `json:"body"`
}

func registerPostUsageDataIncrementCounter(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostUsageDataIncrementCounterRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_usage_data_increment_counter",
		mcp.WithDescription("This feature was introduced in GitLab 13.4."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postUsageDataIncrementCounterHandler))
}

func postUsageDataIncrementCounterHandler(ctx context.Context, request mcp.CallToolRequest, req PostUsageDataIncrementCounterRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4UsageDataIncrementCounter(ctx, req.Body, authorizationHeader))
}

type PostUsageDataIncrementUniqueUsersRequest struct {
	Body client.PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody `json:"body"`
}

func registerPostUsageDataIncrementUniqueUsers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostUsageDataIncrementUniqueUsersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_usage_data_increment_unique_users",
		mcp.WithDescription("Track usage data event for the current user"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postUsageDataIncrementUniqueUsersHandler))
}

func postUsageDataIncrementUniqueUsersHandler(ctx context.Context, request mcp.CallToolRequest, req PostUsageDataIncrementUniqueUsersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4UsageDataIncrementUniqueUsers(ctx, req.Body, authorizationHeader))
}

type PostUsageDataTrackEventsRequest struct {
	Body client.PostApiV4UsageDataTrackEventsJSONRequestBody `json:"body"`
}

func registerPostUsageDataTrackEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostUsageDataTrackEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_usage_data_track_events",
		mcp.WithDescription("This feature was introduced in GitLab 17.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postUsageDataTrackEventsHandler))
}

func postUsageDataTrackEventsHandler(ctx context.Context, request mcp.CallToolRequest, req PostUsageDataTrackEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4UsageDataTrackEvents(ctx, req.Body, authorizationHeader))
}

type GetUsageDataMetricDefinitionsRequest struct {
	Params *client.GetApiV4UsageDataMetricDefinitionsParams `json:"params,omitempty"`
}

func registerGetUsageDataMetricDefinitions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsageDataMetricDefinitionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_usage_data_metric_definitions",
		mcp.WithDescription("This feature was introduced in GitLab 13.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsageDataMetricDefinitionsHandler))
}

func getUsageDataMetricDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsageDataMetricDefinitionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataMetricDefinitions(ctx, req.Params, authorizationHeader))
}

type GetUsageDataServicePingRequest struct {
}

func registerGetUsageDataServicePing(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsageDataServicePingRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_usage_data_service_ping",
		mcp.WithDescription("Introduces in Gitlab 16.9. Requires personal access token with read_service_ping scope."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsageDataServicePingHandler))
}

func getUsageDataServicePingHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsageDataServicePingRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataServicePing(ctx, authorizationHeader))
}

type PostUsageDataTrackEventRequest struct {
	Body client.PostApiV4UsageDataTrackEventJSONRequestBody `json:"body"`
}

func registerPostUsageDataTrackEvent(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostUsageDataTrackEventRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_usage_data_track_event",
		mcp.WithDescription("This feature was introduced in GitLab 16.2."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postUsageDataTrackEventHandler))
}

func postUsageDataTrackEventHandler(ctx context.Context, request mcp.CallToolRequest, req PostUsageDataTrackEventRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4UsageDataTrackEvent(ctx, req.Body, authorizationHeader))
}

type GetUsageDataNonSqlMetricsRequest struct {
}

func registerGetUsageDataNonSqlMetrics(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsageDataNonSqlMetricsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_usage_data_non_sql_metrics",
		mcp.WithDescription("This feature was introduced in GitLab 13.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsageDataNonSqlMetricsHandler))
}

func getUsageDataNonSqlMetricsHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsageDataNonSqlMetricsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataNonSqlMetrics(ctx, authorizationHeader))
}

type GetUsageDataQueriesRequest struct {
}

func registerGetUsageDataQueries(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsageDataQueriesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_usage_data_queries",
		mcp.WithDescription("This feature was introduced in GitLab 13.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsageDataQueriesHandler))
}

func getUsageDataQueriesHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsageDataQueriesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsageDataQueries(ctx, authorizationHeader))
}
