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

func registerGetProjectsIdAccessRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdAccessRequestsHandler)
}

func getProjectsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdAccessRequests(request)
	return toResult(c.GetApiV4ProjectsIdAccessRequests(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdAccessRequests(request mcp.CallToolRequest) client.GetApiV4ProjectsIdAccessRequestsParams {
	params := client.GetApiV4ProjectsIdAccessRequestsParams{}

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

func registerGetProjectsIdAlertManagementAlertsAlertIidMetricImages(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_alert_management_alerts_alert_iid_metric_images",
		mcp.WithDescription("Metric Images for alert"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 17)"),
			mcp.Required(),
		),
		mcp.WithNumber("alert_iid",
			mcp.Description("The IID of the Alert (example: 23)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdAlertManagementAlertsAlertIidMetricImagesHandler)
}

func getProjectsIdAlertManagementAlertsAlertIidMetricImagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	alert_iid := int32(request.GetInt("alert_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages(ctx, id, alert_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
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

	s.AddTool(tool, getProjectsIdIssuesIssueIidAwardEmojiHandler)
}

func getProjectsIdIssuesIssueIidAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := int32(request.GetInt("issue_iid", math.MinInt))
	params := parseGetProjectsIdIssuesIssueIidAwardEmoji(request)
	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidAwardEmoji(ctx, id, issue_iid, &params, authorizationHeader))
}

func parseGetProjectsIdIssuesIssueIidAwardEmoji(request mcp.CallToolRequest) client.GetApiV4ProjectsIdIssuesIssueIidAwardEmojiParams {
	params := client.GetApiV4ProjectsIdIssuesIssueIidAwardEmojiParams{}

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

func registerGetProjectsIdIssuesIssueIidAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidAwardEmojiAwardIdHandler)
}

func getProjectsIdIssuesIssueIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	issue_iid := int32(request.GetInt("issue_iid", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId(ctx, id, issue_iid, award_id, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_notes_note_id_award_emoji",
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
		mcp.WithNumber("issue_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiHandler)
}

func getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	issue_iid := int32(request.GetInt("issue_iid", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	params := parseGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(request)
	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(ctx, id, issue_iid, note_id, &params, authorizationHeader))
}

func parseGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(request mcp.CallToolRequest) client.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiParams {
	params := client.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiParams{}

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

func registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdHandler)
}

func getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	issue_iid := int32(request.GetInt("issue_iid", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(ctx, id, issue_iid, note_id, award_id, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
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

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidAwardEmojiHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmoji(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiParams{}

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

func registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(ctx, id, merge_request_iid, award_id, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_notes_note_id_award_emoji",
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
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(ctx, id, merge_request_iid, note_id, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiParams{}

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

func registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(ctx, id, merge_request_iid, note_id, award_id, authorizationHeader))
}

func registerGetProjectsIdSnippetsSnippetIdAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
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

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdAwardEmojiHandler)
}

func getProjectsIdSnippetsSnippetIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))
	params := parseGetProjectsIdSnippetsSnippetIdAwardEmoji(request)
	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmoji(ctx, id, snippet_id, &params, authorizationHeader))
}

func parseGetProjectsIdSnippetsSnippetIdAwardEmoji(request mcp.CallToolRequest) client.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiParams {
	params := client.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiParams{}

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

func registerGetProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdAwardEmojiAwardIdHandler)
}

func getProjectsIdSnippetsSnippetIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId(ctx, id, snippet_id, award_id, authorizationHeader))
}

func registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_notes_note_id_award_emoji",
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
		mcp.WithNumber("snippet_id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiHandler)
}

func getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	params := parseGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(request)
	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(ctx, id, snippet_id, note_id, &params, authorizationHeader))
}

func parseGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(request mcp.CallToolRequest) client.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiParams {
	params := client.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiParams{}

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

func registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithNumber("award_id",
			mcp.Description("ID of the emoji reaction."),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("note_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdHandler)
}

func getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))
	note_id := int32(request.GetInt("note_id", math.MinInt))
	award_id := int32(request.GetInt("award_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(ctx, id, snippet_id, note_id, award_id, authorizationHeader))
}

func registerGetProjectsIdBadges(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_badges",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user."),
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

	s.AddTool(tool, getProjectsIdBadgesHandler)
}

func getProjectsIdBadgesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdBadges(request)
	return toResult(c.GetApiV4ProjectsIdBadges(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdBadges(request mcp.CallToolRequest) client.GetApiV4ProjectsIdBadgesParams {
	params := client.GetApiV4ProjectsIdBadgesParams{}

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

func registerGetProjectsIdBadgesRender(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_badges_render",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user."),
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

	s.AddTool(tool, getProjectsIdBadgesRenderHandler)
}

func getProjectsIdBadgesRenderHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdBadgesRender(request)
	return toResult(c.GetApiV4ProjectsIdBadgesRender(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdBadgesRender(request mcp.CallToolRequest) client.GetApiV4ProjectsIdBadgesRenderParams {
	params := client.GetApiV4ProjectsIdBadgesRenderParams{}

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

func registerGetProjectsIdBadgesBadgeId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_badges_badge_id",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user."),
			mcp.Required(),
		),
		mcp.WithNumber("badge_id",
			mcp.Description("The badge ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdBadgesBadgeIdHandler)
}

func getProjectsIdBadgesBadgeIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	badge_id := int32(request.GetInt("badge_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdBadgesBadgeId(ctx, id, badge_id, authorizationHeader))
}

func registerGetProjectsIdRepositoryBranches(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_branches",
		mcp.WithDescription("Get a project repository branches"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of branches matching the search criteria"),
		),
		mcp.WithString("regex",
			mcp.Description("Return list of branches matching the regex"),
		),
		mcp.WithString("sort",
			mcp.Description("Return list of branches sorted by the given field"),

			mcp.Enum("name_asc", "updated_asc", "updated_desc"),
		),
		mcp.WithString("page_token",
			mcp.Description("Name of branch to start the pagination from"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryBranchesHandler)
}

func getProjectsIdRepositoryBranchesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryBranches(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryBranches(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryBranches(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryBranchesParams {
	params := client.GetApiV4ProjectsIdRepositoryBranchesParams{}

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

	regex := request.GetString("regex", "")
	if regex != "" {

		params.Regex = &regex
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	page_token := request.GetString("page_token", "")
	if page_token != "" {

		params.PageToken = &page_token
	}

	return params
}

func registerGetProjectsIdRepositoryBranchesBranch(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_branches_branch",
		mcp.WithDescription("Get a single repository branch"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("branch",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryBranchesBranchHandler)
}

func getProjectsIdRepositoryBranchesBranchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	branch := int32(request.GetInt("branch", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdRepositoryBranchesBranch(ctx, id, branch, authorizationHeader))
}

func registerGetProjectsIdJobsArtifactsRefNameDownload(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_jobs_artifacts_ref_name_download",
		mcp.WithDescription("This feature was introduced in GitLab 8.10"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("ref_name",
			mcp.Description("Branch or tag name in repository. `HEAD` or `SHA` references are not supported."),
			mcp.Required(),
		),
		mcp.WithString("job",
			mcp.Description("The name of the job."),
			mcp.Required(),
		),
		mcp.WithString("job_token",
			mcp.Description("To be used with triggers for multi-project pipelines, available only on Premium and Ultimate tiers."),
		),
	)

	s.AddTool(tool, getProjectsIdJobsArtifactsRefNameDownloadHandler)
}

func getProjectsIdJobsArtifactsRefNameDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	ref_name := request.GetString("ref_name", "")
	params := parseGetProjectsIdJobsArtifactsRefNameDownload(request)
	return toResult(c.GetApiV4ProjectsIdJobsArtifactsRefNameDownload(ctx, id, ref_name, &params, authorizationHeader))
}

func parseGetProjectsIdJobsArtifactsRefNameDownload(request mcp.CallToolRequest) client.GetApiV4ProjectsIdJobsArtifactsRefNameDownloadParams {
	params := client.GetApiV4ProjectsIdJobsArtifactsRefNameDownloadParams{}

	job := request.GetString("job", "")
	if job != "" {

		params.Job = job
	}

	job_token := request.GetString("job_token", "")
	if job_token != "" {

		params.JobToken = &job_token
	}

	return params
}

func registerGetProjectsIdJobsJobIdArtifacts(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_jobs_job_id_artifacts",
		mcp.WithDescription("This feature was introduced in GitLab 8.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("job_id",
			mcp.Description("The ID of a job"),
			mcp.Required(),
		),
		mcp.WithString("job_token",
			mcp.Description("To be used with triggers for multi-project pipelines, available only on Premium and Ultimate tiers."),
		),
	)

	s.AddTool(tool, getProjectsIdJobsJobIdArtifactsHandler)
}

func getProjectsIdJobsJobIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	job_id := int32(request.GetInt("job_id", math.MinInt))
	params := parseGetProjectsIdJobsJobIdArtifacts(request)
	return toResult(c.GetApiV4ProjectsIdJobsJobIdArtifacts(ctx, id, job_id, &params, authorizationHeader))
}

func parseGetProjectsIdJobsJobIdArtifacts(request mcp.CallToolRequest) client.GetApiV4ProjectsIdJobsJobIdArtifactsParams {
	params := client.GetApiV4ProjectsIdJobsJobIdArtifactsParams{}

	job_token := request.GetString("job_token", "")
	if job_token != "" {

		params.JobToken = &job_token
	}

	return params
}

func registerGetProjectsIdJobs(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_jobs",
		mcp.WithDescription("Get a projects jobs"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("scope",
			mcp.Description("The scope of builds to show (example: [pending, running])"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdJobsHandler)
}

func getProjectsIdJobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdJobs(request)
	return toResult(c.GetApiV4ProjectsIdJobs(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdJobs(request mcp.CallToolRequest) client.GetApiV4ProjectsIdJobsParams {
	params := client.GetApiV4ProjectsIdJobsParams{}

	scope := request.GetString("scope", "")
	if scope != "" {
		scope := strings.Split(scope, ",")
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

	return params
}

func registerGetProjectsIdJobsJobId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_jobs_job_id",
		mcp.WithDescription("Get a specific job of a project"),
		mcp.WithNumber("job_id",
			mcp.Description("The ID of a job (example: 88)"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdJobsJobIdHandler)
}

func getProjectsIdJobsJobIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	job_id := int32(request.GetInt("job_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdJobsJobId(ctx, id, job_id, authorizationHeader))
}

func registerGetProjectsIdJobsJobIdTrace(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_jobs_job_id_trace",
		mcp.WithDescription("Get a trace of a specific job of a project"),
		mcp.WithNumber("job_id",
			mcp.Description("The ID of a job (example: 88)"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdJobsJobIdTraceHandler)
}

func getProjectsIdJobsJobIdTraceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	job_id := int32(request.GetInt("job_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdJobsJobIdTrace(ctx, id, job_id, authorizationHeader))
}

func registerGetProjectsIdResourceGroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_resource_grps",
		mcp.WithDescription("Get all resource groups for a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdResourceGroupsHandler)
}

func getProjectsIdResourceGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdResourceGroups(request)
	return toResult(c.GetApiV4ProjectsIdResourceGroups(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdResourceGroups(request mcp.CallToolRequest) client.GetApiV4ProjectsIdResourceGroupsParams {
	params := client.GetApiV4ProjectsIdResourceGroupsParams{}

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

func registerGetProjectsIdResourceGroupsKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_resource_grps_key",
		mcp.WithDescription("Get a specific resource group"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithString("key",
			mcp.Description("The key of the resource group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdResourceGroupsKeyHandler)
}

func getProjectsIdResourceGroupsKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	key := request.GetString("key", "")

	return toResult(c.GetApiV4ProjectsIdResourceGroupsKey(ctx, id, key, authorizationHeader))
}

func registerGetProjectsIdResourceGroupsKeyUpcomingJobs(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_resource_grps_key_upcoming_jobs",
		mcp.WithDescription("List upcoming jobs for a specific resource group"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithString("key",
			mcp.Description("The key of the resource group"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdResourceGroupsKeyUpcomingJobsHandler)
}

func getProjectsIdResourceGroupsKeyUpcomingJobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	key := request.GetString("key", "")
	params := parseGetProjectsIdResourceGroupsKeyUpcomingJobs(request)
	return toResult(c.GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs(ctx, id, key, &params, authorizationHeader))
}

func parseGetProjectsIdResourceGroupsKeyUpcomingJobs(request mcp.CallToolRequest) client.GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsParams {
	params := client.GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsParams{}

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

func registerGetProjectsIdRunners(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_runners",
		mcp.WithDescription("List all runners available in the project, including from ancestor groups and any allowed shared runners."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
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

	s.AddTool(tool, getProjectsIdRunnersHandler)
}

func getProjectsIdRunnersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRunners(request)
	return toResult(c.GetApiV4ProjectsIdRunners(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRunners(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRunnersParams {
	params := client.GetApiV4ProjectsIdRunnersParams{}

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

func registerGetProjectsIdSecureFiles(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_secure_files",
		mcp.WithDescription("Get list of secure files in a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdSecureFilesHandler)
}

func getProjectsIdSecureFilesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdSecureFiles(request)
	return toResult(c.GetApiV4ProjectsIdSecureFiles(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdSecureFiles(request mcp.CallToolRequest) client.GetApiV4ProjectsIdSecureFilesParams {
	params := client.GetApiV4ProjectsIdSecureFilesParams{}

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

func registerGetProjectsIdSecureFilesSecureFileId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_secure_files_secure_file_id",
		mcp.WithDescription("Get the details of a specific secure file in a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("secure_file_id",
			mcp.Description("The ID of a secure file"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSecureFilesSecureFileIdHandler)
}

func getProjectsIdSecureFilesSecureFileIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	secure_file_id := int32(request.GetInt("secure_file_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSecureFilesSecureFileId(ctx, id, secure_file_id, authorizationHeader))
}

func registerGetProjectsIdSecureFilesSecureFileIdDownload(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_secure_files_secure_file_id_download",
		mcp.WithDescription("Download secure file"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("secure_file_id",
			mcp.Description("The ID of a secure file"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSecureFilesSecureFileIdDownloadHandler)
}

func getProjectsIdSecureFilesSecureFileIdDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	secure_file_id := int32(request.GetInt("secure_file_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSecureFilesSecureFileIdDownload(ctx, id, secure_file_id, authorizationHeader))
}

func registerGetProjectsIdPipelines(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("scope",
			mcp.Description("The scope of pipelines (example: pending)"),

			mcp.Enum("running", "pending", "finished", "branches", "tags"),
		),
		mcp.WithString("status",
			mcp.Description("The status of pipelines (example: pending)"),

			mcp.Enum("created", "waiting_for_resource", "preparing", "waiting_for_callback", "pending", "running", "success", "failed", "canceling", "canceled", "skipped", "manual", "scheduled"),
		),
		mcp.WithString("ref",
			mcp.Description("The ref of pipelines (example: develop)"),
		),
		mcp.WithString("sha",
			mcp.Description("The sha of pipelines (example: a91957a858320c0e17f3a0eca7cfacbff50ea29a)"),
		),
		mcp.WithBoolean("yaml_errors",
			mcp.Description("Returns pipelines with invalid configurations"),
		),
		mcp.WithString("username",
			mcp.Description("The username of the user who triggered pipelines (example: root)"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return pipelines updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ (example: 2015-12-24T15:51:21.880Z)"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return pipelines updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ (example: 2015-12-24T15:51:21.880Z)"),
		),
		mcp.WithString("created_before",
			mcp.Description("Return pipelines created before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ (example: 2015-12-24T15:51:21.880Z)"),
		),
		mcp.WithString("created_after",
			mcp.Description("Return pipelines created after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ (example: 2015-12-24T15:51:21.880Z)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Order pipelines (example: status) (default: id)"),

			mcp.Enum("id", "status", "ref", "updated_at", "user_id"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort pipelines (example: asc) (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("source",
			mcp.Description("null (example: push)"),

			mcp.Enum("unknown", "push", "web", "trigger", "schedule", "api", "external", "pipeline", "chat", "webide", "merge_request_event", "external_pull_request_event", "parent_pipeline", "ondemand_dast_scan", "ondemand_dast_validation", "security_orchestration_policy", "container_registry_push", "duo_workflow", "pipeline_execution_policy_schedule"),
		),
		mcp.WithString("name",
			mcp.Description("Filter pipelines by name (example: Build pipeline)"),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesHandler)
}

func getProjectsIdPipelinesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPipelines(request)
	return toResult(c.GetApiV4ProjectsIdPipelines(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPipelines(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPipelinesParams {
	params := client.GetApiV4ProjectsIdPipelinesParams{}

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

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	sha := request.GetString("sha", "")
	if sha != "" {

		params.Sha = &sha
	}

	yaml_errors := request.GetBool("yaml_errors", false)
	params.YamlErrors = &yaml_errors

	username := request.GetString("username", "")
	if username != "" {

		params.Username = &username
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

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
	}

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	source := request.GetString("source", "")
	if source != "" {

		params.Source = &source
	}

	name := request.GetString("name", "")
	if name != "" {

		params.Name = &name
	}

	return params
}

func registerGetProjectsIdPipelinesLatest(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_latest",
		mcp.WithDescription("This feature was introduced in GitLab 12.3"),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("Branch ref of pipeline. Uses project default branch if not specified. (example: develop)"),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesLatestHandler)
}

func getProjectsIdPipelinesLatestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPipelinesLatest(request)
	return toResult(c.GetApiV4ProjectsIdPipelinesLatest(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPipelinesLatest(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPipelinesLatestParams {
	params := client.GetApiV4ProjectsIdPipelinesLatestParams{}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	return params
}

func registerGetProjectsIdPipelinesPipelineId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.11"),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("The pipeline ID (example: 18)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesPipelineIdHandler)
}

func getProjectsIdPipelinesPipelineIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_id := int32(request.GetInt("pipeline_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineId(ctx, id, pipeline_id, authorizationHeader))
}

func registerGetProjectsIdPipelinesPipelineIdJobs(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_jobs",
		mcp.WithDescription("Get pipeline jobs"),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("The pipeline ID (example: 18)"),
			mcp.Required(),
		),
		mcp.WithBoolean("include_retried",
			mcp.Description("Includes retried jobs (default: false)"),
		),
		mcp.WithString("scope",
			mcp.Description("The scope of builds to show (example: [pending, running])"),

			mcp.Enum("created", "waiting_for_resource", "preparing", "waiting_for_callback", "pending", "running", "success", "failed", "canceling", "canceled", "skipped", "manual", "scheduled"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesPipelineIdJobsHandler)
}

func getProjectsIdPipelinesPipelineIdJobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_id := int32(request.GetInt("pipeline_id", math.MinInt))
	params := parseGetProjectsIdPipelinesPipelineIdJobs(request)
	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdJobs(ctx, id, pipeline_id, &params, authorizationHeader))
}

func parseGetProjectsIdPipelinesPipelineIdJobs(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPipelinesPipelineIdJobsParams {
	params := client.GetApiV4ProjectsIdPipelinesPipelineIdJobsParams{}

	include_retried := request.GetBool("include_retried", false)
	params.IncludeRetried = &include_retried

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

	return params
}

func registerGetProjectsIdPipelinesPipelineIdBridges(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_bridges",
		mcp.WithDescription("Get pipeline bridge jobs"),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("The pipeline ID (example: 18)"),
			mcp.Required(),
		),
		mcp.WithString("scope",
			mcp.Description("The scope of builds to show (example: [pending, running])"),

			mcp.Enum("created", "waiting_for_resource", "preparing", "waiting_for_callback", "pending", "running", "success", "failed", "canceling", "canceled", "skipped", "manual", "scheduled"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesPipelineIdBridgesHandler)
}

func getProjectsIdPipelinesPipelineIdBridgesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_id := int32(request.GetInt("pipeline_id", math.MinInt))
	params := parseGetProjectsIdPipelinesPipelineIdBridges(request)
	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdBridges(ctx, id, pipeline_id, &params, authorizationHeader))
}

func parseGetProjectsIdPipelinesPipelineIdBridges(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPipelinesPipelineIdBridgesParams {
	params := client.GetApiV4ProjectsIdPipelinesPipelineIdBridgesParams{}

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

	return params
}

func registerGetProjectsIdPipelinesPipelineIdVariables(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_variables",
		mcp.WithDescription("This feature was introduced in GitLab 11.11"),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("The pipeline ID (example: 18)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesPipelineIdVariablesHandler)
}

func getProjectsIdPipelinesPipelineIdVariablesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_id := int32(request.GetInt("pipeline_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdVariables(ctx, id, pipeline_id, authorizationHeader))
}

func registerGetProjectsIdPipelinesPipelineIdTestReport(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_test_report",
		mcp.WithDescription("This feature was introduced in GitLab 13.0."),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("The pipeline ID (example: 18)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesPipelineIdTestReportHandler)
}

func getProjectsIdPipelinesPipelineIdTestReportHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_id := int32(request.GetInt("pipeline_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdTestReport(ctx, id, pipeline_id, authorizationHeader))
}

func registerGetProjectsIdPipelinesPipelineIdTestReportSummary(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_test_report_summary",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithString("id",
			mcp.Description("The project ID or URL-encoded path (example: 11)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("The pipeline ID (example: 18)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPipelinesPipelineIdTestReportSummaryHandler)
}

func getProjectsIdPipelinesPipelineIdTestReportSummaryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_id := int32(request.GetInt("pipeline_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary(ctx, id, pipeline_id, authorizationHeader))
}

func registerGetProjectsIdPipelineSchedules(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pipeline_schedules",
		mcp.WithDescription("Get all pipeline schedules"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 18)"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("scope",
			mcp.Description("The scope of pipeline schedules (example: active)"),

			mcp.Enum("active", "inactive"),
		),
	)

	s.AddTool(tool, getProjectsIdPipelineSchedulesHandler)
}

func getProjectsIdPipelineSchedulesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPipelineSchedules(request)
	return toResult(c.GetApiV4ProjectsIdPipelineSchedules(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPipelineSchedules(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPipelineSchedulesParams {
	params := client.GetApiV4ProjectsIdPipelineSchedulesParams{}

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

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	return params
}

func registerGetProjectsIdPipelineSchedulesPipelineScheduleId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pipeline_schedules_pipeline_schedule_id",
		mcp.WithDescription("Get a single pipeline schedule"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 18)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_schedule_id",
			mcp.Description("The pipeline schedule id (example: 13)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPipelineSchedulesPipelineScheduleIdHandler)
}

func getProjectsIdPipelineSchedulesPipelineScheduleIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_schedule_id := int32(request.GetInt("pipeline_schedule_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx, id, pipeline_schedule_id, authorizationHeader))
}

func registerGetProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pipeline_schedules_pipeline_schedule_id_pls",
		mcp.WithDescription("Get all pipelines triggered from a pipeline schedule"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 18)"),
			mcp.Required(),
		),
		mcp.WithNumber("pipeline_schedule_id",
			mcp.Description("The pipeline schedule ID (example: 13)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesHandler)
}

func getProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	pipeline_schedule_id := int32(request.GetInt("pipeline_schedule_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(ctx, id, pipeline_schedule_id, authorizationHeader))
}

func registerGetProjectsIdTriggers(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_triggers",
		mcp.WithDescription("Get trigger tokens list"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 18)"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdTriggersHandler)
}

func getProjectsIdTriggersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdTriggers(request)
	return toResult(c.GetApiV4ProjectsIdTriggers(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdTriggers(request mcp.CallToolRequest) client.GetApiV4ProjectsIdTriggersParams {
	params := client.GetApiV4ProjectsIdTriggersParams{}

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

func registerGetProjectsIdTriggersTriggerId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_triggers_trigger_id",
		mcp.WithDescription("Get specific trigger token of a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 18)"),
			mcp.Required(),
		),
		mcp.WithNumber("trigger_id",
			mcp.Description("The trigger token ID (example: 10)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdTriggersTriggerIdHandler)
}

func getProjectsIdTriggersTriggerIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	trigger_id := int32(request.GetInt("trigger_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdTriggersTriggerId(ctx, id, trigger_id, authorizationHeader))
}

func registerGetProjectsIdVariables(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_variables",
		mcp.WithDescription("Get project variables"),
		mcp.WithString("id",
			mcp.Description("The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdVariablesHandler)
}

func getProjectsIdVariablesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdVariables(request)
	return toResult(c.GetApiV4ProjectsIdVariables(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdVariables(request mcp.CallToolRequest) client.GetApiV4ProjectsIdVariablesParams {
	params := client.GetApiV4ProjectsIdVariablesParams{}

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

func registerGetProjectsIdVariablesKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_variables_key",
		mcp.WithDescription("Get the details of a single variable from a project"),
		mcp.WithString("id",
			mcp.Description("The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithString("key",
			mcp.Description("The key of a variable"),
			mcp.Required(),
		),
		mcp.WithString("filter[environment_scope]",
			mcp.Description("The environment scope of a variable"),
		),
	)

	s.AddTool(tool, getProjectsIdVariablesKeyHandler)
}

func getProjectsIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	key := request.GetString("key", "")
	params := parseGetProjectsIdVariablesKey(request)
	return toResult(c.GetApiV4ProjectsIdVariablesKey(ctx, id, key, &params, authorizationHeader))
}

func parseGetProjectsIdVariablesKey(request mcp.CallToolRequest) client.GetApiV4ProjectsIdVariablesKeyParams {
	params := client.GetApiV4ProjectsIdVariablesKeyParams{}

	filter_environment_scope_ := request.GetString("filter[environment_scope]", "")
	if filter_environment_scope_ != "" {

		params.FilterEnvironmentScope = &filter_environment_scope_
	}

	return params
}

func registerGetProjectsIdClusterAgentsAgentIdTokens(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_cluster_agents_agent_id_tokens",
		mcp.WithDescription("This feature was introduced in GitLab 15.0. Returns a list of tokens for an agent."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("agent_id",
			mcp.Description("The ID of an agent"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdClusterAgentsAgentIdTokensHandler)
}

func getProjectsIdClusterAgentsAgentIdTokensHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	agent_id := int32(request.GetInt("agent_id", math.MinInt))
	params := parseGetProjectsIdClusterAgentsAgentIdTokens(request)
	return toResult(c.GetApiV4ProjectsIdClusterAgentsAgentIdTokens(ctx, id, agent_id, &params, authorizationHeader))
}

func parseGetProjectsIdClusterAgentsAgentIdTokens(request mcp.CallToolRequest) client.GetApiV4ProjectsIdClusterAgentsAgentIdTokensParams {
	params := client.GetApiV4ProjectsIdClusterAgentsAgentIdTokensParams{}

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

func registerGetProjectsIdClusterAgentsAgentIdTokensTokenId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_cluster_agents_agent_id_tokens_token_id",
		mcp.WithDescription("This feature was introduced in GitLab 15.0. Gets a single agent token."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("agent_id",
			mcp.Description("The ID of an agent"),
			mcp.Required(),
		),
		mcp.WithNumber("token_id",
			mcp.Description("The ID of the agent token"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdClusterAgentsAgentIdTokensTokenIdHandler)
}

func getProjectsIdClusterAgentsAgentIdTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	agent_id := int32(request.GetInt("agent_id", math.MinInt))
	token_id := int32(request.GetInt("token_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx, id, agent_id, token_id, authorizationHeader))
}

func registerGetProjectsIdClusterAgents(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_cluster_agents",
		mcp.WithDescription("This feature was introduced in GitLab 14.10. Returns the list of agents registered for the project."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdClusterAgentsHandler)
}

func getProjectsIdClusterAgentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdClusterAgents(request)
	return toResult(c.GetApiV4ProjectsIdClusterAgents(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdClusterAgents(request mcp.CallToolRequest) client.GetApiV4ProjectsIdClusterAgentsParams {
	params := client.GetApiV4ProjectsIdClusterAgentsParams{}

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

func registerGetProjectsIdClusterAgentsAgentId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_cluster_agents_agent_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.10. Gets a single agent details."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("agent_id",
			mcp.Description("The ID of an agent"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdClusterAgentsAgentIdHandler)
}

func getProjectsIdClusterAgentsAgentIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	agent_id := int32(request.GetInt("agent_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdClusterAgentsAgentId(ctx, id, agent_id, authorizationHeader))
}

func registerGetProjectsIdPackagesCargoConfigJson(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_cargo_config_json",
		mcp.WithDescription("This will be used by cargo for further requests"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesCargoConfigJsonHandler)
}

func getProjectsIdPackagesCargoConfigJsonHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesCargoConfigJson(ctx, id, authorizationHeader))
}

func registerGetProjectsIdRepositoryCommits(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits",
		mcp.WithDescription("Get a project repository commits"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("ref_name",
			mcp.Description("The name of a repository branch or tag, if not given the default branch is used (example: v1.1.0)"),
		),
		mcp.WithString("since",
			mcp.Description("Only commits after or on this date will be returned (example: 2021-09-20T11:50:22.001)"),
		),
		mcp.WithString("until",
			mcp.Description("Only commits before or on this date will be returned (example: 2021-09-20T11:50:22.001)"),
		),
		mcp.WithString("path",
			mcp.Description("The file path (example: README.md)"),
		),
		mcp.WithString("author",
			mcp.Description("Search commits by commit author (example: John Smith)"),
		),
		mcp.WithBoolean("all",
			mcp.Description("Every commit will be returned"),
		),
		mcp.WithBoolean("with_stats",
			mcp.Description("Stats about each commit will be added to the response"),
		),
		mcp.WithBoolean("first_parent",
			mcp.Description("Only include the first parent of merges"),
		),
		mcp.WithString("order",
			mcp.Description("List commits in order (default: default)"),

			mcp.Enum("default", "topo"),
		),
		mcp.WithBoolean("trailers",
			mcp.Description("Parse and include Git trailers for every commit (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsHandler)
}

func getProjectsIdRepositoryCommitsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryCommits(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommits(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommits(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsParams{}

	ref_name := request.GetString("ref_name", "")
	if ref_name != "" {

		params.RefName = &ref_name
	}

	since := request.GetString("since", "")
	if since != "" {
		since, _ := time.Parse(time.RFC3339, since)
		params.Since = &since
	}

	until := request.GetString("until", "")
	if until != "" {
		until, _ := time.Parse(time.RFC3339, until)
		params.Until = &until
	}

	path := request.GetString("path", "")
	if path != "" {

		params.Path = &path
	}

	author := request.GetString("author", "")
	if author != "" {

		params.Author = &author
	}

	all := request.GetBool("all", false)
	params.All = &all

	with_stats := request.GetBool("with_stats", false)
	params.WithStats = &with_stats

	first_parent := request.GetBool("first_parent", false)
	params.FirstParent = &first_parent

	order := request.GetString("order", "")
	if order != "" {

		params.Order = &order
	}

	trailers := request.GetBool("trailers", false)
	params.Trailers = &trailers

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

func registerGetProjectsIdRepositoryCommitsSha(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha",
		mcp.WithDescription("Get a specific commit of a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("A commit sha, or the name of a branch or tag"),
			mcp.Required(),
		),
		mcp.WithBoolean("stats",
			mcp.Description("Include commit stats (default: true)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaHandler)
}

func getProjectsIdRepositoryCommitsShaHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsSha(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsSha(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsSha(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaParams{}

	stats := request.GetBool("stats", true)
	params.Stats = &stats

	return params
}

func registerGetProjectsIdRepositoryCommitsShaDiff(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_diff",
		mcp.WithDescription("Get the diff for a specific commit of a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("A commit sha, or the name of a branch or tag"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("unidiff",
			mcp.Description("A diff in a Unified diff format (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaDiffHandler)
}

func getProjectsIdRepositoryCommitsShaDiffHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsShaDiff(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaDiff(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsShaDiff(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaDiffParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaDiffParams{}

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

	unidiff := request.GetBool("unidiff", false)
	params.Unidiff = &unidiff

	return params
}

func registerGetProjectsIdRepositoryCommitsShaComments(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_comments",
		mcp.WithDescription("Get a commit's comments"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("sha",
			mcp.Description("A commit sha, or the name of a branch or tag"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaCommentsHandler)
}

func getProjectsIdRepositoryCommitsShaCommentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsShaComments(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaComments(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsShaComments(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaCommentsParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaCommentsParams{}

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

func registerGetProjectsIdRepositoryCommitsShaSequence(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_sequence",
		mcp.WithDescription("Get the sequence count of a commit SHA"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("A commit SHA"),
			mcp.Required(),
		),
		mcp.WithBoolean("first_parent",
			mcp.Description("Only include the first parent of merges (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaSequenceHandler)
}

func getProjectsIdRepositoryCommitsShaSequenceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsShaSequence(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaSequence(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsShaSequence(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaSequenceParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaSequenceParams{}

	first_parent := request.GetBool("first_parent", false)
	params.FirstParent = &first_parent

	return params
}

func registerGetProjectsIdRepositoryCommitsShaRefs(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_refs",
		mcp.WithDescription("This feature was introduced in GitLab 10.6"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("A commit sha"),
			mcp.Required(),
		),
		mcp.WithString("type",
			mcp.Description("Scope (default: all)"),

			mcp.Enum("branch", "tag", "all"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaRefsHandler)
}

func getProjectsIdRepositoryCommitsShaRefsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsShaRefs(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaRefs(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsShaRefs(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaRefsParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaRefsParams{}

	_type := request.GetString("type", "")
	if _type != "" {

		params.Type = &_type
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

func registerGetProjectsIdRepositoryCommitsShaMergeRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_mrs",
		mcp.WithDescription("Get Merge Requests associated with a commit"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("A commit sha, or the name of a branch or tag on which to find Merge Requests"),
			mcp.Required(),
		),
		mcp.WithString("state",
			mcp.Description("Filter merge-requests by state (example: merged)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaMergeRequestsHandler)
}

func getProjectsIdRepositoryCommitsShaMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsShaMergeRequests(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsShaMergeRequests(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsParams{}

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

func registerGetProjectsIdRepositoryCommitsShaSignature(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_signature",
		mcp.WithDescription("Get a commit's signature"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("A commit sha, or the name of a branch or tag"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaSignatureHandler)
}

func getProjectsIdRepositoryCommitsShaSignatureHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaSignature(ctx, id, sha, authorizationHeader))
}

func registerGetProjectsIdRepositoryCommitsShaStatuses(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_statuses",
		mcp.WithDescription("Get a commit's statuses"),
		mcp.WithString("id",
			mcp.Description("ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("Hash of the commit. (example: 18f3e63d05582537db6d183d9d557be09e1f90c8)"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("Name of the branch or tag. Default is the default branch. (example: develop)"),
		),
		mcp.WithString("stage",
			mcp.Description("Filter statuses by build stage. (example: test)"),
		),
		mcp.WithString("name",
			mcp.Description("Filter statuses by job name. (example: bundler:audit)"),
		),
		mcp.WithNumber("pipeline_id",
			mcp.Description("Filter statuses by pipeline ID. (example: 1234)"),
		),
		mcp.WithBoolean("all",
			mcp.Description("Include all statuses instead of latest only. Default is `false`. (default: false)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Values for sorting statuses. Valid values are `id` and `pipeline_id`. Default is `id`. (default: id)"),

			mcp.Enum("id", "pipeline_id"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort statuses in ascending or descending order. Valid values are `asc` and `desc`. Default is `asc`. (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCommitsShaStatusesHandler)
}

func getProjectsIdRepositoryCommitsShaStatusesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")
	params := parseGetProjectsIdRepositoryCommitsShaStatuses(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaStatuses(ctx, id, sha, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCommitsShaStatuses(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCommitsShaStatusesParams {
	params := client.GetApiV4ProjectsIdRepositoryCommitsShaStatusesParams{}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	stage := request.GetString("stage", "")
	if stage != "" {

		params.Stage = &stage
	}

	name := request.GetString("name", "")
	if name != "" {

		params.Name = &name
	}

	pipeline_id := request.GetInt("pipeline_id", math.MinInt)
	if pipeline_id != math.MinInt {
		pipeline_id := int32(pipeline_id)
		params.PipelineId = &pipeline_id
	}

	all := request.GetBool("all", false)
	params.All = &all

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

	return params
}

func registerGetProjectsIdPackagesConanV1UsersAuthenticate(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_users_authenticate",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1UsersAuthenticateHandler)
}

func getProjectsIdPackagesConanV1UsersAuthenticateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1UsersCheckCredentials(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_users_check_credentials",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1UsersCheckCredentialsHandler)
}

func getProjectsIdPackagesConanV1UsersCheckCredentialsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansSearch(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_search",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("q",
			mcp.Description("Search query (example: Hello*)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansSearchHandler)
}

func getProjectsIdPackagesConanV1ConansSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPackagesConanV1ConansSearch(request)
	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansSearch(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesConanV1ConansSearch(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesConanV1ConansSearchParams {
	params := client.GetApiV4ProjectsIdPackagesConanV1ConansSearchParams{}

	q := request.GetString("q", "")
	if q != "" {

		params.Q = q
	}

	return params
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1Ping(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_ping",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1PingHandler)
}

func getProjectsIdPackagesConanV1PingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1Ping(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx, id, package_name, package_version, package_username, package_channel, conan_package_reference, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx, id, package_name, package_version, package_username, package_channel, conan_package_reference, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx, id, package_name, package_version, package_username, package_channel, conan_package_reference, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler)
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_export_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Conan Recipe Revision (example: 0)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name (example: conanfile.py)"),
			mcp.Required(),
			mcp.Enum("conanfile.py", "conanmanifest.txt", "conan_sources.tgz", "conan_export.tgz", "conaninfo.txt", "conan_package.tgz"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler)
}

func getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_package_conan_package_reference_package_revision_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Conan Recipe Revision (example: 0)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan Package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
		mcp.WithString("package_revision",
			mcp.Description("Conan Package Revision (example: 0)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name (example: conaninfo.txt)"),
			mcp.Required(),
			mcp.Enum("conanfile.py", "conanmanifest.txt", "conan_sources.tgz", "conan_export.tgz", "conaninfo.txt", "conan_package.tgz"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler)
}

func getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	conan_package_reference := request.GetString("conan_package_reference", "")
	package_revision := request.GetString("package_revision", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, conan_package_reference, package_revision, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2UsersAuthenticate(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_users_authenticate",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2UsersAuthenticateHandler)
}

func getProjectsIdPackagesConanV2UsersAuthenticateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2UsersAuthenticate(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2UsersCheckCredentials(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_users_check_credentials",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2UsersCheckCredentialsHandler)
}

func getProjectsIdPackagesConanV2UsersCheckCredentialsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentials(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansSearch(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_search",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("q",
			mcp.Description("Search query (example: Hello*)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansSearchHandler)
}

func getProjectsIdPackagesConanV2ConansSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPackagesConanV2ConansSearch(request)
	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansSearch(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesConanV2ConansSearch(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesConanV2ConansSearchParams {
	params := client.GetApiV4ProjectsIdPackagesConanV2ConansSearchParams{}

	q := request.GetString("q", "")
	if q != "" {

		params.Q = q
	}

	return params
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_latest",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions(ctx, id, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_files",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_files_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 17.8"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name (example: conanfile.py)"),
			mcp.Required(),
			mcp.Enum("conanfile.py", "conanmanifest.txt", "conan_sources.tgz", "conan_export.tgz", "conaninfo.txt", "conan_package.tgz"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.1"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_latest",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Package reference (example: 5ab84d6acfe1f23c4fae0ab88f26e3a396351ac9)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, conan_package_reference, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Package reference (example: 5ab84d6acfe1f23c4fae0ab88f26e3a396351ac9)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, conan_package_reference, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions_package_revision_files",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Package reference (example: 5ab84d6acfe1f23c4fae0ab88f26e3a396351ac9)"),
			mcp.Required(),
		),
		mcp.WithString("package_revision",
			mcp.Description("Package revision (example: 3bdd2d8c8e76c876ebd1ac0469a4e72c)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	conan_package_reference := request.GetString("conan_package_reference", "")
	package_revision := request.GetString("package_revision", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, conan_package_reference, package_revision, authorizationHeader))
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions_package_revision_files_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Recipe revision (example: df28fd816be3a119de5ce4d374436b25)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Package reference (example: 5ab84d6acfe1f23c4fae0ab88f26e3a396351ac9)"),
			mcp.Required(),
		),
		mcp.WithString("package_revision",
			mcp.Description("Package revision (example: 3bdd2d8c8e76c876ebd1ac0469a4e72c)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name (example: conaninfo.txt)"),
			mcp.Required(),
			mcp.Enum("conanfile.py", "conanmanifest.txt", "conan_sources.tgz", "conan_export.tgz", "conaninfo.txt", "conan_package.tgz"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameHandler)
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	conan_package_reference := request.GetString("conan_package_reference", "")
	package_revision := request.GetString("package_revision", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(ctx, id, package_name, package_version, package_username, package_channel, recipe_revision, conan_package_reference, package_revision, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_debian_pool_distribution_letter_package_name_package_version_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
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

	s.AddTool(tool, getProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameHandler)
}

func getProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	distribution := request.GetString("distribution", "")
	letter := request.GetString("letter", "")
	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(ctx, id, distribution, letter, package_name, package_version, file_name, authorizationHeader))
}

func registerGetProjectsIdDeployKeys(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deploy_keys",
		mcp.WithDescription("Get a list of a project's deploy keys."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdDeployKeysHandler)
}

func getProjectsIdDeployKeysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdDeployKeys(request)
	return toResult(c.GetApiV4ProjectsIdDeployKeys(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdDeployKeys(request mcp.CallToolRequest) client.GetApiV4ProjectsIdDeployKeysParams {
	params := client.GetApiV4ProjectsIdDeployKeysParams{}

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

func registerGetProjectsIdDeployKeysKeyId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deploy_keys_key_id",
		mcp.WithDescription("Get a single key."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("key_id",
			mcp.Description("The ID of the deploy key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdDeployKeysKeyIdHandler)
}

func getProjectsIdDeployKeysKeyIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	key_id := int32(request.GetInt("key_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdDeployKeysKeyId(ctx, id, key_id, authorizationHeader))
}

func registerGetProjectsIdDeployTokens(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deploy_tokens",
		mcp.WithDescription("Get a list of a project's deploy tokens. This feature was introduced in GitLab 12.9."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
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

	s.AddTool(tool, getProjectsIdDeployTokensHandler)
}

func getProjectsIdDeployTokensHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdDeployTokens(request)
	return toResult(c.GetApiV4ProjectsIdDeployTokens(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdDeployTokens(request mcp.CallToolRequest) client.GetApiV4ProjectsIdDeployTokensParams {
	params := client.GetApiV4ProjectsIdDeployTokensParams{}

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

func registerGetProjectsIdDeployTokensTokenId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deploy_tokens_token_id",
		mcp.WithDescription("Get a single project's deploy token by ID. This feature was introduced in GitLab 14.9."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("token_id",
			mcp.Description("The ID of the deploy token"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdDeployTokensTokenIdHandler)
}

func getProjectsIdDeployTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	token_id := int32(request.GetInt("token_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdDeployTokensTokenId(ctx, id, token_id, authorizationHeader))
}

func registerGetProjectsIdDeployments(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deployments",
		mcp.WithDescription("Get a list of deployments in a project. This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return deployments ordered by either one of `id`, `iid`, `created_at`, `updated_at` or `ref` fields. Default is `id` (default: id)"),

			mcp.Enum("id", "iid", "created_at", "updated_at", "finished_at"),
		),
		mcp.WithString("sort",
			mcp.Description("Return deployments sorted in `asc` or `desc` order. Default is `asc` (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return deployments updated after the specified date. Expected in ISO 8601 format (`2019-03-15T08:00:00Z`)"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return deployments updated before the specified date. Expected in ISO 8601 format (`2019-03-15T08:00:00Z`)"),
		),
		mcp.WithString("finished_after",
			mcp.Description("Return deployments finished after the specified date. Expected in ISO 8601 format (`2019-03-15T08:00:00Z`)"),
		),
		mcp.WithString("finished_before",
			mcp.Description("Return deployments finished before the specified date. Expected in ISO 8601 format (`2019-03-15T08:00:00Z`)"),
		),
		mcp.WithString("environment",
			mcp.Description("The name of the environment to filter deployments by"),
		),
		mcp.WithString("status",
			mcp.Description("The status to filter deployments by. One of `created`, `running`, `success`, `failed`, `canceled`, or `blocked`"),

			mcp.Enum("created", "running", "success", "failed", "canceled", "skipped", "blocked"),
		),
	)

	s.AddTool(tool, getProjectsIdDeploymentsHandler)
}

func getProjectsIdDeploymentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdDeployments(request)
	return toResult(c.GetApiV4ProjectsIdDeployments(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdDeployments(request mcp.CallToolRequest) client.GetApiV4ProjectsIdDeploymentsParams {
	params := client.GetApiV4ProjectsIdDeploymentsParams{}

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

	finished_after := request.GetString("finished_after", "")
	if finished_after != "" {
		finished_after, _ := time.Parse(time.RFC3339, finished_after)
		params.FinishedAfter = &finished_after
	}

	finished_before := request.GetString("finished_before", "")
	if finished_before != "" {
		finished_before, _ := time.Parse(time.RFC3339, finished_before)
		params.FinishedBefore = &finished_before
	}

	environment := request.GetString("environment", "")
	if environment != "" {

		params.Environment = &environment
	}

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	return params
}

func registerGetProjectsIdDeploymentsDeploymentId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deployments_deployment_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("deployment_id",
			mcp.Description("The ID of the deployment"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdDeploymentsDeploymentIdHandler)
}

func getProjectsIdDeploymentsDeploymentIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	deployment_id := int32(request.GetInt("deployment_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdDeploymentsDeploymentId(ctx, id, deployment_id, authorizationHeader))
}

func registerGetProjectsIdDeploymentsDeploymentIdMergeRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_deployments_deployment_id_mrs",
		mcp.WithDescription("Retrieves the list of merge requests shipped with a given deployment. This feature was introduced in GitLab 12.7."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("deployment_id",
			mcp.Description("The ID of the deployment"),
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
	)

	s.AddTool(tool, getProjectsIdDeploymentsDeploymentIdMergeRequestsHandler)
}

func getProjectsIdDeploymentsDeploymentIdMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	deployment_id := int32(request.GetInt("deployment_id", math.MinInt))
	params := parseGetProjectsIdDeploymentsDeploymentIdMergeRequests(request)
	return toResult(c.GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests(ctx, id, deployment_id, &params, authorizationHeader))
}

func parseGetProjectsIdDeploymentsDeploymentIdMergeRequests(request mcp.CallToolRequest) client.GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsParams {
	params := client.GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsParams{}

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

	return params
}

func registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotes(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_draft_notes",
		mcp.WithDescription("Get a list of merge request draft notes"),
		mcp.WithString("id",
			mcp.Description("The ID of a project"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The ID of a merge request"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidDraftNotesHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidDraftNotesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_draft_notes_draft_note_id",
		mcp.WithDescription("Get a single draft note"),
		mcp.WithString("id",
			mcp.Description("The ID of a project"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The ID of a merge request"),
			mcp.Required(),
		),
		mcp.WithNumber("draft_note_id",
			mcp.Description("The ID of a draft note"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	draft_note_id := int32(request.GetInt("draft_note_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx, id, merge_request_iid, draft_note_id, authorizationHeader))
}

func registerGetProjectsIdEnvironments(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_environments",
		mcp.WithDescription("Get all environments for a given project. This feature was introduced in GitLab 8.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("name",
			mcp.Description("Return the environment with this name. Mutually exclusive with search"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of environments matching the search criteria. Mutually exclusive with name. Must be at least 3 characters."),
		),
		mcp.WithString("states",
			mcp.Description("List all environments that match a specific state. Accepted values: `available`, `stopping`, or `stopped`. If no state value given, returns all environments"),

			mcp.Enum("stopped", "stopping", "available"),
		),
	)

	s.AddTool(tool, getProjectsIdEnvironmentsHandler)
}

func getProjectsIdEnvironmentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdEnvironments(request)
	return toResult(c.GetApiV4ProjectsIdEnvironments(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdEnvironments(request mcp.CallToolRequest) client.GetApiV4ProjectsIdEnvironmentsParams {
	params := client.GetApiV4ProjectsIdEnvironmentsParams{}

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

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	states := request.GetString("states", "")
	if states != "" {

		params.States = &states
	}

	return params
}

func registerGetProjectsIdEnvironmentsEnvironmentId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_environments_environment_id",
		mcp.WithDescription("Get a specific environment"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("environment_id",
			mcp.Description("The ID of the environment"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdEnvironmentsEnvironmentIdHandler)
}

func getProjectsIdEnvironmentsEnvironmentIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	environment_id := int32(request.GetInt("environment_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdEnvironmentsEnvironmentId(ctx, id, environment_id, authorizationHeader))
}

func registerGetProjectsIdErrorTrackingClientKeys(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_error_tracking_client_keys",
		mcp.WithDescription("List all client keys. This feature was introduced in GitLab 14.3."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdErrorTrackingClientKeysHandler)
}

func getProjectsIdErrorTrackingClientKeysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdErrorTrackingClientKeys(ctx, id, authorizationHeader))
}

func registerGetProjectsIdErrorTrackingSettings(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_error_tracking_settings",
		mcp.WithDescription("Get error tracking settings for the project. This feature was introduced in GitLab 12.7."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdErrorTrackingSettingsHandler)
}

func getProjectsIdErrorTrackingSettingsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdErrorTrackingSettings(ctx, id, authorizationHeader))
}

func registerGetProjectsIdFeatureFlags(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_feature_flags",
		mcp.WithDescription("Gets all feature flags of the requested project. This feature was introduced in GitLab 12.5."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("scope",
			mcp.Description("The scope of feature flags, one of: `enabled`, `disabled`"),

			mcp.Enum("enabled", "disabled"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdFeatureFlagsHandler)
}

func getProjectsIdFeatureFlagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdFeatureFlags(request)
	return toResult(c.GetApiV4ProjectsIdFeatureFlags(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdFeatureFlags(request mcp.CallToolRequest) client.GetApiV4ProjectsIdFeatureFlagsParams {
	params := client.GetApiV4ProjectsIdFeatureFlagsParams{}

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

	return params
}

func registerGetProjectsIdFeatureFlagsFeatureFlagName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_feature_flags_feature_flag_name",
		mcp.WithDescription("Gets a single feature flag. This feature was introduced in GitLab 12.5."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("feature_flag_name",
			mcp.Description("The name of the feature flag"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdFeatureFlagsFeatureFlagNameHandler)
}

func getProjectsIdFeatureFlagsFeatureFlagNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	feature_flag_name := request.GetString("feature_flag_name", "")

	return toResult(c.GetApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx, id, feature_flag_name, authorizationHeader))
}

func registerGetProjectsIdFeatureFlagsUserLists(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_feature_flags_user_lists",
		mcp.WithDescription("Gets all feature flag user lists for the requested project. This feature was introduced in GitLab 12.10."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return user lists matching the search criteria"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdFeatureFlagsUserListsHandler)
}

func getProjectsIdFeatureFlagsUserListsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdFeatureFlagsUserLists(request)
	return toResult(c.GetApiV4ProjectsIdFeatureFlagsUserLists(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdFeatureFlagsUserLists(request mcp.CallToolRequest) client.GetApiV4ProjectsIdFeatureFlagsUserListsParams {
	params := client.GetApiV4ProjectsIdFeatureFlagsUserListsParams{}

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

func registerGetProjectsIdFeatureFlagsUserListsIid(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_feature_flags_user_lists_iid",
		mcp.WithDescription("Gets a feature flag user list. This feature was introduced in GitLab 12.10."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("iid",
			mcp.Description("The internal ID of the project's feature flag user list"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdFeatureFlagsUserListsIidHandler)
}

func getProjectsIdFeatureFlagsUserListsIidHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	iid := request.GetString("iid", "")

	return toResult(c.GetApiV4ProjectsIdFeatureFlagsUserListsIid(ctx, id, iid, authorizationHeader))
}

func registerGetProjectsIdRepositoryFilesFilePathBlame(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_files_file_path_blame",
		mcp.WithDescription("Get blame file from the repository"),
		mcp.WithString("id",
			mcp.Description("The project ID (example: gitlab-org/gitlab)"),
			mcp.Required(),
		),
		mcp.WithString("file_path",
			mcp.Description("The url encoded path to the file. (example: lib%2Fclass%2Erb)"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("The name of branch, tag or commit (example: main)"),
			mcp.Required(),
		),
		mcp.WithNumber("range[start]",
			mcp.Description("The first line of the range to blame"),
			mcp.Required(),
		),
		mcp.WithNumber("range[end]",
			mcp.Description("The last line of the range to blame"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryFilesFilePathBlameHandler)
}

func getProjectsIdRepositoryFilesFilePathBlameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	file_path := request.GetString("file_path", "")
	params := parseGetProjectsIdRepositoryFilesFilePathBlame(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryFilesFilePathBlame(ctx, id, file_path, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryFilesFilePathBlame(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryFilesFilePathBlameParams {
	params := client.GetApiV4ProjectsIdRepositoryFilesFilePathBlameParams{}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = ref
	}

	range_start_ := request.GetInt("range[start]", math.MinInt)
	if range_start_ != math.MinInt {
		range_start_ := int32(range_start_)
		params.RangeStart = range_start_
	}

	range_end_ := request.GetInt("range[end]", math.MinInt)
	if range_end_ != math.MinInt {
		range_end_ := int32(range_end_)
		params.RangeEnd = range_end_
	}

	return params
}

func registerGetProjectsIdRepositoryFilesFilePathRaw(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_files_file_path_raw",
		mcp.WithDescription("Get raw file contents from the repository"),
		mcp.WithString("id",
			mcp.Description("The project ID (example: gitlab-org/gitlab)"),
			mcp.Required(),
		),
		mcp.WithString("file_path",
			mcp.Description("The url encoded path to the file. (example: lib%2Fclass%2Erb)"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("The name of branch, tag or commit (example: main)"),
		),
		mcp.WithBoolean("lfs",
			mcp.Description("Retrieve binary data for a file that is an lfs pointer (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryFilesFilePathRawHandler)
}

func getProjectsIdRepositoryFilesFilePathRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	file_path := request.GetString("file_path", "")
	params := parseGetProjectsIdRepositoryFilesFilePathRaw(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryFilesFilePathRaw(ctx, id, file_path, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryFilesFilePathRaw(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryFilesFilePathRawParams {
	params := client.GetApiV4ProjectsIdRepositoryFilesFilePathRawParams{}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	lfs := request.GetBool("lfs", false)
	params.Lfs = &lfs

	return params
}

func registerGetProjectsIdRepositoryFilesFilePath(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_files_file_path",
		mcp.WithDescription("Get a file from the repository"),
		mcp.WithString("id",
			mcp.Description("The project ID (example: gitlab-org/gitlab)"),
			mcp.Required(),
		),
		mcp.WithString("file_path",
			mcp.Description("The url encoded path to the file. (example: lib%2Fclass%2Erb)"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("The name of branch, tag or commit (example: main)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryFilesFilePathHandler)
}

func getProjectsIdRepositoryFilesFilePathHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	file_path := request.GetString("file_path", "")
	params := parseGetProjectsIdRepositoryFilesFilePath(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryFilesFilePath(ctx, id, file_path, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryFilesFilePath(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryFilesFilePathParams {
	params := client.GetApiV4ProjectsIdRepositoryFilesFilePathParams{}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = ref
	}

	return params
}

func registerGetProjectsIdFreezePeriods(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_freeze_periods",
		mcp.WithDescription("Paginated list of Freeze Periods, sorted by created_at in ascending order. This feature was introduced in GitLab 13.0."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdFreezePeriodsHandler)
}

func getProjectsIdFreezePeriodsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdFreezePeriods(request)
	return toResult(c.GetApiV4ProjectsIdFreezePeriods(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdFreezePeriods(request mcp.CallToolRequest) client.GetApiV4ProjectsIdFreezePeriodsParams {
	params := client.GetApiV4ProjectsIdFreezePeriodsParams{}

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

func registerGetProjectsIdFreezePeriodsFreezePeriodId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_freeze_periods_freeze_period_id",
		mcp.WithDescription("Get a freeze period for the given `freeze_period_id`. This feature was introduced in GitLab 13.0."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("freeze_period_id",
			mcp.Description("The ID of the freeze period"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdFreezePeriodsFreezePeriodIdHandler)
}

func getProjectsIdFreezePeriodsFreezePeriodIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	freeze_period_id := int32(request.GetInt("freeze_period_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx, id, freeze_period_id, authorizationHeader))
}

func registerGetProjectsIdPackagesHelmChannelIndexYaml(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_helm_channel_index_yaml",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithNumber("id",
			mcp.Description("The ID or full path of a project"),
			mcp.Required(),
		),
		mcp.WithString("channel",
			mcp.Description("Helm channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesHelmChannelIndexYamlHandler)
}

func getProjectsIdPackagesHelmChannelIndexYamlHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	channel := request.GetString("channel", "")

	return toResult(c.GetApiV4ProjectsIdPackagesHelmChannelIndexYaml(ctx, id, channel, authorizationHeader))
}

func registerGetProjectsIdPackagesHelmChannelChartsFileNameTgz(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_helm_channel_charts_file_name_tgz",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithNumber("id",
			mcp.Description("The ID or full path of a project"),
			mcp.Required(),
		),
		mcp.WithString("channel",
			mcp.Description("Helm channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Helm package file name (example: mychart)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesHelmChannelChartsFileNameTgzHandler)
}

func getProjectsIdPackagesHelmChannelChartsFileNameTgzHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	channel := request.GetString("channel", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz(ctx, id, channel, file_name, authorizationHeader))
}

func registerGetProjectsIdServices(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_services",
		mcp.WithDescription("Get a list of all active integrations."),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdServicesHandler)
}

func getProjectsIdServicesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdServices(ctx, id, authorizationHeader))
}

func registerGetProjectsIdServicesSlug(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_services_slug",
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

	s.AddTool(tool, getProjectsIdServicesSlugHandler)
}

func getProjectsIdServicesSlugHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	slug := request.GetString("slug", "")

	return toResult(c.GetApiV4ProjectsIdServicesSlug(ctx, id, slug, authorizationHeader))
}

func registerGetProjectsIdIntegrations(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_integrations",
		mcp.WithDescription("Get a list of all active integrations."),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIntegrationsHandler)
}

func getProjectsIdIntegrationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdIntegrations(ctx, id, authorizationHeader))
}

func registerGetProjectsIdIntegrationsSlug(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_integrations_slug",
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

	s.AddTool(tool, getProjectsIdIntegrationsSlugHandler)
}

func getProjectsIdIntegrationsSlugHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	slug := request.GetString("slug", "")

	return toResult(c.GetApiV4ProjectsIdIntegrationsSlug(ctx, id, slug, authorizationHeader))
}

func registerGetProjectsIdInvitations(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_invitations",
		mcp.WithDescription("This feature was introduced in GitLab 13.6"),
		mcp.WithString("id",
			mcp.Description("The project ID"),
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

	s.AddTool(tool, getProjectsIdInvitationsHandler)
}

func getProjectsIdInvitationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdInvitations(request)
	return toResult(c.GetApiV4ProjectsIdInvitations(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdInvitations(request mcp.CallToolRequest) client.GetApiV4ProjectsIdInvitationsParams {
	params := client.GetApiV4ProjectsIdInvitationsParams{}

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

func registerGetProjectsIdIssuesIssueIidLinks(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_links",
		mcp.WithDescription("Get a list of a given issue’s linked issues, sorted by the relationship creation datetime (ascending).Issues are filtered according to the user authorizations."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project’s issue"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidLinksHandler)
}

func getProjectsIdIssuesIssueIidLinksHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := int32(request.GetInt("issue_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidLinks(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidLinksIssueLinkId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_links_issue_link_id",
		mcp.WithDescription("Gets details about an issue link. This feature was introduced in GitLab 15.1."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project’s issue"),
			mcp.Required(),
		),
		mcp.WithString("issue_link_id",
			mcp.Description("ID of an issue relationship"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidLinksIssueLinkIdHandler)
}

func getProjectsIdIssuesIssueIidLinksIssueLinkIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := int32(request.GetInt("issue_iid", math.MinInt))
	issue_link_id := request.GetString("issue_link_id", "")

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx, id, issue_iid, issue_link_id, authorizationHeader))
}

func registerGetProjectsIdCiLint(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_ci_lint",
		mcp.WithDescription("Checks if a project’s .gitlab-ci.yml configuration in a given commit (by default HEAD of the project’s default branch) is valid"),
		mcp.WithString("sha",
			mcp.Description("Deprecated: Use content_ref instead"),
		),
		mcp.WithString("content_ref",
			mcp.Description("The CI/CD configuration content is taken from this commit SHA, branch or tag. Defaults to the HEAD of the project's default branch"),
		),
		mcp.WithBoolean("dry_run",
			mcp.Description("Run pipeline creation simulation, or only do static check. This is false by default (default: false)"),
		),
		mcp.WithBoolean("include_jobs",
			mcp.Description("If the list of jobs that would exist in a static check or pipeline simulation should be included in the response. This is false by default"),
		),
		mcp.WithString("ref",
			mcp.Description("Deprecated: Use dry_run_ref instead"),
		),
		mcp.WithString("dry_run_ref",
			mcp.Description("Branch or tag used as context when executing a dry run. Defaults to the default branch of the project. Only used when dry_run is true"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdCiLintHandler)
}

func getProjectsIdCiLintHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdCiLint(request)
	return toResult(c.GetApiV4ProjectsIdCiLint(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdCiLint(request mcp.CallToolRequest) client.GetApiV4ProjectsIdCiLintParams {
	params := client.GetApiV4ProjectsIdCiLintParams{}

	sha := request.GetString("sha", "")
	if sha != "" {

		params.Sha = &sha
	}

	content_ref := request.GetString("content_ref", "")
	if content_ref != "" {

		params.ContentRef = &content_ref
	}

	dry_run := request.GetBool("dry_run", false)
	params.DryRun = &dry_run

	include_jobs := request.GetBool("include_jobs", false)
	params.IncludeJobs = &include_jobs

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	dry_run_ref := request.GetString("dry_run_ref", "")
	if dry_run_ref != "" {

		params.DryRunRef = &dry_run_ref
	}

	return params
}

func registerGetProjectsIdUploads(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_uploads",
		mcp.WithDescription("Get the list of uploads of a project"),
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

	s.AddTool(tool, getProjectsIdUploadsHandler)
}

func getProjectsIdUploadsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdUploads(request)
	return toResult(c.GetApiV4ProjectsIdUploads(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdUploads(request mcp.CallToolRequest) client.GetApiV4ProjectsIdUploadsParams {
	params := client.GetApiV4ProjectsIdUploadsParams{}

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

func registerGetProjectsIdUploadsUploadId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_uploads_upload_id",
		mcp.WithDescription("Download a single project upload by ID"),
		mcp.WithNumber("upload_id",
			mcp.Description("The ID of a project upload"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdUploadsUploadIdHandler)
}

func getProjectsIdUploadsUploadIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	upload_id := int32(request.GetInt("upload_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdUploadsUploadId(ctx, id, upload_id, authorizationHeader))
}

func registerGetProjectsIdUploadsSecretFilename(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_uploads_secret_filename",
		mcp.WithDescription("Download a single project upload by secret and filename"),
		mcp.WithString("secret",
			mcp.Description("The 32-character secret of a project upload"),
			mcp.Required(),
		),
		mcp.WithString("filename",
			mcp.Description("The filename of a project upload"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdUploadsSecretFilenameHandler)
}

func getProjectsIdUploadsSecretFilenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	secret := request.GetString("secret", "")
	filename := request.GetString("filename", "")

	return toResult(c.GetApiV4ProjectsIdUploadsSecretFilename(ctx, id, secret, filename, authorizationHeader))
}

func registerGetProjectsIdMembers(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_members",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user."),
		mcp.WithString("id",
			mcp.Description("The project ID"),
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

	s.AddTool(tool, getProjectsIdMembersHandler)
}

func getProjectsIdMembersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdMembers(request)
	return toResult(c.GetApiV4ProjectsIdMembers(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdMembers(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMembersParams {
	params := client.GetApiV4ProjectsIdMembersParams{}

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

func registerGetProjectsIdMembersAll(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_members_all",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user, including those who gained membership through ancestor group."),
		mcp.WithString("id",
			mcp.Description("The project ID"),
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

	s.AddTool(tool, getProjectsIdMembersAllHandler)
}

func getProjectsIdMembersAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdMembersAll(request)
	return toResult(c.GetApiV4ProjectsIdMembersAll(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdMembersAll(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMembersAllParams {
	params := client.GetApiV4ProjectsIdMembersAllParams{}

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

func registerGetProjectsIdMembersUserId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_members_user_id",
		mcp.WithDescription("Gets a member of a group or project."),
		mcp.WithString("id",
			mcp.Description("The project ID"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMembersUserIdHandler)
}

func getProjectsIdMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMembersUserId(ctx, id, user_id, authorizationHeader))
}

func registerGetProjectsIdMembersAllUserId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_members_all_user_id",
		mcp.WithDescription("Gets a member of a group or project, including those who gained membership through ancestor group"),
		mcp.WithString("id",
			mcp.Description("The project ID"),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The user ID of the member"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMembersAllUserIdHandler)
}

func getProjectsIdMembersAllUserIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	user_id := int32(request.GetInt("user_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMembersAllUserId(ctx, id, user_id, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidApprovals(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_approvals",
		mcp.WithDescription("List approvals for merge request"),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidApprovalsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidApprovalsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidApprovalState(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_approval_state",
		mcp.WithDescription("Get approval state of merge request"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The IID of a merge request"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidApprovalStateHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidApprovalStateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidTimeStats(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_time_stats",
		mcp.WithDescription("Get time tracking stats"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The internal ID of the merge_request"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidTimeStatsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidTimeStatsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs",
		mcp.WithDescription("Get all merge requests for this project."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
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
		mcp.WithString("iids",
			mcp.Description("Returns the request having the given `iid`."),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsHandler)
}

func getProjectsIdMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdMergeRequests(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequests(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequests(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsParams {
	params := client.GetApiV4ProjectsIdMergeRequestsParams{}

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

	iids := request.GetString("iids", "")
	if iids != "" {
		iids := strings.Split(iids, ",")
		var intSlice []int32
		for _, v := range iids {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, int32(intValue))
		}
		params.Iids = &intSlice
	}

	return params
}

func registerGetProjectsIdMergeRequestsMergeRequestIid(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid",
		mcp.WithDescription("Shows information about a single merge request. Note: the `changes_count` value in the response is a string, not an integer. This is because when an merge request has too many changes to display and store, it is capped at 1,000. In that case, the API returns the string `\"1000+\"` for the changes count."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The internal ID of the merge request."),
			mcp.Required(),
		),
		mcp.WithBoolean("render_html",
			mcp.Description("If `true`, response includes rendered HTML for title and description."),
		),
		mcp.WithBoolean("include_diverged_commits_count",
			mcp.Description("If `true`, response includes the commits behind the target branch."),
		),
		mcp.WithBoolean("include_rebase_in_progress",
			mcp.Description("If `true`, response includes whether a rebase operation is in progress."),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIid(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIid(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidParams{}

	render_html := request.GetBool("render_html", false)
	params.RenderHtml = &render_html

	include_diverged_commits_count := request.GetBool("include_diverged_commits_count", false)
	params.IncludeDivergedCommitsCount = &include_diverged_commits_count

	include_rebase_in_progress := request.GetBool("include_rebase_in_progress", false)
	params.IncludeRebaseInProgress = &include_rebase_in_progress

	return params
}

func registerGetProjectsIdMergeRequestsMergeRequestIidParticipants(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_participants",
		mcp.WithDescription("Get a list of merge request participants."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidParticipantsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidParticipantsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidReviewers(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_reviewers",
		mcp.WithDescription("Get a list of merge request reviewers."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidReviewersHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidReviewersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidCommits(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_commits",
		mcp.WithDescription("Get a list of merge request commits."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidCommitsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidCommitsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidCommits(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidCommits(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsParams{}

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

func registerGetProjectsIdMergeRequestsMergeRequestIidContextCommits(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_context_commits",
		mcp.WithDescription("Get a list of merge request context commits."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidContextCommitsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidContextCommitsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidChanges(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_changes",
		mcp.WithDescription("Shows information about the merge request including its files and changes."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithBoolean("unidiff",
			mcp.Description("A diff in a Unified diff format (default: false)"),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidChangesHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidChangesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidChanges(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidChanges(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesParams{}

	unidiff := request.GetBool("unidiff", false)
	params.Unidiff = &unidiff

	return params
}

func registerGetProjectsIdMergeRequestsMergeRequestIidDiffs(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_diffs",
		mcp.WithDescription("Get a list of merge request diffs."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("unidiff",
			mcp.Description("A diff in a Unified diff format (default: false)"),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidDiffsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidDiffsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidDiffs(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidDiffs(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsParams{}

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

	unidiff := request.GetBool("unidiff", false)
	params.Unidiff = &unidiff

	return params
}

func registerGetProjectsIdMergeRequestsMergeRequestIidRawDiffs(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_raw_diffs",
		mcp.WithDescription("Get the raw diffs of a merge request that can used programmatically."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidRawDiffsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidRawDiffsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffs(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidPipelines(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_pls",
		mcp.WithDescription("Get a list of merge request pipelines."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidPipelinesHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidPipelinesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidMergeRef(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_merge_ref",
		mcp.WithDescription("Returns the up to date merge-ref HEAD commit"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidMergeRefHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidMergeRefHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef(ctx, id, merge_request_iid, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsMergeRequestIidClosesIssues(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_closes_issues",
		mcp.WithDescription("Get all the issues that would be closed by merging the provided merge request."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidClosesIssuesHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidClosesIssuesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidClosesIssues(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidClosesIssues(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesParams{}

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

func registerGetProjectsIdMergeRequestsMergeRequestIidRelatedIssues(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_related_issues",
		mcp.WithDescription("Get all the related issues from title, description, commits, comments and discussions of the merge request."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidRelatedIssuesHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidRelatedIssuesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidRelatedIssues(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssues(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidRelatedIssues(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesParams{}

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

func registerGetProjectsIdMergeRequestsMergeRequestIidVersions(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_versions",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The internal ID of the merge request"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidVersionsHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidVersionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidVersions(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions(ctx, id, merge_request_iid, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidVersions(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsParams{}

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

func registerGetProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_versions_version_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("merge_request_iid",
			mcp.Description("The internal ID of the merge request"),
			mcp.Required(),
		),
		mcp.WithNumber("version_id",
			mcp.Description("The ID of the merge request diff version"),
			mcp.Required(),
		),
		mcp.WithBoolean("unidiff",
			mcp.Description("A diff in a Unified diff format (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdHandler)
}

func getProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	merge_request_iid := int32(request.GetInt("merge_request_iid", math.MinInt))
	version_id := int32(request.GetInt("version_id", math.MinInt))
	params := parseGetProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(ctx, id, merge_request_iid, version_id, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdParams {
	params := client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdParams{}

	unidiff := request.GetBool("unidiff", false)
	params.Unidiff = &unidiff

	return params
}

func registerGetProjectsIdPackagesNugetIndex(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_index",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesNugetIndexHandler)
}

func getProjectsIdPackagesNugetIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesNugetIndex(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesNugetV2(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_v2",
		mcp.WithDescription("This feature was introduced in GitLab 16.2"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesNugetV2Handler)
}

func getProjectsIdPackagesNugetV2Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesNugetV2(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesNugetV2Metadata(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_v2_metadata",
		mcp.WithDescription("This feature was introduced in GitLab 16.3"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesNugetV2MetadataHandler)
}

func getProjectsIdPackagesNugetV2MetadataHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesNugetV2Metadata(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPackagesNugetQuery(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_query",
		mcp.WithDescription("This feature was introduced in GitLab 12.8"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
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

	s.AddTool(tool, getProjectsIdPackagesNugetQueryHandler)
}

func getProjectsIdPackagesNugetQueryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPackagesNugetQuery(request)
	return toResult(c.GetApiV4ProjectsIdPackagesNugetQuery(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesNugetQuery(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesNugetQueryParams {
	params := client.GetApiV4ProjectsIdPackagesNugetQueryParams{}

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

func registerGetProjectsIdPackagesPackageIdPackageFiles(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_package_id_package_files",
		mcp.WithDescription("Get a list of package files of a single package"),
		mcp.WithString("id",
			mcp.Description("ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("package_id",
			mcp.Description("ID of a package"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return package files ordered by `id`, `created_at` or `file_name` (default: id)"),

			mcp.Enum("id", "created_at", "file_name"),
		),
		mcp.WithString("sort",
			mcp.Description("Return package files sorted in `asc` or `desc` order. (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesPackageIdPackageFilesHandler)
}

func getProjectsIdPackagesPackageIdPackageFilesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_id := int32(request.GetInt("package_id", math.MinInt))
	params := parseGetProjectsIdPackagesPackageIdPackageFiles(request)
	return toResult(c.GetApiV4ProjectsIdPackagesPackageIdPackageFiles(ctx, id, package_id, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesPackageIdPackageFiles(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesPackageIdPackageFilesParams {
	params := client.GetApiV4ProjectsIdPackagesPackageIdPackageFilesParams{}

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

	return params
}

func registerGetProjectsIdPages(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pages",
		mcp.WithDescription("Get pages URL and other settings. This feature was introduced in Gitlab 16.8"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPagesHandler)
}

func getProjectsIdPagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPages(ctx, id, authorizationHeader))
}

func registerGetProjectsIdPagesDomains(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pages_domains",
		mcp.WithDescription("Get all pages domains"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdPagesDomainsHandler)
}

func getProjectsIdPagesDomainsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPagesDomains(request)
	return toResult(c.GetApiV4ProjectsIdPagesDomains(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPagesDomains(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPagesDomainsParams {
	params := client.GetApiV4ProjectsIdPagesDomainsParams{}

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

func registerGetProjectsIdPagesDomainsDomain(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pages_domains_domain",
		mcp.WithDescription("Get a single pages domain"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project owned by the authenticated user"),
			mcp.Required(),
		),
		mcp.WithString("domain",
			mcp.Description("The domain"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPagesDomainsDomainHandler)
}

func getProjectsIdPagesDomainsDomainHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	domain := request.GetString("domain", "")

	return toResult(c.GetApiV4ProjectsIdPagesDomainsDomain(ctx, id, domain, authorizationHeader))
}

func registerGetProjectsIdAvatar(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_avatar",
		mcp.WithDescription("This feature was introduced in GitLab 16.9"),
		mcp.WithString("id",
			mcp.Description("ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdAvatarHandler)
}

func getProjectsIdAvatarHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdAvatar(ctx, id, authorizationHeader))
}

func registerGetProjectsIdClusters(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 11.7. Returns a list of project clusters."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdClustersHandler)
}

func getProjectsIdClustersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdClusters(request)
	return toResult(c.GetApiV4ProjectsIdClusters(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdClusters(request mcp.CallToolRequest) client.GetApiV4ProjectsIdClustersParams {
	params := client.GetApiV4ProjectsIdClustersParams{}

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

func registerGetProjectsIdClustersClusterId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.7. Gets a single project cluster."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("cluster_id",
			mcp.Description("The cluster ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdClustersClusterIdHandler)
}

func getProjectsIdClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	cluster_id := int32(request.GetInt("cluster_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdClustersClusterId(ctx, id, cluster_id, authorizationHeader))
}

func registerGetProjectsIdRegistryRepositories(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_registry_repositories",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("tags",
			mcp.Description("Determines if tags should be included (default: false)"),
		),
		mcp.WithBoolean("tags_count",
			mcp.Description("Determines if the tags count should be included (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdRegistryRepositoriesHandler)
}

func getProjectsIdRegistryRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRegistryRepositories(request)
	return toResult(c.GetApiV4ProjectsIdRegistryRepositories(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRegistryRepositories(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRegistryRepositoriesParams {
	params := client.GetApiV4ProjectsIdRegistryRepositoriesParams{}

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

	tags := request.GetBool("tags", false)
	params.Tags = &tags

	tags_count := request.GetBool("tags_count", false)
	params.TagsCount = &tags_count

	return params
}

func registerGetProjectsIdRegistryRepositoriesRepositoryIdTags(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_registry_repositories_repo_id_tags",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("repository_id",
			mcp.Description("The ID of the repository"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRegistryRepositoriesRepositoryIdTagsHandler)
}

func getProjectsIdRegistryRepositoriesRepositoryIdTagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	repository_id := int32(request.GetInt("repository_id", math.MinInt))
	params := parseGetProjectsIdRegistryRepositoriesRepositoryIdTags(request)
	return toResult(c.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx, id, repository_id, &params, authorizationHeader))
}

func parseGetProjectsIdRegistryRepositoriesRepositoryIdTags(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams {
	params := client.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams{}

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

func registerGetProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_registry_repositories_repo_id_tags_tag_name",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("repository_id",
			mcp.Description("The ID of the repository"),
			mcp.Required(),
		),
		mcp.WithString("tag_name",
			mcp.Description("The name of the tag"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameHandler)
}

func getProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	repository_id := int32(request.GetInt("repository_id", math.MinInt))
	tag_name := request.GetString("tag_name", "")

	return toResult(c.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx, id, repository_id, tag_name, authorizationHeader))
}

func registerGetProjectsIdRegistryProtectionRepositoryRules(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_registry_protection_repo_rules",
		mcp.WithDescription("Get list of container registry protection rules for a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRegistryProtectionRepositoryRulesHandler)
}

func getProjectsIdRegistryProtectionRepositoryRulesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdRegistryProtectionRepositoryRules(ctx, id, authorizationHeader))
}

func registerGetProjectsIdDebianDistributions(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_debian_distributions",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
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

	s.AddTool(tool, getProjectsIdDebianDistributionsHandler)
}

func getProjectsIdDebianDistributionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdDebianDistributions(request)
	return toResult(c.GetApiV4ProjectsIdDebianDistributions(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdDebianDistributions(request mcp.CallToolRequest) client.GetApiV4ProjectsIdDebianDistributionsParams {
	params := client.GetApiV4ProjectsIdDebianDistributionsParams{}

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

func registerGetProjectsIdDebianDistributionsCodename(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_debian_distributions_codename",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("codename",
			mcp.Description("The Debian Codename (example: sid)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdDebianDistributionsCodenameHandler)
}

func getProjectsIdDebianDistributionsCodenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	codename := request.GetString("codename", "")

	return toResult(c.GetApiV4ProjectsIdDebianDistributionsCodename(ctx, id, codename, authorizationHeader))
}

func registerGetProjectsIdDebianDistributionsCodenameKeyAsc(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_debian_distributions_codename_key_asc",
		mcp.WithDescription("This feature was introduced in 14.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("codename",
			mcp.Description("The Debian Codename (example: sid)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdDebianDistributionsCodenameKeyAscHandler)
}

func getProjectsIdDebianDistributionsCodenameKeyAscHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	codename := request.GetString("codename", "")

	return toResult(c.GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc(ctx, id, codename, authorizationHeader))
}

func registerGetProjectsIdEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_events",
		mcp.WithDescription("List a project's visible events"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
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
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdEventsHandler)
}

func getProjectsIdEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdEvents(request)
	return toResult(c.GetApiV4ProjectsIdEvents(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdEvents(request mcp.CallToolRequest) client.GetApiV4ProjectsIdEventsParams {
	params := client.GetApiV4ProjectsIdEventsParams{}

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

func registerGetProjectsIdExport(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_export",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdExportHandler)
}

func getProjectsIdExportHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdExport(ctx, id, authorizationHeader))
}

func registerGetProjectsIdExportDownload(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_export_download",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdExportDownloadHandler)
}

func getProjectsIdExportDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdExportDownload(ctx, id, authorizationHeader))
}

func registerGetProjectsIdExportRelationsDownload(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_export_relations_download",
		mcp.WithDescription("This feature was introduced in GitLab 14.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("relation",
			mcp.Description("Project relation name"),
			mcp.Required(),
		),
		mcp.WithBoolean("batched",
			mcp.Description("Whether to download in batches"),
		),
		mcp.WithNumber("batch_number",
			mcp.Description("Batch number to download"),
		),
	)

	s.AddTool(tool, getProjectsIdExportRelationsDownloadHandler)
}

func getProjectsIdExportRelationsDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdExportRelationsDownload(request)
	return toResult(c.GetApiV4ProjectsIdExportRelationsDownload(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdExportRelationsDownload(request mcp.CallToolRequest) client.GetApiV4ProjectsIdExportRelationsDownloadParams {
	params := client.GetApiV4ProjectsIdExportRelationsDownloadParams{}

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

func registerGetProjectsIdExportRelationsStatus(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_export_relations_status",
		mcp.WithDescription("This feature was introduced in GitLab 14.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("relation",
			mcp.Description("Project relation name"),
		),
	)

	s.AddTool(tool, getProjectsIdExportRelationsStatusHandler)
}

func getProjectsIdExportRelationsStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdExportRelationsStatus(request)
	return toResult(c.GetApiV4ProjectsIdExportRelationsStatus(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdExportRelationsStatus(request mcp.CallToolRequest) client.GetApiV4ProjectsIdExportRelationsStatusParams {
	params := client.GetApiV4ProjectsIdExportRelationsStatusParams{}

	relation := request.GetString("relation", "")
	if relation != "" {

		params.Relation = &relation
	}

	return params
}

func registerGetProjectsIdHooks(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_hooks",
		mcp.WithDescription("Get a list of project hooks"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdHooksHandler)
}

func getProjectsIdHooksHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdHooks(request)
	return toResult(c.GetApiV4ProjectsIdHooks(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdHooks(request mcp.CallToolRequest) client.GetApiV4ProjectsIdHooksParams {
	params := client.GetApiV4ProjectsIdHooksParams{}

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

func registerGetProjectsIdHooksHookId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_hooks_hook_id",
		mcp.WithDescription("Get a specific hook for a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("hook_id",
			mcp.Description("The ID of a project hook"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdHooksHookIdHandler)
}

func getProjectsIdHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	hook_id := int32(request.GetInt("hook_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdHooksHookId(ctx, id, hook_id, authorizationHeader))
}

func registerGetProjectsIdHooksHookIdEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_hooks_hook_id_events",
		mcp.WithDescription("List web hook logs by hook id"),
		mcp.WithString("status",
			mcp.Description("null"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
		mcp.WithNumber("hook_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdHooksHookIdEventsHandler)
}

func getProjectsIdHooksHookIdEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	hook_id := int32(request.GetInt("hook_id", math.MinInt))
	params := parseGetProjectsIdHooksHookIdEvents(request)
	return toResult(c.GetApiV4ProjectsIdHooksHookIdEvents(ctx, id, hook_id, &params, authorizationHeader))
}

func parseGetProjectsIdHooksHookIdEvents(request mcp.CallToolRequest) client.GetApiV4ProjectsIdHooksHookIdEventsParams {
	params := client.GetApiV4ProjectsIdHooksHookIdEventsParams{}

	status := request.GetString("status", "")
	if status != "" {
		status := strings.Split(status, ",")
		params.Status = &status
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	return params
}

func registerGetProjectsIdImport(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_import",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdImportHandler)
}

func getProjectsIdImportHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdImport(ctx, id, authorizationHeader))
}

func registerGetProjectsIdRelationImports(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_relation_imports",
		mcp.WithDescription("This feature was introduced in GitLab 16.11."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRelationImportsHandler)
}

func getProjectsIdRelationImportsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdRelationImports(ctx, id, authorizationHeader))
}

func registerGetProjectsIdJobTokenScope(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_job_token_scope",
		mcp.WithDescription("Fetch CI_JOB_TOKEN access settings."),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdJobTokenScopeHandler)
}

func getProjectsIdJobTokenScopeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdJobTokenScope(ctx, id, authorizationHeader))
}

func registerGetProjectsIdJobTokenScopeAllowlist(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_job_token_scope_allowlist",
		mcp.WithDescription("Fetch project inbound allowlist for CI_JOB_TOKEN access settings."),
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

	s.AddTool(tool, getProjectsIdJobTokenScopeAllowlistHandler)
}

func getProjectsIdJobTokenScopeAllowlistHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdJobTokenScopeAllowlist(request)
	return toResult(c.GetApiV4ProjectsIdJobTokenScopeAllowlist(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdJobTokenScopeAllowlist(request mcp.CallToolRequest) client.GetApiV4ProjectsIdJobTokenScopeAllowlistParams {
	params := client.GetApiV4ProjectsIdJobTokenScopeAllowlistParams{}

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

func registerGetProjectsIdJobTokenScopeGroupsAllowlist(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_job_token_scope_grps_allowlist",
		mcp.WithDescription("Fetch project groups allowlist for CI_JOB_TOKEN access settings."),
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

	s.AddTool(tool, getProjectsIdJobTokenScopeGroupsAllowlistHandler)
}

func getProjectsIdJobTokenScopeGroupsAllowlistHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdJobTokenScopeGroupsAllowlist(request)
	return toResult(c.GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdJobTokenScopeGroupsAllowlist(request mcp.CallToolRequest) client.GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistParams {
	params := client.GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistParams{}

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

func registerGetProjectsIdPackages(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs",
		mcp.WithDescription("This feature was introduced in GitLab 11.8"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return packages ordered by `created_at`, `name`, `version` or `type` fields. (default: created_at)"),

			mcp.Enum("created_at", "name", "version", "type"),
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

	s.AddTool(tool, getProjectsIdPackagesHandler)
}

func getProjectsIdPackagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdPackages(request)
	return toResult(c.GetApiV4ProjectsIdPackages(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPackages(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesParams {
	params := client.GetApiV4ProjectsIdPackagesParams{}

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

func registerGetProjectsIdPackagesPackageId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_package_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.9"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("package_id",
			mcp.Description("The ID of a package"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesPackageIdHandler)
}

func getProjectsIdPackagesPackageIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_id := int32(request.GetInt("package_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdPackagesPackageId(ctx, id, package_id, authorizationHeader))
}

func registerGetProjectsIdPackagesPackageIdPipelines(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_package_id_pls",
		mcp.WithDescription("This feature was introduced in GitLab 16.1"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithNumber("package_id",
			mcp.Description("The ID of a package"),
			mcp.Required(),
		),
		mcp.WithString("cursor",
			mcp.Description("Cursor for obtaining the next set of records"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesPackageIdPipelinesHandler)
}

func getProjectsIdPackagesPackageIdPipelinesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	package_id := int32(request.GetInt("package_id", math.MinInt))
	params := parseGetProjectsIdPackagesPackageIdPipelines(request)
	return toResult(c.GetApiV4ProjectsIdPackagesPackageIdPipelines(ctx, id, package_id, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesPackageIdPipelines(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesPackageIdPipelinesParams {
	params := client.GetApiV4ProjectsIdPackagesPackageIdPipelinesParams{}

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

	cursor := request.GetString("cursor", "")
	if cursor != "" {

		params.Cursor = &cursor
	}

	return params
}

func registerGetProjectsIdPackagesProtectionRules(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_protection_rules",
		mcp.WithDescription("Get list of package protection rules for a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesProtectionRulesHandler)
}

func getProjectsIdPackagesProtectionRulesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesProtectionRules(ctx, id, authorizationHeader))
}

func registerGetProjectsIdSnapshot(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snapshot",
		mcp.WithDescription("This feature was introduced in GitLab 10.7"),
		mcp.WithBoolean("wiki",
			mcp.Description("Set to true to receive the wiki repository"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnapshotHandler)
}

func getProjectsIdSnapshotHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdSnapshot(request)
	return toResult(c.GetApiV4ProjectsIdSnapshot(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdSnapshot(request mcp.CallToolRequest) client.GetApiV4ProjectsIdSnapshotParams {
	params := client.GetApiV4ProjectsIdSnapshotParams{}

	wiki := request.GetBool("wiki", false)
	params.Wiki = &wiki

	return params
}

func registerGetProjectsIdSnippets(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets",
		mcp.WithDescription("Get all project snippets"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsHandler)
}

func getProjectsIdSnippetsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdSnippets(request)
	return toResult(c.GetApiV4ProjectsIdSnippets(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdSnippets(request mcp.CallToolRequest) client.GetApiV4ProjectsIdSnippetsParams {
	params := client.GetApiV4ProjectsIdSnippetsParams{}

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

func registerGetProjectsIdSnippetsSnippetId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id",
		mcp.WithDescription("Get a single project snippet"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
			mcp.Description("The ID of a project snippet"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdHandler)
}

func getProjectsIdSnippetsSnippetIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetId(ctx, id, snippet_id, authorizationHeader))
}

func registerGetProjectsIdSnippetsSnippetIdRaw(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_raw",
		mcp.WithDescription("Get a raw project snippet"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
			mcp.Description("The ID of a project snippet"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdRawHandler)
}

func getProjectsIdSnippetsSnippetIdRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdRaw(ctx, id, snippet_id, authorizationHeader))
}

func registerGetProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_files_ref_file_path_raw",
		mcp.WithDescription("Get raw project snippet file contents from the repository"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("file_path",
			mcp.Description("The url encoded path to the file, e.g. lib%2Fclass%2Erb"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("The name of branch, tag or commit"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdFilesRefFilePathRawHandler)
}

func getProjectsIdSnippetsSnippetIdFilesRefFilePathRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))
	ref := request.GetString("ref", "")
	file_path := request.GetString("file_path", "")

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(ctx, id, snippet_id, ref, file_path, authorizationHeader))
}

func registerGetProjectsIdSnippetsSnippetIdUserAgentDetail(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_user_agent_detail",
		mcp.WithDescription("Get the user agent details for a project snippet"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("snippet_id",
			mcp.Description("The ID of a project snippet"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdSnippetsSnippetIdUserAgentDetailHandler)
}

func getProjectsIdSnippetsSnippetIdUserAgentDetailHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	snippet_id := int32(request.GetInt("snippet_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail(ctx, id, snippet_id, authorizationHeader))
}

func registerGetProjectsIdStatistics(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_statistics",
		mcp.WithDescription("Get the list of project fetch statistics for the last 30 days"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdStatisticsHandler)
}

func getProjectsIdStatisticsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdStatistics(ctx, id, authorizationHeader))
}

func registerGetProjectsIdTemplatesType(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_templates_type",
		mcp.WithDescription("This endpoint was introduced in GitLab 11.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("type",
			mcp.Description("The type (dockerfiles|gitignores|gitlab_ci_ymls|licenses|issues|merge_requests) of the template"),
			mcp.Required(),
			mcp.Enum("dockerfiles", "gitignores", "gitlab_ci_ymls", "licenses", "issues", "merge_requests"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdTemplatesTypeHandler)
}

func getProjectsIdTemplatesTypeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	_type := request.GetString("type", "")
	params := parseGetProjectsIdTemplatesType(request)
	return toResult(c.GetApiV4ProjectsIdTemplatesType(ctx, id, _type, &params, authorizationHeader))
}

func parseGetProjectsIdTemplatesType(request mcp.CallToolRequest) client.GetApiV4ProjectsIdTemplatesTypeParams {
	params := client.GetApiV4ProjectsIdTemplatesTypeParams{}

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

func registerGetProjectsIdTemplatesTypeName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_templates_type_name",
		mcp.WithDescription("This endpoint was introduced in GitLab 11.4"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("type",
			mcp.Description("The type (dockerfiles|gitignores|gitlab_ci_ymls|licenses|issues|merge_requests) of the template"),
			mcp.Required(),
			mcp.Enum("dockerfiles", "gitignores", "gitlab_ci_ymls", "licenses", "issues", "merge_requests"),
		),
		mcp.WithString("name",
			mcp.Description("The key of the template, as obtained from the collection endpoint. (example: MIT)"),
			mcp.Required(),
		),
		mcp.WithNumber("source_template_project_id",
			mcp.Description("The project id where a given template is being stored. This is useful when multiple templates from different projects have the same name (example: 1)"),
		),
		mcp.WithString("project",
			mcp.Description("The project name to use when expanding placeholders in the template. Only affects licenses (example: GitLab)"),
		),
		mcp.WithString("fullname",
			mcp.Description("The full name of the copyright holder to use when expanding placeholders in the template. Only affects licenses (example: GitLab B.V.)"),
		),
	)

	s.AddTool(tool, getProjectsIdTemplatesTypeNameHandler)
}

func getProjectsIdTemplatesTypeNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	_type := request.GetString("type", "")
	name := request.GetString("name", "")
	params := parseGetProjectsIdTemplatesTypeName(request)
	return toResult(c.GetApiV4ProjectsIdTemplatesTypeName(ctx, id, _type, name, &params, authorizationHeader))
}

func parseGetProjectsIdTemplatesTypeName(request mcp.CallToolRequest) client.GetApiV4ProjectsIdTemplatesTypeNameParams {
	params := client.GetApiV4ProjectsIdTemplatesTypeNameParams{}

	source_template_project_id := request.GetInt("source_template_project_id", math.MinInt)
	if source_template_project_id != math.MinInt {
		source_template_project_id := int32(source_template_project_id)
		params.SourceTemplateProjectId = &source_template_project_id
	}

	project := request.GetString("project", "")
	if project != "" {

		params.Project = &project
	}

	fullname := request.GetString("fullname", "")
	if fullname != "" {

		params.Fullname = &fullname
	}

	return params
}

func registerGetProjectsIdCustomAttributes(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_custom_attributes",
		mcp.WithDescription("Get all custom attributes on a project"),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdCustomAttributesHandler)
}

func getProjectsIdCustomAttributesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdCustomAttributes(ctx, id, authorizationHeader))
}

func registerGetProjectsIdCustomAttributesKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_custom_attributes_key",
		mcp.WithDescription("Get a custom attribute on a project"),
		mcp.WithString("key",
			mcp.Description("The key of the custom attribute"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdCustomAttributesKeyHandler)
}

func getProjectsIdCustomAttributesKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	key := request.GetString("key", "")

	return toResult(c.GetApiV4ProjectsIdCustomAttributesKey(ctx, id, key, authorizationHeader))
}

func registerGetProjects(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs",
		mcp.WithDescription("Get a list of visible projects for authenticated user"),
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

	s.AddTool(tool, getProjectsHandler)
}

func getProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetProjects(request)
	return toResult(c.GetApiV4Projects(ctx, &params, authorizationHeader))
}

func parseGetProjects(request mcp.CallToolRequest) client.GetApiV4ProjectsParams {
	params := client.GetApiV4ProjectsParams{}

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

func registerGetProjectsIdShareLocations(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_share_locations",
		mcp.WithDescription("Returns group that can be shared with the given project"),
		mcp.WithNumber("id",
			mcp.Description("The id of the project"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return list of groups matching the search criteria"),
		),
	)

	s.AddTool(tool, getProjectsIdShareLocationsHandler)
}

func getProjectsIdShareLocationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdShareLocations(request)
	return toResult(c.GetApiV4ProjectsIdShareLocations(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdShareLocations(request mcp.CallToolRequest) client.GetApiV4ProjectsIdShareLocationsParams {
	params := client.GetApiV4ProjectsIdShareLocationsParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	return params
}

func registerGetProjectsId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id",
		mcp.WithDescription("Get a single project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithBoolean("statistics",
			mcp.Description("Include project statistics (default: false)"),
		),
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
		mcp.WithBoolean("license",
			mcp.Description("Include project license data (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdHandler)
}

func getProjectsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsId(request)
	return toResult(c.GetApiV4ProjectsId(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsId(request mcp.CallToolRequest) client.GetApiV4ProjectsIdParams {
	params := client.GetApiV4ProjectsIdParams{}

	statistics := request.GetBool("statistics", false)
	params.Statistics = &statistics

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	license := request.GetBool("license", false)
	params.License = &license

	return params
}

func registerGetProjectsIdForks(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_forks",
		mcp.WithDescription("List forks of this project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
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
		mcp.WithBoolean("with_custom_attributes",
			mcp.Description("Include custom attributes in the response (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdForksHandler)
}

func getProjectsIdForksHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdForks(request)
	return toResult(c.GetApiV4ProjectsIdForks(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdForks(request mcp.CallToolRequest) client.GetApiV4ProjectsIdForksParams {
	params := client.GetApiV4ProjectsIdForksParams{}

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

	with_custom_attributes := request.GetBool("with_custom_attributes", false)
	params.WithCustomAttributes = &with_custom_attributes

	return params
}

func registerGetProjectsIdPagesAccess(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pages_access",
		mcp.WithDescription("Check pages access of this project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPagesAccessHandler)
}

func getProjectsIdPagesAccessHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPagesAccess(ctx, id, authorizationHeader))
}

func registerGetProjectsIdStarrers(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_starrers",
		mcp.WithDescription("Get the users who starred a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return list of users matching the search criteria (example: user)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdStarrersHandler)
}

func getProjectsIdStarrersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdStarrers(request)
	return toResult(c.GetApiV4ProjectsIdStarrers(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdStarrers(request mcp.CallToolRequest) client.GetApiV4ProjectsIdStarrersParams {
	params := client.GetApiV4ProjectsIdStarrersParams{}

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

func registerGetProjectsIdLanguages(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_languages",
		mcp.WithDescription("Get languages in project repository"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdLanguagesHandler)
}

func getProjectsIdLanguagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdLanguages(ctx, id, authorizationHeader))
}

func registerGetProjectsIdUsers(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_users",
		mcp.WithDescription("Get the users list of a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return list of users matching the search criteria (example: user)"),
		),
		mcp.WithString("skip_users",
			mcp.Description("Filter out users with the specified IDs"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdUsersHandler)
}

func getProjectsIdUsersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdUsers(request)
	return toResult(c.GetApiV4ProjectsIdUsers(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdUsers(request mcp.CallToolRequest) client.GetApiV4ProjectsIdUsersParams {
	params := client.GetApiV4ProjectsIdUsersParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
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

func registerGetProjectsIdGroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_grps",
		mcp.WithDescription("Get ancestor and shared groups for a project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return list of groups matching the search criteria (example: group)"),
		),
		mcp.WithString("skip_groups",
			mcp.Description("Array of group ids to exclude from list"),
		),
		mcp.WithBoolean("with_shared",
			mcp.Description("Include shared groups (default: false)"),
		),
		mcp.WithBoolean("shared_visible_only",
			mcp.Description("Limit to shared groups user has access to (default: false)"),
		),
		mcp.WithNumber("shared_min_access_level",
			mcp.Description("Limit returned shared groups by minimum access level to the project"),

			mcp.Enum("10", "15", "20", "30", "40", "50"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdGroupsHandler)
}

func getProjectsIdGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdGroups(request)
	return toResult(c.GetApiV4ProjectsIdGroups(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdGroups(request mcp.CallToolRequest) client.GetApiV4ProjectsIdGroupsParams {
	params := client.GetApiV4ProjectsIdGroupsParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

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

	with_shared := request.GetBool("with_shared", false)
	params.WithShared = &with_shared

	shared_visible_only := request.GetBool("shared_visible_only", false)
	params.SharedVisibleOnly = &shared_visible_only

	shared_min_access_level := request.GetInt("shared_min_access_level", math.MinInt)
	if shared_min_access_level != math.MinInt {
		shared_min_access_level := int32(shared_min_access_level)
		params.SharedMinAccessLevel = &shared_min_access_level
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

func registerGetProjectsIdInvitedGroups(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_invited_grps",
		mcp.WithDescription("Get a list of invited groups in this project"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("relation",
			mcp.Description("Filter by group relation"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a specific group"),
		),
		mcp.WithNumber("min_access_level",
			mcp.Description("Limit by minimum access level of authenticated user"),

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

	s.AddTool(tool, getProjectsIdInvitedGroupsHandler)
}

func getProjectsIdInvitedGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdInvitedGroups(request)
	return toResult(c.GetApiV4ProjectsIdInvitedGroups(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdInvitedGroups(request mcp.CallToolRequest) client.GetApiV4ProjectsIdInvitedGroupsParams {
	params := client.GetApiV4ProjectsIdInvitedGroupsParams{}

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

func registerGetProjectsIdTransferLocations(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_transfer_locations",
		mcp.WithDescription("Get the namespaces to where the project can be transferred"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("search",
			mcp.Description("Return list of namespaces matching the search criteria (example: search)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdTransferLocationsHandler)
}

func getProjectsIdTransferLocationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdTransferLocations(request)
	return toResult(c.GetApiV4ProjectsIdTransferLocations(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdTransferLocations(request mcp.CallToolRequest) client.GetApiV4ProjectsIdTransferLocationsParams {
	params := client.GetApiV4ProjectsIdTransferLocationsParams{}

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

func registerGetProjectsIdStorage(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_storage",
		mcp.WithDescription("Show the storage information"),
		mcp.WithString("id",
			mcp.Description("ID of a project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdStorageHandler)
}

func getProjectsIdStorageHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdStorage(ctx, id, authorizationHeader))
}

func registerGetProjectsIdAuditEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_audit_events",
		mcp.WithDescription("Get a list of audit events in this project."),
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

	s.AddTool(tool, getProjectsIdAuditEventsHandler)
}

func getProjectsIdAuditEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdAuditEvents(request)
	return toResult(c.GetApiV4ProjectsIdAuditEvents(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdAuditEvents(request mcp.CallToolRequest) client.GetApiV4ProjectsIdAuditEventsParams {
	params := client.GetApiV4ProjectsIdAuditEventsParams{}

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

func registerGetProjectsIdAuditEventsAuditEventId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_audit_events_audit_event_id",
		mcp.WithDescription("Get a specific audit event in this project."),
		mcp.WithNumber("audit_event_id",
			mcp.Description("The ID of the audit event"),
			mcp.Required(),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdAuditEventsAuditEventIdHandler)
}

func getProjectsIdAuditEventsAuditEventIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	audit_event_id := int32(request.GetInt("audit_event_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdAuditEventsAuditEventId(ctx, id, audit_event_id, authorizationHeader))
}

func registerGetProjectsIdProtectedBranches(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_protected_branches",
		mcp.WithDescription("Get a project's protected branches"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: gitlab-org/gitlab)"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("search",
			mcp.Description("Search for a protected branch by name (example: mai)"),
		),
	)

	s.AddTool(tool, getProjectsIdProtectedBranchesHandler)
}

func getProjectsIdProtectedBranchesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdProtectedBranches(request)
	return toResult(c.GetApiV4ProjectsIdProtectedBranches(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdProtectedBranches(request mcp.CallToolRequest) client.GetApiV4ProjectsIdProtectedBranchesParams {
	params := client.GetApiV4ProjectsIdProtectedBranchesParams{}

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

	return params
}

func registerGetProjectsIdProtectedBranchesName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_protected_branches_name",
		mcp.WithDescription("Get a single protected branch"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: gitlab-org/gitlab)"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("The name of the branch or wildcard (example: main)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdProtectedBranchesNameHandler)
}

func getProjectsIdProtectedBranchesNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	name := request.GetString("name", "")

	return toResult(c.GetApiV4ProjectsIdProtectedBranchesName(ctx, id, name, authorizationHeader))
}

func registerGetProjectsIdProtectedTags(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_protected_tags",
		mcp.WithDescription("This feature was introduced in GitLab 11.3."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdProtectedTagsHandler)
}

func getProjectsIdProtectedTagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdProtectedTags(request)
	return toResult(c.GetApiV4ProjectsIdProtectedTags(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdProtectedTags(request mcp.CallToolRequest) client.GetApiV4ProjectsIdProtectedTagsParams {
	params := client.GetApiV4ProjectsIdProtectedTagsParams{}

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

func registerGetProjectsIdProtectedTagsName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_protected_tags_name",
		mcp.WithDescription("This feature was introduced in GitLab 11.3."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("The name of the tag or wildcard (example: release*)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdProtectedTagsNameHandler)
}

func getProjectsIdProtectedTagsNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	name := request.GetString("name", "")

	return toResult(c.GetApiV4ProjectsIdProtectedTagsName(ctx, id, name, authorizationHeader))
}

func registerGetProjectsIdPackagesPypiSimple(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_pypi_simple",
		mcp.WithDescription("This feature was introduced in GitLab 15.1"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesPypiSimpleHandler)
}

func getProjectsIdPackagesPypiSimpleHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4ProjectsIdPackagesPypiSimple(ctx, id, authorizationHeader))
}

func registerGetProjectsIdReleases(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_releases",
		mcp.WithDescription("Returns a paginated list of releases. This feature was introduced in GitLab 11.7."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("order_by",
			mcp.Description("The field to use as order. Either `released_at` (default) or `created_at` (default: released_at)"),

			mcp.Enum("released_at", "created_at"),
		),
		mcp.WithString("sort",
			mcp.Description("The direction of the order. Either `desc` (default) for descending order or `asc` for ascending order (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("include_html_description",
			mcp.Description("If `true`, a response includes HTML rendered markdown of the release description"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return releases updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return releases updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ"),
		),
	)

	s.AddTool(tool, getProjectsIdReleasesHandler)
}

func getProjectsIdReleasesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdReleases(request)
	return toResult(c.GetApiV4ProjectsIdReleases(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdReleases(request mcp.CallToolRequest) client.GetApiV4ProjectsIdReleasesParams {
	params := client.GetApiV4ProjectsIdReleasesParams{}

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

	include_html_description := request.GetBool("include_html_description", false)
	params.IncludeHtmlDescription = &include_html_description

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

	return params
}

func registerGetProjectsIdReleasesTagName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_releases_tag_name",
		mcp.WithDescription("Gets a release for the given tag. This feature was introduced in GitLab 11.7."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("tag_name",
			mcp.Description("The Git tag the release is associated with"),
			mcp.Required(),
		),
		mcp.WithBoolean("include_html_description",
			mcp.Description("If `true`, a response includes HTML rendered markdown of the release description"),
		),
	)

	s.AddTool(tool, getProjectsIdReleasesTagNameHandler)
}

func getProjectsIdReleasesTagNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	tag_name := request.GetString("tag_name", "")
	params := parseGetProjectsIdReleasesTagName(request)
	return toResult(c.GetApiV4ProjectsIdReleasesTagName(ctx, id, tag_name, &params, authorizationHeader))
}

func parseGetProjectsIdReleasesTagName(request mcp.CallToolRequest) client.GetApiV4ProjectsIdReleasesTagNameParams {
	params := client.GetApiV4ProjectsIdReleasesTagNameParams{}

	include_html_description := request.GetBool("include_html_description", false)
	params.IncludeHtmlDescription = &include_html_description

	return params
}

func registerGetProjectsIdReleasesTagNameAssetsLinks(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_releases_tag_name_assets_links",
		mcp.WithDescription("Get assets as links from a release. This feature was introduced in GitLab 11.7."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("tag_name",
			mcp.Description("The tag associated with the release"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdReleasesTagNameAssetsLinksHandler)
}

func getProjectsIdReleasesTagNameAssetsLinksHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	tag_name := request.GetString("tag_name", "")
	params := parseGetProjectsIdReleasesTagNameAssetsLinks(request)
	return toResult(c.GetApiV4ProjectsIdReleasesTagNameAssetsLinks(ctx, id, tag_name, &params, authorizationHeader))
}

func parseGetProjectsIdReleasesTagNameAssetsLinks(request mcp.CallToolRequest) client.GetApiV4ProjectsIdReleasesTagNameAssetsLinksParams {
	params := client.GetApiV4ProjectsIdReleasesTagNameAssetsLinksParams{}

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

func registerGetProjectsIdReleasesTagNameAssetsLinksLinkId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_releases_tag_name_assets_links_link_id",
		mcp.WithDescription("Get an asset as a link from a release. This feature was introduced in GitLab 11.7."),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("tag_name",
			mcp.Description("The tag associated with the release"),
			mcp.Required(),
		),
		mcp.WithNumber("link_id",
			mcp.Description("The ID of the link"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdReleasesTagNameAssetsLinksLinkIdHandler)
}

func getProjectsIdReleasesTagNameAssetsLinksLinkIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	tag_name := request.GetString("tag_name", "")
	link_id := int32(request.GetInt("link_id", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx, id, tag_name, link_id, authorizationHeader))
}

func registerGetProjectsIdRemoteMirrors(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_remote_mirrors",
		mcp.WithDescription("List the project's remote mirrors"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRemoteMirrorsHandler)
}

func getProjectsIdRemoteMirrorsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRemoteMirrors(request)
	return toResult(c.GetApiV4ProjectsIdRemoteMirrors(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRemoteMirrors(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRemoteMirrorsParams {
	params := client.GetApiV4ProjectsIdRemoteMirrorsParams{}

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

func registerGetProjectsIdRemoteMirrorsMirrorId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_remote_mirrors_mirror_id",
		mcp.WithDescription("Get a single remote mirror"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("mirror_id",
			mcp.Description("The ID of a remote mirror"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRemoteMirrorsMirrorIdHandler)
}

func getProjectsIdRemoteMirrorsMirrorIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	mirror_id := request.GetString("mirror_id", "")

	return toResult(c.GetApiV4ProjectsIdRemoteMirrorsMirrorId(ctx, id, mirror_id, authorizationHeader))
}

func registerGetProjectsIdRemoteMirrorsMirrorIdPublicKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_remote_mirrors_mirror_id_public_key",
		mcp.WithDescription("Get the public key of a single remote mirror"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("mirror_id",
			mcp.Description("The ID of a remote mirror"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRemoteMirrorsMirrorIdPublicKeyHandler)
}

func getProjectsIdRemoteMirrorsMirrorIdPublicKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	mirror_id := request.GetString("mirror_id", "")

	return toResult(c.GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKey(ctx, id, mirror_id, authorizationHeader))
}

func registerGetProjectsIdRepositoryTree(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_tree",
		mcp.WithDescription("Get a project repository tree"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("ref",
			mcp.Description("The name of a repository branch or tag, if not given the default branch is used (example: main)"),
		),
		mcp.WithString("path",
			mcp.Description("The path of the tree (example: files/html)"),
		),
		mcp.WithBoolean("recursive",
			mcp.Description("Used to get a recursive tree (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("pagination",
			mcp.Description("Specify the pagination method (\"none\" is only valid if \"recursive\" is true) (default: legacy)"),

			mcp.Enum("legacy", "keyset", "none"),
		),
		mcp.WithString("page_token",
			mcp.Description("Record from which to start the keyset pagination (example: a1e8f8d745cc87e3a9248358d9352bb7f9a0aeba)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryTreeHandler)
}

func getProjectsIdRepositoryTreeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryTree(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryTree(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryTree(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryTreeParams {
	params := client.GetApiV4ProjectsIdRepositoryTreeParams{}

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	path := request.GetString("path", "")
	if path != "" {

		params.Path = &path
	}

	recursive := request.GetBool("recursive", false)
	params.Recursive = &recursive

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

	pagination := request.GetString("pagination", "")
	if pagination != "" {

		params.Pagination = &pagination
	}

	page_token := request.GetString("page_token", "")
	if page_token != "" {

		params.PageToken = &page_token
	}

	return params
}

func registerGetProjectsIdRepositoryBlobsShaRaw(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_blobs_sha_raw",
		mcp.WithDescription("Get raw blob contents from the repository"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("The commit hash (example: 7d70e02340bac451f281cecf0a980907974bd8be)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryBlobsShaRawHandler)
}

func getProjectsIdRepositoryBlobsShaRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")

	return toResult(c.GetApiV4ProjectsIdRepositoryBlobsShaRaw(ctx, id, sha, authorizationHeader))
}

func registerGetProjectsIdRepositoryBlobsSha(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_blobs_sha",
		mcp.WithDescription("Get a blob from the repository"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("The commit hash (example: 7d70e02340bac451f281cecf0a980907974bd8be)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryBlobsShaHandler)
}

func getProjectsIdRepositoryBlobsShaHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")

	return toResult(c.GetApiV4ProjectsIdRepositoryBlobsSha(ctx, id, sha, authorizationHeader))
}

func registerGetProjectsIdRepositoryArchive(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_archive",
		mcp.WithDescription("Get an archive of the repository"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("The commit sha of the archive to be downloaded (example: 7d70e02340bac451f281cecf0a980907974bd8be)"),
		),
		mcp.WithString("format",
			mcp.Description("The archive format (example: tar.gz)"),
		),
		mcp.WithString("path",
			mcp.Description("Subfolder of the repository to be downloaded (example: files/archives)"),
		),
		mcp.WithBoolean("include_lfs_blobs",
			mcp.Description("Used to exclude LFS objects from archive (default: true)"),
		),
		mcp.WithString("exclude_paths",
			mcp.Description("Comma-separated list of paths to exclude from the archive"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryArchiveHandler)
}

func getProjectsIdRepositoryArchiveHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryArchive(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryArchive(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryArchive(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryArchiveParams {
	params := client.GetApiV4ProjectsIdRepositoryArchiveParams{}

	sha := request.GetString("sha", "")
	if sha != "" {

		params.Sha = &sha
	}

	format := request.GetString("format", "")
	if format != "" {

		params.Format = &format
	}

	path := request.GetString("path", "")
	if path != "" {

		params.Path = &path
	}

	include_lfs_blobs := request.GetBool("include_lfs_blobs", true)
	params.IncludeLfsBlobs = &include_lfs_blobs

	exclude_paths := request.GetString("exclude_paths", "")
	if exclude_paths != "" {
		exclude_paths := strings.Split(exclude_paths, ",")
		params.ExcludePaths = &exclude_paths
	}

	return params
}

func registerGetProjectsIdRepositoryCompare(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_compare",
		mcp.WithDescription("Compare two branches, tags, or commits"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("from",
			mcp.Description("The commit, branch name, or tag name to start comparison (example: main)"),
			mcp.Required(),
		),
		mcp.WithString("to",
			mcp.Description("The commit, branch name, or tag name to stop comparison (example: feature)"),
			mcp.Required(),
		),
		mcp.WithNumber("from_project_id",
			mcp.Description("The project to compare from (example: 1)"),
		),
		mcp.WithBoolean("straight",
			mcp.Description("Comparison method, `true` for direct comparison between `from` and `to` (`from`..`to`), `false` to compare using merge base (`from`...`to`) (default: false)"),
		),
		mcp.WithBoolean("unidiff",
			mcp.Description("A diff in a Unified diff format (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryCompareHandler)
}

func getProjectsIdRepositoryCompareHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryCompare(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryCompare(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryCompare(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryCompareParams {
	params := client.GetApiV4ProjectsIdRepositoryCompareParams{}

	from := request.GetString("from", "")
	if from != "" {

		params.From = from
	}

	to := request.GetString("to", "")
	if to != "" {

		params.To = to
	}

	from_project_id := request.GetInt("from_project_id", math.MinInt)
	if from_project_id != math.MinInt {
		from_project_id := int32(from_project_id)
		params.FromProjectId = &from_project_id
	}

	straight := request.GetBool("straight", false)
	params.Straight = &straight

	unidiff := request.GetBool("unidiff", false)
	params.Unidiff = &unidiff

	return params
}

func registerGetProjectsIdRepositoryHealth(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_health",
		mcp.WithDescription("Get repository health"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithBoolean("generate",
			mcp.Description("Triggers a new health report to be generated (default: false)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryHealthHandler)
}

func getProjectsIdRepositoryHealthHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryHealth(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryHealth(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryHealth(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryHealthParams {
	params := client.GetApiV4ProjectsIdRepositoryHealthParams{}

	generate := request.GetBool("generate", false)
	params.Generate = &generate

	return params
}

func registerGetProjectsIdRepositoryContributors(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_contributors",
		mcp.WithDescription("Get repository contributors"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("ref",
			mcp.Description("The name of a repository branch or tag, if not given the default branch is used (example: main)"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return contributors ordered by `name` or `email` or `commits` (default: commits)"),

			mcp.Enum("email", "name", "commits"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort by asc (ascending) or desc (descending) (default: asc)"),

			mcp.Enum("asc", "desc"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryContributorsHandler)
}

func getProjectsIdRepositoryContributorsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryContributors(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryContributors(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryContributors(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryContributorsParams {
	params := client.GetApiV4ProjectsIdRepositoryContributorsParams{}

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

	ref := request.GetString("ref", "")
	if ref != "" {

		params.Ref = &ref
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	return params
}

func registerGetProjectsIdRepositoryMergeBase(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_merge_base",
		mcp.WithDescription("Get the common ancestor between commits"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("refs",
			mcp.Description("The refs to find the common ancestor of, multiple refs can be passed (example: main)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryMergeBaseHandler)
}

func getProjectsIdRepositoryMergeBaseHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryMergeBase(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryMergeBase(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryMergeBase(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryMergeBaseParams {
	params := client.GetApiV4ProjectsIdRepositoryMergeBaseParams{}

	refs := request.GetString("refs", "")
	if refs != "" {
		refs := strings.Split(refs, ",")
		params.Refs = refs
	}

	return params
}

func registerGetProjectsIdRepositoryChangelog(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_changelog",
		mcp.WithDescription("This feature was introduced in GitLab 14.6"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project (example: 1)"),
			mcp.Required(),
		),
		mcp.WithString("version",
			mcp.Description("The version of the release, using the semantic versioning format (example: 1.0.0)"),
			mcp.Required(),
		),
		mcp.WithString("from",
			mcp.Description("The first commit in the range of commits to use for the changelog (example: ed899a2f4b50b4370feeea94676502b42383c746)"),
		),
		mcp.WithString("to",
			mcp.Description("The last commit in the range of commits to use for the changelog (example: 6104942438c14ec7bd21c6cd5bd995272b3faff6)"),
		),
		mcp.WithString("date",
			mcp.Description("The date and time of the release (example: 2021-09-20T11:50:22.001+00:00)"),
		),
		mcp.WithString("trailer",
			mcp.Description("The Git trailer to use for determining if commits are to be included in the changelog (example: Changelog) (default: Changelog)"),
		),
		mcp.WithString("config_file",
			mcp.Description("The file path to the configuration file as stored in the project's Git repository. Defaults to '.gitlab/changelog_config.yml' (example: .gitlab/changelog_config.yml)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryChangelogHandler)
}

func getProjectsIdRepositoryChangelogHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryChangelog(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryChangelog(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryChangelog(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryChangelogParams {
	params := client.GetApiV4ProjectsIdRepositoryChangelogParams{}

	version := request.GetString("version", "")
	if version != "" {

		params.Version = version
	}

	from := request.GetString("from", "")
	if from != "" {

		params.From = &from
	}

	to := request.GetString("to", "")
	if to != "" {

		params.To = &to
	}

	date := request.GetString("date", "")
	if date != "" {
		date, _ := time.Parse(time.RFC3339, date)
		params.Date = &date
	}

	trailer := request.GetString("trailer", "")
	if trailer != "" {

		params.Trailer = &trailer
	}

	config_file := request.GetString("config_file", "")
	if config_file != "" {

		params.ConfigFile = &config_file
	}

	return params
}

func registerGetProjectsIdIssuesEventableIdResourceMilestoneEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_eventable_id_resource_milestone_events",
		mcp.WithDescription("Gets a list of all milestone events for a single Issue"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("eventable_id",
			mcp.Description("The ID of the eventable"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesEventableIdResourceMilestoneEventsHandler)
}

func getProjectsIdIssuesEventableIdResourceMilestoneEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	eventable_id := int32(request.GetInt("eventable_id", math.MinInt))
	params := parseGetProjectsIdIssuesEventableIdResourceMilestoneEvents(request)
	return toResult(c.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents(ctx, id, eventable_id, &params, authorizationHeader))
}

func parseGetProjectsIdIssuesEventableIdResourceMilestoneEvents(request mcp.CallToolRequest) client.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsParams {
	params := client.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsParams{}

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

func registerGetProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_eventable_id_resource_milestone_events_event_id",
		mcp.WithDescription("Returns a single milestone event for a specific project Issue"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("event_id",
			mcp.Description("The ID of a resource milestone event"),
			mcp.Required(),
		),
		mcp.WithNumber("eventable_id",
			mcp.Description("The ID of the eventable"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdHandler)
}

func getProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	eventable_id := int32(request.GetInt("eventable_id", math.MinInt))
	event_id := request.GetString("event_id", "")

	return toResult(c.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(ctx, id, eventable_id, event_id, authorizationHeader))
}

func registerGetProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_eventable_id_resource_milestone_events",
		mcp.WithDescription("Gets a list of all milestone events for a single Merge request"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithNumber("eventable_id",
			mcp.Description("The ID of the eventable"),
			mcp.Required(),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsHandler)
}

func getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	eventable_id := int32(request.GetInt("eventable_id", math.MinInt))
	params := parseGetProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(request)
	return toResult(c.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(ctx, id, eventable_id, &params, authorizationHeader))
}

func parseGetProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(request mcp.CallToolRequest) client.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsParams {
	params := client.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsParams{}

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

func registerGetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_mrs_eventable_id_resource_milestone_events_event_id",
		mcp.WithDescription("Returns a single milestone event for a specific project Merge request"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("event_id",
			mcp.Description("The ID of a resource milestone event"),
			mcp.Required(),
		),
		mcp.WithNumber("eventable_id",
			mcp.Description("The ID of the eventable"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdHandler)
}

func getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	eventable_id := int32(request.GetInt("eventable_id", math.MinInt))
	event_id := request.GetString("event_id", "")

	return toResult(c.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(ctx, id, eventable_id, event_id, authorizationHeader))
}

func registerGetProjectsIdPackagesRubygemsFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithNumber("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Spec file name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesRubygemsFileNameHandler)
}

func getProjectsIdPackagesRubygemsFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsFileName(ctx, id, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesRubygemsQuickMarshal48FileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_quick_marshal_4_8_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithNumber("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Gemspec file name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesRubygemsQuickMarshal48FileNameHandler)
}

func getProjectsIdPackagesRubygemsQuickMarshal48FileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName(ctx, id, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesRubygemsGemsFileName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_gems_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithNumber("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesRubygemsGemsFileNameHandler)
}

func getProjectsIdPackagesRubygemsGemsFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsGemsFileName(ctx, id, file_name, authorizationHeader))
}

func registerGetProjectsIdPackagesRubygemsApiV1Dependencies(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_api_v1_dependencies",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithNumber("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("gems",
			mcp.Description("Comma delimited gem names"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesRubygemsApiV1DependenciesHandler)
}

func getProjectsIdPackagesRubygemsApiV1DependenciesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdPackagesRubygemsApiV1Dependencies(request)
	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesRubygemsApiV1Dependencies(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesParams {
	params := client.GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesParams{}

	gems := request.GetString("gems", "")
	if gems != "" {
		gems := strings.Split(gems, ",")
		params.Gems = &gems
	}

	return params
}

func registerGetProjectsIdRepositoryTags(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_tags",
		mcp.WithDescription("Get a project repository tags"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("sort",
			mcp.Description("Return tags sorted in updated by `asc` or `desc` order. (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("order_by",
			mcp.Description("Return tags ordered by `name`, `updated`, `version` fields. (default: updated)"),

			mcp.Enum("name", "updated", "version"),
		),
		mcp.WithString("search",
			mcp.Description("Return list of tags matching the search criteria"),
		),
		mcp.WithString("page_token",
			mcp.Description("Name of tag to start the paginaition from"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryTagsHandler)
}

func getProjectsIdRepositoryTagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdRepositoryTags(request)
	return toResult(c.GetApiV4ProjectsIdRepositoryTags(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdRepositoryTags(request mcp.CallToolRequest) client.GetApiV4ProjectsIdRepositoryTagsParams {
	params := client.GetApiV4ProjectsIdRepositoryTagsParams{}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	page_token := request.GetString("page_token", "")
	if page_token != "" {

		params.PageToken = &page_token
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

func registerGetProjectsIdRepositoryTagsTagName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_tags_tag_name",
		mcp.WithDescription("Get a single repository tag"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("tag_name",
			mcp.Description("The name of the tag"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryTagsTagNameHandler)
}

func getProjectsIdRepositoryTagsTagNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	tag_name := request.GetString("tag_name", "")

	return toResult(c.GetApiV4ProjectsIdRepositoryTagsTagName(ctx, id, tag_name, authorizationHeader))
}

func registerGetProjectsIdRepositoryTagsTagNameSignature(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_repo_tags_tag_name_signature",
		mcp.WithDescription("Get a tag's signature"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("tag_name",
			mcp.Description("The name of the tag"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdRepositoryTagsTagNameSignatureHandler)
}

func getProjectsIdRepositoryTagsTagNameSignatureHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	tag_name := request.GetString("tag_name", "")

	return toResult(c.GetApiV4ProjectsIdRepositoryTagsTagNameSignature(ctx, id, tag_name, authorizationHeader))
}

func registerGetProjectsIdPackagesTerraformModulesModuleNameModuleSystem(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_pkgs_terraform_modules_module_name_module_system",
		mcp.WithDescription("This feature was introduced in GitLab 16.7"),
		mcp.WithString("id",
			mcp.Description("The ID or full path of a project"),
			mcp.Required(),
		),
		mcp.WithString("module_name",
			mcp.Description("Module name (example: infra-registry)"),
			mcp.Required(),
		),
		mcp.WithString("module_system",
			mcp.Description("Module system (example: aws)"),
			mcp.Required(),
		),
		mcp.WithString("terraform-get",
			mcp.Description("Terraform get redirection flag"),

			mcp.Enum("1"),
		),
	)

	s.AddTool(tool, getProjectsIdPackagesTerraformModulesModuleNameModuleSystemHandler)
}

func getProjectsIdPackagesTerraformModulesModuleNameModuleSystemHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	module_name := request.GetString("module_name", "")
	module_system := request.GetString("module_system", "")
	params := parseGetProjectsIdPackagesTerraformModulesModuleNameModuleSystem(request)
	return toResult(c.GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem(ctx, id, module_name, module_system, &params, authorizationHeader))
}

func parseGetProjectsIdPackagesTerraformModulesModuleNameModuleSystem(request mcp.CallToolRequest) client.GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemParams {
	params := client.GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemParams{}

	terraform_get := request.GetString("terraform-get", "")
	if terraform_get != "" {

		params.TerraformGet = &terraform_get
	}

	return params
}

func registerGetProjectsIdTerraformStateName(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_terraform_state_name",
		mcp.WithDescription("Get a Terraform state by its name"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("The name of a Terraform state"),
			mcp.Required(),
		),
		mcp.WithString("ID",
			mcp.Description("Terraform state lock ID"),
		),
	)

	s.AddTool(tool, getProjectsIdTerraformStateNameHandler)
}

func getProjectsIdTerraformStateNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	name := request.GetString("name", "")
	params := parseGetProjectsIdTerraformStateName(request)
	return toResult(c.GetApiV4ProjectsIdTerraformStateName(ctx, id, name, &params, authorizationHeader))
}

func parseGetProjectsIdTerraformStateName(request mcp.CallToolRequest) client.GetApiV4ProjectsIdTerraformStateNameParams {
	params := client.GetApiV4ProjectsIdTerraformStateNameParams{}

	ID := request.GetString("ID", "")
	if ID != "" {

		params.ID = &ID
	}

	return params
}

func registerGetProjectsIdTerraformStateNameVersionsSerial(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_terraform_state_name_versions_serial",
		mcp.WithDescription("Get a Terraform state version"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of the project"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("The name of a Terraform state"),
			mcp.Required(),
		),
		mcp.WithNumber("serial",
			mcp.Description("The version number of the state"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdTerraformStateNameVersionsSerialHandler)
}

func getProjectsIdTerraformStateNameVersionsSerialHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	name := request.GetString("name", "")
	serial := int32(request.GetInt("serial", math.MinInt))

	return toResult(c.GetApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx, id, name, serial, authorizationHeader))
}

func registerGetProjectsIdWikis(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_wikis",
		mcp.WithDescription("Get a list of wiki pages"),
		mcp.WithBoolean("with_content",
			mcp.Description("Include pages' content (default: false)"),
		),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdWikisHandler)
}

func getProjectsIdWikisHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetProjectsIdWikis(request)
	return toResult(c.GetApiV4ProjectsIdWikis(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdWikis(request mcp.CallToolRequest) client.GetApiV4ProjectsIdWikisParams {
	params := client.GetApiV4ProjectsIdWikisParams{}

	with_content := request.GetBool("with_content", false)
	params.WithContent = &with_content

	return params
}

func registerGetProjectsIdWikisSlug(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_wikis_slug",
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

	s.AddTool(tool, getProjectsIdWikisSlugHandler)
}

func getProjectsIdWikisSlugHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	slug := request.GetString("slug", "")
	params := parseGetProjectsIdWikisSlug(request)
	return toResult(c.GetApiV4ProjectsIdWikisSlug(ctx, id, slug, &params, authorizationHeader))
}

func parseGetProjectsIdWikisSlug(request mcp.CallToolRequest) client.GetApiV4ProjectsIdWikisSlugParams {
	params := client.GetApiV4ProjectsIdWikisSlugParams{}

	version := request.GetString("version", "")
	if version != "" {

		params.Version = &version
	}

	render_html := request.GetBool("render_html", false)
	params.RenderHtml = &render_html

	return params
}

func registerGetProjectsIdIssues(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues",
		mcp.WithDescription("List project issues"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
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
		mcp.WithString("not",
			mcp.Description("Return issues that do not match the parameters supplied. Accepts: labels, milestone, author_id, author_username, assignee_id, assignee_username, my_reaction_emoji, search, in."),
		),
		mcp.WithString("order_by",
			mcp.Description("Return issues ordered by created_at, updated_at, priority, due_date, relative_position, label_priority, milestone_due, popularity, weight fields. Default is created_at."),
		),
		mcp.WithString("scope",
			mcp.Description("Return issues for the given scope: created_by_me, assigned_to_me or all. Defaults to all."),
		),
		mcp.WithString("search",
			mcp.Description("Search project issues against their title and description."),
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

	s.AddTool(tool, getProjectsIdIssuesHandler)
}

func getProjectsIdIssuesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetProjectsIdIssues(request)
	return toResult(c.GetApiV4ProjectsIdIssues(ctx, id, &params, authorizationHeader))
}

func parseGetProjectsIdIssues(request mcp.CallToolRequest) client.GetApiV4ProjectsIdIssuesParams {
	params := client.GetApiV4ProjectsIdIssuesParams{}

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

func registerGetProjectsIdIssuesIssueIid(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid",
		mcp.WithDescription("Single project issue"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project's issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidHandler)
}

func getProjectsIdIssuesIssueIidHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIid(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidTimeStats(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_time_stats",
		mcp.WithDescription("Get time tracking stats"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project's issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidTimeStatsHandler)
}

func getProjectsIdIssuesIssueIidTimeStatsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidTimeStats(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidRelatedMergeRequests(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_related_mrs",
		mcp.WithDescription("List merge requests related to issue"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project's issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidRelatedMergeRequestsHandler)
}

func getProjectsIdIssuesIssueIidRelatedMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequests(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidClosedBy(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_closed_by",
		mcp.WithDescription("List merge requests that close a particular issue on merge"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidClosedByHandler)
}

func getProjectsIdIssuesIssueIidClosedByHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidClosedBy(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidParticipants(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_participants",
		mcp.WithDescription("List participants in an issue"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project's issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidParticipantsHandler)
}

func getProjectsIdIssuesIssueIidParticipantsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidParticipants(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidUserAgentDetail(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_user_agent_detail",
		mcp.WithDescription("Get user agent details"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project's issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidUserAgentDetailHandler)
}

func getProjectsIdIssuesIssueIidUserAgentDetailHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidUserAgentDetail(ctx, id, issue_iid, authorizationHeader))
}

func registerGetProjectsIdIssuesIssueIidMetricImages(s *server.MCPServer) {
	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_metric_images",
		mcp.WithDescription("List metric images"),
		mcp.WithString("id",
			mcp.Description("The global ID or URL-encoded path of the project."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_iid",
			mcp.Description("The internal ID of a project's issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getProjectsIdIssuesIssueIidMetricImagesHandler)
}

func getProjectsIdIssuesIssueIidMetricImagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	issue_iid := request.GetInt("issue_iid", math.MinInt)

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidMetricImages(ctx, id, issue_iid, authorizationHeader))
}
