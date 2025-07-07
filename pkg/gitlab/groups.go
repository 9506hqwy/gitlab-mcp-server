package gitlab

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func registerPostGroupsIdAccessRequests(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdAccessRequestsHandler)
}

func postGroupsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdAccessRequests(ctx, id, authorizationHeader))
}

func registerGetGroupsIdAccessRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdAccessRequestsHandler)
}

func getGroupsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdAccessRequests(request)
	return toResult(c.GetApiV4GroupsIdAccessRequests(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdAccessRequests(request mcp.CallToolRequest) client.GetApiV4GroupsIdAccessRequestsParams {
	params := client.GetApiV4GroupsIdAccessRequestsParams{}

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

func registerGetGroupsIdEpicsEpicIidAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_epics_epic_iid_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group"),
			mcp.Required(),
		),
		mcp.WithNumber("epic_iid",
			mcp.Description("ID (`iid` for merge requests/issues/epics, `id` for snippets) of an awardable."),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdEpicsEpicIidAwardEmojiHandler)
}

func getGroupsIdEpicsEpicIidAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	epic_iid := int32(request.GetInt("epic_iid", math.MinInt))
	params := parseGetGroupsIdEpicsEpicIidAwardEmoji(request)
	return toResult(c.GetApiV4GroupsIdEpicsEpicIidAwardEmoji(ctx, id, epic_iid, &params, authorizationHeader))
}

func parseGetGroupsIdEpicsEpicIidAwardEmoji(request mcp.CallToolRequest) client.GetApiV4GroupsIdEpicsEpicIidAwardEmojiParams {
	params := client.GetApiV4GroupsIdEpicsEpicIidAwardEmojiParams{}

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

func registerGetGroupsIdEpicsEpicIidAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_epics_epic_iid_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("epic_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdEpicsEpicIidAwardEmojiAwardIdHandler)
}

func getGroupsIdEpicsEpicIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	epic_iid := int32(request.GetInt("epic_iid", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId(ctx, id, epic_iid, award_id, authorizationHeader))
}

func registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_epics_epic_iid_notes_note_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("epic_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiHandler)
}

func getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	epic_iid := int32(request.GetInt("epic_iid", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	params := parseGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(request)
	return toResult(c.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(ctx, id, epic_iid, note_id, &params, authorizationHeader))
}

func parseGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(request mcp.CallToolRequest) client.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiParams {
	params := client.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiParams{}

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

func registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_epics_epic_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("epic_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdHandler)
}

func getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	epic_iid := int32(request.GetInt("epic_iid", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(ctx, id, epic_iid, note_id, award_id, authorizationHeader))
}

func registerGetGroupsIdBadges(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_badges",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user."),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("name",
			mcp.Description("Name for the badge"),
		),
	)

	s.AddTool(tool, getGroupsIdBadgesHandler)
}

func getGroupsIdBadgesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdBadges(request)
	return toResult(c.GetApiV4GroupsIdBadges(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdBadges(request mcp.CallToolRequest) client.GetApiV4GroupsIdBadgesParams {
	params := client.GetApiV4GroupsIdBadgesParams{}

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

	name := request.GetString("name", "")
	if name != "" {

		params.Name = &name
	}

	return params
}

func registerGetGroupsIdBadgesRender(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_badges_render",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user."),
			mcp.Required(),
		),
		mcp.WithString("link_url",
			mcp.Description("URL of the badge link"),
			mcp.Required(),
		),
		mcp.WithString("image_url",
			mcp.Description("URL of the badge image"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdBadgesRenderHandler)
}

func getGroupsIdBadgesRenderHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdBadgesRender(request)
	return toResult(c.GetApiV4GroupsIdBadgesRender(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdBadgesRender(request mcp.CallToolRequest) client.GetApiV4GroupsIdBadgesRenderParams {
	params := client.GetApiV4GroupsIdBadgesRenderParams{}

	link_url := request.GetString("link_url", "")
	if link_url != "" {

		params.LinkUrl = link_url
	}

	image_url := request.GetString("image_url", "")
	if image_url != "" {

		params.ImageUrl = image_url
	}

	return params
}

func registerGetGroupsIdBadgesBadgeId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_badges_badge_id",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user."),
			mcp.Required(),
		),
		mcp.WithNumber("badge_id",
			mcp.Description("The badge ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdBadgesBadgeIdHandler)
}

func getGroupsIdBadgesBadgeIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	badge_id := int32(request.GetInt("badge_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdBadgesBadgeId(ctx, id, badge_id, authorizationHeader))
}

func registerGetGroupsIdCustomAttributes(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_custom_attributes",
		mcp.WithDescription("Get all custom attributes on a group"),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdCustomAttributesHandler)
}

func getGroupsIdCustomAttributesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdCustomAttributes(ctx, id, authorizationHeader))
}

func registerGetGroupsIdCustomAttributesKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_custom_attributes_key",
		mcp.WithDescription("Get a custom attribute on a group"),
		mcp.WithString("key",
			mcp.Description("The key of the custom attribute"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdCustomAttributesKeyHandler)
}

func getGroupsIdCustomAttributesKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	key := request.GetString("key", "")

	return toResult(c.GetApiV4GroupsIdCustomAttributesKey(ctx, id, key, authorizationHeader))
}

func registerGetGroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps",
		mcp.WithDescription("Get a groups list"),
		mcp.WithBoolean("statistics",
			mcp.Description("Include project statistics (default: false)"),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("skip_groups",
			mcp.Description("Array of group ids to exclude from list"),
		),
		mcp.WithBoolean("all_available",
			mcp.Description("When `true`, returns all accessible groups. When `false`, returns only groups where the user is a member."),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a specific group"),
		),
		mcp.WithBoolean("owned",
			mcp.Description("Limit by owned by authenticated user (default: false)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Order by name, path, id or similarity if searching (default: name)"),

			mcp.Enum("name", "path", "id", "similarity"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort by asc (ascending) or desc (descending) (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithBoolean("top_level_only",
			mcp.Description("Only include top-level groups"),
		),
		mcp.WithString("marked_for_deletion_on",
			mcp.Description("Return groups that are marked for deletion on this date"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by groups that are not archived and not marked for deletion"),
		),
		mcp.WithString("repository_storage",
			mcp.Description("Filter by repository storage used by the group"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsHandler)
}

func getGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetGroups(request)
	return toResult(c.GetApiV4Groups(ctx, &params, authorizationHeader))
}

func parseGetGroups(request mcp.CallToolRequest) client.GetApiV4GroupsParams {
	params := client.GetApiV4GroupsParams{}

	statistics := request.GetBool("statistics", false)
	params.Statistics = &statistics

	archived := request.GetBool("archived", false)
	params.Archived = &archived

	skip_groups := request.GetString("skip_groups", "")
	if skip_groups != "" {
		skip_groups := strings.Split(skip_groups, ",")
		var intSlice []int32
		for _, v := range skip_groups {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.SkipGroups = &intSlice
	}

	all_available := request.GetBool("all_available", false)
	params.AllAvailable = &all_available

	visibility := request.GetString("visibility", "")
	if visibility != "" {

		params.Visibility = &visibility
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	owned := request.GetBool("owned", false)
	params.Owned = &owned

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

	top_level_only := request.GetBool("top_level_only", false)
	params.TopLevelOnly = &top_level_only

	marked_for_deletion_on := request.GetString("marked_for_deletion_on", "")
	if marked_for_deletion_on != "" {
		marked_for_deletion_on, _ := time.Parse(time.DateOnly, marked_for_deletion_on)
		params.MarkedForDeletionOn = &openapi_types.Date{Time: marked_for_deletion_on}
	}

	active := request.GetBool("active", false)
	params.Active = &active

	repository_storage := request.GetString("repository_storage", "")
	if repository_storage != "" {

		params.RepositoryStorage = &repository_storage
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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetGroupsId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id",
		mcp.WithDescription("Get a single group, with containing projects."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
		mcp.WithBoolean("with_projects",
			mcp.Description("Omit project details (default: true)"),
		),
	)

	s.AddTool(tool, getGroupsIdHandler)
}

func getGroupsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsId(request)
	return toResult(c.GetApiV4GroupsId(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsId(request mcp.CallToolRequest) client.GetApiV4GroupsIdParams {
	params := client.GetApiV4GroupsIdParams{}

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	with_projects := request.GetBool("with_projects", true)
	params.WithProjects = &with_projects

	return params
}

func registerPostGroupsIdArchive(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_archive",
		mcp.WithDescription("Archive a group"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdArchiveHandler)
}

func postGroupsIdArchiveHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdArchive(ctx, id, authorizationHeader))
}

func registerPostGroupsIdUnarchive(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_unarchive",
		mcp.WithDescription("Unarchive a group"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdUnarchiveHandler)
}

func postGroupsIdUnarchiveHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdUnarchive(ctx, id, authorizationHeader))
}

func registerPostGroupsIdRestore(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_restore",
		mcp.WithDescription("Restore a group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdRestoreHandler)
}

func postGroupsIdRestoreHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdRestore(ctx, id, authorizationHeader))
}

func registerGetGroupsIdGroupsShared(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_groups_shared",
		mcp.WithDescription("Get a list of shared groups this group was invited to"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithString("skip_groups",
			mcp.Description("Array of group ids to exclude from list"),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a specific group"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithString("order_by",
			mcp.Description("Order by name, path, id or similarity if searching (default: name)"),

			mcp.Enum("name", "path", "id", "similarity"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort by asc (ascending) or desc (descending) (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsIdGroupsSharedHandler)
}

func getGroupsIdGroupsSharedHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdGroupsShared(request)
	return toResult(c.GetApiV4GroupsIdGroupsShared(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdGroupsShared(request mcp.CallToolRequest) client.GetApiV4GroupsIdGroupsSharedParams {
	params := client.GetApiV4GroupsIdGroupsSharedParams{}

	skip_groups := request.GetString("skip_groups", "")
	if skip_groups != "" {
		skip_groups := strings.Split(skip_groups, ",")
		var intSlice []int32
		for _, v := range skip_groups {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.SkipGroups = &intSlice
	}

	visibility := request.GetString("visibility", "")
	if visibility != "" {

		params.Visibility = &visibility
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetGroupsIdInvitedGroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_invited_groups",
		mcp.WithDescription("Get a list of invited groups in this group"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithString("relation",
			mcp.Description("Include group relations"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a specific group"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsIdInvitedGroupsHandler)
}

func getGroupsIdInvitedGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdInvitedGroups(request)
	return toResult(c.GetApiV4GroupsIdInvitedGroups(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdInvitedGroups(request mcp.CallToolRequest) client.GetApiV4GroupsIdInvitedGroupsParams {
	params := client.GetApiV4GroupsIdInvitedGroupsParams{}

	relation := request.GetString("relation", "")
	if relation != "" {
		relation := strings.Split(relation, ",")
		params.Relation = &relation
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetGroupsIdProjects(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pjs",
		mcp.WithDescription("Get a list of projects in this group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of authorized projects matching the search criteria"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return projects ordered by field (default: created_at)"),

			mcp.Enum("id", "name", "path", "created_at", "updated_at", "last_activity_at", "similarity", "star_count"),
		),
		mcp.WithString("sort",
			mcp.Description("Return projects sorted in ascending and descending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("simple",
			mcp.Description("Return only the ID, URL, name, and path of each project (default: false)"),
		),
		mcp.WithBoolean("owned",
			mcp.Description("Limit by owned by authenticated user (default: false)"),
		),
		mcp.WithBoolean("starred",
			mcp.Description("Limit by starred status (default: false)"),
		),
		mcp.WithBoolean("with_issues_enabled",
			mcp.Description("Limit by enabled issues feature (default: false)"),
		),
		mcp.WithBoolean("with_merge_requests_enabled",
			mcp.Description("Limit by enabled merge requests feature (default: false)"),
		),
		mcp.WithBoolean("with_shared",
			mcp.Description("Include projects shared to this group (default: true)"),
		),
		mcp.WithBoolean("include_subgroups",
			mcp.Description("Includes projects in subgroups of this group (default: false)"),
		),
		mcp.WithBoolean("include_ancestor_groups",
			mcp.Description("Includes projects in ancestors of this group (default: false)"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Limit by minimum access level of authenticated user on projects"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
		mcp.WithBoolean("with_security_reports",
			mcp.Description("Return only projects having security report artifacts present (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsIdProjectsHandler)
}

func getGroupsIdProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdProjects(request)
	return toResult(c.GetApiV4GroupsIdProjects(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdProjects(request mcp.CallToolRequest) client.GetApiV4GroupsIdProjectsParams {
	params := client.GetApiV4GroupsIdProjectsParams{}

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

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	simple := request.GetBool("simple", false)
	params.Simple = &simple

	owned := request.GetBool("owned", false)
	params.Owned = &owned

	starred := request.GetBool("starred", false)
	params.Starred = &starred

	with_issues_enabled := request.GetBool("with_issues_enabled", false)
	params.WithIssuesEnabled = &with_issues_enabled

	with_merge_requests_enabled := request.GetBool("with_merge_requests_enabled", false)
	params.WithMergeRequestsEnabled = &with_merge_requests_enabled

	with_shared := request.GetBool("with_shared", true)
	params.WithShared = &with_shared

	include_subgroups := request.GetBool("include_subgroups", false)
	params.IncludeSubgroups = &include_subgroups

	include_ancestor_groups := request.GetBool("include_ancestor_groups", false)
	params.IncludeAncestorGroups = &include_ancestor_groups

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	with_security_reports := request.GetBool("with_security_reports", false)
	params.WithSecurityReports = &with_security_reports

	return params
}

func registerGetGroupsIdProjectsShared(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pjs_shared",
		mcp.WithDescription("Get a list of shared projects in this group"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of authorized projects matching the search criteria"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return projects ordered by field (default: created_at)"),

			mcp.Enum("id", "name", "path", "created_at", "updated_at", "last_activity_at", "star_count"),
		),
		mcp.WithString("sort",
			mcp.Description("Return projects sorted in ascending and descending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("simple",
			mcp.Description("Return only the ID, URL, name, and path of each project (default: false)"),
		),
		mcp.WithBoolean("starred",
			mcp.Description("Limit by starred status (default: false)"),
		),
		mcp.WithBoolean("with_issues_enabled",
			mcp.Description("Limit by enabled issues feature (default: false)"),
		),
		mcp.WithBoolean("with_merge_requests_enabled",
			mcp.Description("Limit by enabled merge requests feature (default: false)"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Limit by minimum access level of authenticated user on projects"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsIdProjectsSharedHandler)
}

func getGroupsIdProjectsSharedHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdProjectsShared(request)
	return toResult(c.GetApiV4GroupsIdProjectsShared(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdProjectsShared(request mcp.CallToolRequest) client.GetApiV4GroupsIdProjectsSharedParams {
	params := client.GetApiV4GroupsIdProjectsSharedParams{}

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

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	simple := request.GetBool("simple", false)
	params.Simple = &simple

	starred := request.GetBool("starred", false)
	params.Starred = &starred

	with_issues_enabled := request.GetBool("with_issues_enabled", false)
	params.WithIssuesEnabled = &with_issues_enabled

	with_merge_requests_enabled := request.GetBool("with_merge_requests_enabled", false)
	params.WithMergeRequestsEnabled = &with_merge_requests_enabled

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetGroupsIdSubgroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_subgroups",
		mcp.WithDescription("Get a list of subgroups in this group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithBoolean("statistics",
			mcp.Description("Include project statistics (default: false)"),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("skip_groups",
			mcp.Description("Array of group ids to exclude from list"),
		),
		mcp.WithBoolean("all_available",
			mcp.Description("When `true`, returns all accessible groups. When `false`, returns only groups where the user is a member."),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a specific group"),
		),
		mcp.WithBoolean("owned",
			mcp.Description("Limit by owned by authenticated user (default: false)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Order by name, path, id or similarity if searching (default: name)"),

			mcp.Enum("name", "path", "id", "similarity"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort by asc (ascending) or desc (descending) (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithBoolean("top_level_only",
			mcp.Description("Only include top-level groups"),
		),
		mcp.WithString("marked_for_deletion_on",
			mcp.Description("Return groups that are marked for deletion on this date"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by groups that are not archived and not marked for deletion"),
		),
		mcp.WithString("repository_storage",
			mcp.Description("Filter by repository storage used by the group"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsIdSubgroupsHandler)
}

func getGroupsIdSubgroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdSubgroups(request)
	return toResult(c.GetApiV4GroupsIdSubgroups(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdSubgroups(request mcp.CallToolRequest) client.GetApiV4GroupsIdSubgroupsParams {
	params := client.GetApiV4GroupsIdSubgroupsParams{}

	statistics := request.GetBool("statistics", false)
	params.Statistics = &statistics

	archived := request.GetBool("archived", false)
	params.Archived = &archived

	skip_groups := request.GetString("skip_groups", "")
	if skip_groups != "" {
		skip_groups := strings.Split(skip_groups, ",")
		var intSlice []int32
		for _, v := range skip_groups {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.SkipGroups = &intSlice
	}

	all_available := request.GetBool("all_available", false)
	params.AllAvailable = &all_available

	visibility := request.GetString("visibility", "")
	if visibility != "" {

		params.Visibility = &visibility
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	owned := request.GetBool("owned", false)
	params.Owned = &owned

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

	top_level_only := request.GetBool("top_level_only", false)
	params.TopLevelOnly = &top_level_only

	marked_for_deletion_on := request.GetString("marked_for_deletion_on", "")
	if marked_for_deletion_on != "" {
		marked_for_deletion_on, _ := time.Parse(time.DateOnly, marked_for_deletion_on)
		params.MarkedForDeletionOn = &openapi_types.Date{Time: marked_for_deletion_on}
	}

	active := request.GetBool("active", false)
	params.Active = &active

	repository_storage := request.GetString("repository_storage", "")
	if repository_storage != "" {

		params.RepositoryStorage = &repository_storage
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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetGroupsIdDescendantGroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_descendant_groups",
		mcp.WithDescription("Get a list of descendant groups of this group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithBoolean("statistics",
			mcp.Description("Include project statistics (default: false)"),
		),
		mcp.WithBoolean("archived",
			mcp.Description("Limit by archived status"),
		),
		mcp.WithString("skip_groups",
			mcp.Description("Array of group ids to exclude from list"),
		),
		mcp.WithBoolean("all_available",
			mcp.Description("When `true`, returns all accessible groups. When `false`, returns only groups where the user is a member."),
		),
		mcp.WithString("visibility",
			mcp.Description("Limit by visibility"),

			mcp.Enum("private", "internal", "public"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a specific group"),
		),
		mcp.WithBoolean("owned",
			mcp.Description("Limit by owned by authenticated user (default: false)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Order by name, path, id or similarity if searching (default: name)"),

			mcp.Enum("name", "path", "id", "similarity"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort by asc (ascending) or desc (descending) (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Minimum access level of authenticated user"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithBoolean("top_level_only",
			mcp.Description("Only include top-level groups"),
		),
		mcp.WithString("marked_for_deletion_on",
			mcp.Description("Return groups that are marked for deletion on this date"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by groups that are not archived and not marked for deletion"),
		),
		mcp.WithString("repository_storage",
			mcp.Description("Filter by repository storage used by the group"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getGroupsIdDescendantGroupsHandler)
}

func getGroupsIdDescendantGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdDescendantGroups(request)
	return toResult(c.GetApiV4GroupsIdDescendantGroups(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdDescendantGroups(request mcp.CallToolRequest) client.GetApiV4GroupsIdDescendantGroupsParams {
	params := client.GetApiV4GroupsIdDescendantGroupsParams{}

	statistics := request.GetBool("statistics", false)
	params.Statistics = &statistics

	archived := request.GetBool("archived", false)
	params.Archived = &archived

	skip_groups := request.GetString("skip_groups", "")
	if skip_groups != "" {
		skip_groups := strings.Split(skip_groups, ",")
		var intSlice []int32
		for _, v := range skip_groups {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.SkipGroups = &intSlice
	}

	all_available := request.GetBool("all_available", false)
	params.AllAvailable = &all_available

	visibility := request.GetString("visibility", "")
	if visibility != "" {

		params.Visibility = &visibility
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	owned := request.GetBool("owned", false)
	params.Owned = &owned

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	min_access_level := request.GetInt("min_access_level", math.MinInt)
	if min_access_level != math.MinInt {
		min_access_level := int32(min_access_level)
		params.MinAccessLevel = &min_access_level
	}

	top_level_only := request.GetBool("top_level_only", false)
	params.TopLevelOnly = &top_level_only

	marked_for_deletion_on := request.GetString("marked_for_deletion_on", "")
	if marked_for_deletion_on != "" {
		marked_for_deletion_on, _ := time.Parse(time.DateOnly, marked_for_deletion_on)
		params.MarkedForDeletionOn = &openapi_types.Date{Time: marked_for_deletion_on}
	}

	active := request.GetBool("active", false)
	params.Active = &active

	repository_storage := request.GetString("repository_storage", "")
	if repository_storage != "" {

		params.RepositoryStorage = &repository_storage
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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerPostGroupsIdProjectsProjectId(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_pjs_project_id",
		mcp.WithDescription("Transfer a project to the group namespace. Available only for admin."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithString("project_id",
			mcp.Description("The ID or path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdProjectsProjectIdHandler)
}

func postGroupsIdProjectsProjectIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	project_id := request.GetString("project_id", "")

	return toResult(c.PostApiV4GroupsIdProjectsProjectId(ctx, id, project_id, authorizationHeader))
}

func registerGetGroupsIdTransferLocations(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_transfer_locations",
		mcp.WithDescription("Get the groups to where the current group can be transferred to"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return list of namespaces matching the search criteria"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdTransferLocationsHandler)
}

func getGroupsIdTransferLocationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdTransferLocations(request)
	return toResult(c.GetApiV4GroupsIdTransferLocations(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdTransferLocations(request mcp.CallToolRequest) client.GetApiV4GroupsIdTransferLocationsParams {
	params := client.GetApiV4GroupsIdTransferLocationsParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
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

func registerPostGroupsIdLdapSync(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_ldap_sync",
		mcp.WithDescription("Sync a group with LDAP."),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdLdapSyncHandler)
}

func postGroupsIdLdapSyncHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.PostApiV4GroupsIdLdapSync(ctx, id, authorizationHeader))
}

func registerGetGroupsIdAuditEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_audit_events",
		mcp.WithDescription("Get a list of audit events in this group."),
		mcp.WithString("created_after",
			mcp.Description("Return audit events created after the specified time (example: 2016-01-19T09:05:50.355Z)"),
		),
		mcp.WithString("created_before",
			mcp.Description("Return audit events created before the specified time (example: 2016-01-19T09:05:50.355Z)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdAuditEventsHandler)
}

func getGroupsIdAuditEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdAuditEvents(request)
	return toResult(c.GetApiV4GroupsIdAuditEvents(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdAuditEvents(request mcp.CallToolRequest) client.GetApiV4GroupsIdAuditEventsParams {
	params := client.GetApiV4GroupsIdAuditEventsParams{}

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
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

func registerGetGroupsIdAuditEventsAuditEventId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_audit_events_audit_event_id",
		mcp.WithDescription("Get a specific audit event in this group."),
		mcp.WithNumber("audit_event_id",
			mcp.Description("The ID of the audit event"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdAuditEventsAuditEventIdHandler)
}

func getGroupsIdAuditEventsAuditEventIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	audit_event_id := int32(request.GetInt("audit_event_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdAuditEventsAuditEventId(ctx, id, audit_event_id, authorizationHeader))
}

func registerGetGroupsIdSamlUsers(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_saml_users",
		mcp.WithDescription("Get a list of SAML users of the group"),
		mcp.WithString("username",
			mcp.Description("Return single user with a specific username."),
		),
		mcp.WithString("search",
			mcp.Description("Search users by name, email, username."),
		),
		mcp.WithBoolean("active",
			mcp.Description("Return only active users. (default: false)"),
		),
		mcp.WithBoolean("blocked",
			mcp.Description("Return only blocked users. (default: false)"),
		),
		mcp.WithString("created_after",
			mcp.Description("Return users created after the specified time."),
		),
		mcp.WithString("created_before",
			mcp.Description("Return users created before the specified time."),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdSamlUsersHandler)
}

func getGroupsIdSamlUsersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdSamlUsers(request)
	return toResult(c.GetApiV4GroupsIdSamlUsers(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdSamlUsers(request mcp.CallToolRequest) client.GetApiV4GroupsIdSamlUsersParams {
	params := client.GetApiV4GroupsIdSamlUsersParams{}

	username := request.GetString("username", "")
	if username != "" {

		params.Username = &username
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	active := request.GetBool("active", false)
	params.Active = &active

	blocked := request.GetBool("blocked", false)
	params.Blocked = &blocked

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
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

func registerGetGroupsIdProvisionedUsers(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_provisioned_users",
		mcp.WithDescription("Get a list of users provisioned by the group"),
		mcp.WithString("username",
			mcp.Description("Return a single user with a specific username"),
		),
		mcp.WithString("search",
			mcp.Description("Search users by name, email or username"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Return only active users (default: false)"),
		),
		mcp.WithBoolean("blocked",
			mcp.Description("Return only blocked users (default: false)"),
		),
		mcp.WithString("created_after",
			mcp.Description("Return users created after the specified time"),
		),
		mcp.WithString("created_before",
			mcp.Description("Return users created before the specified time"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdProvisionedUsersHandler)
}

func getGroupsIdProvisionedUsersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdProvisionedUsers(request)
	return toResult(c.GetApiV4GroupsIdProvisionedUsers(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdProvisionedUsers(request mcp.CallToolRequest) client.GetApiV4GroupsIdProvisionedUsersParams {
	params := client.GetApiV4GroupsIdProvisionedUsersParams{}

	username := request.GetString("username", "")
	if username != "" {

		params.Username = &username
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	active := request.GetBool("active", false)
	params.Active = &active

	blocked := request.GetBool("blocked", false)
	params.Blocked = &blocked

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
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

func registerGetGroupsIdUsers(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_users",
		mcp.WithDescription("Get a list of users for the group"),
		mcp.WithString("search",
			mcp.Description("Search users by name, email or username"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Filters only active users (default: false)"),
		),
		mcp.WithBoolean("include_saml_users",
			mcp.Description("Return users with a SAML identity in this group"),
		),
		mcp.WithBoolean("include_service_accounts",
			mcp.Description("Return service accounts owned by this group"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdUsersHandler)
}

func getGroupsIdUsersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdUsers(request)
	return toResult(c.GetApiV4GroupsIdUsers(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdUsers(request mcp.CallToolRequest) client.GetApiV4GroupsIdUsersParams {
	params := client.GetApiV4GroupsIdUsersParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	active := request.GetBool("active", false)
	params.Active = &active

	include_saml_users := request.GetBool("include_saml_users", false)
	params.IncludeSamlUsers = &include_saml_users

	include_service_accounts := request.GetBool("include_service_accounts", false)
	params.IncludeServiceAccounts = &include_service_accounts

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

func registerGetGroupsIdSshCertificates(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_ssh_certificates",
		mcp.WithDescription("Get a list of ssh certificates created for a group."),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdSshCertificatesHandler)
}

func getGroupsIdSshCertificatesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdSshCertificates(request)
	return toResult(c.GetApiV4GroupsIdSshCertificates(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdSshCertificates(request mcp.CallToolRequest) client.GetApiV4GroupsIdSshCertificatesParams {
	params := client.GetApiV4GroupsIdSshCertificatesParams{}

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

func registerGetGroupsIdRunners(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_runners",
		mcp.WithDescription("List all runners available in the group as well as its ancestor groups, including any allowed shared runners."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
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

	s.AddTool(tool, getGroupsIdRunnersHandler)
}

func getGroupsIdRunnersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdRunners(request)
	return toResult(c.GetApiV4GroupsIdRunners(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdRunners(request mcp.CallToolRequest) client.GetApiV4GroupsIdRunnersParams {
	params := client.GetApiV4GroupsIdRunnersParams{}

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

func registerPostGroupsIdRunnersResetRegistrationToken(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_runners_reset_registration_token",
		mcp.WithDescription("Reset runner registration token"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdRunnersResetRegistrationTokenHandler)
}

func postGroupsIdRunnersResetRegistrationTokenHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdRunnersResetRegistrationToken(ctx, id, authorizationHeader))
}

func registerGetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs_debian_pool_distribution_project_id_letter_package_name_package_version_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithString("id",
			mcp.Description("The group ID or full group path."),
			mcp.Required(),
		),
		mcp.WithNumber("project_id",
			mcp.Description("The Project Id"),
			mcp.Required(),
		),
		mcp.WithString("distribution",
			mcp.Description("The Debian Codename or Suite (example: my-distro)"),
			mcp.Required(),
		),
		mcp.WithString("letter",
			mcp.Description("The Debian Classification (first-letter or lib-first-letter) (example: a)"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("The Debian Source Package Name (example: my-pkg)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("The Debian Source Package Version (example: 1.0.0)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("The Debian File Name (example: example_1.0.0~alpha2_amd64.deb)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameHandler)
}

func getGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	distribution := request.GetString("distribution", "")
	project_id := int32(request.GetInt("project_id", math.MinInt))
	letter := request.GetString("letter", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(ctx, id, distribution, project_id, letter, package_name, package_version, file_name, authorizationHeader))
}

func registerGetGroupsIdDeployTokens(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_deploy_tokens",
		mcp.WithDescription("Get a list of a group's deploy tokens. This feature was introduced in GitLab 12.9."),
		mcp.WithNumber("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by active status"),
		),
	)

	s.AddTool(tool, getGroupsIdDeployTokensHandler)
}

func getGroupsIdDeployTokensHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdDeployTokens(request)
	return toResult(c.GetApiV4GroupsIdDeployTokens(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdDeployTokens(request mcp.CallToolRequest) client.GetApiV4GroupsIdDeployTokensParams {
	params := client.GetApiV4GroupsIdDeployTokensParams{}

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

	active := request.GetBool("active", false)
	params.Active = &active

	return params
}

func registerGetGroupsIdDeployTokensTokenId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_deploy_tokens_token_id",
		mcp.WithDescription("Get a single group's deploy token by ID. This feature was introduced in GitLab 14.9."),
		mcp.WithNumber("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("token_id",
			mcp.Description("The ID of the deploy token"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdDeployTokensTokenIdHandler)
}

func getGroupsIdDeployTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	token_id := int32(request.GetInt("token_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdDeployTokensTokenId(ctx, id, token_id, authorizationHeader))
}

func registerGetGroupsIdAvatar(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_avatar",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithString("id",
			mcp.Description("The ID of the group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdAvatarHandler)
}

func getGroupsIdAvatarHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4GroupsIdAvatar(ctx, id, authorizationHeader))
}

func registerGetGroupsIdClusters(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 12.1. Returns a list of group clusters."),
		mcp.WithString("id",
			mcp.Description("The ID of the group"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdClustersHandler)
}

func getGroupsIdClustersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdClusters(request)
	return toResult(c.GetApiV4GroupsIdClusters(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdClusters(request mcp.CallToolRequest) client.GetApiV4GroupsIdClustersParams {
	params := client.GetApiV4GroupsIdClustersParams{}

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

func registerGetGroupsIdClustersClusterId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 12.1. Gets a single group cluster."),
		mcp.WithString("id",
			mcp.Description("The ID of the group"),
			mcp.Required(),
		),
		mcp.WithNumber("cluster_id",
			mcp.Description("The cluster ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdClustersClusterIdHandler)
}

func getGroupsIdClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	cluster_id := int32(request.GetInt("cluster_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdClustersClusterId(ctx, id, cluster_id, authorizationHeader))
}

func registerGetGroupsIdRegistryRepositories(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_registry_repositories",
		mcp.WithDescription("Get a list of registry repositories in a group. This feature was introduced in GitLab 12.2."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group accessible by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdRegistryRepositoriesHandler)
}

func getGroupsIdRegistryRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdRegistryRepositories(request)
	return toResult(c.GetApiV4GroupsIdRegistryRepositories(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdRegistryRepositories(request mcp.CallToolRequest) client.GetApiV4GroupsIdRegistryRepositoriesParams {
	params := client.GetApiV4GroupsIdRegistryRepositoriesParams{}

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

func registerGetGroupsIdDebianDistributions(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_debian_distributions",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("codename",
			mcp.Description("The Debian Codename (example: sid)"),
		),
		mcp.WithString("suite",
			mcp.Description("The Debian Suite (example: unstable)"),
		),
		mcp.WithString("origin",
			mcp.Description("The Debian Origin (example: Grep)"),
		),
		mcp.WithString("label",
			mcp.Description("The Debian Label (example: grep.be)"),
		),
		mcp.WithString("version",
			mcp.Description("The Debian Version (example: 12)"),
		),
		mcp.WithString("description",
			mcp.Description("The Debian Description (example: My description)"),
		),
		mcp.WithNumber("valid_time_duration_seconds",
			mcp.Description("The duration before the Release file should be considered expired by the client (example: 604800)"),
		),
		mcp.WithString("components",
			mcp.Description("The list of Components (example: main)"),
		),
		mcp.WithString("architectures",
			mcp.Description("The list of Architectures (example: amd64)"),
		),
	)

	s.AddTool(tool, getGroupsIdDebianDistributionsHandler)
}

func getGroupsIdDebianDistributionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdDebianDistributions(request)
	return toResult(c.GetApiV4GroupsIdDebianDistributions(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdDebianDistributions(request mcp.CallToolRequest) client.GetApiV4GroupsIdDebianDistributionsParams {
	params := client.GetApiV4GroupsIdDebianDistributionsParams{}

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

	codename := request.GetString("codename", "")
	if codename != "" {

		params.Codename = &codename
	}

	suite := request.GetString("suite", "")
	if suite != "" {

		params.Suite = &suite
	}

	origin := request.GetString("origin", "")
	if origin != "" {

		params.Origin = &origin
	}

	label := request.GetString("label", "")
	if label != "" {

		params.Label = &label
	}

	version := request.GetString("version", "")
	if version != "" {

		params.Version = &version
	}

	description := request.GetString("description", "")
	if description != "" {

		params.Description = &description
	}

	valid_time_duration_seconds := request.GetInt("valid_time_duration_seconds", math.MinInt)
	if valid_time_duration_seconds != math.MinInt {
		valid_time_duration_seconds := int32(valid_time_duration_seconds)
		params.ValidTimeDurationSeconds = &valid_time_duration_seconds
	}

	components := request.GetString("components", "")
	if components != "" {
		components := strings.Split(components, ",")
		params.Components = &components
	}

	architectures := request.GetString("architectures", "")
	if architectures != "" {
		architectures := strings.Split(architectures, ",")
		params.Architectures = &architectures
	}

	return params
}

func registerGetGroupsIdDebianDistributionsCodename(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_debian_distributions_codename",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group"),
			mcp.Required(),
		),
		mcp.WithString("codename",
			mcp.Description("The Debian Codename (example: sid)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdDebianDistributionsCodenameHandler)
}

func getGroupsIdDebianDistributionsCodenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	codename := request.GetString("codename", "")

	return toResult(c.GetApiV4GroupsIdDebianDistributionsCodename(ctx, id, codename, authorizationHeader))
}

func registerGetGroupsIdDebianDistributionsCodenameKeyAsc(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_debian_distributions_codename_key_asc",
		mcp.WithDescription("This feature was introduced in 14.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group"),
			mcp.Required(),
		),
		mcp.WithString("codename",
			mcp.Description("The Debian Codename (example: sid)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdDebianDistributionsCodenameKeyAscHandler)
}

func getGroupsIdDebianDistributionsCodenameKeyAscHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	codename := request.GetString("codename", "")

	return toResult(c.GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc(ctx, id, codename, authorizationHeader))
}

func registerGetGroupsIdExportDownload(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_export_download",
		mcp.WithDescription("This feature was introduced in GitLab 12.5."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdExportDownloadHandler)
}

func getGroupsIdExportDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4GroupsIdExportDownload(ctx, id, authorizationHeader))
}

func registerPostGroupsIdExport(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_export",
		mcp.WithDescription("This feature was introduced in GitLab 12.5."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdExportHandler)
}

func postGroupsIdExportHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdExport(ctx, id, authorizationHeader))
}

func registerGetGroupsIdExportRelationsDownload(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_export_relations_download",
		mcp.WithDescription("This feature was introduced in GitLab 13.12"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithString("relation",
			mcp.Description("Group relation name"),
			mcp.Required(),
		),
		mcp.WithBoolean("batched",
			mcp.Description("Whether to download in batches"),
		),
		mcp.WithNumber("batch_number",
			mcp.Description("Batch number to download"),
		),
	)

	s.AddTool(tool, getGroupsIdExportRelationsDownloadHandler)
}

func getGroupsIdExportRelationsDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdExportRelationsDownload(request)
	return toResult(c.GetApiV4GroupsIdExportRelationsDownload(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdExportRelationsDownload(request mcp.CallToolRequest) client.GetApiV4GroupsIdExportRelationsDownloadParams {
	params := client.GetApiV4GroupsIdExportRelationsDownloadParams{}

	relation := request.GetString("relation", "")
	if relation != "" {

		params.Relation = relation
	}

	batched := request.GetBool("batched", false)
	params.Batched = &batched

	batch_number := request.GetInt("batch_number", math.MinInt)
	if batch_number != math.MinInt {
		batch_number := int32(batch_number)
		params.BatchNumber = &batch_number
	}

	return params
}

func registerGetGroupsIdExportRelationsStatus(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_export_relations_status",
		mcp.WithDescription("This feature was introduced in GitLab 13.12"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithString("relation",
			mcp.Description("Group relation name"),
		),
	)

	s.AddTool(tool, getGroupsIdExportRelationsStatusHandler)
}

func getGroupsIdExportRelationsStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdExportRelationsStatus(request)
	return toResult(c.GetApiV4GroupsIdExportRelationsStatus(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdExportRelationsStatus(request mcp.CallToolRequest) client.GetApiV4GroupsIdExportRelationsStatusParams {
	params := client.GetApiV4GroupsIdExportRelationsStatusParams{}

	relation := request.GetString("relation", "")
	if relation != "" {

		params.Relation = &relation
	}

	return params
}

func registerPostGroupsImportAuthorize(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_import_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 12.8"),
	)

	s.AddTool(tool, postGroupsImportAuthorizeHandler)
}

func postGroupsImportAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsImportAuthorize(ctx, authorizationHeader))
}

func registerGetGroupsIdPackages(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs",
		mcp.WithDescription("Get a list of project packages at the group level. This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("ID or URL-encoded path of the group"),
			mcp.Required(),
		),
		mcp.WithBoolean("exclude_subgroups",
			mcp.Description("Determines if subgroups should be excluded (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return packages ordered by `created_at`, `name`, `version` or `type` fields. (default: created_at)"),

			mcp.Enum("created_at", "name", "version", "type", "project_path"),
		),
		mcp.WithString("sort",
			mcp.Description("Return packages sorted in `asc` or `desc` order. (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("package_type",
			mcp.Description("Return packages of a certain type"),

			mcp.Enum("maven", "npm", "conan", "nuget", "pypi", "composer", "generic", "golang", "debian", "rubygems", "helm", "terraform_module", "rpm", "ml_model"),
		),
		mcp.WithString("package_name",
			mcp.Description("Return packages with this name"),
		),
		mcp.WithString("package_version",
			mcp.Description("Return packages with this version"),
		),
		mcp.WithBoolean("include_versionless",
			mcp.Description("Returns packages without a version"),
		),
		mcp.WithString("status",
			mcp.Description("Return packages with specified status"),

			mcp.Enum("default", "hidden", "processing", "error", "pending_destruction", "deprecated"),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesHandler)
}

func getGroupsIdPackagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdPackages(request)
	return toResult(c.GetApiV4GroupsIdPackages(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdPackages(request mcp.CallToolRequest) client.GetApiV4GroupsIdPackagesParams {
	params := client.GetApiV4GroupsIdPackagesParams{}

	exclude_subgroups := request.GetBool("exclude_subgroups", false)
	params.ExcludeSubgroups = &exclude_subgroups

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

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	package_type := request.GetString("package_type", "")
	if package_type != "" {

		params.PackageType = &package_type
	}

	package_name := request.GetString("package_name", "")
	if package_name != "" {

		params.PackageName = &package_name
	}

	package_version := request.GetString("package_version", "")
	if package_version != "" {

		params.PackageVersion = &package_version
	}

	include_versionless := request.GetBool("include_versionless", false)
	params.IncludeVersionless = &include_versionless

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	return params
}

func registerGetGroupsIdPlaceholderReassignments(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_placeholder_reassignments",
		mcp.WithDescription("This feature was added in GitLab 17.10"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdPlaceholderReassignmentsHandler)
}

func getGroupsIdPlaceholderReassignmentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4GroupsIdPlaceholderReassignments(ctx, id, authorizationHeader))
}

func registerPostGroupsIdPlaceholderReassignmentsAuthorize(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_placeholder_reassignments_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 17.10"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdPlaceholderReassignmentsAuthorizeHandler)
}

func postGroupsIdPlaceholderReassignmentsAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdPlaceholderReassignmentsAuthorize(ctx, id, authorizationHeader))
}

func registerGetGroupsIdVariables(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_variables",
		mcp.WithDescription("Get a list of group-level variables"),
		mcp.WithString("id",
			mcp.Description("The ID of a group or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdVariablesHandler)
}

func getGroupsIdVariablesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdVariables(request)
	return toResult(c.GetApiV4GroupsIdVariables(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdVariables(request mcp.CallToolRequest) client.GetApiV4GroupsIdVariablesParams {
	params := client.GetApiV4GroupsIdVariablesParams{}

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

func registerGetGroupsIdVariablesKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_variables_key",
		mcp.WithDescription("Get the details of a group’s specific variable"),
		mcp.WithString("id",
			mcp.Description("The ID of a group or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithString("key",
			mcp.Description("The key of the variable"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdVariablesKeyHandler)
}

func getGroupsIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	key := request.GetString("key", "")

	return toResult(c.GetApiV4GroupsIdVariablesKey(ctx, id, key, authorizationHeader))
}

func registerGetGroupsIdIntegrations(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_integrations",
		mcp.WithDescription("Get a list of all active integrations."),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdIntegrationsHandler)
}

func getGroupsIdIntegrationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdIntegrations(ctx, id, authorizationHeader))
}

func registerGetGroupsIdIntegrationsSlug(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_integrations_slug",
		mcp.WithDescription("Get the integration settings."),
		mcp.WithString("slug",
			mcp.Description("The name of the integration"),
			mcp.Required(),
			mcp.Enum("apple-app-store", "asana", "assembla", "bamboo", "bugzilla", "buildkite", "campfire", "confluence", "custom-issue-tracker", "datadog", "diffblue-cover", "discord", "drone-ci", "emails-on-push", "external-wiki", "gitlab-slack-application", "google-play", "hangouts-chat", "harbor", "irker", "jenkins", "jira", "jira-cloud-app", "matrix", "mattermost-slash-commands", "slack-slash-commands", "packagist", "phorge", "pipelines-email", "pivotaltracker", "pumble", "pushover", "redmine", "ewm", "youtrack", "clickup", "slack", "microsoft-teams", "mattermost", "teamcity", "telegram", "unify-circuit", "webex-teams", "zentao", "squash-tm", "github", "git-guardian", "google-cloud-platform-artifact-registry", "google-cloud-platform-workload-identity-federation", "mock-ci", "mock-monitoring"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdIntegrationsSlugHandler)
}

func getGroupsIdIntegrationsSlugHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	slug := request.GetString("slug", "")

	return toResult(c.GetApiV4GroupsIdIntegrationsSlug(ctx, id, slug, authorizationHeader))
}

func registerGetGroupsIdInvitations(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_invitations",
		mcp.WithDescription("This feature was introduced in GitLab 13.6"),
		mcp.WithString("id",
			mcp.Description("The group ID"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("query",
			mcp.Description("A query string to search for members"),
		),
	)

	s.AddTool(tool, getGroupsIdInvitationsHandler)
}

func getGroupsIdInvitationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdInvitations(request)
	return toResult(c.GetApiV4GroupsIdInvitations(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdInvitations(request mcp.CallToolRequest) client.GetApiV4GroupsIdInvitationsParams {
	params := client.GetApiV4GroupsIdInvitationsParams{}

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

	query := request.GetString("query", "")
	if query != "" {

		params.Query = &query
	}

	return params
}

func registerGetGroupsIdUploads(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_uploads",
		mcp.WithDescription("Get the list of uploads of a group"),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdUploadsHandler)
}

func getGroupsIdUploadsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdUploads(request)
	return toResult(c.GetApiV4GroupsIdUploads(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdUploads(request mcp.CallToolRequest) client.GetApiV4GroupsIdUploadsParams {
	params := client.GetApiV4GroupsIdUploadsParams{}

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

func registerGetGroupsIdUploadsUploadId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_uploads_upload_id",
		mcp.WithDescription("Download a single group upload by ID"),
		mcp.WithNumber("upload_id",
			mcp.Description("The ID of a group upload"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdUploadsUploadIdHandler)
}

func getGroupsIdUploadsUploadIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	upload_id := int32(request.GetInt("upload_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdUploadsUploadId(ctx, id, upload_id, authorizationHeader))
}

func registerGetGroupsIdUploadsSecretFilename(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_uploads_secret_filename",
		mcp.WithDescription("Download a single project upload by secret and filename"),
		mcp.WithString("secret",
			mcp.Description("The 32-character secret of a group upload"),
			mcp.Required(),
		),
		mcp.WithString("filename",
			mcp.Description("The filename of a group upload"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdUploadsSecretFilenameHandler)
}

func getGroupsIdUploadsSecretFilenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	secret := request.GetString("secret", "")
	filename := request.GetString("filename", "")

	return toResult(c.GetApiV4GroupsIdUploadsSecretFilename(ctx, id, secret, filename, authorizationHeader))
}

func registerGetGroupsIdMembers(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_members",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user."),
		mcp.WithString("id",
			mcp.Description("The group ID"),
			mcp.Required(),
		),
		mcp.WithString("query",
			mcp.Description("A query string to search for members"),
		),
		mcp.WithString("user_ids",
			mcp.Description("Array of user ids to look up for membership"),
		),
		mcp.WithString("skip_users",
			mcp.Description("Array of user ids to be skipped for membership"),
		),
		mcp.WithBoolean("show_seat_info",
			mcp.Description("Show seat information for members"),
		),
		mcp.WithBoolean("with_saml_identity",
			mcp.Description("List only members with linked SAML identity"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdMembersHandler)
}

func getGroupsIdMembersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdMembers(request)
	return toResult(c.GetApiV4GroupsIdMembers(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdMembers(request mcp.CallToolRequest) client.GetApiV4GroupsIdMembersParams {
	params := client.GetApiV4GroupsIdMembersParams{}

	query := request.GetString("query", "")
	if query != "" {

		params.Query = &query
	}

	user_ids := request.GetString("user_ids", "")
	if user_ids != "" {
		user_ids := strings.Split(user_ids, ",")
		var intSlice []int32
		for _, v := range user_ids {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.UserIds = &intSlice
	}

	skip_users := request.GetString("skip_users", "")
	if skip_users != "" {
		skip_users := strings.Split(skip_users, ",")
		var intSlice []int32
		for _, v := range skip_users {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.SkipUsers = &intSlice
	}

	show_seat_info := request.GetBool("show_seat_info", false)
	params.ShowSeatInfo = &show_seat_info

	with_saml_identity := request.GetBool("with_saml_identity", false)
	params.WithSamlIdentity = &with_saml_identity

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

func registerGetGroupsIdMembersAll(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_members_all",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user, including those who gained membership through ancestor group."),
		mcp.WithString("id",
			mcp.Description("The group ID"),
			mcp.Required(),
		),
		mcp.WithString("query",
			mcp.Description("A query string to search for members"),
		),
		mcp.WithString("user_ids",
			mcp.Description("Array of user ids to look up for membership"),
		),
		mcp.WithBoolean("show_seat_info",
			mcp.Description("Show seat information for members"),
		),
		mcp.WithString("state",
			mcp.Description("Filter results by member state"),

			mcp.Enum("awaiting", "active"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdMembersAllHandler)
}

func getGroupsIdMembersAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdMembersAll(request)
	return toResult(c.GetApiV4GroupsIdMembersAll(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdMembersAll(request mcp.CallToolRequest) client.GetApiV4GroupsIdMembersAllParams {
	params := client.GetApiV4GroupsIdMembersAllParams{}

	query := request.GetString("query", "")
	if query != "" {

		params.Query = &query
	}

	user_ids := request.GetString("user_ids", "")
	if user_ids != "" {
		user_ids := strings.Split(user_ids, ",")
		var intSlice []int32
		for _, v := range user_ids {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.UserIds = &intSlice
	}

	show_seat_info := request.GetBool("show_seat_info", false)
	params.ShowSeatInfo = &show_seat_info

	state := request.GetString("state", "")
	if state != "" {

		params.State = &state
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

func registerGetGroupsIdMembersUserId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_members_user_id",
		mcp.WithDescription("Gets a member of a group or project."),
		mcp.WithString("id",
			mcp.Description("The group ID"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdMembersUserIdHandler)
}

func getGroupsIdMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdMembersUserId(ctx, id, user_id, authorizationHeader))
}

func registerGetGroupsIdMembersAllUserId(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_members_all_user_id",
		mcp.WithDescription("Gets a member of a group or project, including those who gained membership through ancestor group"),
		mcp.WithString("id",
			mcp.Description("The group ID"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdMembersAllUserIdHandler)
}

func getGroupsIdMembersAllUserIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdMembersAllUserId(ctx, id, user_id, authorizationHeader))
}

func registerPostGroupsIdMembersUserIdOverride(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_members_user_id_override",
		mcp.WithDescription("Overrides the access level of an LDAP group member."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdMembersUserIdOverrideHandler)
}

func postGroupsIdMembersUserIdOverrideHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))

	return toResult(c.PostApiV4GroupsIdMembersUserIdOverride(ctx, id, user_id, authorizationHeader))
}

func registerPutGroupsIdMembersMemberIdApprove(s *server.MCPServer) {
	tool := mcp.NewTool("put_grps_id_members_member_id_approve",
		mcp.WithDescription("Approves a pending member"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithNumber("member_id",
			mcp.Description("The ID of the member requiring approval"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, putGroupsIdMembersMemberIdApproveHandler)
}

func putGroupsIdMembersMemberIdApproveHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	member_id := int32(request.GetInt("member_id", math.MinInt))

	return toResult(c.PutApiV4GroupsIdMembersMemberIdApprove(ctx, id, member_id, authorizationHeader))
}

func registerPostGroupsIdMembersApproveAll(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_members_approve_all",
		mcp.WithDescription("Approves all pending members"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdMembersApproveAllHandler)
}

func postGroupsIdMembersApproveAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdMembersApproveAll(ctx, id, authorizationHeader))
}

func registerGetGroupsIdPendingMembers(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pending_members",
		mcp.WithDescription("Lists all pending members for a group including invited users"),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdPendingMembersHandler)
}

func getGroupsIdPendingMembersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdPendingMembers(request)
	return toResult(c.GetApiV4GroupsIdPendingMembers(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdPendingMembers(request mcp.CallToolRequest) client.GetApiV4GroupsIdPendingMembersParams {
	params := client.GetApiV4GroupsIdPendingMembersParams{}

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

func registerGetGroupsIdBillableMembers(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_billable_members",
		mcp.WithDescription("Gets a list of billable users of top-level group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("search",
			mcp.Description("The exact name of the subscribed member"),
		),
		mcp.WithString("sort",
			mcp.Description("The sorting option"),

			mcp.Enum("access_level_asc", "access_level_desc", "last_joined", "name_asc", "name_desc", "oldest_joined", "oldest_sign_in", "recent_sign_in", "last_activity_on_asc", "last_activity_on_desc"),
		),
	)

	s.AddTool(tool, getGroupsIdBillableMembersHandler)
}

func getGroupsIdBillableMembersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdBillableMembers(request)
	return toResult(c.GetApiV4GroupsIdBillableMembers(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdBillableMembers(request mcp.CallToolRequest) client.GetApiV4GroupsIdBillableMembersParams {
	params := client.GetApiV4GroupsIdBillableMembersParams{}

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

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	return params
}

func registerGetGroupsIdBillableMembersUserIdMemberships(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_billable_members_user_id_memberships",
		mcp.WithDescription("Get the direct memberships of a billable user of a top-level group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdBillableMembersUserIdMembershipsHandler)
}

func getGroupsIdBillableMembersUserIdMembershipsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))
	params := parseGetGroupsIdBillableMembersUserIdMemberships(request)
	return toResult(c.GetApiV4GroupsIdBillableMembersUserIdMemberships(ctx, id, user_id, &params, authorizationHeader))
}

func parseGetGroupsIdBillableMembersUserIdMemberships(request mcp.CallToolRequest) client.GetApiV4GroupsIdBillableMembersUserIdMembershipsParams {
	params := client.GetApiV4GroupsIdBillableMembersUserIdMembershipsParams{}

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

func registerGetGroupsIdBillableMembersUserIdIndirect(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_billable_members_user_id_indirect",
		mcp.WithDescription("Get the indirect memberships of a billable user of a top-level group."),
		mcp.WithString("id",
			mcp.Description("The ID of a group"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdBillableMembersUserIdIndirectHandler)
}

func getGroupsIdBillableMembersUserIdIndirectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))
	params := parseGetGroupsIdBillableMembersUserIdIndirect(request)
	return toResult(c.GetApiV4GroupsIdBillableMembersUserIdIndirect(ctx, id, user_id, &params, authorizationHeader))
}

func parseGetGroupsIdBillableMembersUserIdIndirect(request mcp.CallToolRequest) client.GetApiV4GroupsIdBillableMembersUserIdIndirectParams {
	params := client.GetApiV4GroupsIdBillableMembersUserIdIndirectParams{}

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

func registerGetGroupsIdMergeRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_mrs",
		mcp.WithDescription("Get all merge requests for this group and its subgroups."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user."),
			mcp.Required(),
		),
		mcp.WithNumber("author_id",
			mcp.Description("Returns merge requests created by the given user `id`. Mutually exclusive with `author_username`. Combine with `scope=all` or `scope=assigned_to_me`."),
		),
		mcp.WithString("author_username",
			mcp.Description("Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithNumber("assignee_id",
			mcp.Description("Returns merge requests assigned to the given user `id`. `None` returns unassigned merge requests. `Any` returns merge requests with an assignee."),
		),
		mcp.WithString("assignee_username",
			mcp.Description("Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithString("reviewer_username",
			mcp.Description("Returns merge requests which have the user as a reviewer with the given `username`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_id`. Introduced in GitLab 13.8."),
		),
		mcp.WithString("labels",
			mcp.Description("Returns merge requests matching a comma-separated list of labels. `None` lists all merge requests with no labels. `Any` lists all merge requests with at least one label. Predefined names are case-insensitive."),
		),
		mcp.WithString("milestone",
			mcp.Description("Returns merge requests for a specific milestone. `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone."),
		),
		mcp.WithString("my_reaction_emoji",
			mcp.Description("Returns merge requests reacted by the authenticated user by the given `emoji`. `None` returns issues not given a reaction. `Any` returns issues given at least one reaction."),
		),
		mcp.WithNumber("reviewer_id",
			mcp.Description("Returns merge requests which have the user as a reviewer with the given user `id`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_username`."),
		),
		mcp.WithString("state",
			mcp.Description("Returns `all` merge requests or just those that are `opened`, `closed`, `locked`, or `merged`. (default: all)"),

			mcp.Enum("opened", "closed", "locked", "merged", "all"),
		),
		mcp.WithString("order_by",
			mcp.Description("Returns merge requests ordered by `created_at`, `label_priority`, `milestone_due`, `popularity`, `priority`, `title`, `updated_at` or `merged_at` fields. Introduced in GitLab 14.8. (default: created_at)"),

			mcp.Enum("created_at", "label_priority", "milestone_due", "popularity", "priority", "title", "updated_at", "merged_at"),
		),
		mcp.WithString("sort",
			mcp.Description("Returns merge requests sorted in `asc` or `desc` order. (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("with_labels_details",
			mcp.Description("If `true`, response returns more details for each label in labels field: `:name`,`:color`, `:description`, `:description_html`, `:text_color` (default: false)"),
		),
		mcp.WithBoolean("with_merge_status_recheck",
			mcp.Description("If `true`, this projection requests (but does not guarantee) that the `merge_status` field be recalculated asynchronously. Introduced in GitLab 13.0. (default: false)"),
		),
		mcp.WithString("created_after",
			mcp.Description("Returns merge requests created on or after the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("created_before",
			mcp.Description("Returns merge requests created on or before the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Returns merge requests updated on or after the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Returns merge requests updated on or before the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("view",
			mcp.Description("If simple, returns the `iid`, URL, title, description, and basic state of merge request"),

			mcp.Enum("simple"),
		),
		mcp.WithString("scope",
			mcp.Description("Returns merge requests for the given scope: `created_by_me`, `assigned_to_me` or `all`"),

			mcp.Enum("created-by-me", "assigned-to-me", "created_by_me", "assigned_to_me", "all"),
		),
		mcp.WithString("source_branch",
			mcp.Description("Returns merge requests with the given source branch"),
		),
		mcp.WithNumber("source_project_id",
			mcp.Description("Returns merge requests with the given source project id"),
		),
		mcp.WithString("target_branch",
			mcp.Description("Returns merge requests with the given target branch"),
		),
		mcp.WithString("search",
			mcp.Description("Search merge requests against their `title` and `description`."),
		),
		mcp.WithString("in",
			mcp.Description("Modify the scope of the search attribute. `title`, `description`, or a string joining them with comma. (example: title,description)"),
		),
		mcp.WithString("wip",
			mcp.Description("Filter merge requests against their `wip` status. `yes` to return only draft merge requests, `no` to return non-draft merge requests."),

			mcp.Enum("yes", "no"),
		),
		mcp.WithNumber("not[author_id]",
			mcp.Description("`<Negated>` Returns merge requests created by the given user `id`. Mutually exclusive with `author_username`. Combine with `scope=all` or `scope=assigned_to_me`."),
		),
		mcp.WithString("not[author_username]",
			mcp.Description("`<Negated>` Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithNumber("not[assignee_id]",
			mcp.Description("`<Negated>` Returns merge requests assigned to the given user `id`. `None` returns unassigned merge requests. `Any` returns merge requests with an assignee."),
		),
		mcp.WithString("not[assignee_username]",
			mcp.Description("`<Negated>` Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithString("not[reviewer_username]",
			mcp.Description("`<Negated>` Returns merge requests which have the user as a reviewer with the given `username`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_id`. Introduced in GitLab 13.8."),
		),
		mcp.WithString("not[labels]",
			mcp.Description("`<Negated>` Returns merge requests matching a comma-separated list of labels. `None` lists all merge requests with no labels. `Any` lists all merge requests with at least one label. Predefined names are case-insensitive."),
		),
		mcp.WithString("not[milestone]",
			mcp.Description("`<Negated>` Returns merge requests for a specific milestone. `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone."),
		),
		mcp.WithString("not[my_reaction_emoji]",
			mcp.Description("`<Negated>` Returns merge requests reacted by the authenticated user by the given `emoji`. `None` returns issues not given a reaction. `Any` returns issues given at least one reaction."),
		),
		mcp.WithNumber("not[reviewer_id]",
			mcp.Description("`<Negated>` Returns merge requests which have the user as a reviewer with the given user `id`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_username`."),
		),
		mcp.WithString("deployed_before",
			mcp.Description("Returns merge requests deployed before the given date/time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("deployed_after",
			mcp.Description("Returns merge requests deployed after the given date/time. Expected in ISO 8601 format (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("environment",
			mcp.Description("Returns merge requests deployed to the given environment (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("approved",
			mcp.Description("Filters merge requests by their `approved` status. `yes` returns only approved merge requests. `no` returns only non-approved merge requests."),

			mcp.Enum("yes", "no"),
		),
		mcp.WithNumber("merge_user_id",
			mcp.Description("Returns merge requests which have been merged by the user with the given user `id`. Mutually exclusive with `merge_user_username`."),
		),
		mcp.WithString("merge_user_username",
			mcp.Description("Returns merge requests which have been merged by the user with the given `username`. Mutually exclusive with `merge_user_id`."),
		),
		mcp.WithString("approver_ids",
			mcp.Description("Return merge requests which have specified the users with the given IDs as an individual approver"),
		),
		mcp.WithString("approved_by_ids",
			mcp.Description("Return merge requests which have been approved by the specified users with the given IDs"),
		),
		mcp.WithString("approved_by_usernames",
			mcp.Description("Return merge requests which have been approved by the specified users with the given usernames"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("non_archived",
			mcp.Description("Returns merge requests from non archived projects only. (default: true)"),
		),
	)

	s.AddTool(tool, getGroupsIdMergeRequestsHandler)
}

func getGroupsIdMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdMergeRequests(request)
	return toResult(c.GetApiV4GroupsIdMergeRequests(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdMergeRequests(request mcp.CallToolRequest) client.GetApiV4GroupsIdMergeRequestsParams {
	params := client.GetApiV4GroupsIdMergeRequestsParams{}

	author_id := request.GetInt("author_id", math.MinInt)
	if author_id != math.MinInt {
		author_id := int32(author_id)
		params.AuthorId = &author_id
	}

	author_username := request.GetString("author_username", "")
	if author_username != "" {

		params.AuthorUsername = &author_username
	}

	assignee_id := request.GetInt("assignee_id", math.MinInt)
	if assignee_id != math.MinInt {
		assignee_id := int32(assignee_id)
		params.AssigneeId = &assignee_id
	}

	assignee_username := request.GetString("assignee_username", "")
	if assignee_username != "" {
		assignee_username := strings.Split(assignee_username, ",")
		params.AssigneeUsername = &assignee_username
	}

	reviewer_username := request.GetString("reviewer_username", "")
	if reviewer_username != "" {

		params.ReviewerUsername = &reviewer_username
	}

	labels := request.GetString("labels", "")
	if labels != "" {
		labels := strings.Split(labels, ",")
		params.Labels = &labels
	}

	milestone := request.GetString("milestone", "")
	if milestone != "" {

		params.Milestone = &milestone
	}

	my_reaction_emoji := request.GetString("my_reaction_emoji", "")
	if my_reaction_emoji != "" {

		params.MyReactionEmoji = &my_reaction_emoji
	}

	reviewer_id := request.GetInt("reviewer_id", math.MinInt)
	if reviewer_id != math.MinInt {
		reviewer_id := int32(reviewer_id)
		params.ReviewerId = &reviewer_id
	}

	state := request.GetString("state", "")
	if state != "" {

		params.State = &state
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	with_labels_details := request.GetBool("with_labels_details", false)
	params.WithLabelsDetails = &with_labels_details

	with_merge_status_recheck := request.GetBool("with_merge_status_recheck", false)
	params.WithMergeStatusRecheck = &with_merge_status_recheck

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
	}

	updated_after := request.GetString("updated_after", "")
	if updated_after != "" {
		updated_after, _ := time.Parse(time.RFC3339, updated_after)
		params.UpdatedAfter = &updated_after
	}

	updated_before := request.GetString("updated_before", "")
	if updated_before != "" {
		updated_before, _ := time.Parse(time.RFC3339, updated_before)
		params.UpdatedBefore = &updated_before
	}

	view := request.GetString("view", "")
	if view != "" {

		params.View = &view
	}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	source_branch := request.GetString("source_branch", "")
	if source_branch != "" {

		params.SourceBranch = &source_branch
	}

	source_project_id := request.GetInt("source_project_id", math.MinInt)
	if source_project_id != math.MinInt {
		source_project_id := int32(source_project_id)
		params.SourceProjectId = &source_project_id
	}

	target_branch := request.GetString("target_branch", "")
	if target_branch != "" {

		params.TargetBranch = &target_branch
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	in := request.GetString("in", "")
	if in != "" {

		params.In = &in
	}

	wip := request.GetString("wip", "")
	if wip != "" {

		params.Wip = &wip
	}

	not_author_id_ := request.GetInt("not[author_id]", math.MinInt)
	if not_author_id_ != math.MinInt {
		not_author_id_ := int32(not_author_id_)
		params.NotAuthorId = &not_author_id_
	}

	not_author_username_ := request.GetString("not[author_username]", "")
	if not_author_username_ != "" {

		params.NotAuthorUsername = &not_author_username_
	}

	not_assignee_id_ := request.GetInt("not[assignee_id]", math.MinInt)
	if not_assignee_id_ != math.MinInt {
		not_assignee_id_ := int32(not_assignee_id_)
		params.NotAssigneeId = &not_assignee_id_
	}

	not_assignee_username_ := request.GetString("not[assignee_username]", "")
	if not_assignee_username_ != "" {
		not_assignee_username_ := strings.Split(not_assignee_username_, ",")
		params.NotAssigneeUsername = &not_assignee_username_
	}

	not_reviewer_username_ := request.GetString("not[reviewer_username]", "")
	if not_reviewer_username_ != "" {

		params.NotReviewerUsername = &not_reviewer_username_
	}

	not_labels_ := request.GetString("not[labels]", "")
	if not_labels_ != "" {
		not_labels_ := strings.Split(not_labels_, ",")
		params.NotLabels = &not_labels_
	}

	not_milestone_ := request.GetString("not[milestone]", "")
	if not_milestone_ != "" {

		params.NotMilestone = &not_milestone_
	}

	not_my_reaction_emoji_ := request.GetString("not[my_reaction_emoji]", "")
	if not_my_reaction_emoji_ != "" {

		params.NotMyReactionEmoji = &not_my_reaction_emoji_
	}

	not_reviewer_id_ := request.GetInt("not[reviewer_id]", math.MinInt)
	if not_reviewer_id_ != math.MinInt {
		not_reviewer_id_ := int32(not_reviewer_id_)
		params.NotReviewerId = &not_reviewer_id_
	}

	deployed_before := request.GetString("deployed_before", "")
	if deployed_before != "" {

		params.DeployedBefore = &deployed_before
	}

	deployed_after := request.GetString("deployed_after", "")
	if deployed_after != "" {

		params.DeployedAfter = &deployed_after
	}

	environment := request.GetString("environment", "")
	if environment != "" {

		params.Environment = &environment
	}

	approved := request.GetString("approved", "")
	if approved != "" {

		params.Approved = &approved
	}

	merge_user_id := request.GetInt("merge_user_id", math.MinInt)
	if merge_user_id != math.MinInt {
		merge_user_id := int32(merge_user_id)
		params.MergeUserId = &merge_user_id
	}

	merge_user_username := request.GetString("merge_user_username", "")
	if merge_user_username != "" {

		params.MergeUserUsername = &merge_user_username
	}

	approver_ids := request.GetString("approver_ids", "")
	if approver_ids != "" {

		params.ApproverIds = &approver_ids
	}

	approved_by_ids := request.GetString("approved_by_ids", "")
	if approved_by_ids != "" {

		params.ApprovedByIds = &approved_by_ids
	}

	approved_by_usernames := request.GetString("approved_by_usernames", "")
	if approved_by_usernames != "" {

		params.ApprovedByUsernames = &approved_by_usernames
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

	non_archived := request.GetBool("non_archived", true)
	params.NonArchived = &non_archived

	return params
}

func registerPostGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_pkgs_npm_npm_v1_security_advisories_bulk",
		mcp.WithDescription("This feature was introduced in GitLab 15.6"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkHandler)
}

func postGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx, id, authorizationHeader))
}

func registerPostGroupsIdPackagesNpmNpmV1SecurityAuditsQuick(s *server.MCPServer) {
	tool := mcp.NewTool("post_grps_id_pkgs_npm_npm_v1_security_audits_quick",
		mcp.WithDescription("This feature was introduced in GitLab 15.6"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postGroupsIdPackagesNpmNpmV1SecurityAuditsQuickHandler)
}

func postGroupsIdPackagesNpmNpmV1SecurityAuditsQuickHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx, id, authorizationHeader))
}

func registerGetGroupsIdPackagesNugetIndex(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs_nuget_index",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithNumber("id",
			mcp.Description("The group ID or full group path."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesNugetIndexHandler)
}

func getGroupsIdPackagesNugetIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdPackagesNugetIndex(ctx, id, authorizationHeader))
}

func registerGetGroupsIdPackagesNugetV2(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs_nuget_v2",
		mcp.WithDescription("This feature was introduced in GitLab 16.2"),
		mcp.WithNumber("id",
			mcp.Description("The group ID or full group path."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesNugetV2Handler)
}

func getGroupsIdPackagesNugetV2Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdPackagesNugetV2(ctx, id, authorizationHeader))
}

func registerGetGroupsIdPackagesNugetV2Metadata(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs_nuget_v2_metadata",
		mcp.WithDescription("This feature was introduced in GitLab 16.3"),
		mcp.WithNumber("id",
			mcp.Description("The group ID or full group path."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesNugetV2MetadataHandler)
}

func getGroupsIdPackagesNugetV2MetadataHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdPackagesNugetV2Metadata(ctx, id, authorizationHeader))
}

func registerGetGroupsIdPackagesNugetQuery(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs_nuget_query",
		mcp.WithDescription("This feature was introduced in GitLab 12.8"),
		mcp.WithNumber("id",
			mcp.Description("The group ID or full group path."),
			mcp.Required(),
		),
		mcp.WithString("q",
			mcp.Description("The search term (example: MyNuGet)"),
		),
		mcp.WithNumber("skip",
			mcp.Description("The number of results to skip (example: 1) (default: 0)"),
		),
		mcp.WithNumber("take",
			mcp.Description("The number of results to return (example: 1) (default: 20)"),
		),
		mcp.WithBoolean("prerelease",
			mcp.Description("Include prerelease versions (default: true)"),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesNugetQueryHandler)
}

func getGroupsIdPackagesNugetQueryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdPackagesNugetQuery(request)
	return toResult(c.GetApiV4GroupsIdPackagesNugetQuery(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdPackagesNugetQuery(request mcp.CallToolRequest) client.GetApiV4GroupsIdPackagesNugetQueryParams {
	params := client.GetApiV4GroupsIdPackagesNugetQueryParams{}

	q := request.GetString("q", "")
	if q != "" {

		params.Q = &q
	}

	skip := request.GetInt("skip", 0)
	if skip != math.MinInt {
		skip := int32(skip)
		params.Skip = &skip
	}

	take := request.GetInt("take", 20)
	if take != math.MinInt {
		take := int32(take)
		params.Take = &take
	}

	prerelease := request.GetBool("prerelease", true)
	params.Prerelease = &prerelease

	return params
}

func registerGetGroupsIdPackagesPypiSimple(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_pkgs_pypi_simple",
		mcp.WithDescription("This feature was introduced in GitLab 15.1"),
		mcp.WithNumber("id",
			mcp.Description("The ID or full path of the group."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdPackagesPypiSimpleHandler)
}

func getGroupsIdPackagesPypiSimpleHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4GroupsIdPackagesPypiSimple(ctx, id, authorizationHeader))
}

func registerGetGroupsIdReleases(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_releases",
		mcp.WithDescription("Returns a list of group releases."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the group owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithString("sort",
			mcp.Description("The direction of the order. Either `desc` (default) for descending order or `asc` for ascending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("simple",
			mcp.Description("Return only limited fields for each release (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getGroupsIdReleasesHandler)
}

func getGroupsIdReleasesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdReleases(request)
	return toResult(c.GetApiV4GroupsIdReleases(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdReleases(request mcp.CallToolRequest) client.GetApiV4GroupsIdReleasesParams {
	params := client.GetApiV4GroupsIdReleasesParams{}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	simple := request.GetBool("simple", false)
	params.Simple = &simple

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

func registerGetGroupsIdWikis(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_wikis",
		mcp.WithDescription("Get a list of wiki pages"),
		mcp.WithBoolean("with_content",
			mcp.Description("Include pages' content (default: false)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdWikisHandler)
}

func getGroupsIdWikisHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetGroupsIdWikis(request)
	return toResult(c.GetApiV4GroupsIdWikis(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdWikis(request mcp.CallToolRequest) client.GetApiV4GroupsIdWikisParams {
	params := client.GetApiV4GroupsIdWikisParams{}

	with_content := request.GetBool("with_content", false)
	params.WithContent = &with_content

	return params
}

func registerGetGroupsIdWikisSlug(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_wikis_slug",
		mcp.WithDescription("Get a wiki page"),
		mcp.WithString("slug",
			mcp.Description("The slug of a wiki page"),
			mcp.Required(),
		),
		mcp.WithString("version",
			mcp.Description("The version hash of a wiki page"),
		),
		mcp.WithBoolean("render_html",
			mcp.Description("Render content to HTML (default: false)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupsIdWikisSlugHandler)
}

func getGroupsIdWikisSlugHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	slug := request.GetString("slug", "")
	params := parseGetGroupsIdWikisSlug(request)
	return toResult(c.GetApiV4GroupsIdWikisSlug(ctx, id, slug, &params, authorizationHeader))
}

func parseGetGroupsIdWikisSlug(request mcp.CallToolRequest) client.GetApiV4GroupsIdWikisSlugParams {
	params := client.GetApiV4GroupsIdWikisSlugParams{}

	version := request.GetString("version", "")
	if version != "" {

		params.Version = &version
	}

	render_html := request.GetBool("render_html", false)
	params.RenderHtml = &render_html

	return params
}

func registerGetGroupsIdIssues(s *server.MCPServer) {
	tool := mcp.NewTool("get_grps_id_issues",
		mcp.WithDescription("List group issues"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the group."),
			mcp.Required(),
		),
		mcp.WithNumber("assignee_id",
			mcp.Description("Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee."),
		),
		mcp.WithString("assignee_username",
			mcp.Description("Return issues assigned to the given username. Similar to assignee_id and mutually exclusive with assignee_id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned."),
		),
		mcp.WithNumber("author_id",
			mcp.Description("Return issues created by the given user id. Mutually exclusive with author_username. Combine with scope=all or scope=assigned_to_me."),
		),
		mcp.WithString("author_username",
			mcp.Description("Return issues created by the given username. Similar to author_id and mutually exclusive with author_id."),
		),
		mcp.WithBoolean("confidential",
			mcp.Description("Filter confidential or public issues."),
		),
		mcp.WithString("created_after",
			mcp.Description("Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithString("created_before",
			mcp.Description("Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithString("due_date",
			mcp.Description("Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next_month_and_previous_two_weeks."),
		),
		mcp.WithNumber("epic_id",
			mcp.Description("Return issues associated with the given epic ID. None returns issues that are not associated with an epic. Any returns issues that are associated with an epic. Premium and Ultimate only."),
		),
		mcp.WithString("iids[]",
			mcp.Description("Return only the issues having the given iid."),
		),
		mcp.WithString("issue_type",
			mcp.Description("Filter to a given type of issue. One of issue, incident, test_case or task."),
		),
		mcp.WithNumber("iteration_id",
			mcp.Description("Return issues assigned to the given iteration ID. None returns issues that do not belong to an iteration. Any returns issues that belong to an iteration. Mutually exclusive with iteration_title. Premium and Ultimate only."),
		),
		mcp.WithString("iteration_title",
			mcp.Description("Return issues assigned to the iteration with the given title. Similar to iteration_id and mutually exclusive with iteration_id. Premium and Ultimate only."),
		),
		mcp.WithString("labels",
			mcp.Description("Comma-separated list of label names, issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive."),
		),
		mcp.WithString("milestone",
			mcp.Description("The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone."),
		),
		mcp.WithString("my_reaction_emoji",
			mcp.Description("Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction."),
		),
		mcp.WithBoolean("non_archived",
			mcp.Description("Return issues from non archived projects. Default is true."),
		),
		mcp.WithString("not",
			mcp.Description("Return issues that do not match the parameters supplied. Accepts: labels, milestone, author_id, author_username, assignee_id, assignee_username, my_reaction_emoji, search, in."),
		),
		mcp.WithString("order_by",
			mcp.Description("Return issues ordered by created_at, updated_at, priority, due_date, relative_position, label_priority, milestone_due, popularity, weight fields. Default is created_at"),
		),
		mcp.WithString("scope",
			mcp.Description("Return issues for the given scope: created_by_me, assigned_to_me or all. Defaults to all."),
		),
		mcp.WithString("search",
			mcp.Description("Search group issues against their title and description."),
		),
		mcp.WithString("sort",
			mcp.Description("Return issues sorted in asc or desc order. Default is desc."),
		),
		mcp.WithString("state",
			mcp.Description("Return all issues or just those that are opened or closed."),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithNumber("weight",
			mcp.Description("Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned. Premium and Ultimate only."),
		),
		mcp.WithBoolean("with_labels_details",
			mcp.Description("If true, the response returns more details for each label in labels field: :name, :color, :description, :description_html, :text_color. Default is false."),
		),
	)

	s.AddTool(tool, getGroupsIdIssuesHandler)
}

func getGroupsIdIssuesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetGroupsIdIssues(request)
	return toResult(c.GetApiV4GroupsIdIssues(ctx, id, &params, authorizationHeader))
}

func parseGetGroupsIdIssues(request mcp.CallToolRequest) client.GetApiV4GroupsIdIssuesParams {
	params := client.GetApiV4GroupsIdIssuesParams{}

	assignee_id := request.GetInt("assignee_id", math.MinInt)
	if assignee_id != math.MinInt {

		params.AssigneeId = &assignee_id
	}

	assignee_username := request.GetString("assignee_username", "")
	if assignee_username != "" {
		assignee_username := strings.Split(assignee_username, ",")
		params.AssigneeUsername = &assignee_username
	}

	author_id := request.GetInt("author_id", math.MinInt)
	if author_id != math.MinInt {

		params.AuthorId = &author_id
	}

	author_username := request.GetString("author_username", "")
	if author_username != "" {

		params.AuthorUsername = &author_username
	}

	confidential := request.GetBool("confidential", false)
	params.Confidential = &confidential

	created_after := request.GetString("created_after", "")
	if created_after != "" {

		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {

		params.CreatedBefore = &created_before
	}

	due_date := request.GetString("due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	epic_id := request.GetInt("epic_id", math.MinInt)
	if epic_id != math.MinInt {

		params.EpicId = &epic_id
	}

	iids__ := request.GetString("iids[]", "")
	if iids__ != "" {
		iids__ := strings.Split(iids__, ",")
		var intSlice []int
		for _, v := range iids__ {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, intValue)
		}
		params.Iids = &intSlice
	}

	issue_type := request.GetString("issue_type", "")
	if issue_type != "" {

		params.IssueType = &issue_type
	}

	iteration_id := request.GetInt("iteration_id", math.MinInt)
	if iteration_id != math.MinInt {

		params.IterationId = &iteration_id
	}

	iteration_title := request.GetString("iteration_title", "")
	if iteration_title != "" {

		params.IterationTitle = &iteration_title
	}

	labels := request.GetString("labels", "")
	if labels != "" {

		params.Labels = &labels
	}

	milestone := request.GetString("milestone", "")
	if milestone != "" {

		params.Milestone = &milestone
	}

	my_reaction_emoji := request.GetString("my_reaction_emoji", "")
	if my_reaction_emoji != "" {

		params.MyReactionEmoji = &my_reaction_emoji
	}

	non_archived := request.GetBool("non_archived", false)
	params.NonArchived = &non_archived

	not := request.GetString("not", "")
	if not != "" {

		params.Not = &not
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	state := request.GetString("state", "")
	if state != "" {

		params.State = &state
	}

	updated_after := request.GetString("updated_after", "")
	if updated_after != "" {

		params.UpdatedAfter = &updated_after
	}

	updated_before := request.GetString("updated_before", "")
	if updated_before != "" {

		params.UpdatedBefore = &updated_before
	}

	weight := request.GetInt("weight", math.MinInt)
	if weight != math.MinInt {

		params.Weight = &weight
	}

	with_labels_details := request.GetBool("with_labels_details", false)
	params.WithLabelsDetails = &with_labels_details

	return params
}
