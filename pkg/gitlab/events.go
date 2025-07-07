package gitlab

import (
	"context"
	"math"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func registerGetEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_events",
		mcp.WithDescription("This feature was introduced in GitLab 9.3."),
		mcp.WithString("scope",
			mcp.Description("Include all events across a user’s projects (example: all)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("action",
			mcp.Description("Event action to filter on"),
		),
		mcp.WithString("target_type",
			mcp.Description("Event target type to filter on"),

			mcp.Enum("issue", "milestone", "merge_request", "note", "project", "snippet", "user", "wiki", "design"),
		),
		mcp.WithString("before",
			mcp.Description("Include only events created before this date"),
		),
		mcp.WithString("after",
			mcp.Description("Include only events created after this date"),
		),
		mcp.WithString("sort",
			mcp.Description("Return events sorted in ascending and descending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
	)

	s.AddTool(tool, getEventsHandler)
}

func getEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetEvents(request)
	return toResult(c.GetApiV4Events(ctx, &params, authorizationHeader))
}

func parseGetEvents(request mcp.CallToolRequest) client.GetApiV4EventsParams {
	params := client.GetApiV4EventsParams{}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	action := request.GetString("action", "")
	if action != "" {

		params.Action = &action
	}

	target_type := request.GetString("target_type", "")
	if target_type != "" {

		params.TargetType = &target_type
	}

	before := request.GetString("before", "")
	if before != "" {
		before, _ := time.Parse(time.DateOnly, before)
		params.Before = &openapi_types.Date{Time: before}
	}

	after := request.GetString("after", "")
	if after != "" {
		after, _ := time.Parse(time.DateOnly, after)
		params.After = &openapi_types.Date{Time: after}
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	return params
}
