package gitlab

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetRunners(s *server.MCPServer) {
	tool := mcp.NewTool("runners_",
		mcp.WithDescription("Get runners available for user"),
		mcp.WithString("scope",
			mcp.Description("Deprecated: Use `type` or `status` instead. The scope of runners to return"),

			mcp.Enum("specific", "shared", "instance_type", "group_type", "project_type", "active", "paused", "online", "offline", "never_contacted", "stale"),
		),
		mcp.WithString("type",
			mcp.Description("The type of runners to return"),

			mcp.Enum("instance_type", "group_type", "project_type"),
		),
		mcp.WithBoolean("paused",
			mcp.Description("Whether to include only runners that are accepting or ignoring new jobs"),
		),
		mcp.WithString("status",
			mcp.Description("The status of runners to return"),

			mcp.Enum("active", "paused", "online", "offline", "never_contacted", "stale"),
		),
		mcp.WithString("tag_list",
			mcp.Description("A list of runner tags (example: ['macos', 'shell'])"),
		),
		mcp.WithString("version_prefix",
			mcp.Description("The version prefix of runners to return (example: '15.1.' or '16.')"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getRunnersHandler)
}

func getRunnersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetRunners(request)
	return toResult(c.GetApiV4Runners(ctx, &params, authorizationHeader))
}

func parseGetRunners(request mcp.CallToolRequest) client.GetApiV4RunnersParams {
	params := client.GetApiV4RunnersParams{}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	_type := request.GetString("type", "")
	if _type != "" {

		params.Type = &_type
	}

	paused := request.GetBool("paused", false)
	params.Paused = &paused

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	tag_list := request.GetString("tag_list", "")
	if tag_list != "" {
		tag_list := strings.Split(tag_list, ",")
		params.TagList = &tag_list
	}

	version_prefix := request.GetString("version_prefix", "")
	if version_prefix != "" {

		params.VersionPrefix = &version_prefix
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

	return params
}

func registerGetRunnersAll(s *server.MCPServer) {
	tool := mcp.NewTool("runners_all",
		mcp.WithDescription("Get a list of all runners in the GitLab instance (shared and project). Access is restricted to users with either administrator access or auditor access."),
		mcp.WithString("scope",
			mcp.Description("Deprecated: Use `type` or `status` instead. The scope of runners to return"),

			mcp.Enum("specific", "shared", "instance_type", "group_type", "project_type", "active", "paused", "online", "offline", "never_contacted", "stale"),
		),
		mcp.WithString("type",
			mcp.Description("The type of runners to return"),

			mcp.Enum("instance_type", "group_type", "project_type"),
		),
		mcp.WithBoolean("paused",
			mcp.Description("Whether to include only runners that are accepting or ignoring new jobs"),
		),
		mcp.WithString("status",
			mcp.Description("The status of runners to return"),

			mcp.Enum("active", "paused", "online", "offline", "never_contacted", "stale"),
		),
		mcp.WithString("tag_list",
			mcp.Description("A list of runner tags (example: ['macos', 'shell'])"),
		),
		mcp.WithString("version_prefix",
			mcp.Description("The version prefix of runners to return (example: '15.1.' or '16.')"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getRunnersAllHandler)
}

func getRunnersAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetRunnersAll(request)
	return toResult(c.GetApiV4RunnersAll(ctx, &params, authorizationHeader))
}

func parseGetRunnersAll(request mcp.CallToolRequest) client.GetApiV4RunnersAllParams {
	params := client.GetApiV4RunnersAllParams{}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	_type := request.GetString("type", "")
	if _type != "" {

		params.Type = &_type
	}

	paused := request.GetBool("paused", false)
	params.Paused = &paused

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	tag_list := request.GetString("tag_list", "")
	if tag_list != "" {
		tag_list := strings.Split(tag_list, ",")
		params.TagList = &tag_list
	}

	version_prefix := request.GetString("version_prefix", "")
	if version_prefix != "" {

		params.VersionPrefix = &version_prefix
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

	return params
}

func registerGetRunnersId(s *server.MCPServer) {
	tool := mcp.NewTool("runners_id",
		mcp.WithDescription("At least the Maintainer role is required to get runner details at the project and group level. Instance-level runner details via this endpoint are available to all signed in users."),
		mcp.WithNumber("id",
			mcp.Description("The ID of a runner"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getRunnersIdHandler)
}

func getRunnersIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4RunnersId(ctx, id, authorizationHeader))
}

func registerGetRunnersIdManagers(s *server.MCPServer) {
	tool := mcp.NewTool("runners_id_managers",
		mcp.WithDescription("Get a list of all runner's managers"),
		mcp.WithNumber("id",
			mcp.Description("The ID of a runner"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getRunnersIdManagersHandler)
}

func getRunnersIdManagersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4RunnersIdManagers(ctx, id, authorizationHeader))
}

func registerGetRunnersIdJobs(s *server.MCPServer) {
	tool := mcp.NewTool("runners_id_jobs",
		mcp.WithDescription("List jobs that are being processed or were processed by the specified runner. The list of jobs is limited to projects where the user has at least the Reporter role."),
		mcp.WithNumber("id",
			mcp.Description("The ID of a runner"),
			mcp.Required(),
		),
		mcp.WithString("system_id",
			mcp.Description("System ID associated with the runner manager"),
		),
		mcp.WithString("status",
			mcp.Description("Status of the job"),

			mcp.Enum("created", "waiting_for_resource", "preparing", "waiting_for_callback", "pending", "running", "success", "failed", "canceling", "canceled", "skipped", "manual", "scheduled"),
		),
		mcp.WithString("order_by",
			mcp.Description("Order by `id`"),

			mcp.Enum("id"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort by `asc` or `desc` order. Specify `order_by` as well, including for `id` (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("cursor",
			mcp.Description("Cursor for obtaining the next set of records"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getRunnersIdJobsHandler)
}

func getRunnersIdJobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetRunnersIdJobs(request)
	return toResult(c.GetApiV4RunnersIdJobs(ctx, id, &params, authorizationHeader))
}

func parseGetRunnersIdJobs(request mcp.CallToolRequest) client.GetApiV4RunnersIdJobsParams {
	params := client.GetApiV4RunnersIdJobsParams{}

	system_id := request.GetString("system_id", "")
	if system_id != "" {

		params.SystemId = &system_id
	}

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	cursor := request.GetString("cursor", "")
	if cursor != "" {

		params.Cursor = &cursor
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

	return params
}
