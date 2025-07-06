package gitlab

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func registerGetUsersIdEvents(s *server.MCPServer) {
	tool := mcp.NewTool("users_id_events",
		mcp.WithDescription("This feature was introduced in GitLab 8.13."),
		mcp.WithString("id",
			mcp.Description("The ID or username of the user"),
			mcp.Required(),
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

	s.AddTool(tool, getUsersIdEventsHandler)
}

func getUsersIdEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetUsersIdEvents(request)
	return toResult(c.GetApiV4UsersIdEvents(ctx, id, &params, authorizationHeader))
}

func parseGetUsersIdEvents(request mcp.CallToolRequest) client.GetApiV4UsersIdEventsParams {
	params := client.GetApiV4UsersIdEventsParams{}

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

func registerGetUsersUserIdProjects(s *server.MCPServer) {
	tool := mcp.NewTool("users_user_id_pjs",
		mcp.WithDescription("Get a user projects"),
		mcp.WithString("user_id",
			mcp.Description("The ID or username of the user"),
			mcp.Required(),
		),
		mcp.WithString("order_by",
			mcp.Description("Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. (default: created_at)"),

			mcp.Enum("id", "name", "path", "created_at", "updated_at", "last_activity_at", "similarity", "star_count", "storage_size", "repository_size", "wiki_size", "packages_size"),
		),
		mcp.WithString("sort",
			mcp.Description("Return projects sorted in ascending and descending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of projects matching the search criteria"),
		),
		mcp.WithBoolean("search_namespaces",
			mcp.Description("Include ancestor namespaces when matching search criteria"),
		),
		mcp.WithBoolean("owned",
			mcp.Description("Limit by owned by authenticated user (default: false)"),
		),
		mcp.WithBoolean("starred",
			mcp.Description("Limit by starred status (default: false)"),
		),
		mcp.WithBoolean("imported",
			mcp.Description("Limit by imported by authenticated user (default: false)"),
		),
		mcp.WithBoolean("membership",
			mcp.Description("Limit by projects that the current user is a member of (default: false)"),
		),
		mcp.WithBoolean("with_issues_enabled",
			mcp.Description("Limit by enabled issues feature (default: false)"),
		),
		mcp.WithBoolean("with_merge_requests_enabled",
			mcp.Description("Limit by enabled merge requests feature (default: false)"),
		),
		mcp.WithString("with_programming_language",
			mcp.Description("Limit to repositories which use the given programming language"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Limit by minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("id_after",
			mcp.Description("Limit results to projects with IDs greater than the specified ID"),
		),
		mcp.WithNumber("id_before",
			mcp.Description("Limit results to projects with IDs less than the specified ID"),
		),
		mcp.WithString("last_activity_after",
			mcp.Description("Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("last_activity_before",
			mcp.Description("Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("repository_storage",
			mcp.Description("Which storage shard the repository is on. Available only to admins"),
		),
		mcp.WithString("topic",
			mcp.Description("Comma-separated list of topics. Limit results to projects having all topics"),
		),
		mcp.WithNumber("topic_id",
			mcp.Description("Limit results to projects with the assigned topic given by the topic ID"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithBoolean("include_pending_delete",
			mcp.Description("Include projects in pending delete state. Can only be set by admins"),
		),
		mcp.WithString("marked_for_deletion_on",
			mcp.Description("Date when the project was marked for deletion"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by projects that are not archived and not marked for deletion"),
		),
		mcp.WithBoolean("wiki_checksum_failed",
			mcp.Description("Limit by projects where wiki checksum is failed (default: false)"),
		),
		mcp.WithBoolean("repository_checksum_failed",
			mcp.Description("Limit by projects where repository checksum is failed (default: false)"),
		),
		mcp.WithBoolean("include_hidden",
			mcp.Description("Include hidden projects. Can only be set by admins (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("simple",
			mcp.Description("Return only the ID, URL, name, and path of each project (default: false)"),
		),
		mcp.WithBoolean("statistics",
			mcp.Description("Include project statistics (default: false)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getUsersUserIdProjectsHandler)
}

func getUsersUserIdProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	user_id := request.GetString("user_id", "")
	params := parseGetUsersUserIdProjects(request)
	return toResult(c.GetApiV4UsersUserIdProjects(ctx, user_id, &params, authorizationHeader))
}

func parseGetUsersUserIdProjects(request mcp.CallToolRequest) client.GetApiV4UsersUserIdProjectsParams {
	params := client.GetApiV4UsersUserIdProjectsParams{}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	archived := request.GetBool("archived", false)
	params.Archived = &archived

	visibility := request.GetString("visibility", "")
	if visibility != "" {

		params.Visibility = &visibility
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	search_namespaces := request.GetBool("search_namespaces", false)
	params.SearchNamespaces = &search_namespaces

	owned := request.GetBool("owned", false)
	params.Owned = &owned

	starred := request.GetBool("starred", false)
	params.Starred = &starred

	imported := request.GetBool("imported", false)
	params.Imported = &imported

	membership := request.GetBool("membership", false)
	params.Membership = &membership

	with_issues_enabled := request.GetBool("with_issues_enabled", false)
	params.WithIssuesEnabled = &with_issues_enabled

	with_merge_requests_enabled := request.GetBool("with_merge_requests_enabled", false)
	params.WithMergeRequestsEnabled = &with_merge_requests_enabled

	with_programming_language := request.GetString("with_programming_language", "")
	if with_programming_language != "" {

		params.WithProgrammingLanguage = &with_programming_language
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

	id_after := request.GetInt("id_after", math.MinInt)
	if id_after != math.MinInt {
		id_after := int32(id_after)
		params.IdAfter = &id_after
	}

	id_before := request.GetInt("id_before", math.MinInt)
	if id_before != math.MinInt {
		id_before := int32(id_before)
		params.IdBefore = &id_before
	}

	last_activity_after := request.GetString("last_activity_after", "")
	if last_activity_after != "" {
		last_activity_after, _ := time.Parse(time.RFC3339, last_activity_after)
		params.LastActivityAfter = &last_activity_after
	}

	last_activity_before := request.GetString("last_activity_before", "")
	if last_activity_before != "" {
		last_activity_before, _ := time.Parse(time.RFC3339, last_activity_before)
		params.LastActivityBefore = &last_activity_before
	}

	repository_storage := request.GetString("repository_storage", "")
	if repository_storage != "" {

		params.RepositoryStorage = &repository_storage
	}

	topic := request.GetString("topic", "")
	if topic != "" {
		topic := strings.Split(topic, ",")
		params.Topic = &topic
	}

	topic_id := request.GetInt("topic_id", math.MinInt)
	if topic_id != math.MinInt {
		topic_id := int32(topic_id)
		params.TopicId = &topic_id
	}

	updated_before := request.GetString("updated_before", "")
	if updated_before != "" {
		updated_before, _ := time.Parse(time.RFC3339, updated_before)
		params.UpdatedBefore = &updated_before
	}

	updated_after := request.GetString("updated_after", "")
	if updated_after != "" {
		updated_after, _ := time.Parse(time.RFC3339, updated_after)
		params.UpdatedAfter = &updated_after
	}

	include_pending_delete := request.GetBool("include_pending_delete", false)
	params.IncludePendingDelete = &include_pending_delete

	marked_for_deletion_on := request.GetString("marked_for_deletion_on", "")
	if marked_for_deletion_on != "" {
		marked_for_deletion_on, _ := time.Parse(time.DateOnly, marked_for_deletion_on)
		params.MarkedForDeletionOn = &openapi_types.Date{Time: marked_for_deletion_on}
	}

	active := request.GetBool("active", false)
	params.Active = &active

	wiki_checksum_failed := request.GetBool("wiki_checksum_failed", false)
	params.WikiChecksumFailed = &wiki_checksum_failed

	repository_checksum_failed := request.GetBool("repository_checksum_failed", false)
	params.RepositoryChecksumFailed = &repository_checksum_failed

	include_hidden := request.GetBool("include_hidden", false)
	params.IncludeHidden = &include_hidden

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

	simple := request.GetBool("simple", false)
	params.Simple = &simple

	statistics := request.GetBool("statistics", false)
	params.Statistics = &statistics

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetUsersUserIdContributedProjects(s *server.MCPServer) {
	tool := mcp.NewTool("users_user_id_contributed_pjs",
		mcp.WithDescription("Get projects that a user has contributed to"),
		mcp.WithString("user_id",
			mcp.Description("The ID or username of the user"),
			mcp.Required(),
		),
		mcp.WithString("order_by",
			mcp.Description("Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. (default: created_at)"),

			mcp.Enum("id", "name", "path", "created_at", "updated_at", "last_activity_at", "similarity", "star_count", "storage_size", "repository_size", "wiki_size", "packages_size"),
		),
		mcp.WithString("sort",
			mcp.Description("Return projects sorted in ascending and descending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("simple",
			mcp.Description("Return only the ID, URL, name, and path of each project (default: false)"),
		),
	)

	s.AddTool(tool, getUsersUserIdContributedProjectsHandler)
}

func getUsersUserIdContributedProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	user_id := request.GetString("user_id", "")
	params := parseGetUsersUserIdContributedProjects(request)
	return toResult(c.GetApiV4UsersUserIdContributedProjects(ctx, user_id, &params, authorizationHeader))
}

func parseGetUsersUserIdContributedProjects(request mcp.CallToolRequest) client.GetApiV4UsersUserIdContributedProjectsParams {
	params := client.GetApiV4UsersUserIdContributedProjectsParams{}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
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

	simple := request.GetBool("simple", false)
	params.Simple = &simple

	return params
}

func registerGetUsersUserIdStarredProjects(s *server.MCPServer) {
	tool := mcp.NewTool("users_user_id_starred_pjs",
		mcp.WithDescription("Get projects starred by a user"),
		mcp.WithString("user_id",
			mcp.Description("The ID or username of the user"),
			mcp.Required(),
		),
		mcp.WithString("order_by",
			mcp.Description("Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to. (default: created_at)"),

			mcp.Enum("id", "name", "path", "created_at", "updated_at", "last_activity_at", "similarity", "star_count", "storage_size", "repository_size", "wiki_size", "packages_size"),
		),
		mcp.WithString("sort",
			mcp.Description("Return projects sorted in ascending and descending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of projects matching the search criteria"),
		),
		mcp.WithBoolean("search_namespaces",
			mcp.Description("Include ancestor namespaces when matching search criteria"),
		),
		mcp.WithBoolean("owned",
			mcp.Description("Limit by owned by authenticated user (default: false)"),
		),
		mcp.WithBoolean("starred",
			mcp.Description("Limit by starred status (default: false)"),
		),
		mcp.WithBoolean("imported",
			mcp.Description("Limit by imported by authenticated user (default: false)"),
		),
		mcp.WithBoolean("membership",
			mcp.Description("Limit by projects that the current user is a member of (default: false)"),
		),
		mcp.WithBoolean("with_issues_enabled",
			mcp.Description("Limit by enabled issues feature (default: false)"),
		),
		mcp.WithBoolean("with_merge_requests_enabled",
			mcp.Description("Limit by enabled merge requests feature (default: false)"),
		),
		mcp.WithString("with_programming_language",
			mcp.Description("Limit to repositories which use the given programming language"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Limit by minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("id_after",
			mcp.Description("Limit results to projects with IDs greater than the specified ID"),
		),
		mcp.WithNumber("id_before",
			mcp.Description("Limit results to projects with IDs less than the specified ID"),
		),
		mcp.WithString("last_activity_after",
			mcp.Description("Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("last_activity_before",
			mcp.Description("Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("repository_storage",
			mcp.Description("Which storage shard the repository is on. Available only to admins"),
		),
		mcp.WithString("topic",
			mcp.Description("Comma-separated list of topics. Limit results to projects having all topics"),
		),
		mcp.WithNumber("topic_id",
			mcp.Description("Limit results to projects with the assigned topic given by the topic ID"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithBoolean("include_pending_delete",
			mcp.Description("Include projects in pending delete state. Can only be set by admins"),
		),
		mcp.WithString("marked_for_deletion_on",
			mcp.Description("Date when the project was marked for deletion"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by projects that are not archived and not marked for deletion"),
		),
		mcp.WithBoolean("wiki_checksum_failed",
			mcp.Description("Limit by projects where wiki checksum is failed (default: false)"),
		),
		mcp.WithBoolean("repository_checksum_failed",
			mcp.Description("Limit by projects where repository checksum is failed (default: false)"),
		),
		mcp.WithBoolean("include_hidden",
			mcp.Description("Include hidden projects. Can only be set by admins (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("simple",
			mcp.Description("Return only the ID, URL, name, and path of each project (default: false)"),
		),
		mcp.WithBoolean("statistics",
			mcp.Description("Include project statistics (default: false)"),
		),
	)

	s.AddTool(tool, getUsersUserIdStarredProjectsHandler)
}

func getUsersUserIdStarredProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	user_id := request.GetString("user_id", "")
	params := parseGetUsersUserIdStarredProjects(request)
	return toResult(c.GetApiV4UsersUserIdStarredProjects(ctx, user_id, &params, authorizationHeader))
}

func parseGetUsersUserIdStarredProjects(request mcp.CallToolRequest) client.GetApiV4UsersUserIdStarredProjectsParams {
	params := client.GetApiV4UsersUserIdStarredProjectsParams{}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	archived := request.GetBool("archived", false)
	params.Archived = &archived

	visibility := request.GetString("visibility", "")
	if visibility != "" {

		params.Visibility = &visibility
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	search_namespaces := request.GetBool("search_namespaces", false)
	params.SearchNamespaces = &search_namespaces

	owned := request.GetBool("owned", false)
	params.Owned = &owned

	starred := request.GetBool("starred", false)
	params.Starred = &starred

	imported := request.GetBool("imported", false)
	params.Imported = &imported

	membership := request.GetBool("membership", false)
	params.Membership = &membership

	with_issues_enabled := request.GetBool("with_issues_enabled", false)
	params.WithIssuesEnabled = &with_issues_enabled

	with_merge_requests_enabled := request.GetBool("with_merge_requests_enabled", false)
	params.WithMergeRequestsEnabled = &with_merge_requests_enabled

	with_programming_language := request.GetString("with_programming_language", "")
	if with_programming_language != "" {

		params.WithProgrammingLanguage = &with_programming_language
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

	id_after := request.GetInt("id_after", math.MinInt)
	if id_after != math.MinInt {
		id_after := int32(id_after)
		params.IdAfter = &id_after
	}

	id_before := request.GetInt("id_before", math.MinInt)
	if id_before != math.MinInt {
		id_before := int32(id_before)
		params.IdBefore = &id_before
	}

	last_activity_after := request.GetString("last_activity_after", "")
	if last_activity_after != "" {
		last_activity_after, _ := time.Parse(time.RFC3339, last_activity_after)
		params.LastActivityAfter = &last_activity_after
	}

	last_activity_before := request.GetString("last_activity_before", "")
	if last_activity_before != "" {
		last_activity_before, _ := time.Parse(time.RFC3339, last_activity_before)
		params.LastActivityBefore = &last_activity_before
	}

	repository_storage := request.GetString("repository_storage", "")
	if repository_storage != "" {

		params.RepositoryStorage = &repository_storage
	}

	topic := request.GetString("topic", "")
	if topic != "" {
		topic := strings.Split(topic, ",")
		params.Topic = &topic
	}

	topic_id := request.GetInt("topic_id", math.MinInt)
	if topic_id != math.MinInt {
		topic_id := int32(topic_id)
		params.TopicId = &topic_id
	}

	updated_before := request.GetString("updated_before", "")
	if updated_before != "" {
		updated_before, _ := time.Parse(time.RFC3339, updated_before)
		params.UpdatedBefore = &updated_before
	}

	updated_after := request.GetString("updated_after", "")
	if updated_after != "" {
		updated_after, _ := time.Parse(time.RFC3339, updated_after)
		params.UpdatedAfter = &updated_after
	}

	include_pending_delete := request.GetBool("include_pending_delete", false)
	params.IncludePendingDelete = &include_pending_delete

	marked_for_deletion_on := request.GetString("marked_for_deletion_on", "")
	if marked_for_deletion_on != "" {
		marked_for_deletion_on, _ := time.Parse(time.DateOnly, marked_for_deletion_on)
		params.MarkedForDeletionOn = &openapi_types.Date{Time: marked_for_deletion_on}
	}

	active := request.GetBool("active", false)
	params.Active = &active

	wiki_checksum_failed := request.GetBool("wiki_checksum_failed", false)
	params.WikiChecksumFailed = &wiki_checksum_failed

	repository_checksum_failed := request.GetBool("repository_checksum_failed", false)
	params.RepositoryChecksumFailed = &repository_checksum_failed

	include_hidden := request.GetBool("include_hidden", false)
	params.IncludeHidden = &include_hidden

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

	simple := request.GetBool("simple", false)
	params.Simple = &simple

	statistics := request.GetBool("statistics", false)
	params.Statistics = &statistics

	return params
}
