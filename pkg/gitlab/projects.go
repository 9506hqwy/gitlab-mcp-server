package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostProjectsIdAccessRequestsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
}

func registerPostProjectsIdAccessRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdAccessRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdAccessRequestsHandler))
}

func postProjectsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdAccessRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdAccessRequests(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdAccessRequestsRequest struct {
	Id     string                                         `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdAccessRequestsParams `json:"params,omitempty"`
}

func registerGetProjectsIdAccessRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdAccessRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdAccessRequestsHandler))
}

func getProjectsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdAccessRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdAccessRequests(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdAccessRequestsUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the access requester"`
}

func registerDeleteProjectsIdAccessRequestsUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdAccessRequestsUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_access_requests_user_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdAccessRequestsUserIdHandler))
}

func deleteProjectsIdAccessRequestsUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdAccessRequestsUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdAccessRequestsUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type PostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AlertIid int32  `json:"alert_iid" jsonschema:"description=The IID of the Alert"`
}

func registerPostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_alert_management_alerts_alert_iid_metric_images_authorize",
		mcp.WithDescription("Workhorse authorize metric image file upload"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeHandler))
}

func postProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(ctx, req.Id, req.AlertIid, authorizationHeader))
}

type GetProjectsIdAlertManagementAlertsAlertIidMetricImagesRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AlertIid int32  `json:"alert_iid" jsonschema:"description=The IID of the Alert"`
}

func registerGetProjectsIdAlertManagementAlertsAlertIidMetricImages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdAlertManagementAlertsAlertIidMetricImagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_alert_management_alerts_alert_iid_metric_images",
		mcp.WithDescription("Metric Images for alert"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdAlertManagementAlertsAlertIidMetricImagesHandler))
}

func getProjectsIdAlertManagementAlertsAlertIidMetricImagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdAlertManagementAlertsAlertIidMetricImagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages(ctx, req.Id, req.AlertIid, authorizationHeader))
}

type DeleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdRequest struct {
	Id            string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AlertIid      int32  `json:"alert_iid" jsonschema:"description=The IID of the Alert"`
	MetricImageId int32  `json:"metric_image_id" jsonschema:"description=The ID of metric image"`
}

func registerDeleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_alert_management_alerts_alert_iid_metric_images_metric_image_id",
		mcp.WithDescription("Remove a metric image for an alert"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdHandler))
}

func deleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(ctx, req.Id, req.AlertIid, req.MetricImageId, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidAwardEmojiRequest struct {
	Id       string                                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	IssueIid int32                                                    `json:"issue_iid" jsonschema:"description=ID ('iid' for merge requests/issues/epics, 'id' for snippets) of an awardable."`
	Params   *client.GetApiV4ProjectsIdIssuesIssueIidAwardEmojiParams `json:"params,omitempty"`
}

func registerGetProjectsIdIssuesIssueIidAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidAwardEmojiHandler))
}

func getProjectsIdIssuesIssueIidAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidAwardEmoji(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type DeleteProjectsIdIssuesIssueIidAwardEmojiAwardIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	IssueIid int32 `json:"issue_iid" jsonschema:"description=null"`
	AwardId  int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteProjectsIdIssuesIssueIidAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdIssuesIssueIidAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_issues_issue_iid_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdIssuesIssueIidAwardEmojiAwardIdHandler))
}

func deleteProjectsIdIssuesIssueIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdIssuesIssueIidAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId(ctx, req.Id, req.IssueIid, req.AwardId, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidAwardEmojiAwardIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	IssueIid int32 `json:"issue_iid" jsonschema:"description=null"`
	AwardId  int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetProjectsIdIssuesIssueIidAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidAwardEmojiAwardIdHandler))
}

func getProjectsIdIssuesIssueIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId(ctx, req.Id, req.IssueIid, req.AwardId, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiRequest struct {
	Id       int32                                                               `json:"id" jsonschema:"description=null"`
	IssueIid int32                                                               `json:"issue_iid" jsonschema:"description=null"`
	NoteId   int32                                                               `json:"note_id" jsonschema:"description=null"`
	Params   *client.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiParams `json:"params,omitempty"`
}

func registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_notes_note_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiHandler))
}

func getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(ctx, req.Id, req.IssueIid, req.NoteId, req.Params, authorizationHeader))
}

type DeleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	IssueIid int32 `json:"issue_iid" jsonschema:"description=null"`
	NoteId   int32 `json:"note_id" jsonschema:"description=null"`
	AwardId  int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_issues_issue_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdHandler))
}

func deleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.IssueIid, req.NoteId, req.AwardId, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	IssueIid int32 `json:"issue_iid" jsonschema:"description=null"`
	NoteId   int32 `json:"note_id" jsonschema:"description=null"`
	AwardId  int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdHandler))
}

func getProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.IssueIid, req.NoteId, req.AwardId, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidAwardEmojiRequest struct {
	Id              string                                                                 `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MergeRequestIid int32                                                                  `json:"merge_request_iid" jsonschema:"description=ID ('iid' for merge requests/issues/epics, 'id' for snippets) of an awardable."`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidAwardEmojiHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmoji(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type DeleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
	AwardId         int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_mrs_merge_request_iid_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdHandler))
}

func deleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(ctx, req.Id, req.MergeRequestIid, req.AwardId, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
	AwardId         int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(ctx, req.Id, req.MergeRequestIid, req.AwardId, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiRequest struct {
	Id              int32                                                                             `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32                                                                             `json:"merge_request_iid" jsonschema:"description=null"`
	NoteId          int32                                                                             `json:"note_id" jsonschema:"description=null"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_notes_note_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(ctx, req.Id, req.MergeRequestIid, req.NoteId, req.Params, authorizationHeader))
}

type DeleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
	NoteId          int32 `json:"note_id" jsonschema:"description=null"`
	AwardId         int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_mrs_merge_request_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdHandler))
}

func deleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.MergeRequestIid, req.NoteId, req.AwardId, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
	NoteId          int32 `json:"note_id" jsonschema:"description=null"`
	AwardId         int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.MergeRequestIid, req.NoteId, req.AwardId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdAwardEmojiRequest struct {
	Id        string                                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	SnippetId int32                                                       `json:"snippet_id" jsonschema:"description=ID ('iid' for merge requests/issues/epics, 'id' for snippets) of an awardable."`
	Params    *client.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiParams `json:"params,omitempty"`
}

func registerGetProjectsIdSnippetsSnippetIdAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdAwardEmojiHandler))
}

func getProjectsIdSnippetsSnippetIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmoji(ctx, req.Id, req.SnippetId, req.Params, authorizationHeader))
}

type DeleteProjectsIdSnippetsSnippetIdAwardEmojiAwardIdRequest struct {
	Id        int32 `json:"id" jsonschema:"description=null"`
	SnippetId int32 `json:"snippet_id" jsonschema:"description=null"`
	AwardId   int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdSnippetsSnippetIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_snippets_snippet_id_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdSnippetsSnippetIdAwardEmojiAwardIdHandler))
}

func deleteProjectsIdSnippetsSnippetIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdSnippetsSnippetIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId(ctx, req.Id, req.SnippetId, req.AwardId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdAwardEmojiAwardIdRequest struct {
	Id        int32 `json:"id" jsonschema:"description=null"`
	SnippetId int32 `json:"snippet_id" jsonschema:"description=null"`
	AwardId   int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdAwardEmojiAwardIdHandler))
}

func getProjectsIdSnippetsSnippetIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId(ctx, req.Id, req.SnippetId, req.AwardId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiRequest struct {
	Id        int32                                                                  `json:"id" jsonschema:"description=null"`
	SnippetId int32                                                                  `json:"snippet_id" jsonschema:"description=null"`
	NoteId    int32                                                                  `json:"note_id" jsonschema:"description=null"`
	Params    *client.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiParams `json:"params,omitempty"`
}

func registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_notes_note_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiHandler))
}

func getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(ctx, req.Id, req.SnippetId, req.NoteId, req.Params, authorizationHeader))
}

type DeleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id        int32 `json:"id" jsonschema:"description=null"`
	SnippetId int32 `json:"snippet_id" jsonschema:"description=null"`
	NoteId    int32 `json:"note_id" jsonschema:"description=null"`
	AwardId   int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_snippets_snippet_id_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdHandler))
}

func deleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.SnippetId, req.NoteId, req.AwardId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id        int32 `json:"id" jsonschema:"description=null"`
	SnippetId int32 `json:"snippet_id" jsonschema:"description=null"`
	NoteId    int32 `json:"note_id" jsonschema:"description=null"`
	AwardId   int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdHandler))
}

func getProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.SnippetId, req.NoteId, req.AwardId, authorizationHeader))
}

type GetProjectsIdBadgesRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user."`
	Params *client.GetApiV4ProjectsIdBadgesParams `json:"params,omitempty"`
}

func registerGetProjectsIdBadges(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdBadgesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_badges",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdBadgesHandler))
}

func getProjectsIdBadgesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdBadgesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdBadges(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdBadgesRenderRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user."`
	Params *client.GetApiV4ProjectsIdBadgesRenderParams `json:"params,omitempty"`
}

func registerGetProjectsIdBadgesRender(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdBadgesRenderRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_badges_render",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdBadgesRenderHandler))
}

func getProjectsIdBadgesRenderHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdBadgesRenderRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdBadgesRender(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdBadgesBadgeIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user."`
	BadgeId int32  `json:"badge_id" jsonschema:"description=The badge ID"`
}

func registerDeleteProjectsIdBadgesBadgeId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdBadgesBadgeIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_badges_badge_id",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdBadgesBadgeIdHandler))
}

func deleteProjectsIdBadgesBadgeIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdBadgesBadgeIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdBadgesBadgeId(ctx, req.Id, req.BadgeId, authorizationHeader))
}

type GetProjectsIdBadgesBadgeIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user."`
	BadgeId int32  `json:"badge_id" jsonschema:"description=The badge ID"`
}

func registerGetProjectsIdBadgesBadgeId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdBadgesBadgeIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_badges_badge_id",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdBadgesBadgeIdHandler))
}

func getProjectsIdBadgesBadgeIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdBadgesBadgeIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdBadgesBadgeId(ctx, req.Id, req.BadgeId, authorizationHeader))
}

type GetProjectsIdRepositoryBranchesRequest struct {
	Id     string                                             `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryBranchesParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryBranches(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryBranchesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_branches",
		mcp.WithDescription("Get a project repository branches"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryBranchesHandler))
}

func getProjectsIdRepositoryBranchesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryBranchesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryBranches(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdRepositoryBranchesBranchRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Branch string `json:"branch" jsonschema:"description=The name of the branch"`
}

func registerDeleteProjectsIdRepositoryBranchesBranch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRepositoryBranchesBranchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_repo_branches_branch",
		mcp.WithDescription("Delete a branch"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRepositoryBranchesBranchHandler))
}

func deleteProjectsIdRepositoryBranchesBranchHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRepositoryBranchesBranchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRepositoryBranchesBranch(ctx, req.Id, req.Branch, authorizationHeader))
}

type GetProjectsIdRepositoryBranchesBranchRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Branch int32  `json:"branch" jsonschema:"description=null"`
}

func registerGetProjectsIdRepositoryBranchesBranch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryBranchesBranchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_branches_branch",
		mcp.WithDescription("Get a single repository branch"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryBranchesBranchHandler))
}

func getProjectsIdRepositoryBranchesBranchHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryBranchesBranchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryBranchesBranch(ctx, req.Id, req.Branch, authorizationHeader))
}

type PutProjectsIdRepositoryBranchesBranchUnprotectRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Branch string `json:"branch" jsonschema:"description=The name of the branch"`
}

func registerPutProjectsIdRepositoryBranchesBranchUnprotect(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdRepositoryBranchesBranchUnprotectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_repo_branches_branch_unprotect",
		mcp.WithDescription("Unprotect a single branch"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdRepositoryBranchesBranchUnprotectHandler))
}

func putProjectsIdRepositoryBranchesBranchUnprotectHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdRepositoryBranchesBranchUnprotectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect(ctx, req.Id, req.Branch, authorizationHeader))
}

type DeleteProjectsIdRepositoryMergedBranchesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerDeleteProjectsIdRepositoryMergedBranches(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRepositoryMergedBranchesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_repo_merged_branches",
		mcp.WithDescription("Delete all merged branches"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRepositoryMergedBranchesHandler))
}

func deleteProjectsIdRepositoryMergedBranchesHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRepositoryMergedBranchesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRepositoryMergedBranches(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdJobsArtifactsRefNameDownloadRequest struct {
	Id      string                                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	RefName string                                                       `json:"ref_name" jsonschema:"description=Branch or tag name in repository. 'HEAD' or 'SHA' references are not supported."`
	Params  *client.GetApiV4ProjectsIdJobsArtifactsRefNameDownloadParams `json:"params,omitempty"`
}

func registerGetProjectsIdJobsArtifactsRefNameDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobsArtifactsRefNameDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_jobs_artifacts_ref_name_download",
		mcp.WithDescription("This feature was introduced in GitLab 8.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobsArtifactsRefNameDownloadHandler))
}

func getProjectsIdJobsArtifactsRefNameDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobsArtifactsRefNameDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobsArtifactsRefNameDownload(ctx, req.Id, req.RefName, req.Params, authorizationHeader))
}

type DeleteProjectsIdJobsJobIdArtifactsRequest struct {
	Id    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	JobId int32  `json:"job_id" jsonschema:"description=The ID of a job"`
}

func registerDeleteProjectsIdJobsJobIdArtifacts(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdJobsJobIdArtifactsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_jobs_job_id_artifacts",
		mcp.WithDescription("This feature was introduced in GitLab 11.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdJobsJobIdArtifactsHandler))
}

func deleteProjectsIdJobsJobIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdJobsJobIdArtifactsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdJobsJobIdArtifacts(ctx, req.Id, req.JobId, authorizationHeader))
}

type GetProjectsIdJobsJobIdArtifactsRequest struct {
	Id     string                                             `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	JobId  int32                                              `json:"job_id" jsonschema:"description=The ID of a job"`
	Params *client.GetApiV4ProjectsIdJobsJobIdArtifactsParams `json:"params,omitempty"`
}

func registerGetProjectsIdJobsJobIdArtifacts(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobsJobIdArtifactsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_jobs_job_id_artifacts",
		mcp.WithDescription("This feature was introduced in GitLab 8.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobsJobIdArtifactsHandler))
}

func getProjectsIdJobsJobIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobsJobIdArtifactsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobsJobIdArtifacts(ctx, req.Id, req.JobId, req.Params, authorizationHeader))
}

type PostProjectsIdJobsJobIdArtifactsKeepRequest struct {
	Id    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	JobId int32  `json:"job_id" jsonschema:"description=The ID of a job"`
}

func registerPostProjectsIdJobsJobIdArtifactsKeep(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdJobsJobIdArtifactsKeepRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_jobs_job_id_artifacts_keep",
		mcp.WithDescription("Keep the artifacts to prevent them from being deleted"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdJobsJobIdArtifactsKeepHandler))
}

func postProjectsIdJobsJobIdArtifactsKeepHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdJobsJobIdArtifactsKeepRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdJobsJobIdArtifactsKeep(ctx, req.Id, req.JobId, authorizationHeader))
}

type DeleteProjectsIdArtifactsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerDeleteProjectsIdArtifacts(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdArtifactsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_artifacts",
		mcp.WithDescription("Expire the artifacts files from a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdArtifactsHandler))
}

func deleteProjectsIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdArtifactsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdArtifacts(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdJobsRequest struct {
	Id     string                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdJobsParams `json:"params,omitempty"`
}

func registerGetProjectsIdJobs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_jobs",
		mcp.WithDescription("Get a projects jobs"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobsHandler))
}

func getProjectsIdJobsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobs(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdJobsJobIdRequest struct {
	Id    int32 `json:"id" jsonschema:"description=null"`
	JobId int32 `json:"job_id" jsonschema:"description=The ID of a job"`
}

func registerGetProjectsIdJobsJobId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobsJobIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_jobs_job_id",
		mcp.WithDescription("Get a specific job of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobsJobIdHandler))
}

func getProjectsIdJobsJobIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobsJobIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobsJobId(ctx, req.Id, req.JobId, authorizationHeader))
}

type GetProjectsIdJobsJobIdTraceRequest struct {
	Id    int32 `json:"id" jsonschema:"description=null"`
	JobId int32 `json:"job_id" jsonschema:"description=The ID of a job"`
}

func registerGetProjectsIdJobsJobIdTrace(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobsJobIdTraceRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_jobs_job_id_trace",
		mcp.WithDescription("Get a trace of a specific job of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobsJobIdTraceHandler))
}

func getProjectsIdJobsJobIdTraceHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobsJobIdTraceRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobsJobIdTrace(ctx, req.Id, req.JobId, authorizationHeader))
}

type PostProjectsIdJobsJobIdRetryRequest struct {
	Id    int32 `json:"id" jsonschema:"description=null"`
	JobId int32 `json:"job_id" jsonschema:"description=The ID of a job"`
}

func registerPostProjectsIdJobsJobIdRetry(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdJobsJobIdRetryRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_jobs_job_id_retry",
		mcp.WithDescription("Retry a specific job of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdJobsJobIdRetryHandler))
}

func postProjectsIdJobsJobIdRetryHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdJobsJobIdRetryRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdJobsJobIdRetry(ctx, req.Id, req.JobId, authorizationHeader))
}

type PostProjectsIdJobsJobIdEraseRequest struct {
	Id    int32 `json:"id" jsonschema:"description=null"`
	JobId int32 `json:"job_id" jsonschema:"description=The ID of a build"`
}

func registerPostProjectsIdJobsJobIdErase(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdJobsJobIdEraseRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_jobs_job_id_erase",
		mcp.WithDescription("Erase job (remove artifacts and the trace)"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdJobsJobIdEraseHandler))
}

func postProjectsIdJobsJobIdEraseHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdJobsJobIdEraseRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdJobsJobIdErase(ctx, req.Id, req.JobId, authorizationHeader))
}

type GetProjectsIdResourceGroupsRequest struct {
	Id     string                                         `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdResourceGroupsParams `json:"params,omitempty"`
}

func registerGetProjectsIdResourceGroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdResourceGroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_resource_grps",
		mcp.WithDescription("Get all resource groups for a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdResourceGroupsHandler))
}

func getProjectsIdResourceGroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdResourceGroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdResourceGroups(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdResourceGroupsKeyRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Key string `json:"key" jsonschema:"description=The key of the resource group"`
}

func registerGetProjectsIdResourceGroupsKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdResourceGroupsKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_resource_grps_key",
		mcp.WithDescription("Get a specific resource group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdResourceGroupsKeyHandler))
}

func getProjectsIdResourceGroupsKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdResourceGroupsKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdResourceGroupsKey(ctx, req.Id, req.Key, authorizationHeader))
}

type GetProjectsIdResourceGroupsKeyUpcomingJobsRequest struct {
	Id     string                                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Key    string                                                        `json:"key" jsonschema:"description=The key of the resource group"`
	Params *client.GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsParams `json:"params,omitempty"`
}

func registerGetProjectsIdResourceGroupsKeyUpcomingJobs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdResourceGroupsKeyUpcomingJobsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_resource_grps_key_upcoming_jobs",
		mcp.WithDescription("List upcoming jobs for a specific resource group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdResourceGroupsKeyUpcomingJobsHandler))
}

func getProjectsIdResourceGroupsKeyUpcomingJobsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdResourceGroupsKeyUpcomingJobsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs(ctx, req.Id, req.Key, req.Params, authorizationHeader))
}

type GetProjectsIdRunnersRequest struct {
	Id     string                                  `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdRunnersParams `json:"params,omitempty"`
}

func registerGetProjectsIdRunners(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRunnersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_runners",
		mcp.WithDescription("List all runners available in the project, including from ancestor groups and any allowed shared runners."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRunnersHandler))
}

func getProjectsIdRunnersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRunnersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRunners(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdRunnersRunnerIdRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	RunnerId int32  `json:"runner_id" jsonschema:"description=The ID of a runner"`
}

func registerDeleteProjectsIdRunnersRunnerId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRunnersRunnerIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_runners_runner_id",
		mcp.WithDescription("It is not possible to unassign a runner from the owner project. If so, an error is returned. Use the call to delete a runner instead."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRunnersRunnerIdHandler))
}

func deleteProjectsIdRunnersRunnerIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRunnersRunnerIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRunnersRunnerId(ctx, req.Id, req.RunnerId, authorizationHeader))
}

type PostProjectsIdRunnersResetRegistrationTokenRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a project"`
}

func registerPostProjectsIdRunnersResetRegistrationToken(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdRunnersResetRegistrationTokenRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_runners_reset_registration_token",
		mcp.WithDescription("Reset runner registration token"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdRunnersResetRegistrationTokenHandler))
}

func postProjectsIdRunnersResetRegistrationTokenHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdRunnersResetRegistrationTokenRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdRunnersResetRegistrationToken(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdSecureFilesRequest struct {
	Id     string                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdSecureFilesParams `json:"params,omitempty"`
}

func registerGetProjectsIdSecureFiles(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSecureFilesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_secure_files",
		mcp.WithDescription("Get list of secure files in a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSecureFilesHandler))
}

func getProjectsIdSecureFilesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSecureFilesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSecureFiles(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdSecureFilesSecureFileIdRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	SecureFileId int32  `json:"secure_file_id" jsonschema:"description=null"`
}

func registerDeleteProjectsIdSecureFilesSecureFileId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdSecureFilesSecureFileIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_secure_files_secure_file_id",
		mcp.WithDescription("Remove a secure file"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdSecureFilesSecureFileIdHandler))
}

func deleteProjectsIdSecureFilesSecureFileIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdSecureFilesSecureFileIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdSecureFilesSecureFileId(ctx, req.Id, req.SecureFileId, authorizationHeader))
}

type GetProjectsIdSecureFilesSecureFileIdRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	SecureFileId int32  `json:"secure_file_id" jsonschema:"description=The ID of a secure file"`
}

func registerGetProjectsIdSecureFilesSecureFileId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSecureFilesSecureFileIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_secure_files_secure_file_id",
		mcp.WithDescription("Get the details of a specific secure file in a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSecureFilesSecureFileIdHandler))
}

func getProjectsIdSecureFilesSecureFileIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSecureFilesSecureFileIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSecureFilesSecureFileId(ctx, req.Id, req.SecureFileId, authorizationHeader))
}

type GetProjectsIdSecureFilesSecureFileIdDownloadRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	SecureFileId int32  `json:"secure_file_id" jsonschema:"description=The ID of a secure file"`
}

func registerGetProjectsIdSecureFilesSecureFileIdDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSecureFilesSecureFileIdDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_secure_files_secure_file_id_download",
		mcp.WithDescription("Download secure file"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSecureFilesSecureFileIdDownloadHandler))
}

func getProjectsIdSecureFilesSecureFileIdDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSecureFilesSecureFileIdDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSecureFilesSecureFileIdDownload(ctx, req.Id, req.SecureFileId, authorizationHeader))
}

type GetProjectsIdPipelinesRequest struct {
	Id     string                                    `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	Params *client.GetApiV4ProjectsIdPipelinesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPipelines(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesHandler))
}

func getProjectsIdPipelinesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelines(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdPipelinesLatestRequest struct {
	Id     string                                          `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	Params *client.GetApiV4ProjectsIdPipelinesLatestParams `json:"params,omitempty"`
}

func registerGetProjectsIdPipelinesLatest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesLatestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_latest",
		mcp.WithDescription("This feature was introduced in GitLab 12.3"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesLatestHandler))
}

func getProjectsIdPipelinesLatestHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesLatestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesLatest(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdPipelinesPipelineIdRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerDeleteProjectsIdPipelinesPipelineId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPipelinesPipelineIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pls_pipeline_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPipelinesPipelineIdHandler))
}

func deleteProjectsIdPipelinesPipelineIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPipelinesPipelineIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPipelinesPipelineId(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type GetProjectsIdPipelinesPipelineIdRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerGetProjectsIdPipelinesPipelineId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesPipelineIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesPipelineIdHandler))
}

func getProjectsIdPipelinesPipelineIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesPipelineIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineId(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type GetProjectsIdPipelinesPipelineIdJobsRequest struct {
	Id         string                                                  `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32                                                   `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
	Params     *client.GetApiV4ProjectsIdPipelinesPipelineIdJobsParams `json:"params,omitempty"`
}

func registerGetProjectsIdPipelinesPipelineIdJobs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesPipelineIdJobsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_jobs",
		mcp.WithDescription("Get pipeline jobs"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesPipelineIdJobsHandler))
}

func getProjectsIdPipelinesPipelineIdJobsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesPipelineIdJobsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdJobs(ctx, req.Id, req.PipelineId, req.Params, authorizationHeader))
}

type GetProjectsIdPipelinesPipelineIdBridgesRequest struct {
	Id         string                                                     `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32                                                      `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
	Params     *client.GetApiV4ProjectsIdPipelinesPipelineIdBridgesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPipelinesPipelineIdBridges(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesPipelineIdBridgesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_bridges",
		mcp.WithDescription("Get pipeline bridge jobs"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesPipelineIdBridgesHandler))
}

func getProjectsIdPipelinesPipelineIdBridgesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesPipelineIdBridgesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdBridges(ctx, req.Id, req.PipelineId, req.Params, authorizationHeader))
}

type GetProjectsIdPipelinesPipelineIdVariablesRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerGetProjectsIdPipelinesPipelineIdVariables(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesPipelineIdVariablesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_variables",
		mcp.WithDescription("This feature was introduced in GitLab 11.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesPipelineIdVariablesHandler))
}

func getProjectsIdPipelinesPipelineIdVariablesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesPipelineIdVariablesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdVariables(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type GetProjectsIdPipelinesPipelineIdTestReportRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerGetProjectsIdPipelinesPipelineIdTestReport(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesPipelineIdTestReportRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_test_report",
		mcp.WithDescription("This feature was introduced in GitLab 13.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesPipelineIdTestReportHandler))
}

func getProjectsIdPipelinesPipelineIdTestReportHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesPipelineIdTestReportRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdTestReport(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type GetProjectsIdPipelinesPipelineIdTestReportSummaryRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerGetProjectsIdPipelinesPipelineIdTestReportSummary(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelinesPipelineIdTestReportSummaryRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pls_pipeline_id_test_report_summary",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelinesPipelineIdTestReportSummaryHandler))
}

func getProjectsIdPipelinesPipelineIdTestReportSummaryHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelinesPipelineIdTestReportSummaryRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type PostProjectsIdPipelinesPipelineIdRetryRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerPostProjectsIdPipelinesPipelineIdRetry(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPipelinesPipelineIdRetryRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pls_pipeline_id_retry",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPipelinesPipelineIdRetryHandler))
}

func postProjectsIdPipelinesPipelineIdRetryHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPipelinesPipelineIdRetryRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPipelinesPipelineIdRetry(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type PostProjectsIdPipelinesPipelineIdCancelRequest struct {
	Id         string `json:"id" jsonschema:"description=The project ID or URL-encoded path"`
	PipelineId int32  `json:"pipeline_id" jsonschema:"description=The pipeline ID"`
}

func registerPostProjectsIdPipelinesPipelineIdCancel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPipelinesPipelineIdCancelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pls_pipeline_id_cancel",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPipelinesPipelineIdCancelHandler))
}

func postProjectsIdPipelinesPipelineIdCancelHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPipelinesPipelineIdCancelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPipelinesPipelineIdCancel(ctx, req.Id, req.PipelineId, authorizationHeader))
}

type GetProjectsIdPipelineSchedulesRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdPipelineSchedulesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPipelineSchedules(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelineSchedulesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pipeline_schedules",
		mcp.WithDescription("Get all pipeline schedules"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelineSchedulesHandler))
}

func getProjectsIdPipelineSchedulesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelineSchedulesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelineSchedules(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdPipelineSchedulesPipelineScheduleIdRequest struct {
	Id                 string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PipelineScheduleId int32  `json:"pipeline_schedule_id" jsonschema:"description=The pipeline schedule id"`
}

func registerDeleteProjectsIdPipelineSchedulesPipelineScheduleId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPipelineSchedulesPipelineScheduleIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pipeline_schedules_pipeline_schedule_id",
		mcp.WithDescription("Delete a pipeline schedule"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPipelineSchedulesPipelineScheduleIdHandler))
}

func deleteProjectsIdPipelineSchedulesPipelineScheduleIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPipelineSchedulesPipelineScheduleIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx, req.Id, req.PipelineScheduleId, authorizationHeader))
}

type GetProjectsIdPipelineSchedulesPipelineScheduleIdRequest struct {
	Id                 string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PipelineScheduleId int32  `json:"pipeline_schedule_id" jsonschema:"description=The pipeline schedule id"`
}

func registerGetProjectsIdPipelineSchedulesPipelineScheduleId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelineSchedulesPipelineScheduleIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pipeline_schedules_pipeline_schedule_id",
		mcp.WithDescription("Get a single pipeline schedule"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelineSchedulesPipelineScheduleIdHandler))
}

func getProjectsIdPipelineSchedulesPipelineScheduleIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelineSchedulesPipelineScheduleIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx, req.Id, req.PipelineScheduleId, authorizationHeader))
}

type GetProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesRequest struct {
	Id                 string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PipelineScheduleId int32  `json:"pipeline_schedule_id" jsonschema:"description=The pipeline schedule ID"`
}

func registerGetProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pipeline_schedules_pipeline_schedule_id_pls",
		mcp.WithDescription("Get all pipelines triggered from a pipeline schedule"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesHandler))
}

func getProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(ctx, req.Id, req.PipelineScheduleId, authorizationHeader))
}

type PostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipRequest struct {
	Id                 string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PipelineScheduleId int32  `json:"pipeline_schedule_id" jsonschema:"description=The pipeline schedule id"`
}

func registerPostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pipeline_schedules_pipeline_schedule_id_take_ownership",
		mcp.WithDescription("Take ownership of a pipeline schedule"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipHandler))
}

func postProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(ctx, req.Id, req.PipelineScheduleId, authorizationHeader))
}

type PostProjectsIdPipelineSchedulesPipelineScheduleIdPlayRequest struct {
	Id                 string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PipelineScheduleId int32  `json:"pipeline_schedule_id" jsonschema:"description=The pipeline schedule id"`
}

func registerPostProjectsIdPipelineSchedulesPipelineScheduleIdPlay(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPipelineSchedulesPipelineScheduleIdPlayRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pipeline_schedules_pipeline_schedule_id_play",
		mcp.WithDescription("This feature was added in GitLab 12.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPipelineSchedulesPipelineScheduleIdPlayHandler))
}

func postProjectsIdPipelineSchedulesPipelineScheduleIdPlayHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPipelineSchedulesPipelineScheduleIdPlayRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay(ctx, req.Id, req.PipelineScheduleId, authorizationHeader))
}

type DeleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyRequest struct {
	Id                 string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PipelineScheduleId int32  `json:"pipeline_schedule_id" jsonschema:"description=The pipeline schedule id"`
	Key                string `json:"key" jsonschema:"description=The key of the variable"`
}

func registerDeleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pipeline_schedules_pipeline_schedule_id_variables_key",
		mcp.WithDescription("Delete a pipeline schedule variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyHandler))
}

func deleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(ctx, req.Id, req.PipelineScheduleId, req.Key, authorizationHeader))
}

type GetProjectsIdTriggersRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdTriggersParams `json:"params,omitempty"`
}

func registerGetProjectsIdTriggers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTriggersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_triggers",
		mcp.WithDescription("Get trigger tokens list"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTriggersHandler))
}

func getProjectsIdTriggersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTriggersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTriggers(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdTriggersTriggerIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TriggerId int32  `json:"trigger_id" jsonschema:"description=The trigger token ID"`
}

func registerDeleteProjectsIdTriggersTriggerId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdTriggersTriggerIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_triggers_trigger_id",
		mcp.WithDescription("Delete a trigger token"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdTriggersTriggerIdHandler))
}

func deleteProjectsIdTriggersTriggerIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdTriggersTriggerIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdTriggersTriggerId(ctx, req.Id, req.TriggerId, authorizationHeader))
}

type GetProjectsIdTriggersTriggerIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TriggerId int32  `json:"trigger_id" jsonschema:"description=The trigger token ID"`
}

func registerGetProjectsIdTriggersTriggerId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTriggersTriggerIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_triggers_trigger_id",
		mcp.WithDescription("Get specific trigger token of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTriggersTriggerIdHandler))
}

func getProjectsIdTriggersTriggerIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTriggersTriggerIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTriggersTriggerId(ctx, req.Id, req.TriggerId, authorizationHeader))
}

type GetProjectsIdVariablesRequest struct {
	Id     string                                    `json:"id" jsonschema:"description=The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdVariablesParams `json:"params,omitempty"`
}

func registerGetProjectsIdVariables(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdVariablesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_variables",
		mcp.WithDescription("Get project variables"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdVariablesHandler))
}

func getProjectsIdVariablesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdVariablesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdVariables(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdVariablesKeyRequest struct {
	Id     string                                          `json:"id" jsonschema:"description=The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user"`
	Key    string                                          `json:"key" jsonschema:"description=The key of a variable"`
	Params *client.DeleteApiV4ProjectsIdVariablesKeyParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_variables_key",
		mcp.WithDescription("Delete an existing variable from a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdVariablesKeyHandler))
}

func deleteProjectsIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdVariablesKey(ctx, req.Id, req.Key, req.Params, authorizationHeader))
}

type GetProjectsIdVariablesKeyRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID of a project or URL-encoded NAMESPACE/PROJECT_NAME of the project owned by the authenticated user"`
	Key    string                                       `json:"key" jsonschema:"description=The key of a variable"`
	Params *client.GetApiV4ProjectsIdVariablesKeyParams `json:"params,omitempty"`
}

func registerGetProjectsIdVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_variables_key",
		mcp.WithDescription("Get the details of a single variable from a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdVariablesKeyHandler))
}

func getProjectsIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdVariablesKey(ctx, req.Id, req.Key, req.Params, authorizationHeader))
}

type GetProjectsIdClusterAgentsAgentIdTokensRequest struct {
	Id      string                                                     `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AgentId int32                                                      `json:"agent_id" jsonschema:"description=The ID of an agent"`
	Params  *client.GetApiV4ProjectsIdClusterAgentsAgentIdTokensParams `json:"params,omitempty"`
}

func registerGetProjectsIdClusterAgentsAgentIdTokens(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdClusterAgentsAgentIdTokensRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_cluster_agents_agent_id_tokens",
		mcp.WithDescription("This feature was introduced in GitLab 15.0. Returns a list of tokens for an agent."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdClusterAgentsAgentIdTokensHandler))
}

func getProjectsIdClusterAgentsAgentIdTokensHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdClusterAgentsAgentIdTokensRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdClusterAgentsAgentIdTokens(ctx, req.Id, req.AgentId, req.Params, authorizationHeader))
}

type DeleteProjectsIdClusterAgentsAgentIdTokensTokenIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AgentId int32  `json:"agent_id" jsonschema:"description=The ID of an agent"`
	TokenId int32  `json:"token_id" jsonschema:"description=The ID of the agent token"`
}

func registerDeleteProjectsIdClusterAgentsAgentIdTokensTokenId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdClusterAgentsAgentIdTokensTokenIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_cluster_agents_agent_id_tokens_token_id",
		mcp.WithDescription("This feature was introduced in GitLab 15.0. Revokes an agent token."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdClusterAgentsAgentIdTokensTokenIdHandler))
}

func deleteProjectsIdClusterAgentsAgentIdTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdClusterAgentsAgentIdTokensTokenIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx, req.Id, req.AgentId, req.TokenId, authorizationHeader))
}

type GetProjectsIdClusterAgentsAgentIdTokensTokenIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AgentId int32  `json:"agent_id" jsonschema:"description=The ID of an agent"`
	TokenId int32  `json:"token_id" jsonschema:"description=The ID of the agent token"`
}

func registerGetProjectsIdClusterAgentsAgentIdTokensTokenId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdClusterAgentsAgentIdTokensTokenIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_cluster_agents_agent_id_tokens_token_id",
		mcp.WithDescription("This feature was introduced in GitLab 15.0. Gets a single agent token."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdClusterAgentsAgentIdTokensTokenIdHandler))
}

func getProjectsIdClusterAgentsAgentIdTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdClusterAgentsAgentIdTokensTokenIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx, req.Id, req.AgentId, req.TokenId, authorizationHeader))
}

type GetProjectsIdClusterAgentsRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdClusterAgentsParams `json:"params,omitempty"`
}

func registerGetProjectsIdClusterAgents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdClusterAgentsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_cluster_agents",
		mcp.WithDescription("This feature was introduced in GitLab 14.10. Returns the list of agents registered for the project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdClusterAgentsHandler))
}

func getProjectsIdClusterAgentsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdClusterAgentsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdClusterAgents(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdClusterAgentsAgentIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AgentId int32  `json:"agent_id" jsonschema:"description=The ID of an agent"`
}

func registerDeleteProjectsIdClusterAgentsAgentId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdClusterAgentsAgentIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_cluster_agents_agent_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.10. Deletes an existing agent registration."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdClusterAgentsAgentIdHandler))
}

func deleteProjectsIdClusterAgentsAgentIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdClusterAgentsAgentIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdClusterAgentsAgentId(ctx, req.Id, req.AgentId, authorizationHeader))
}

type GetProjectsIdClusterAgentsAgentIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	AgentId int32  `json:"agent_id" jsonschema:"description=The ID of an agent"`
}

func registerGetProjectsIdClusterAgentsAgentId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdClusterAgentsAgentIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_cluster_agents_agent_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.10. Gets a single agent details."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdClusterAgentsAgentIdHandler))
}

func getProjectsIdClusterAgentsAgentIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdClusterAgentsAgentIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdClusterAgentsAgentId(ctx, req.Id, req.AgentId, authorizationHeader))
}

type GetProjectsIdPackagesCargoConfigJsonRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesCargoConfigJson(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesCargoConfigJsonRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_cargo_config_json",
		mcp.WithDescription("This will be used by cargo for further requests"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesCargoConfigJsonHandler))
}

func getProjectsIdPackagesCargoConfigJsonHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesCargoConfigJsonRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesCargoConfigJson(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommits(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits",
		mcp.WithDescription("Get a project repository commits"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsHandler))
}

func getProjectsIdRepositoryCommitsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommits(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaRequest struct {
	Id     string                                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha    string                                               `json:"sha" jsonschema:"description=A commit sha, or the name of a branch or tag"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsSha(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha",
		mcp.WithDescription("Get a specific commit of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaHandler))
}

func getProjectsIdRepositoryCommitsShaHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsSha(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaDiffRequest struct {
	Id     string                                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha    string                                                   `json:"sha" jsonschema:"description=A commit sha, or the name of a branch or tag"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaDiffParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsShaDiff(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaDiffRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_diff",
		mcp.WithDescription("Get the diff for a specific commit of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaDiffHandler))
}

func getProjectsIdRepositoryCommitsShaDiffHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaDiffRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaDiff(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaCommentsRequest struct {
	Id     string                                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha    string                                                       `json:"sha" jsonschema:"description=A commit sha, or the name of a branch or tag"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaCommentsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsShaComments(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaCommentsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_comments",
		mcp.WithDescription("Get a commit's comments"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaCommentsHandler))
}

func getProjectsIdRepositoryCommitsShaCommentsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaCommentsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaComments(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaSequenceRequest struct {
	Id     string                                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha    string                                                       `json:"sha" jsonschema:"description=A commit SHA"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaSequenceParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsShaSequence(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaSequenceRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_sequence",
		mcp.WithDescription("Get the sequence count of a commit SHA"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaSequenceHandler))
}

func getProjectsIdRepositoryCommitsShaSequenceHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaSequenceRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaSequence(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaRefsRequest struct {
	Id     string                                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha    string                                                   `json:"sha" jsonschema:"description=A commit sha"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaRefsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsShaRefs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaRefsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_refs",
		mcp.WithDescription("This feature was introduced in GitLab 10.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaRefsHandler))
}

func getProjectsIdRepositoryCommitsShaRefsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaRefsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaRefs(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaMergeRequestsRequest struct {
	Id     string                                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha    string                                                            `json:"sha" jsonschema:"description=A commit sha, or the name of a branch or tag on which to find Merge Requests"`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsShaMergeRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaMergeRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_mrs",
		mcp.WithDescription("Get Merge Requests associated with a commit"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaMergeRequestsHandler))
}

func getProjectsIdRepositoryCommitsShaMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaMergeRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaSignatureRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha string `json:"sha" jsonschema:"description=A commit sha, or the name of a branch or tag"`
}

func registerGetProjectsIdRepositoryCommitsShaSignature(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaSignatureRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_signature",
		mcp.WithDescription("Get a commit's signature"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaSignatureHandler))
}

func getProjectsIdRepositoryCommitsShaSignatureHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaSignatureRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaSignature(ctx, req.Id, req.Sha, authorizationHeader))
}

type GetProjectsIdRepositoryCommitsShaStatusesRequest struct {
	Id     string                                                       `json:"id" jsonschema:"description=ID or URL-encoded path of the project."`
	Sha    string                                                       `json:"sha" jsonschema:"description=Hash of the commit."`
	Params *client.GetApiV4ProjectsIdRepositoryCommitsShaStatusesParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCommitsShaStatuses(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCommitsShaStatusesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_commits_sha_statuses",
		mcp.WithDescription("Get a commit's statuses"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCommitsShaStatusesHandler))
}

func getProjectsIdRepositoryCommitsShaStatusesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCommitsShaStatusesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCommitsShaStatuses(ctx, req.Id, req.Sha, req.Params, authorizationHeader))
}

type GetProjectsIdPackagesConanV1UsersAuthenticateRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesConanV1UsersAuthenticate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1UsersAuthenticateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_users_authenticate",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1UsersAuthenticateHandler))
}

func getProjectsIdPackagesConanV1UsersAuthenticateHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1UsersAuthenticateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesConanV1UsersCheckCredentialsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesConanV1UsersCheckCredentials(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1UsersCheckCredentialsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_users_check_credentials",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1UsersCheckCredentialsHandler))
}

func getProjectsIdPackagesConanV1UsersCheckCredentialsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1UsersCheckCredentialsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansSearchRequest struct {
	Id     string                                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdPackagesConanV1ConansSearchParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesConanV1ConansSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_search",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansSearchHandler))
}

func getProjectsIdPackagesConanV1ConansSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansSearch(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV1PingRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesConanV1Ping(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1PingRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_ping",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1PingHandler))
}

func getProjectsIdPackagesConanV1PingHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1PingRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1Ping(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type DeleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerDeleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler))
}

func deleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler))
}

func getProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type PostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerPostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_upload_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsHandler))
}

func postProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type PostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerPostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_upload_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsHandler))
}

func postProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	FileName        string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_export_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler))
}

func getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.FileName, authorizationHeader))
}

type PutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	FileName        string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_export_file_name_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeHandler))
}

func putProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.FileName, authorizationHeader))
}

type GetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan Package ID"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Conan Package Revision"`
	FileName              string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_package_conan_package_reference_package_revision_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler))
}

func getProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, req.FileName, authorizationHeader))
}

type PutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan Package ID"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Conan Package Revision"`
	FileName              string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_package_conan_package_reference_package_revision_file_name_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeHandler))
}

func putProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, req.FileName, authorizationHeader))
}

type GetProjectsIdPackagesConanV2UsersAuthenticateRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesConanV2UsersAuthenticate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2UsersAuthenticateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_users_authenticate",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2UsersAuthenticateHandler))
}

func getProjectsIdPackagesConanV2UsersAuthenticateHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2UsersAuthenticateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2UsersAuthenticate(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesConanV2UsersCheckCredentialsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesConanV2UsersCheckCredentials(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2UsersCheckCredentialsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_users_check_credentials",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2UsersCheckCredentialsHandler))
}

func getProjectsIdPackagesConanV2UsersCheckCredentialsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2UsersCheckCredentialsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentials(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansSearchRequest struct {
	Id     string                                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdPackagesConanV2ConansSearchParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesConanV2ConansSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_search",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansSearchHandler))
}

func getProjectsIdPackagesConanV2ConansSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansSearch(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_latest",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type DeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
}

func registerDeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevision(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision",
		mcp.WithDescription("This feature was introduced in GitLab 18.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionHandler))
}

func deleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevision(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_files",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	FileName        string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_files_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 17.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.FileName, authorizationHeader))
}

type PutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	FileName        string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_files_file_name_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 17.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeHandler))
}

func putProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.FileName, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Package reference"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_latest",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Package reference"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, authorizationHeader))
}

type DeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Package reference"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Package revision"`
}

func registerDeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevision(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions_package_revision",
		mcp.WithDescription("This feature was introduced in GitLab 18.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionHandler))
}

func deleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevision(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Package reference"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Package revision"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions_package_revision_files",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, authorizationHeader))
}

type GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Package reference"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Package revision"`
	FileName              string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions_package_revision_files_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameHandler))
}

func getProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, req.FileName, authorizationHeader))
}

type PutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeRequest struct {
	Id                    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Recipe revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Package reference"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Package revision"`
	FileName              string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_conan_v2_conans_package_name_package_version_package_username_package_channel_revisions_recipe_revision_packages_conan_package_reference_revisions_package_revision_files_file_name_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 17.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeHandler))
}

func putProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize(ctx, req.Id, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, req.FileName, authorizationHeader))
}

type GetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameRequest struct {
	Id             string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Distribution   string `json:"distribution" jsonschema:"description=The Debian Codename or Suite"`
	Letter         string `json:"letter" jsonschema:"description=The Debian Classification (first-letter or lib-first-letter)"`
	PackageName    string `json:"package_name" jsonschema:"description=The Debian Source Package Name"`
	PackageVersion string `json:"package_version" jsonschema:"description=The Debian Source Package Version"`
	FileName       string `json:"file_name" jsonschema:"description=The Debian File Name"`
}

func registerGetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_debian_pool_distribution_letter_package_name_package_version_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameHandler))
}

func getProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(ctx, req.Id, req.Distribution, req.Letter, req.PackageName, req.PackageVersion, req.FileName, authorizationHeader))
}

type GetProjectsIdDeployKeysRequest struct {
	Id     string                                     `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdDeployKeysParams `json:"params,omitempty"`
}

func registerGetProjectsIdDeployKeys(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeployKeysRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deploy_keys",
		mcp.WithDescription("Get a list of a project's deploy keys."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeployKeysHandler))
}

func getProjectsIdDeployKeysHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeployKeysRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeployKeys(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdDeployKeysKeyIdRequest struct {
	Id    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	KeyId int32  `json:"key_id" jsonschema:"description=The ID of the deploy key"`
}

func registerDeleteProjectsIdDeployKeysKeyId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdDeployKeysKeyIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_deploy_keys_key_id",
		mcp.WithDescription("Removes a deploy key from the project. If the deploy key is used only for this project, it's deleted from the system."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdDeployKeysKeyIdHandler))
}

func deleteProjectsIdDeployKeysKeyIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdDeployKeysKeyIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdDeployKeysKeyId(ctx, req.Id, req.KeyId, authorizationHeader))
}

type GetProjectsIdDeployKeysKeyIdRequest struct {
	Id    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	KeyId int32  `json:"key_id" jsonschema:"description=The ID of the deploy key"`
}

func registerGetProjectsIdDeployKeysKeyId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeployKeysKeyIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deploy_keys_key_id",
		mcp.WithDescription("Get a single key."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeployKeysKeyIdHandler))
}

func getProjectsIdDeployKeysKeyIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeployKeysKeyIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeployKeysKeyId(ctx, req.Id, req.KeyId, authorizationHeader))
}

type PostProjectsIdDeployKeysKeyIdEnableRequest struct {
	Id    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	KeyId int32  `json:"key_id" jsonschema:"description=The ID of the deploy key"`
}

func registerPostProjectsIdDeployKeysKeyIdEnable(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdDeployKeysKeyIdEnableRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_deploy_keys_key_id_enable",
		mcp.WithDescription("Enables a deploy key for a project so this can be used. Returns the enabled key, with a status code 201 when successful. This feature was added in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdDeployKeysKeyIdEnableHandler))
}

func postProjectsIdDeployKeysKeyIdEnableHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdDeployKeysKeyIdEnableRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdDeployKeysKeyIdEnable(ctx, req.Id, req.KeyId, authorizationHeader))
}

type GetProjectsIdDeployTokensRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdDeployTokensParams `json:"params,omitempty"`
}

func registerGetProjectsIdDeployTokens(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeployTokensRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deploy_tokens",
		mcp.WithDescription("Get a list of a project's deploy tokens. This feature was introduced in GitLab 12.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeployTokensHandler))
}

func getProjectsIdDeployTokensHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeployTokensRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeployTokens(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdDeployTokensTokenIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	TokenId int32  `json:"token_id" jsonschema:"description=The ID of the deploy token"`
}

func registerDeleteProjectsIdDeployTokensTokenId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdDeployTokensTokenIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_deploy_tokens_token_id",
		mcp.WithDescription("This feature was introduced in GitLab 12.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdDeployTokensTokenIdHandler))
}

func deleteProjectsIdDeployTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdDeployTokensTokenIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdDeployTokensTokenId(ctx, req.Id, req.TokenId, authorizationHeader))
}

type GetProjectsIdDeployTokensTokenIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	TokenId int32  `json:"token_id" jsonschema:"description=The ID of the deploy token"`
}

func registerGetProjectsIdDeployTokensTokenId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeployTokensTokenIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deploy_tokens_token_id",
		mcp.WithDescription("Get a single project's deploy token by ID. This feature was introduced in GitLab 14.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeployTokensTokenIdHandler))
}

func getProjectsIdDeployTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeployTokensTokenIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeployTokensTokenId(ctx, req.Id, req.TokenId, authorizationHeader))
}

type GetProjectsIdDeploymentsRequest struct {
	Id     string                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdDeploymentsParams `json:"params,omitempty"`
}

func registerGetProjectsIdDeployments(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeploymentsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deployments",
		mcp.WithDescription("Get a list of deployments in a project. This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeploymentsHandler))
}

func getProjectsIdDeploymentsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeploymentsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeployments(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdDeploymentsDeploymentIdRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	DeploymentId int32  `json:"deployment_id" jsonschema:"description=The ID of the deployment"`
}

func registerDeleteProjectsIdDeploymentsDeploymentId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdDeploymentsDeploymentIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_deployments_deployment_id",
		mcp.WithDescription("Delete a specific deployment that is not currently the last deployment for an environment or in a running state. This feature was introduced in GitLab 15.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdDeploymentsDeploymentIdHandler))
}

func deleteProjectsIdDeploymentsDeploymentIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdDeploymentsDeploymentIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdDeploymentsDeploymentId(ctx, req.Id, req.DeploymentId, authorizationHeader))
}

type GetProjectsIdDeploymentsDeploymentIdRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	DeploymentId int32  `json:"deployment_id" jsonschema:"description=The ID of the deployment"`
}

func registerGetProjectsIdDeploymentsDeploymentId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeploymentsDeploymentIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deployments_deployment_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeploymentsDeploymentIdHandler))
}

func getProjectsIdDeploymentsDeploymentIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeploymentsDeploymentIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeploymentsDeploymentId(ctx, req.Id, req.DeploymentId, authorizationHeader))
}

type GetProjectsIdDeploymentsDeploymentIdMergeRequestsRequest struct {
	Id           string                                                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	DeploymentId int32                                                                `json:"deployment_id" jsonschema:"description=The ID of the deployment"`
	Params       *client.GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsParams `json:"params,omitempty"`
}

func registerGetProjectsIdDeploymentsDeploymentIdMergeRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDeploymentsDeploymentIdMergeRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_deployments_deployment_id_mrs",
		mcp.WithDescription("Retrieves the list of merge requests shipped with a given deployment. This feature was introduced in GitLab 12.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDeploymentsDeploymentIdMergeRequestsHandler))
}

func getProjectsIdDeploymentsDeploymentIdMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDeploymentsDeploymentIdMergeRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests(ctx, req.Id, req.DeploymentId, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidDraftNotesRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID of a project"`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The ID of a merge request"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotes(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidDraftNotesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_draft_notes",
		mcp.WithDescription("Get a list of merge request draft notes"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidDraftNotesHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidDraftNotesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidDraftNotesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type DeleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID of a project"`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The ID of a merge request"`
	DraftNoteId     int32  `json:"draft_note_id" jsonschema:"description=The ID of a draft note"`
}

func registerDeleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_mrs_merge_request_iid_draft_notes_draft_note_id",
		mcp.WithDescription("Delete a draft note"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdHandler))
}

func deleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx, req.Id, req.MergeRequestIid, req.DraftNoteId, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID of a project"`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The ID of a merge request"`
	DraftNoteId     int32  `json:"draft_note_id" jsonschema:"description=The ID of a draft note"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_draft_notes_draft_note_id",
		mcp.WithDescription("Get a single draft note"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx, req.Id, req.MergeRequestIid, req.DraftNoteId, authorizationHeader))
}

type PutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID of a project"`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The ID of a merge request"`
	DraftNoteId     int32  `json:"draft_note_id" jsonschema:"description=The ID of a draft note"`
}

func registerPutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_mrs_merge_request_iid_draft_notes_draft_note_id_publish",
		mcp.WithDescription("Publish a pending draft note"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishHandler))
}

func putProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(ctx, req.Id, req.MergeRequestIid, req.DraftNoteId, authorizationHeader))
}

type PostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID of a project"`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The ID of a merge request"`
}

func registerPostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_mrs_merge_request_iid_draft_notes_bulk_publish",
		mcp.WithDescription("Bulk publish all pending draft notes"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishHandler))
}

func postProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdEnvironmentsRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdEnvironmentsParams `json:"params,omitempty"`
}

func registerGetProjectsIdEnvironments(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdEnvironmentsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_environments",
		mcp.WithDescription("Get all environments for a given project. This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdEnvironmentsHandler))
}

func getProjectsIdEnvironmentsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdEnvironmentsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdEnvironments(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdEnvironmentsEnvironmentIdRequest struct {
	Id            string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	EnvironmentId int32  `json:"environment_id" jsonschema:"description=The ID of the environment"`
}

func registerDeleteProjectsIdEnvironmentsEnvironmentId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdEnvironmentsEnvironmentIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_environments_environment_id",
		mcp.WithDescription("It returns 204 if the environment was successfully deleted, and 404 if the environment does not exist. This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdEnvironmentsEnvironmentIdHandler))
}

func deleteProjectsIdEnvironmentsEnvironmentIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdEnvironmentsEnvironmentIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdEnvironmentsEnvironmentId(ctx, req.Id, req.EnvironmentId, authorizationHeader))
}

type GetProjectsIdEnvironmentsEnvironmentIdRequest struct {
	Id            string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	EnvironmentId int32  `json:"environment_id" jsonschema:"description=The ID of the environment"`
}

func registerGetProjectsIdEnvironmentsEnvironmentId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdEnvironmentsEnvironmentIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_environments_environment_id",
		mcp.WithDescription("Get a specific environment"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdEnvironmentsEnvironmentIdHandler))
}

func getProjectsIdEnvironmentsEnvironmentIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdEnvironmentsEnvironmentIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdEnvironmentsEnvironmentId(ctx, req.Id, req.EnvironmentId, authorizationHeader))
}

type DeleteProjectsIdEnvironmentsReviewAppsRequest struct {
	Id     string                                                    `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.DeleteApiV4ProjectsIdEnvironmentsReviewAppsParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdEnvironmentsReviewApps(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdEnvironmentsReviewAppsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_environments_review_apps",
		mcp.WithDescription("It schedules for deletion multiple environments that have already been stopped and are in the review app folder. The actual deletion is performed after 1 week from the time of execution. By default, it only deletes environments 30 days or older. You can change this default using the `before` parameter."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdEnvironmentsReviewAppsHandler))
}

func deleteProjectsIdEnvironmentsReviewAppsHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdEnvironmentsReviewAppsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdEnvironmentsReviewApps(ctx, req.Id, req.Params, authorizationHeader))
}

type PostProjectsIdErrorTrackingClientKeysRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
}

func registerPostProjectsIdErrorTrackingClientKeys(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdErrorTrackingClientKeysRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_error_tracking_client_keys",
		mcp.WithDescription("Creates a new client key for a project. The public key attribute is generated automatically.This feature was introduced in GitLab 14.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdErrorTrackingClientKeysHandler))
}

func postProjectsIdErrorTrackingClientKeysHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdErrorTrackingClientKeysRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdErrorTrackingClientKeys(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdErrorTrackingClientKeysRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
}

func registerGetProjectsIdErrorTrackingClientKeys(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdErrorTrackingClientKeysRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_error_tracking_client_keys",
		mcp.WithDescription("List all client keys. This feature was introduced in GitLab 14.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdErrorTrackingClientKeysHandler))
}

func getProjectsIdErrorTrackingClientKeysHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdErrorTrackingClientKeysRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdErrorTrackingClientKeys(ctx, req.Id, authorizationHeader))
}

type DeleteProjectsIdErrorTrackingClientKeysKeyIdRequest struct {
	Id    string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	KeyId int32  `json:"key_id" jsonschema:"description=null"`
}

func registerDeleteProjectsIdErrorTrackingClientKeysKeyId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdErrorTrackingClientKeysKeyIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_error_tracking_client_keys_key_id",
		mcp.WithDescription("Removes a client key from the project. This feature was introduced in GitLab 14.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdErrorTrackingClientKeysKeyIdHandler))
}

func deleteProjectsIdErrorTrackingClientKeysKeyIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdErrorTrackingClientKeysKeyIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId(ctx, req.Id, req.KeyId, authorizationHeader))
}

type GetProjectsIdErrorTrackingSettingsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
}

func registerGetProjectsIdErrorTrackingSettings(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdErrorTrackingSettingsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_error_tracking_settings",
		mcp.WithDescription("Get error tracking settings for the project. This feature was introduced in GitLab 12.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdErrorTrackingSettingsHandler))
}

func getProjectsIdErrorTrackingSettingsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdErrorTrackingSettingsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdErrorTrackingSettings(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdFeatureFlagsRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdFeatureFlagsParams `json:"params,omitempty"`
}

func registerGetProjectsIdFeatureFlags(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdFeatureFlagsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_feature_flags",
		mcp.WithDescription("Gets all feature flags of the requested project. This feature was introduced in GitLab 12.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdFeatureFlagsHandler))
}

func getProjectsIdFeatureFlagsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdFeatureFlagsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdFeatureFlags(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdFeatureFlagsFeatureFlagNameRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FeatureFlagName string `json:"feature_flag_name" jsonschema:"description=The name of the feature flag"`
}

func registerDeleteProjectsIdFeatureFlagsFeatureFlagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdFeatureFlagsFeatureFlagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_feature_flags_feature_flag_name",
		mcp.WithDescription("Deletes a feature flag. This feature was introduced in GitLab 12.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdFeatureFlagsFeatureFlagNameHandler))
}

func deleteProjectsIdFeatureFlagsFeatureFlagNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdFeatureFlagsFeatureFlagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx, req.Id, req.FeatureFlagName, authorizationHeader))
}

type GetProjectsIdFeatureFlagsFeatureFlagNameRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FeatureFlagName string `json:"feature_flag_name" jsonschema:"description=The name of the feature flag"`
}

func registerGetProjectsIdFeatureFlagsFeatureFlagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdFeatureFlagsFeatureFlagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_feature_flags_feature_flag_name",
		mcp.WithDescription("Gets a single feature flag. This feature was introduced in GitLab 12.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdFeatureFlagsFeatureFlagNameHandler))
}

func getProjectsIdFeatureFlagsFeatureFlagNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdFeatureFlagsFeatureFlagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx, req.Id, req.FeatureFlagName, authorizationHeader))
}

type GetProjectsIdFeatureFlagsUserListsRequest struct {
	Id     string                                                `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdFeatureFlagsUserListsParams `json:"params,omitempty"`
}

func registerGetProjectsIdFeatureFlagsUserLists(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdFeatureFlagsUserListsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_feature_flags_user_lists",
		mcp.WithDescription("Gets all feature flag user lists for the requested project. This feature was introduced in GitLab 12.10."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdFeatureFlagsUserListsHandler))
}

func getProjectsIdFeatureFlagsUserListsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdFeatureFlagsUserListsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdFeatureFlagsUserLists(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdFeatureFlagsUserListsIidRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Iid string `json:"iid" jsonschema:"description=The internal ID of the project's feature flag user list"`
}

func registerDeleteProjectsIdFeatureFlagsUserListsIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdFeatureFlagsUserListsIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_feature_flags_user_lists_iid",
		mcp.WithDescription("Deletes a feature flag user list. This feature was introduced in GitLab 12.10."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdFeatureFlagsUserListsIidHandler))
}

func deleteProjectsIdFeatureFlagsUserListsIidHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdFeatureFlagsUserListsIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdFeatureFlagsUserListsIid(ctx, req.Id, req.Iid, authorizationHeader))
}

type GetProjectsIdFeatureFlagsUserListsIidRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Iid string `json:"iid" jsonschema:"description=The internal ID of the project's feature flag user list"`
}

func registerGetProjectsIdFeatureFlagsUserListsIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdFeatureFlagsUserListsIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_feature_flags_user_lists_iid",
		mcp.WithDescription("Gets a feature flag user list. This feature was introduced in GitLab 12.10."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdFeatureFlagsUserListsIidHandler))
}

func getProjectsIdFeatureFlagsUserListsIidHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdFeatureFlagsUserListsIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdFeatureFlagsUserListsIid(ctx, req.Id, req.Iid, authorizationHeader))
}

type GetProjectsIdRepositoryFilesFilePathBlameRequest struct {
	Id       string                                                       `json:"id" jsonschema:"description=The project ID"`
	FilePath string                                                       `json:"file_path" jsonschema:"description=The url encoded path to the file."`
	Params   *client.GetApiV4ProjectsIdRepositoryFilesFilePathBlameParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryFilesFilePathBlame(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryFilesFilePathBlameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_files_file_path_blame",
		mcp.WithDescription("Get blame file from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryFilesFilePathBlameHandler))
}

func getProjectsIdRepositoryFilesFilePathBlameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryFilesFilePathBlameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryFilesFilePathBlame(ctx, req.Id, req.FilePath, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryFilesFilePathRawRequest struct {
	Id       string                                                     `json:"id" jsonschema:"description=The project ID"`
	FilePath string                                                     `json:"file_path" jsonschema:"description=The url encoded path to the file."`
	Params   *client.GetApiV4ProjectsIdRepositoryFilesFilePathRawParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryFilesFilePathRaw(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryFilesFilePathRawRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_files_file_path_raw",
		mcp.WithDescription("Get raw file contents from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryFilesFilePathRawHandler))
}

func getProjectsIdRepositoryFilesFilePathRawHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryFilesFilePathRawRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryFilesFilePathRaw(ctx, req.Id, req.FilePath, req.Params, authorizationHeader))
}

type DeleteProjectsIdRepositoryFilesFilePathRequest struct {
	Id       string                                                     `json:"id" jsonschema:"description=The project ID"`
	FilePath string                                                     `json:"file_path" jsonschema:"description=The url encoded path to the file."`
	Params   *client.DeleteApiV4ProjectsIdRepositoryFilesFilePathParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdRepositoryFilesFilePath(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRepositoryFilesFilePathRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_repo_files_file_path",
		mcp.WithDescription("Delete an existing file in repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRepositoryFilesFilePathHandler))
}

func deleteProjectsIdRepositoryFilesFilePathHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRepositoryFilesFilePathRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRepositoryFilesFilePath(ctx, req.Id, req.FilePath, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryFilesFilePathRequest struct {
	Id       string                                                  `json:"id" jsonschema:"description=The project ID"`
	FilePath string                                                  `json:"file_path" jsonschema:"description=The url encoded path to the file."`
	Params   *client.GetApiV4ProjectsIdRepositoryFilesFilePathParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryFilesFilePath(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryFilesFilePathRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_files_file_path",
		mcp.WithDescription("Get a file from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryFilesFilePathHandler))
}

func getProjectsIdRepositoryFilesFilePathHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryFilesFilePathRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryFilesFilePath(ctx, req.Id, req.FilePath, req.Params, authorizationHeader))
}

type GetProjectsIdFreezePeriodsRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdFreezePeriodsParams `json:"params,omitempty"`
}

func registerGetProjectsIdFreezePeriods(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdFreezePeriodsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_freeze_periods",
		mcp.WithDescription("Paginated list of Freeze Periods, sorted by created_at in ascending order. This feature was introduced in GitLab 13.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdFreezePeriodsHandler))
}

func getProjectsIdFreezePeriodsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdFreezePeriodsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdFreezePeriods(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdFreezePeriodsFreezePeriodIdRequest struct {
	Id             string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FreezePeriodId int32  `json:"freeze_period_id" jsonschema:"description=The ID of the freeze period"`
}

func registerDeleteProjectsIdFreezePeriodsFreezePeriodId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdFreezePeriodsFreezePeriodIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_freeze_periods_freeze_period_id",
		mcp.WithDescription("Deletes a freeze period for the given `freeze_period_id`. This feature was introduced in GitLab 13.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdFreezePeriodsFreezePeriodIdHandler))
}

func deleteProjectsIdFreezePeriodsFreezePeriodIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdFreezePeriodsFreezePeriodIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx, req.Id, req.FreezePeriodId, authorizationHeader))
}

type GetProjectsIdFreezePeriodsFreezePeriodIdRequest struct {
	Id             string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FreezePeriodId int32  `json:"freeze_period_id" jsonschema:"description=The ID of the freeze period"`
}

func registerGetProjectsIdFreezePeriodsFreezePeriodId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdFreezePeriodsFreezePeriodIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_freeze_periods_freeze_period_id",
		mcp.WithDescription("Get a freeze period for the given `freeze_period_id`. This feature was introduced in GitLab 13.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdFreezePeriodsFreezePeriodIdHandler))
}

func getProjectsIdFreezePeriodsFreezePeriodIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdFreezePeriodsFreezePeriodIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx, req.Id, req.FreezePeriodId, authorizationHeader))
}

type GetProjectsIdPackagesHelmChannelIndexYamlRequest struct {
	Id      int32  `json:"id" jsonschema:"description=The ID or full path of a project"`
	Channel string `json:"channel" jsonschema:"description=Helm channel"`
}

func registerGetProjectsIdPackagesHelmChannelIndexYaml(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesHelmChannelIndexYamlRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_helm_channel_index_yaml",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesHelmChannelIndexYamlHandler))
}

func getProjectsIdPackagesHelmChannelIndexYamlHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesHelmChannelIndexYamlRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesHelmChannelIndexYaml(ctx, req.Id, req.Channel, authorizationHeader))
}

type GetProjectsIdPackagesHelmChannelChartsFileNameTgzRequest struct {
	Id       int32  `json:"id" jsonschema:"description=The ID or full path of a project"`
	Channel  string `json:"channel" jsonschema:"description=Helm channel"`
	FileName string `json:"file_name" jsonschema:"description=Helm package file name"`
}

func registerGetProjectsIdPackagesHelmChannelChartsFileNameTgz(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesHelmChannelChartsFileNameTgzRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_helm_channel_charts_file_name_tgz",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesHelmChannelChartsFileNameTgzHandler))
}

func getProjectsIdPackagesHelmChannelChartsFileNameTgzHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesHelmChannelChartsFileNameTgzRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz(ctx, req.Id, req.Channel, req.FileName, authorizationHeader))
}

type PostProjectsIdPackagesHelmApiChannelChartsAuthorizeRequest struct {
	Id      int32  `json:"id" jsonschema:"description=The ID or full path of a project"`
	Channel string `json:"channel" jsonschema:"description=Helm channel"`
}

func registerPostProjectsIdPackagesHelmApiChannelChartsAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesHelmApiChannelChartsAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_helm_api_channel_charts_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesHelmApiChannelChartsAuthorizeHandler))
}

func postProjectsIdPackagesHelmApiChannelChartsAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesHelmApiChannelChartsAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize(ctx, req.Id, req.Channel, authorizationHeader))
}

type GetProjectsIdServicesRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetProjectsIdServices(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdServicesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_services",
		mcp.WithDescription("Get a list of all active integrations."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdServicesHandler))
}

func getProjectsIdServicesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdServicesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdServices(ctx, req.Id, authorizationHeader))
}

type DeleteProjectsIdServicesSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The name of the integration"`
}

func registerDeleteProjectsIdServicesSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdServicesSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_services_slug",
		mcp.WithDescription("Disable the integration. Integration settings are preserved."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdServicesSlugHandler))
}

func deleteProjectsIdServicesSlugHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdServicesSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdServicesSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetProjectsIdServicesSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The name of the integration"`
}

func registerGetProjectsIdServicesSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdServicesSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_services_slug",
		mcp.WithDescription("Get the integration settings."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdServicesSlugHandler))
}

func getProjectsIdServicesSlugHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdServicesSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdServicesSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetProjectsIdIntegrationsRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetProjectsIdIntegrations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIntegrationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_integrations",
		mcp.WithDescription("Get a list of all active integrations."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIntegrationsHandler))
}

func getProjectsIdIntegrationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIntegrationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIntegrations(ctx, req.Id, authorizationHeader))
}

type DeleteProjectsIdIntegrationsSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The name of the integration"`
}

func registerDeleteProjectsIdIntegrationsSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdIntegrationsSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_integrations_slug",
		mcp.WithDescription("Disable the integration. Integration settings are preserved."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdIntegrationsSlugHandler))
}

func deleteProjectsIdIntegrationsSlugHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdIntegrationsSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdIntegrationsSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetProjectsIdIntegrationsSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The name of the integration"`
}

func registerGetProjectsIdIntegrationsSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIntegrationsSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_integrations_slug",
		mcp.WithDescription("Get the integration settings."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIntegrationsSlugHandler))
}

func getProjectsIdIntegrationsSlugHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIntegrationsSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIntegrationsSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetProjectsIdInvitationsRequest struct {
	Id     string                                      `json:"id" jsonschema:"description=The project ID"`
	Params *client.GetApiV4ProjectsIdInvitationsParams `json:"params,omitempty"`
}

func registerGetProjectsIdInvitations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdInvitationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_invitations",
		mcp.WithDescription("This feature was introduced in GitLab 13.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdInvitationsHandler))
}

func getProjectsIdInvitationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdInvitationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdInvitations(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdInvitationsEmailRequest struct {
	Id    string `json:"id" jsonschema:"description=The project ID"`
	Email string `json:"email" jsonschema:"description=The email address of the invitation"`
}

func registerDeleteProjectsIdInvitationsEmail(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdInvitationsEmailRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_invitations_email",
		mcp.WithDescription("Removes an invitation from a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdInvitationsEmailHandler))
}

func deleteProjectsIdInvitationsEmailHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdInvitationsEmailRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdInvitationsEmail(ctx, req.Id, req.Email, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidLinksRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	IssueIid int32  `json:"issue_iid" jsonschema:"description=The internal ID of a project’s issue"`
}

func registerGetProjectsIdIssuesIssueIidLinks(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidLinksRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_links",
		mcp.WithDescription("Get a list of a given issue’s linked issues, sorted by the relationship creation datetime (ascending).Issues are filtered according to the user authorizations."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidLinksHandler))
}

func getProjectsIdIssuesIssueIidLinksHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidLinksRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidLinks(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type DeleteProjectsIdIssuesIssueIidLinksIssueLinkIdRequest struct {
	Id          string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	IssueIid    int32  `json:"issue_iid" jsonschema:"description=The internal ID of a project’s issue"`
	IssueLinkId string `json:"issue_link_id" jsonschema:"description=The ID of an issue relationship"`
}

func registerDeleteProjectsIdIssuesIssueIidLinksIssueLinkId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdIssuesIssueIidLinksIssueLinkIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_issues_issue_iid_links_issue_link_id",
		mcp.WithDescription("Deletes an issue link, thus removes the two-way relationship."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdIssuesIssueIidLinksIssueLinkIdHandler))
}

func deleteProjectsIdIssuesIssueIidLinksIssueLinkIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdIssuesIssueIidLinksIssueLinkIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx, req.Id, req.IssueIid, req.IssueLinkId, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidLinksIssueLinkIdRequest struct {
	Id          string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	IssueIid    int32  `json:"issue_iid" jsonschema:"description=The internal ID of a project’s issue"`
	IssueLinkId string `json:"issue_link_id" jsonschema:"description=ID of an issue relationship"`
}

func registerGetProjectsIdIssuesIssueIidLinksIssueLinkId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidLinksIssueLinkIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_links_issue_link_id",
		mcp.WithDescription("Gets details about an issue link. This feature was introduced in GitLab 15.1."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidLinksIssueLinkIdHandler))
}

func getProjectsIdIssuesIssueIidLinksIssueLinkIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidLinksIssueLinkIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx, req.Id, req.IssueIid, req.IssueLinkId, authorizationHeader))
}

type GetProjectsIdCiLintRequest struct {
	Id     int32                                  `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdCiLintParams `json:"params,omitempty"`
}

func registerGetProjectsIdCiLint(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdCiLintRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_ci_lint",
		mcp.WithDescription("Checks if a project’s .gitlab-ci.yml configuration in a given commit (by default HEAD of the project’s default branch) is valid"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdCiLintHandler))
}

func getProjectsIdCiLintHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdCiLintRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdCiLint(ctx, req.Id, req.Params, authorizationHeader))
}

type PostProjectsIdUploadsAuthorizeRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerPostProjectsIdUploadsAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdUploadsAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_uploads_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 13.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdUploadsAuthorizeHandler))
}

func postProjectsIdUploadsAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdUploadsAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdUploadsAuthorize(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdUploadsRequest struct {
	Id     int32                                   `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdUploadsParams `json:"params,omitempty"`
}

func registerGetProjectsIdUploads(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdUploadsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_uploads",
		mcp.WithDescription("Get the list of uploads of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdUploadsHandler))
}

func getProjectsIdUploadsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdUploadsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdUploads(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdUploadsUploadIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	UploadId int32 `json:"upload_id" jsonschema:"description=The ID of a project upload"`
}

func registerDeleteProjectsIdUploadsUploadId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdUploadsUploadIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_uploads_upload_id",
		mcp.WithDescription("Delete a single project upload by ID"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdUploadsUploadIdHandler))
}

func deleteProjectsIdUploadsUploadIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdUploadsUploadIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdUploadsUploadId(ctx, req.Id, req.UploadId, authorizationHeader))
}

type GetProjectsIdUploadsUploadIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	UploadId int32 `json:"upload_id" jsonschema:"description=The ID of a project upload"`
}

func registerGetProjectsIdUploadsUploadId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdUploadsUploadIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_uploads_upload_id",
		mcp.WithDescription("Download a single project upload by ID"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdUploadsUploadIdHandler))
}

func getProjectsIdUploadsUploadIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdUploadsUploadIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdUploadsUploadId(ctx, req.Id, req.UploadId, authorizationHeader))
}

type DeleteProjectsIdUploadsSecretFilenameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=null"`
	Secret   string `json:"secret" jsonschema:"description=The 32-character secret of a project upload"`
	Filename string `json:"filename" jsonschema:"description=The filename of a project upload"`
}

func registerDeleteProjectsIdUploadsSecretFilename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdUploadsSecretFilenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_uploads_secret_filename",
		mcp.WithDescription("Delete a single project upload by secret and filename"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdUploadsSecretFilenameHandler))
}

func deleteProjectsIdUploadsSecretFilenameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdUploadsSecretFilenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdUploadsSecretFilename(ctx, req.Id, req.Secret, req.Filename, authorizationHeader))
}

type GetProjectsIdUploadsSecretFilenameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=null"`
	Secret   string `json:"secret" jsonschema:"description=The 32-character secret of a project upload"`
	Filename string `json:"filename" jsonschema:"description=The filename of a project upload"`
}

func registerGetProjectsIdUploadsSecretFilename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdUploadsSecretFilenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_uploads_secret_filename",
		mcp.WithDescription("Download a single project upload by secret and filename"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdUploadsSecretFilenameHandler))
}

func getProjectsIdUploadsSecretFilenameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdUploadsSecretFilenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdUploadsSecretFilename(ctx, req.Id, req.Secret, req.Filename, authorizationHeader))
}

type GetProjectsIdMembersRequest struct {
	Id     string                                  `json:"id" jsonschema:"description=The project ID"`
	Params *client.GetApiV4ProjectsIdMembersParams `json:"params,omitempty"`
}

func registerGetProjectsIdMembers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMembersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_members",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMembersHandler))
}

func getProjectsIdMembersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMembersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMembers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdMembersAllRequest struct {
	Id     string                                     `json:"id" jsonschema:"description=The project ID"`
	Params *client.GetApiV4ProjectsIdMembersAllParams `json:"params,omitempty"`
}

func registerGetProjectsIdMembersAll(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMembersAllRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_members_all",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user, including those who gained membership through ancestor group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMembersAllHandler))
}

func getProjectsIdMembersAllHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMembersAllRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMembersAll(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdMembersUserIdRequest struct {
	Id     string                                           `json:"id" jsonschema:"description=The project ID"`
	UserId int32                                            `json:"user_id" jsonschema:"description=The user ID of the member"`
	Params *client.DeleteApiV4ProjectsIdMembersUserIdParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdMembersUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdMembersUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_members_user_id",
		mcp.WithDescription("Removes a user from a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdMembersUserIdHandler))
}

func deleteProjectsIdMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdMembersUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdMembersUserId(ctx, req.Id, req.UserId, req.Params, authorizationHeader))
}

type GetProjectsIdMembersUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The project ID"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerGetProjectsIdMembersUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMembersUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_members_user_id",
		mcp.WithDescription("Gets a member of a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMembersUserIdHandler))
}

func getProjectsIdMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMembersUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMembersUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type GetProjectsIdMembersAllUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The project ID"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerGetProjectsIdMembersAllUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMembersAllUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_members_all_user_id",
		mcp.WithDescription("Gets a member of a group or project, including those who gained membership through ancestor group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMembersAllUserIdHandler))
}

func getProjectsIdMembersAllUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMembersAllUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMembersAllUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidApprovalsRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidApprovals(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidApprovalsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_approvals",
		mcp.WithDescription("List approvals for merge request"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidApprovalsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidApprovalsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidApprovalsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type PostProjectsIdMergeRequestsMergeRequestIidUnapproveRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerPostProjectsIdMergeRequestsMergeRequestIidUnapprove(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdMergeRequestsMergeRequestIidUnapproveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_mrs_merge_request_iid_unapprove",
		mcp.WithDescription("Remove an approval from a merge request"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdMergeRequestsMergeRequestIidUnapproveHandler))
}

func postProjectsIdMergeRequestsMergeRequestIidUnapproveHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdMergeRequestsMergeRequestIidUnapproveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type PutProjectsIdMergeRequestsMergeRequestIidResetApprovalsRequest struct {
	Id              int32 `json:"id" jsonschema:"description=null"`
	MergeRequestIid int32 `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerPutProjectsIdMergeRequestsMergeRequestIidResetApprovals(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdMergeRequestsMergeRequestIidResetApprovalsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_mrs_merge_request_iid_reset_approvals",
		mcp.WithDescription("Clear all approvals of merge request. This feature was added in GitLab 15.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdMergeRequestsMergeRequestIidResetApprovalsHandler))
}

func putProjectsIdMergeRequestsMergeRequestIidResetApprovalsHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdMergeRequestsMergeRequestIidResetApprovalsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidApprovalStateRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The IID of a merge request"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidApprovalState(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidApprovalStateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_approval_state",
		mcp.WithDescription("Get approval state of merge request"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidApprovalStateHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidApprovalStateHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidApprovalStateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type PostProjectsIdCreateCiConfigRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerPostProjectsIdCreateCiConfig(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdCreateCiConfigRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_create_ci_config",
		mcp.WithDescription("Creates merge request for missing ci config in project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdCreateCiConfigHandler))
}

func postProjectsIdCreateCiConfigHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdCreateCiConfigRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdCreateCiConfig(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge_request."`
}

func registerPostProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_mrs_merge_request_iid_reset_time_estimate",
		mcp.WithDescription("Resets the estimated time for this merge_request to 0 seconds."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateHandler))
}

func postProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type PostProjectsIdMergeRequestsMergeRequestIidResetSpentTimeRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge_request"`
}

func registerPostProjectsIdMergeRequestsMergeRequestIidResetSpentTime(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdMergeRequestsMergeRequestIidResetSpentTimeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_mrs_merge_request_iid_reset_spent_time",
		mcp.WithDescription("Resets the total spent time for this merge_request to 0 seconds."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdMergeRequestsMergeRequestIidResetSpentTimeHandler))
}

func postProjectsIdMergeRequestsMergeRequestIidResetSpentTimeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdMergeRequestsMergeRequestIidResetSpentTimeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidTimeStatsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge_request"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidTimeStats(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidTimeStatsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_time_stats",
		mcp.WithDescription("Get time tracking stats"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidTimeStatsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidTimeStatsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidTimeStatsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	Params *client.GetApiV4ProjectsIdMergeRequestsParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs",
		mcp.WithDescription("Get all merge requests for this project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsHandler))
}

func getProjectsIdMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequests(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdMergeRequestsMergeRequestIidRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge request."`
}

func registerDeleteProjectsIdMergeRequestsMergeRequestIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdMergeRequestsMergeRequestIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_mrs_merge_request_iid",
		mcp.WithDescription("Only for administrators and project owners. Deletes the merge request in question."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdMergeRequestsMergeRequestIidHandler))
}

func deleteProjectsIdMergeRequestsMergeRequestIidHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdMergeRequestsMergeRequestIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidRequest struct {
	Id              string                                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                        `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge request."`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid",
		mcp.WithDescription("Shows information about a single merge request. Note: the `changes_count` value in the response is a string, not an integer. This is because when an merge request has too many changes to display and store, it is capped at 1,000. In that case, the API returns the string `\"1000+\"` for the changes count."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidParticipantsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidParticipants(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidParticipantsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_participants",
		mcp.WithDescription("Get a list of merge request participants."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidParticipantsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidParticipantsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidParticipantsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidReviewersRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidReviewers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidReviewersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_reviewers",
		mcp.WithDescription("Get a list of merge request reviewers."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidReviewersHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidReviewersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidReviewersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidCommitsRequest struct {
	Id              string                                                              `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                               `json:"merge_request_iid" jsonschema:"description=null"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidCommits(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidCommitsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_commits",
		mcp.WithDescription("Get a list of merge request commits."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidCommitsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidCommitsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidCommitsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type DeleteProjectsIdMergeRequestsMergeRequestIidContextCommitsRequest struct {
	Id              string                                                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                                         `json:"merge_request_iid" jsonschema:"description=null"`
	Params          *client.DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdMergeRequestsMergeRequestIidContextCommits(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdMergeRequestsMergeRequestIidContextCommitsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_mrs_merge_request_iid_context_commits",
		mcp.WithDescription("Delete a list of merge request context commits."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdMergeRequestsMergeRequestIidContextCommitsHandler))
}

func deleteProjectsIdMergeRequestsMergeRequestIidContextCommitsHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdMergeRequestsMergeRequestIidContextCommitsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidContextCommitsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidContextCommits(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidContextCommitsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_context_commits",
		mcp.WithDescription("Get a list of merge request context commits."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidContextCommitsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidContextCommitsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidContextCommitsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidChangesRequest struct {
	Id              string                                                              `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                               `json:"merge_request_iid" jsonschema:"description=null"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidChanges(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidChangesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_changes",
		mcp.WithDescription("Shows information about the merge request including its files and changes."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidChangesHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidChangesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidChangesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidDiffsRequest struct {
	Id              string                                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                             `json:"merge_request_iid" jsonschema:"description=null"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidDiffs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidDiffsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_diffs",
		mcp.WithDescription("Get a list of merge request diffs."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidDiffsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidDiffsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidDiffsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidRawDiffsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidRawDiffs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidRawDiffsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_raw_diffs",
		mcp.WithDescription("Get the raw diffs of a merge request that can used programmatically."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidRawDiffsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidRawDiffsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidRawDiffsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffs(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidPipelinesRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidPipelines(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidPipelinesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_pls",
		mcp.WithDescription("Get a list of merge request pipelines."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidPipelinesHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidPipelinesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidPipelinesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidMergeRefRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidMergeRef(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidMergeRefRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_merge_ref",
		mcp.WithDescription("Returns the up to date merge-ref HEAD commit"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidMergeRefHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidMergeRefHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidMergeRefRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type PostProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsRequest struct {
	Id              string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32  `json:"merge_request_iid" jsonschema:"description=null"`
}

func registerPostProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_mrs_merge_request_iid_cancel_merge_when_pipeline_succeeds",
		mcp.WithDescription("Cancel merge if \"Merge When Pipeline Succeeds\" is enabled"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsHandler))
}

func postProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds(ctx, req.Id, req.MergeRequestIid, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidClosesIssuesRequest struct {
	Id              string                                                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                                    `json:"merge_request_iid" jsonschema:"description=null"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidClosesIssues(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidClosesIssuesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_closes_issues",
		mcp.WithDescription("Get all the issues that would be closed by merging the provided merge request."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidClosesIssuesHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidClosesIssuesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidClosesIssuesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidRelatedIssuesRequest struct {
	Id              string                                                                    `json:"id" jsonschema:"description=The ID or URL-encoded path of the project."`
	MergeRequestIid int32                                                                     `json:"merge_request_iid" jsonschema:"description=null"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidRelatedIssues(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidRelatedIssuesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_related_issues",
		mcp.WithDescription("Get all the related issues from title, description, commits, comments and discussions of the merge request."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidRelatedIssuesHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidRelatedIssuesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidRelatedIssuesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssues(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidVersionsRequest struct {
	Id              string                                                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MergeRequestIid int32                                                                `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge request"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidVersions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidVersionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_versions",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidVersionsHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidVersionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidVersionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions(ctx, req.Id, req.MergeRequestIid, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdRequest struct {
	Id              string                                                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MergeRequestIid int32                                                                         `json:"merge_request_iid" jsonschema:"description=The internal ID of the merge request"`
	VersionId       int32                                                                         `json:"version_id" jsonschema:"description=The ID of the merge request diff version"`
	Params          *client.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_merge_request_iid_versions_version_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdHandler))
}

func getProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(ctx, req.Id, req.MergeRequestIid, req.VersionId, req.Params, authorizationHeader))
}

type PostProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_npm_npm_v1_security_advisories_bulk",
		mcp.WithDescription("This feature was introduced in GitLab 15.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkHandler))
}

func postProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdPackagesNpmNpmV1SecurityAuditsQuickRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdPackagesNpmNpmV1SecurityAuditsQuick(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesNpmNpmV1SecurityAuditsQuickRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_npm_npm_v1_security_audits_quick",
		mcp.WithDescription("This feature was introduced in GitLab 15.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesNpmNpmV1SecurityAuditsQuickHandler))
}

func postProjectsIdPackagesNpmNpmV1SecurityAuditsQuickHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesNpmNpmV1SecurityAuditsQuickRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesNugetIndexRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesNugetIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesNugetIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_index",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesNugetIndexHandler))
}

func getProjectsIdPackagesNugetIndexHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesNugetIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesNugetIndex(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesNugetV2Request struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesNugetV2(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesNugetV2Request{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_v2",
		mcp.WithDescription("This feature was introduced in GitLab 16.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesNugetV2Handler))
}

func getProjectsIdPackagesNugetV2Handler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesNugetV2Request) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesNugetV2(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesNugetV2MetadataRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesNugetV2Metadata(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesNugetV2MetadataRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_v2_metadata",
		mcp.WithDescription("This feature was introduced in GitLab 16.3"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesNugetV2MetadataHandler))
}

func getProjectsIdPackagesNugetV2MetadataHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesNugetV2MetadataRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesNugetV2Metadata(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesNugetQueryRequest struct {
	Id     string                                             `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdPackagesNugetQueryParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesNugetQuery(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesNugetQueryRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_nuget_query",
		mcp.WithDescription("This feature was introduced in GitLab 12.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesNugetQueryHandler))
}

func getProjectsIdPackagesNugetQueryHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesNugetQueryRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesNugetQuery(ctx, req.Id, req.Params, authorizationHeader))
}

type PutProjectsIdPackagesNugetAuthorizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPutProjectsIdPackagesNugetAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesNugetAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_nuget_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 14.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesNugetAuthorizeHandler))
}

func putProjectsIdPackagesNugetAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesNugetAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesNugetAuthorize(ctx, req.Id, authorizationHeader))
}

type PutProjectsIdPackagesNugetSymbolpackageAuthorizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPutProjectsIdPackagesNugetSymbolpackageAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesNugetSymbolpackageAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_nuget_symbolpackage_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 14.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesNugetSymbolpackageAuthorizeHandler))
}

func putProjectsIdPackagesNugetSymbolpackageAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesNugetSymbolpackageAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize(ctx, req.Id, authorizationHeader))
}

type PutProjectsIdPackagesNugetV2AuthorizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPutProjectsIdPackagesNugetV2Authorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPackagesNugetV2AuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pkgs_nuget_v2_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 16.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPackagesNugetV2AuthorizeHandler))
}

func putProjectsIdPackagesNugetV2AuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPackagesNugetV2AuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPackagesNugetV2Authorize(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesPackageIdPackageFilesRequest struct {
	Id        string                                                        `json:"id" jsonschema:"description=ID or URL-encoded path of the project"`
	PackageId int32                                                         `json:"package_id" jsonschema:"description=ID of a package"`
	Params    *client.GetApiV4ProjectsIdPackagesPackageIdPackageFilesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesPackageIdPackageFiles(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesPackageIdPackageFilesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_package_id_package_files",
		mcp.WithDescription("Get a list of package files of a single package"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesPackageIdPackageFilesHandler))
}

func getProjectsIdPackagesPackageIdPackageFilesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesPackageIdPackageFilesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesPackageIdPackageFiles(ctx, req.Id, req.PackageId, req.Params, authorizationHeader))
}

type DeleteProjectsIdPackagesPackageIdPackageFilesPackageFileIdRequest struct {
	Id            string `json:"id" jsonschema:"description=ID or URL-encoded path of the project"`
	PackageId     int32  `json:"package_id" jsonschema:"description=ID of a package"`
	PackageFileId int32  `json:"package_file_id" jsonschema:"description=ID of a package file"`
}

func registerDeleteProjectsIdPackagesPackageIdPackageFilesPackageFileId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPackagesPackageIdPackageFilesPackageFileIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pkgs_package_id_package_files_package_file_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.12"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPackagesPackageIdPackageFilesPackageFileIdHandler))
}

func deleteProjectsIdPackagesPackageIdPackageFilesPackageFileIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPackagesPackageIdPackageFilesPackageFileIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId(ctx, req.Id, req.PackageId, req.PackageFileId, authorizationHeader))
}

type DeleteProjectsIdPagesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
}

func registerDeleteProjectsIdPages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pages",
		mcp.WithDescription("Remove pages. The user must have administrator access. This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPagesHandler))
}

func deleteProjectsIdPagesHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPages(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPagesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
}

func registerGetProjectsIdPages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pages",
		mcp.WithDescription("Get pages URL and other settings. This feature was introduced in Gitlab 16.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPagesHandler))
}

func getProjectsIdPagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPages(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPagesDomainsRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Params *client.GetApiV4ProjectsIdPagesDomainsParams `json:"params,omitempty"`
}

func registerGetProjectsIdPagesDomains(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPagesDomainsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pages_domains",
		mcp.WithDescription("Get all pages domains"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPagesDomainsHandler))
}

func getProjectsIdPagesDomainsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPagesDomainsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPagesDomains(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdPagesDomainsDomainRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Domain string `json:"domain" jsonschema:"description=The domain"`
}

func registerDeleteProjectsIdPagesDomainsDomain(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPagesDomainsDomainRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pages_domains_domain",
		mcp.WithDescription("Delete a pages domain"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPagesDomainsDomainHandler))
}

func deleteProjectsIdPagesDomainsDomainHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPagesDomainsDomainRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPagesDomainsDomain(ctx, req.Id, req.Domain, authorizationHeader))
}

type GetProjectsIdPagesDomainsDomainRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Domain string `json:"domain" jsonschema:"description=The domain"`
}

func registerGetProjectsIdPagesDomainsDomain(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPagesDomainsDomainRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pages_domains_domain",
		mcp.WithDescription("Get a single pages domain"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPagesDomainsDomainHandler))
}

func getProjectsIdPagesDomainsDomainHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPagesDomainsDomainRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPagesDomainsDomain(ctx, req.Id, req.Domain, authorizationHeader))
}

type PutProjectsIdPagesDomainsDomainVerifyRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project owned by the authenticated user"`
	Domain string `json:"domain" jsonschema:"description=The domain to verify"`
}

func registerPutProjectsIdPagesDomainsDomainVerify(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdPagesDomainsDomainVerifyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_pages_domains_domain_verify",
		mcp.WithDescription("Verify a pages domain"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdPagesDomainsDomainVerifyHandler))
}

func putProjectsIdPagesDomainsDomainVerifyHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdPagesDomainsDomainVerifyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdPagesDomainsDomainVerify(ctx, req.Id, req.Domain, authorizationHeader))
}

type GetProjectsIdAvatarRequest struct {
	Id string `json:"id" jsonschema:"description=ID or URL-encoded path of the project"`
}

func registerGetProjectsIdAvatar(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdAvatarRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_avatar",
		mcp.WithDescription("This feature was introduced in GitLab 16.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdAvatarHandler))
}

func getProjectsIdAvatarHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdAvatarRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdAvatar(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdClustersRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdClustersParams `json:"params,omitempty"`
}

func registerGetProjectsIdClusters(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdClustersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 11.7. Returns a list of project clusters."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdClustersHandler))
}

func getProjectsIdClustersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdClustersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdClusters(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdClustersClusterIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	ClusterId int32  `json:"cluster_id" jsonschema:"description=The Cluster ID"`
}

func registerDeleteProjectsIdClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.7. Deletes an existing project cluster. Does not remove existing resources within the connected Kubernetes cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdClustersClusterIdHandler))
}

func deleteProjectsIdClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdClustersClusterId(ctx, req.Id, req.ClusterId, authorizationHeader))
}

type GetProjectsIdClustersClusterIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	ClusterId int32  `json:"cluster_id" jsonschema:"description=The cluster ID"`
}

func registerGetProjectsIdClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.7. Gets a single project cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdClustersClusterIdHandler))
}

func getProjectsIdClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdClustersClusterId(ctx, req.Id, req.ClusterId, authorizationHeader))
}

type GetProjectsIdRegistryRepositoriesRequest struct {
	Id     string                                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRegistryRepositoriesParams `json:"params,omitempty"`
}

func registerGetProjectsIdRegistryRepositories(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRegistryRepositoriesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_registry_repositories",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRegistryRepositoriesHandler))
}

func getProjectsIdRegistryRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRegistryRepositoriesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRegistryRepositories(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdRegistryRepositoriesRepositoryIdRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	RepositoryId int32  `json:"repository_id" jsonschema:"description=The ID of the repository"`
}

func registerDeleteProjectsIdRegistryRepositoriesRepositoryId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRegistryRepositoriesRepositoryIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_registry_repositories_repo_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRegistryRepositoriesRepositoryIdHandler))
}

func deleteProjectsIdRegistryRepositoriesRepositoryIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRegistryRepositoriesRepositoryIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId(ctx, req.Id, req.RepositoryId, authorizationHeader))
}

type DeleteProjectsIdRegistryRepositoriesRepositoryIdTagsRequest struct {
	Id           string                                                                  `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	RepositoryId int32                                                                   `json:"repository_id" jsonschema:"description=The ID of the repository"`
	Params       *client.DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdRegistryRepositoriesRepositoryIdTags(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRegistryRepositoriesRepositoryIdTagsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_registry_repositories_repo_id_tags",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRegistryRepositoriesRepositoryIdTagsHandler))
}

func deleteProjectsIdRegistryRepositoriesRepositoryIdTagsHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRegistryRepositoriesRepositoryIdTagsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx, req.Id, req.RepositoryId, req.Params, authorizationHeader))
}

type GetProjectsIdRegistryRepositoriesRepositoryIdTagsRequest struct {
	Id           string                                                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	RepositoryId int32                                                                `json:"repository_id" jsonschema:"description=The ID of the repository"`
	Params       *client.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRegistryRepositoriesRepositoryIdTags(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRegistryRepositoriesRepositoryIdTagsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_registry_repositories_repo_id_tags",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRegistryRepositoriesRepositoryIdTagsHandler))
}

func getProjectsIdRegistryRepositoriesRepositoryIdTagsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRegistryRepositoriesRepositoryIdTagsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx, req.Id, req.RepositoryId, req.Params, authorizationHeader))
}

type DeleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	RepositoryId int32  `json:"repository_id" jsonschema:"description=The ID of the repository"`
	TagName      string `json:"tag_name" jsonschema:"description=The name of the tag"`
}

func registerDeleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_registry_repositories_repo_id_tags_tag_name",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameHandler))
}

func deleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx, req.Id, req.RepositoryId, req.TagName, authorizationHeader))
}

type GetProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	RepositoryId int32  `json:"repository_id" jsonschema:"description=The ID of the repository"`
	TagName      string `json:"tag_name" jsonschema:"description=The name of the tag"`
}

func registerGetProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_registry_repositories_repo_id_tags_tag_name",
		mcp.WithDescription("This feature was introduced in GitLab 11.8."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameHandler))
}

func getProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx, req.Id, req.RepositoryId, req.TagName, authorizationHeader))
}

type GetProjectsIdRegistryProtectionRepositoryRulesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdRegistryProtectionRepositoryRules(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRegistryProtectionRepositoryRulesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_registry_protection_repo_rules",
		mcp.WithDescription("Get list of container registry protection rules for a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRegistryProtectionRepositoryRulesHandler))
}

func getProjectsIdRegistryProtectionRepositoryRulesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRegistryProtectionRepositoryRulesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRegistryProtectionRepositoryRules(ctx, req.Id, authorizationHeader))
}

type DeleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdRequest struct {
	Id               string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	ProtectionRuleId int32  `json:"protection_rule_id" jsonschema:"description=The ID of the container protection rule"`
}

func registerDeleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_registry_protection_repo_rules_protection_rule_id",
		mcp.WithDescription("Delete container protection rule"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdHandler))
}

func deleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId(ctx, req.Id, req.ProtectionRuleId, authorizationHeader))
}

type GetProjectsIdDebianDistributionsRequest struct {
	Id     string                                              `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdDebianDistributionsParams `json:"params,omitempty"`
}

func registerGetProjectsIdDebianDistributions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDebianDistributionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_debian_distributions",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDebianDistributionsHandler))
}

func getProjectsIdDebianDistributionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDebianDistributionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDebianDistributions(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdDebianDistributionsCodenameRequest struct {
	Id       string                                                         `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Codename string                                                         `json:"codename" jsonschema:"description=The Debian Codename"`
	Params   *client.DeleteApiV4ProjectsIdDebianDistributionsCodenameParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdDebianDistributionsCodename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdDebianDistributionsCodenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_debian_distributions_codename",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdDebianDistributionsCodenameHandler))
}

func deleteProjectsIdDebianDistributionsCodenameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdDebianDistributionsCodenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdDebianDistributionsCodename(ctx, req.Id, req.Codename, req.Params, authorizationHeader))
}

type GetProjectsIdDebianDistributionsCodenameRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Codename string `json:"codename" jsonschema:"description=The Debian Codename"`
}

func registerGetProjectsIdDebianDistributionsCodename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDebianDistributionsCodenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_debian_distributions_codename",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDebianDistributionsCodenameHandler))
}

func getProjectsIdDebianDistributionsCodenameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDebianDistributionsCodenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDebianDistributionsCodename(ctx, req.Id, req.Codename, authorizationHeader))
}

type GetProjectsIdDebianDistributionsCodenameKeyAscRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Codename string `json:"codename" jsonschema:"description=The Debian Codename"`
}

func registerGetProjectsIdDebianDistributionsCodenameKeyAsc(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdDebianDistributionsCodenameKeyAscRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_debian_distributions_codename_key_asc",
		mcp.WithDescription("This feature was introduced in 14.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdDebianDistributionsCodenameKeyAscHandler))
}

func getProjectsIdDebianDistributionsCodenameKeyAscHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdDebianDistributionsCodenameKeyAscRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc(ctx, req.Id, req.Codename, authorizationHeader))
}

type GetProjectsIdEventsRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdEventsParams `json:"params,omitempty"`
}

func registerGetProjectsIdEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_events",
		mcp.WithDescription("List a project's visible events"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdEventsHandler))
}

func getProjectsIdEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdEvents(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdExportRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdExport(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdExportRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_export",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdExportHandler))
}

func getProjectsIdExportHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdExportRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdExport(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdExportDownloadRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdExportDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdExportDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_export_download",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdExportDownloadHandler))
}

func getProjectsIdExportDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdExportDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdExportDownload(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdExportRelationsDownloadRequest struct {
	Id     string                                                  `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdExportRelationsDownloadParams `json:"params,omitempty"`
}

func registerGetProjectsIdExportRelationsDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdExportRelationsDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_export_relations_download",
		mcp.WithDescription("This feature was introduced in GitLab 14.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdExportRelationsDownloadHandler))
}

func getProjectsIdExportRelationsDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdExportRelationsDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdExportRelationsDownload(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdExportRelationsStatusRequest struct {
	Id     string                                                `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdExportRelationsStatusParams `json:"params,omitempty"`
}

func registerGetProjectsIdExportRelationsStatus(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdExportRelationsStatusRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_export_relations_status",
		mcp.WithDescription("This feature was introduced in GitLab 14.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdExportRelationsStatusHandler))
}

func getProjectsIdExportRelationsStatusHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdExportRelationsStatusRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdExportRelationsStatus(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdHooksHookIdUrlVariablesKeyRequest struct {
	Id     int32  `json:"id" jsonschema:"description=null"`
	HookId int32  `json:"hook_id" jsonschema:"description=The ID of the hook"`
	Key    string `json:"key" jsonschema:"description=The key of the variable"`
}

func registerDeleteProjectsIdHooksHookIdUrlVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdHooksHookIdUrlVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_hooks_hook_id_url_variables_key",
		mcp.WithDescription("Un-Set a url variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdHooksHookIdUrlVariablesKeyHandler))
}

func deleteProjectsIdHooksHookIdUrlVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdHooksHookIdUrlVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey(ctx, req.Id, req.HookId, req.Key, authorizationHeader))
}

type DeleteProjectsIdHooksHookIdCustomHeadersKeyRequest struct {
	Id     int32  `json:"id" jsonschema:"description=null"`
	HookId int32  `json:"hook_id" jsonschema:"description=The ID of the hook"`
	Key    string `json:"key" jsonschema:"description=The key of the custom header"`
}

func registerDeleteProjectsIdHooksHookIdCustomHeadersKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdHooksHookIdCustomHeadersKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_hooks_hook_id_custom_headers_key",
		mcp.WithDescription("Un-Set a custom header"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdHooksHookIdCustomHeadersKeyHandler))
}

func deleteProjectsIdHooksHookIdCustomHeadersKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdHooksHookIdCustomHeadersKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdHooksHookIdCustomHeadersKey(ctx, req.Id, req.HookId, req.Key, authorizationHeader))
}

type GetProjectsIdHooksRequest struct {
	Id     string                                `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdHooksParams `json:"params,omitempty"`
}

func registerGetProjectsIdHooks(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdHooksRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_hooks",
		mcp.WithDescription("Get a list of project hooks"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdHooksHandler))
}

func getProjectsIdHooksHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdHooksRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdHooks(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdHooksHookIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	HookId int32  `json:"hook_id" jsonschema:"description=The ID of the project hook"`
}

func registerDeleteProjectsIdHooksHookId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdHooksHookIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_hooks_hook_id",
		mcp.WithDescription("Removes a hook from a project. This is an idempotent method and can be called multiple times. Either the hook is available or not."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdHooksHookIdHandler))
}

func deleteProjectsIdHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdHooksHookIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdHooksHookId(ctx, req.Id, req.HookId, authorizationHeader))
}

type GetProjectsIdHooksHookIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	HookId int32  `json:"hook_id" jsonschema:"description=The ID of a project hook"`
}

func registerGetProjectsIdHooksHookId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdHooksHookIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_hooks_hook_id",
		mcp.WithDescription("Get a specific hook for a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdHooksHookIdHandler))
}

func getProjectsIdHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdHooksHookIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdHooksHookId(ctx, req.Id, req.HookId, authorizationHeader))
}

type GetProjectsIdHooksHookIdEventsRequest struct {
	Id     int32                                             `json:"id" jsonschema:"description=null"`
	HookId int32                                             `json:"hook_id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdHooksHookIdEventsParams `json:"params,omitempty"`
}

func registerGetProjectsIdHooksHookIdEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdHooksHookIdEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_hooks_hook_id_events",
		mcp.WithDescription("List web hook logs by hook id"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdHooksHookIdEventsHandler))
}

func getProjectsIdHooksHookIdEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdHooksHookIdEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdHooksHookIdEvents(ctx, req.Id, req.HookId, req.Params, authorizationHeader))
}

type PostProjectsIdHooksHookIdTestTriggerRequest struct {
	Id      int32  `json:"id" jsonschema:"description=null"`
	HookId  int32  `json:"hook_id" jsonschema:"description=The ID of the hook"`
	Trigger string `json:"trigger" jsonschema:"description=The type of trigger hook"`
}

func registerPostProjectsIdHooksHookIdTestTrigger(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdHooksHookIdTestTriggerRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_hooks_hook_id_test_trigger",
		mcp.WithDescription("Triggers a hook test"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdHooksHookIdTestTriggerHandler))
}

func postProjectsIdHooksHookIdTestTriggerHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdHooksHookIdTestTriggerRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdHooksHookIdTestTrigger(ctx, req.Id, req.HookId, req.Trigger, authorizationHeader))
}

type PostProjectsIdHooksHookIdEventsHookLogIdResendRequest struct {
	Id        int32 `json:"id" jsonschema:"description=null"`
	HookId    int32 `json:"hook_id" jsonschema:"description=null"`
	HookLogId int32 `json:"hook_log_id" jsonschema:"description=null"`
}

func registerPostProjectsIdHooksHookIdEventsHookLogIdResend(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdHooksHookIdEventsHookLogIdResendRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_hooks_hook_id_events_hook_log_id_resend",
		mcp.WithDescription("Resend a webhook event"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdHooksHookIdEventsHookLogIdResendHandler))
}

func postProjectsIdHooksHookIdEventsHookLogIdResendHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdHooksHookIdEventsHookLogIdResendRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdHooksHookIdEventsHookLogIdResend(ctx, req.Id, req.HookId, req.HookLogId, authorizationHeader))
}

type PostProjectsImportAuthorizeRequest struct {
}

func registerPostProjectsImportAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsImportAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_import_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 12.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsImportAuthorizeHandler))
}

func postProjectsImportAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsImportAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsImportAuthorize(ctx, authorizationHeader))
}

type GetProjectsIdImportRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdImport(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdImportRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_import",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdImportHandler))
}

func getProjectsIdImportHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdImportRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdImport(ctx, req.Id, authorizationHeader))
}

type PostProjectsImportRelationAuthorizeRequest struct {
}

func registerPostProjectsImportRelationAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsImportRelationAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_import_relation_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 16.11"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsImportRelationAuthorizeHandler))
}

func postProjectsImportRelationAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsImportRelationAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsImportRelationAuthorize(ctx, authorizationHeader))
}

type GetProjectsIdRelationImportsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdRelationImports(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRelationImportsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_relation_imports",
		mcp.WithDescription("This feature was introduced in GitLab 16.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRelationImportsHandler))
}

func getProjectsIdRelationImportsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRelationImportsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRelationImports(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdJobTokenScopeRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetProjectsIdJobTokenScope(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobTokenScopeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_job_token_scope",
		mcp.WithDescription("Fetch CI_JOB_TOKEN access settings."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobTokenScopeHandler))
}

func getProjectsIdJobTokenScopeHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobTokenScopeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobTokenScope(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdJobTokenScopeAllowlistRequest struct {
	Id     int32                                                  `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdJobTokenScopeAllowlistParams `json:"params,omitempty"`
}

func registerGetProjectsIdJobTokenScopeAllowlist(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobTokenScopeAllowlistRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_job_token_scope_allowlist",
		mcp.WithDescription("Fetch project inbound allowlist for CI_JOB_TOKEN access settings."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobTokenScopeAllowlistHandler))
}

func getProjectsIdJobTokenScopeAllowlistHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobTokenScopeAllowlistRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobTokenScopeAllowlist(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdJobTokenScopeGroupsAllowlistRequest struct {
	Id     int32                                                        `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistParams `json:"params,omitempty"`
}

func registerGetProjectsIdJobTokenScopeGroupsAllowlist(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdJobTokenScopeGroupsAllowlistRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_job_token_scope_grps_allowlist",
		mcp.WithDescription("Fetch project groups allowlist for CI_JOB_TOKEN access settings."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdJobTokenScopeGroupsAllowlistHandler))
}

func getProjectsIdJobTokenScopeGroupsAllowlistHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdJobTokenScopeGroupsAllowlistRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdRequest struct {
	Id            int32 `json:"id" jsonschema:"description=ID of user project"`
	TargetGroupId int32 `json:"target_group_id" jsonschema:"description=ID of the group to be removed from the allowlist"`
}

func registerDeleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_job_token_scope_grps_allowlist_target_group_id",
		mcp.WithDescription("Delete target group from allowlist."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdHandler))
}

func deleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId(ctx, req.Id, req.TargetGroupId, authorizationHeader))
}

type DeleteProjectsIdJobTokenScopeAllowlistTargetProjectIdRequest struct {
	Id              int32 `json:"id" jsonschema:"description=ID of user project"`
	TargetProjectId int32 `json:"target_project_id" jsonschema:"description=ID of the project to be removed from the allowlist"`
}

func registerDeleteProjectsIdJobTokenScopeAllowlistTargetProjectId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdJobTokenScopeAllowlistTargetProjectIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_job_token_scope_allowlist_target_project_id",
		mcp.WithDescription("Delete project from allowlist."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdJobTokenScopeAllowlistTargetProjectIdHandler))
}

func deleteProjectsIdJobTokenScopeAllowlistTargetProjectIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdJobTokenScopeAllowlistTargetProjectIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId(ctx, req.Id, req.TargetProjectId, authorizationHeader))
}

type GetProjectsIdPackagesRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdPackagesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs",
		mcp.WithDescription("This feature was introduced in GitLab 11.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesHandler))
}

func getProjectsIdPackagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackages(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdPackagesPackageIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageId int32  `json:"package_id" jsonschema:"description=The ID of a package"`
}

func registerDeleteProjectsIdPackagesPackageId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPackagesPackageIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pkgs_package_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPackagesPackageIdHandler))
}

func deleteProjectsIdPackagesPackageIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPackagesPackageIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPackagesPackageId(ctx, req.Id, req.PackageId, authorizationHeader))
}

type GetProjectsIdPackagesPackageIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageId int32  `json:"package_id" jsonschema:"description=The ID of a package"`
}

func registerGetProjectsIdPackagesPackageId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesPackageIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_package_id",
		mcp.WithDescription("This feature was introduced in GitLab 11.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesPackageIdHandler))
}

func getProjectsIdPackagesPackageIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesPackageIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesPackageId(ctx, req.Id, req.PackageId, authorizationHeader))
}

type GetProjectsIdPackagesPackageIdPipelinesRequest struct {
	Id        string                                                     `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageId int32                                                      `json:"package_id" jsonschema:"description=The ID of a package"`
	Params    *client.GetApiV4ProjectsIdPackagesPackageIdPipelinesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesPackageIdPipelines(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesPackageIdPipelinesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_package_id_pls",
		mcp.WithDescription("This feature was introduced in GitLab 16.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesPackageIdPipelinesHandler))
}

func getProjectsIdPackagesPackageIdPipelinesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesPackageIdPipelinesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesPackageIdPipelines(ctx, req.Id, req.PackageId, req.Params, authorizationHeader))
}

type GetProjectsIdPackagesProtectionRulesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesProtectionRules(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesProtectionRulesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_protection_rules",
		mcp.WithDescription("Get list of package protection rules for a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesProtectionRulesHandler))
}

func getProjectsIdPackagesProtectionRulesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesProtectionRulesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesProtectionRules(ctx, req.Id, authorizationHeader))
}

type DeleteProjectsIdPackagesProtectionRulesPackageProtectionRuleIdRequest struct {
	Id                      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	PackageProtectionRuleId int32  `json:"package_protection_rule_id" jsonschema:"description=The ID of the package protection rule"`
}

func registerDeleteProjectsIdPackagesProtectionRulesPackageProtectionRuleId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdPackagesProtectionRulesPackageProtectionRuleIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_pkgs_protection_rules_package_protection_rule_id",
		mcp.WithDescription("Delete package protection rule"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdPackagesProtectionRulesPackageProtectionRuleIdHandler))
}

func deleteProjectsIdPackagesProtectionRulesPackageProtectionRuleIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdPackagesProtectionRulesPackageProtectionRuleIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleId(ctx, req.Id, req.PackageProtectionRuleId, authorizationHeader))
}

type GetProjectsIdSnapshotRequest struct {
	Id     int32                                    `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdSnapshotParams `json:"params,omitempty"`
}

func registerGetProjectsIdSnapshot(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnapshotRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snapshot",
		mcp.WithDescription("This feature was introduced in GitLab 10.7"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnapshotHandler))
}

func getProjectsIdSnapshotHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnapshotRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnapshot(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdSnippetsRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdSnippetsParams `json:"params,omitempty"`
}

func registerGetProjectsIdSnippets(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets",
		mcp.WithDescription("Get all project snippets"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsHandler))
}

func getProjectsIdSnippetsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippets(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdSnippetsSnippetIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	SnippetId int32  `json:"snippet_id" jsonschema:"description=The ID of a project snippet"`
}

func registerDeleteProjectsIdSnippetsSnippetId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdSnippetsSnippetIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_snippets_snippet_id",
		mcp.WithDescription("Delete a project snippet"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdSnippetsSnippetIdHandler))
}

func deleteProjectsIdSnippetsSnippetIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdSnippetsSnippetIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdSnippetsSnippetId(ctx, req.Id, req.SnippetId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	SnippetId int32  `json:"snippet_id" jsonschema:"description=The ID of a project snippet"`
}

func registerGetProjectsIdSnippetsSnippetId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id",
		mcp.WithDescription("Get a single project snippet"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdHandler))
}

func getProjectsIdSnippetsSnippetIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetId(ctx, req.Id, req.SnippetId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdRawRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	SnippetId int32  `json:"snippet_id" jsonschema:"description=The ID of a project snippet"`
}

func registerGetProjectsIdSnippetsSnippetIdRaw(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdRawRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_raw",
		mcp.WithDescription("Get a raw project snippet"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdRawHandler))
}

func getProjectsIdSnippetsSnippetIdRawHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdRawRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdRaw(ctx, req.Id, req.SnippetId, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdFilesRefFilePathRawRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	SnippetId int32  `json:"snippet_id" jsonschema:"description=null"`
	Ref       string `json:"ref" jsonschema:"description=The name of branch, tag or commit"`
	FilePath  string `json:"file_path" jsonschema:"description=The url encoded path to the file, e.g. lib%2Fclass%2Erb"`
}

func registerGetProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdFilesRefFilePathRawRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_files_ref_file_path_raw",
		mcp.WithDescription("Get raw project snippet file contents from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdFilesRefFilePathRawHandler))
}

func getProjectsIdSnippetsSnippetIdFilesRefFilePathRawHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdFilesRefFilePathRawRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(ctx, req.Id, req.SnippetId, req.Ref, req.FilePath, authorizationHeader))
}

type GetProjectsIdSnippetsSnippetIdUserAgentDetailRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	SnippetId int32  `json:"snippet_id" jsonschema:"description=The ID of a project snippet"`
}

func registerGetProjectsIdSnippetsSnippetIdUserAgentDetail(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdSnippetsSnippetIdUserAgentDetailRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_snippets_snippet_id_user_agent_detail",
		mcp.WithDescription("Get the user agent details for a project snippet"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdSnippetsSnippetIdUserAgentDetailHandler))
}

func getProjectsIdSnippetsSnippetIdUserAgentDetailHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdSnippetsSnippetIdUserAgentDetailRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail(ctx, req.Id, req.SnippetId, authorizationHeader))
}

type GetProjectsIdStatisticsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdStatistics(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdStatisticsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_statistics",
		mcp.WithDescription("Get the list of project fetch statistics for the last 30 days"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdStatisticsHandler))
}

func getProjectsIdStatisticsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdStatisticsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdStatistics(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdTemplatesTypeRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Type   string                                        `json:"type" jsonschema:"description=The type (dockerfiles|gitignores|gitlab_ci_ymls|licenses|issues|merge_requests) of the template"`
	Params *client.GetApiV4ProjectsIdTemplatesTypeParams `json:"params,omitempty"`
}

func registerGetProjectsIdTemplatesType(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTemplatesTypeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_templates_type",
		mcp.WithDescription("This endpoint was introduced in GitLab 11.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTemplatesTypeHandler))
}

func getProjectsIdTemplatesTypeHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTemplatesTypeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTemplatesType(ctx, req.Id, req.Type, req.Params, authorizationHeader))
}

type GetProjectsIdTemplatesTypeNameRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Type   string                                            `json:"type" jsonschema:"description=The type (dockerfiles|gitignores|gitlab_ci_ymls|licenses|issues|merge_requests) of the template"`
	Name   string                                            `json:"name" jsonschema:"description=The key of the template, as obtained from the collection endpoint."`
	Params *client.GetApiV4ProjectsIdTemplatesTypeNameParams `json:"params,omitempty"`
}

func registerGetProjectsIdTemplatesTypeName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTemplatesTypeNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_templates_type_name",
		mcp.WithDescription("This endpoint was introduced in GitLab 11.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTemplatesTypeNameHandler))
}

func getProjectsIdTemplatesTypeNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTemplatesTypeNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTemplatesTypeName(ctx, req.Id, req.Type, req.Name, req.Params, authorizationHeader))
}

type GetProjectsIdCustomAttributesRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetProjectsIdCustomAttributes(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdCustomAttributesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_custom_attributes",
		mcp.WithDescription("Get all custom attributes on a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdCustomAttributesHandler))
}

func getProjectsIdCustomAttributesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdCustomAttributesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdCustomAttributes(ctx, req.Id, authorizationHeader))
}

type DeleteProjectsIdCustomAttributesKeyRequest struct {
	Id  int32  `json:"id" jsonschema:"description=null"`
	Key string `json:"key" jsonschema:"description=The key of the custom attribute"`
}

func registerDeleteProjectsIdCustomAttributesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdCustomAttributesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_custom_attributes_key",
		mcp.WithDescription("Delete a custom attribute on a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdCustomAttributesKeyHandler))
}

func deleteProjectsIdCustomAttributesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdCustomAttributesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdCustomAttributesKey(ctx, req.Id, req.Key, authorizationHeader))
}

type GetProjectsIdCustomAttributesKeyRequest struct {
	Id  int32  `json:"id" jsonschema:"description=null"`
	Key string `json:"key" jsonschema:"description=The key of the custom attribute"`
}

func registerGetProjectsIdCustomAttributesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdCustomAttributesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_custom_attributes_key",
		mcp.WithDescription("Get a custom attribute on a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdCustomAttributesKeyHandler))
}

func getProjectsIdCustomAttributesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdCustomAttributesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdCustomAttributesKey(ctx, req.Id, req.Key, authorizationHeader))
}

type PostProjectsIdRestoreRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerPostProjectsIdRestore(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdRestoreRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_restore",
		mcp.WithDescription("Restore a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdRestoreHandler))
}

func postProjectsIdRestoreHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdRestoreRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdRestore(ctx, req.Id, authorizationHeader))
}

type GetProjectsRequest struct {
	Params *client.GetApiV4ProjectsParams `json:"params,omitempty"`
}

func registerGetProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs",
		mcp.WithDescription("Get a list of visible projects for authenticated user"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsHandler))
}

func getProjectsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Projects(ctx, req.Params, authorizationHeader))
}

type GetProjectsIdShareLocationsRequest struct {
	Id     int32                                          `json:"id" jsonschema:"description=The id of the project"`
	Params *client.GetApiV4ProjectsIdShareLocationsParams `json:"params,omitempty"`
}

func registerGetProjectsIdShareLocations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdShareLocationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_share_locations",
		mcp.WithDescription("Returns group that can be shared with the given project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdShareLocationsHandler))
}

func getProjectsIdShareLocationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdShareLocationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdShareLocations(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerDeleteProjectsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id",
		mcp.WithDescription("Delete a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdHandler))
}

func deleteProjectsIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsId(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdRequest struct {
	Id     string                           `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdParams `json:"params,omitempty"`
}

func registerGetProjectsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id",
		mcp.WithDescription("Get a single project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdHandler))
}

func getProjectsIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsId(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdForkRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerDeleteProjectsIdFork(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdForkRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_fork",
		mcp.WithDescription("Remove a forked_from relationship"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdForkHandler))
}

func deleteProjectsIdForkHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdForkRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdFork(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdForksRequest struct {
	Id     string                                `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdForksParams `json:"params,omitempty"`
}

func registerGetProjectsIdForks(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdForksRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_forks",
		mcp.WithDescription("List forks of this project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdForksHandler))
}

func getProjectsIdForksHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdForksRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdForks(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdPagesAccessRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPagesAccess(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPagesAccessRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pages_access",
		mcp.WithDescription("Check pages access of this project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPagesAccessHandler))
}

func getProjectsIdPagesAccessHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPagesAccessRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPagesAccess(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdArchiveRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdArchive(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdArchiveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_archive",
		mcp.WithDescription("Archive a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdArchiveHandler))
}

func postProjectsIdArchiveHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdArchiveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdArchive(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdUnarchiveRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdUnarchive(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdUnarchiveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_unarchive",
		mcp.WithDescription("Unarchive a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdUnarchiveHandler))
}

func postProjectsIdUnarchiveHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdUnarchiveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdUnarchive(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdStarRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdStar(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdStarRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_star",
		mcp.WithDescription("Star a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdStarHandler))
}

func postProjectsIdStarHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdStarRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdStar(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdUnstarRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdUnstar(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdUnstarRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_unstar",
		mcp.WithDescription("Unstar a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdUnstarHandler))
}

func postProjectsIdUnstarHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdUnstarRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdUnstar(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdStarrersRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdStarrersParams `json:"params,omitempty"`
}

func registerGetProjectsIdStarrers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdStarrersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_starrers",
		mcp.WithDescription("Get the users who starred a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdStarrersHandler))
}

func getProjectsIdStarrersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdStarrersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdStarrers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdLanguagesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdLanguages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdLanguagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_languages",
		mcp.WithDescription("Get languages in project repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdLanguagesHandler))
}

func getProjectsIdLanguagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdLanguagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdLanguages(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdForkForkedFromIdRequest struct {
	Id           string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	ForkedFromId string `json:"forked_from_id" jsonschema:"description=The ID of the project it was forked from"`
}

func registerPostProjectsIdForkForkedFromId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdForkForkedFromIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_fork_forked_from_id",
		mcp.WithDescription("Mark this project as forked from another"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdForkForkedFromIdHandler))
}

func postProjectsIdForkForkedFromIdHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdForkForkedFromIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdForkForkedFromId(ctx, req.Id, req.ForkedFromId, authorizationHeader))
}

type DeleteProjectsIdShareGroupIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	GroupId int32  `json:"group_id" jsonschema:"description=The ID of the group"`
}

func registerDeleteProjectsIdShareGroupId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdShareGroupIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_share_group_id",
		mcp.WithDescription("Remove a group share"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdShareGroupIdHandler))
}

func deleteProjectsIdShareGroupIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdShareGroupIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdShareGroupId(ctx, req.Id, req.GroupId, authorizationHeader))
}

type PostProjectsIdImportProjectMembersProjectIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	ProjectId int32  `json:"project_id" jsonschema:"description=The ID of the source project to import the members from."`
}

func registerPostProjectsIdImportProjectMembersProjectId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdImportProjectMembersProjectIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_import_project_members_project_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdImportProjectMembersProjectIdHandler))
}

func postProjectsIdImportProjectMembersProjectIdHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdImportProjectMembersProjectIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdImportProjectMembersProjectId(ctx, req.Id, req.ProjectId, authorizationHeader))
}

type GetProjectsIdUsersRequest struct {
	Id     string                                `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdUsersParams `json:"params,omitempty"`
}

func registerGetProjectsIdUsers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdUsersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_users",
		mcp.WithDescription("Get the users list of a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdUsersHandler))
}

func getProjectsIdUsersHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdUsersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdUsers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdGroupsRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdGroupsParams `json:"params,omitempty"`
}

func registerGetProjectsIdGroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdGroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_grps",
		mcp.WithDescription("Get ancestor and shared groups for a project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdGroupsHandler))
}

func getProjectsIdGroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdGroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdGroups(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdInvitedGroupsRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdInvitedGroupsParams `json:"params,omitempty"`
}

func registerGetProjectsIdInvitedGroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdInvitedGroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_invited_grps",
		mcp.WithDescription("Get a list of invited groups in this project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdInvitedGroupsHandler))
}

func getProjectsIdInvitedGroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdInvitedGroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdInvitedGroups(ctx, req.Id, req.Params, authorizationHeader))
}

type PostProjectsIdRepositorySizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdRepositorySize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdRepositorySizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_repo_size",
		mcp.WithDescription("This feature was introduced in GitLab 15.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdRepositorySizeHandler))
}

func postProjectsIdRepositorySizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdRepositorySizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdRepositorySize(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdTransferLocationsRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdTransferLocationsParams `json:"params,omitempty"`
}

func registerGetProjectsIdTransferLocations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTransferLocationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_transfer_locations",
		mcp.WithDescription("Get the namespaces to where the project can be transferred"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTransferLocationsHandler))
}

func getProjectsIdTransferLocationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTransferLocationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTransferLocations(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdStorageRequest struct {
	Id string `json:"id" jsonschema:"description=ID of a project"`
}

func registerGetProjectsIdStorage(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdStorageRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_storage",
		mcp.WithDescription("Show the storage information"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdStorageHandler))
}

func getProjectsIdStorageHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdStorageRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdStorage(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdAuditEventsRequest struct {
	Id     int32                                       `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdAuditEventsParams `json:"params,omitempty"`
}

func registerGetProjectsIdAuditEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdAuditEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_audit_events",
		mcp.WithDescription("Get a list of audit events in this project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdAuditEventsHandler))
}

func getProjectsIdAuditEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdAuditEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdAuditEvents(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdAuditEventsAuditEventIdRequest struct {
	Id           int32 `json:"id" jsonschema:"description=null"`
	AuditEventId int32 `json:"audit_event_id" jsonschema:"description=The ID of the audit event"`
}

func registerGetProjectsIdAuditEventsAuditEventId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdAuditEventsAuditEventIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_audit_events_audit_event_id",
		mcp.WithDescription("Get a specific audit event in this project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdAuditEventsAuditEventIdHandler))
}

func getProjectsIdAuditEventsAuditEventIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdAuditEventsAuditEventIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdAuditEventsAuditEventId(ctx, req.Id, req.AuditEventId, authorizationHeader))
}

type GetProjectsIdProtectedBranchesRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdProtectedBranchesParams `json:"params,omitempty"`
}

func registerGetProjectsIdProtectedBranches(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdProtectedBranchesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_protected_branches",
		mcp.WithDescription("Get a project's protected branches"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdProtectedBranchesHandler))
}

func getProjectsIdProtectedBranchesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdProtectedBranchesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdProtectedBranches(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdProtectedBranchesNameRequest struct {
	Id   string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name string `json:"name" jsonschema:"description=The name of the protected branch"`
}

func registerDeleteProjectsIdProtectedBranchesName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdProtectedBranchesNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_protected_branches_name",
		mcp.WithDescription("Unprotect a single branch"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdProtectedBranchesNameHandler))
}

func deleteProjectsIdProtectedBranchesNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdProtectedBranchesNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdProtectedBranchesName(ctx, req.Id, req.Name, authorizationHeader))
}

type GetProjectsIdProtectedBranchesNameRequest struct {
	Id   string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name string `json:"name" jsonschema:"description=The name of the branch or wildcard"`
}

func registerGetProjectsIdProtectedBranchesName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdProtectedBranchesNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_protected_branches_name",
		mcp.WithDescription("Get a single protected branch"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdProtectedBranchesNameHandler))
}

func getProjectsIdProtectedBranchesNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdProtectedBranchesNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdProtectedBranchesName(ctx, req.Id, req.Name, authorizationHeader))
}

type GetProjectsIdProtectedTagsRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdProtectedTagsParams `json:"params,omitempty"`
}

func registerGetProjectsIdProtectedTags(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdProtectedTagsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_protected_tags",
		mcp.WithDescription("This feature was introduced in GitLab 11.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdProtectedTagsHandler))
}

func getProjectsIdProtectedTagsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdProtectedTagsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdProtectedTags(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdProtectedTagsNameRequest struct {
	Id   string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name string `json:"name" jsonschema:"description=The name of the protected tag"`
}

func registerDeleteProjectsIdProtectedTagsName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdProtectedTagsNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_protected_tags_name",
		mcp.WithDescription("This feature was introduced in GitLab 11.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdProtectedTagsNameHandler))
}

func deleteProjectsIdProtectedTagsNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdProtectedTagsNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdProtectedTagsName(ctx, req.Id, req.Name, authorizationHeader))
}

type GetProjectsIdProtectedTagsNameRequest struct {
	Id   string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name string `json:"name" jsonschema:"description=The name of the tag or wildcard"`
}

func registerGetProjectsIdProtectedTagsName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdProtectedTagsNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_protected_tags_name",
		mcp.WithDescription("This feature was introduced in GitLab 11.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdProtectedTagsNameHandler))
}

func getProjectsIdProtectedTagsNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdProtectedTagsNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdProtectedTagsName(ctx, req.Id, req.Name, authorizationHeader))
}

type GetProjectsIdPackagesPypiSimpleRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerGetProjectsIdPackagesPypiSimple(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesPypiSimpleRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_pypi_simple",
		mcp.WithDescription("This feature was introduced in GitLab 15.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesPypiSimpleHandler))
}

func getProjectsIdPackagesPypiSimpleHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesPypiSimpleRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesPypiSimple(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdPackagesPypiAuthorizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdPackagesPypiAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesPypiAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_pypi_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 12.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesPypiAuthorizeHandler))
}

func postProjectsIdPackagesPypiAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesPypiAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesPypiAuthorize(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdReleasesRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdReleasesParams `json:"params,omitempty"`
}

func registerGetProjectsIdReleases(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdReleasesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_releases",
		mcp.WithDescription("Returns a paginated list of releases. This feature was introduced in GitLab 11.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdReleasesHandler))
}

func getProjectsIdReleasesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdReleasesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdReleases(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdReleasesTagNameRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string `json:"tag_name" jsonschema:"description=The Git tag the release is associated with"`
}

func registerDeleteProjectsIdReleasesTagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdReleasesTagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_releases_tag_name",
		mcp.WithDescription("Delete a release. Deleting a release doesn't delete the associated tag. Maintainer level access to the project is required to delete a release. This feature was introduced in GitLab 11.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdReleasesTagNameHandler))
}

func deleteProjectsIdReleasesTagNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdReleasesTagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdReleasesTagName(ctx, req.Id, req.TagName, authorizationHeader))
}

type GetProjectsIdReleasesTagNameRequest struct {
	Id      string                                          `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string                                          `json:"tag_name" jsonschema:"description=The Git tag the release is associated with"`
	Params  *client.GetApiV4ProjectsIdReleasesTagNameParams `json:"params,omitempty"`
}

func registerGetProjectsIdReleasesTagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdReleasesTagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_releases_tag_name",
		mcp.WithDescription("Gets a release for the given tag. This feature was introduced in GitLab 11.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdReleasesTagNameHandler))
}

func getProjectsIdReleasesTagNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdReleasesTagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdReleasesTagName(ctx, req.Id, req.TagName, req.Params, authorizationHeader))
}

type PostProjectsIdReleasesTagNameEvidenceRequest struct {
	Id      int32  `json:"id" jsonschema:"description=null"`
	TagName string `json:"tag_name" jsonschema:"description=The Git tag the release is associated with"`
}

func registerPostProjectsIdReleasesTagNameEvidence(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdReleasesTagNameEvidenceRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_releases_tag_name_evidence",
		mcp.WithDescription("Creates an evidence for an existing Release. This feature was introduced in GitLab 12.10."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdReleasesTagNameEvidenceHandler))
}

func postProjectsIdReleasesTagNameEvidenceHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdReleasesTagNameEvidenceRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdReleasesTagNameEvidence(ctx, req.Id, req.TagName, authorizationHeader))
}

type GetProjectsIdReleasesTagNameAssetsLinksRequest struct {
	Id      string                                                     `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string                                                     `json:"tag_name" jsonschema:"description=The tag associated with the release"`
	Params  *client.GetApiV4ProjectsIdReleasesTagNameAssetsLinksParams `json:"params,omitempty"`
}

func registerGetProjectsIdReleasesTagNameAssetsLinks(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdReleasesTagNameAssetsLinksRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_releases_tag_name_assets_links",
		mcp.WithDescription("Get assets as links from a release. This feature was introduced in GitLab 11.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdReleasesTagNameAssetsLinksHandler))
}

func getProjectsIdReleasesTagNameAssetsLinksHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdReleasesTagNameAssetsLinksRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdReleasesTagNameAssetsLinks(ctx, req.Id, req.TagName, req.Params, authorizationHeader))
}

type DeleteProjectsIdReleasesTagNameAssetsLinksLinkIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string `json:"tag_name" jsonschema:"description=The tag associated with the release"`
	LinkId  int32  `json:"link_id" jsonschema:"description=The ID of the link"`
}

func registerDeleteProjectsIdReleasesTagNameAssetsLinksLinkId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdReleasesTagNameAssetsLinksLinkIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_releases_tag_name_assets_links_link_id",
		mcp.WithDescription("Deletes an asset as a link from a release. This feature was introduced in GitLab 11.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdReleasesTagNameAssetsLinksLinkIdHandler))
}

func deleteProjectsIdReleasesTagNameAssetsLinksLinkIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdReleasesTagNameAssetsLinksLinkIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx, req.Id, req.TagName, req.LinkId, authorizationHeader))
}

type GetProjectsIdReleasesTagNameAssetsLinksLinkIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string `json:"tag_name" jsonschema:"description=The tag associated with the release"`
	LinkId  int32  `json:"link_id" jsonschema:"description=The ID of the link"`
}

func registerGetProjectsIdReleasesTagNameAssetsLinksLinkId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdReleasesTagNameAssetsLinksLinkIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_releases_tag_name_assets_links_link_id",
		mcp.WithDescription("Get an asset as a link from a release. This feature was introduced in GitLab 11.7."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdReleasesTagNameAssetsLinksLinkIdHandler))
}

func getProjectsIdReleasesTagNameAssetsLinksLinkIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdReleasesTagNameAssetsLinksLinkIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx, req.Id, req.TagName, req.LinkId, authorizationHeader))
}

type GetProjectsIdRemoteMirrorsRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRemoteMirrorsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRemoteMirrors(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRemoteMirrorsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_remote_mirrors",
		mcp.WithDescription("List the project's remote mirrors"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRemoteMirrorsHandler))
}

func getProjectsIdRemoteMirrorsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRemoteMirrorsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRemoteMirrors(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdRemoteMirrorsMirrorIdRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MirrorId string `json:"mirror_id" jsonschema:"description=The ID of a remote mirror"`
}

func registerDeleteProjectsIdRemoteMirrorsMirrorId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRemoteMirrorsMirrorIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_remote_mirrors_mirror_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRemoteMirrorsMirrorIdHandler))
}

func deleteProjectsIdRemoteMirrorsMirrorIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRemoteMirrorsMirrorIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRemoteMirrorsMirrorId(ctx, req.Id, req.MirrorId, authorizationHeader))
}

type GetProjectsIdRemoteMirrorsMirrorIdRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MirrorId string `json:"mirror_id" jsonschema:"description=The ID of a remote mirror"`
}

func registerGetProjectsIdRemoteMirrorsMirrorId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRemoteMirrorsMirrorIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_remote_mirrors_mirror_id",
		mcp.WithDescription("Get a single remote mirror"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRemoteMirrorsMirrorIdHandler))
}

func getProjectsIdRemoteMirrorsMirrorIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRemoteMirrorsMirrorIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRemoteMirrorsMirrorId(ctx, req.Id, req.MirrorId, authorizationHeader))
}

type PostProjectsIdRemoteMirrorsMirrorIdSyncRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MirrorId string `json:"mirror_id" jsonschema:"description=The ID of a remote mirror"`
}

func registerPostProjectsIdRemoteMirrorsMirrorIdSync(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdRemoteMirrorsMirrorIdSyncRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_remote_mirrors_mirror_id_sync",
		mcp.WithDescription("Triggers a push mirror operation"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdRemoteMirrorsMirrorIdSyncHandler))
}

func postProjectsIdRemoteMirrorsMirrorIdSyncHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdRemoteMirrorsMirrorIdSyncRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync(ctx, req.Id, req.MirrorId, authorizationHeader))
}

type GetProjectsIdRemoteMirrorsMirrorIdPublicKeyRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	MirrorId string `json:"mirror_id" jsonschema:"description=The ID of a remote mirror"`
}

func registerGetProjectsIdRemoteMirrorsMirrorIdPublicKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRemoteMirrorsMirrorIdPublicKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_remote_mirrors_mirror_id_public_key",
		mcp.WithDescription("Get the public key of a single remote mirror"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRemoteMirrorsMirrorIdPublicKeyHandler))
}

func getProjectsIdRemoteMirrorsMirrorIdPublicKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRemoteMirrorsMirrorIdPublicKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKey(ctx, req.Id, req.MirrorId, authorizationHeader))
}

type GetProjectsIdRepositoryTreeRequest struct {
	Id     string                                         `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryTreeParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryTree(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryTreeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_tree",
		mcp.WithDescription("Get a project repository tree"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryTreeHandler))
}

func getProjectsIdRepositoryTreeHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryTreeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryTree(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryBlobsShaRawRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha string `json:"sha" jsonschema:"description=The commit hash"`
}

func registerGetProjectsIdRepositoryBlobsShaRaw(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryBlobsShaRawRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_blobs_sha_raw",
		mcp.WithDescription("Get raw blob contents from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryBlobsShaRawHandler))
}

func getProjectsIdRepositoryBlobsShaRawHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryBlobsShaRawRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryBlobsShaRaw(ctx, req.Id, req.Sha, authorizationHeader))
}

type GetProjectsIdRepositoryBlobsShaRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Sha string `json:"sha" jsonschema:"description=The commit hash"`
}

func registerGetProjectsIdRepositoryBlobsSha(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryBlobsShaRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_blobs_sha",
		mcp.WithDescription("Get a blob from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryBlobsShaHandler))
}

func getProjectsIdRepositoryBlobsShaHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryBlobsShaRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryBlobsSha(ctx, req.Id, req.Sha, authorizationHeader))
}

type GetProjectsIdRepositoryArchiveRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryArchiveParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryArchive(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryArchiveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_archive",
		mcp.WithDescription("Get an archive of the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryArchiveHandler))
}

func getProjectsIdRepositoryArchiveHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryArchiveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryArchive(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryCompareRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryCompareParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryCompare(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryCompareRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_compare",
		mcp.WithDescription("Compare two branches, tags, or commits"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryCompareHandler))
}

func getProjectsIdRepositoryCompareHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryCompareRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryCompare(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryHealthRequest struct {
	Id     string                                           `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryHealthParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryHealth(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryHealthRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_health",
		mcp.WithDescription("Get repository health"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryHealthHandler))
}

func getProjectsIdRepositoryHealthHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryHealthRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryHealth(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryContributorsRequest struct {
	Id     string                                                 `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryContributorsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryContributors(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryContributorsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_contributors",
		mcp.WithDescription("Get repository contributors"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryContributorsHandler))
}

func getProjectsIdRepositoryContributorsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryContributorsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryContributors(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryMergeBaseRequest struct {
	Id     string                                              `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryMergeBaseParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryMergeBase(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryMergeBaseRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_merge_base",
		mcp.WithDescription("Get the common ancestor between commits"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryMergeBaseHandler))
}

func getProjectsIdRepositoryMergeBaseHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryMergeBaseRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryMergeBase(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryChangelogRequest struct {
	Id     string                                              `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryChangelogParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryChangelog(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryChangelogRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_changelog",
		mcp.WithDescription("This feature was introduced in GitLab 14.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryChangelogHandler))
}

func getProjectsIdRepositoryChangelogHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryChangelogRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryChangelog(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdIssuesEventableIdResourceMilestoneEventsRequest struct {
	Id          string                                                                   `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	EventableId int32                                                                    `json:"eventable_id" jsonschema:"description=The ID of the eventable"`
	Params      *client.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsParams `json:"params,omitempty"`
}

func registerGetProjectsIdIssuesEventableIdResourceMilestoneEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesEventableIdResourceMilestoneEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_eventable_id_resource_milestone_events",
		mcp.WithDescription("Gets a list of all milestone events for a single Issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesEventableIdResourceMilestoneEventsHandler))
}

func getProjectsIdIssuesEventableIdResourceMilestoneEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesEventableIdResourceMilestoneEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents(ctx, req.Id, req.EventableId, req.Params, authorizationHeader))
}

type GetProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdRequest struct {
	Id          string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	EventableId int32  `json:"eventable_id" jsonschema:"description=The ID of the eventable"`
	EventId     string `json:"event_id" jsonschema:"description=The ID of a resource milestone event"`
}

func registerGetProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_eventable_id_resource_milestone_events_event_id",
		mcp.WithDescription("Returns a single milestone event for a specific project Issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdHandler))
}

func getProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(ctx, req.Id, req.EventableId, req.EventId, authorizationHeader))
}

type GetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsRequest struct {
	Id          string                                                                          `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	EventableId int32                                                                           `json:"eventable_id" jsonschema:"description=The ID of the eventable"`
	Params      *client.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsParams `json:"params,omitempty"`
}

func registerGetProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_eventable_id_resource_milestone_events",
		mcp.WithDescription("Gets a list of all milestone events for a single Merge request"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsHandler))
}

func getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(ctx, req.Id, req.EventableId, req.Params, authorizationHeader))
}

type GetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdRequest struct {
	Id          string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	EventableId int32  `json:"eventable_id" jsonschema:"description=The ID of the eventable"`
	EventId     string `json:"event_id" jsonschema:"description=The ID of a resource milestone event"`
}

func registerGetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_mrs_eventable_id_resource_milestone_events_event_id",
		mcp.WithDescription("Returns a single milestone event for a specific project Merge request"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdHandler))
}

func getProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(ctx, req.Id, req.EventableId, req.EventId, authorizationHeader))
}

type PostProjectsIdPackagesRpmRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdPackagesRpm(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesRpmRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_rpm",
		mcp.WithDescription("This feature was introduced in GitLab 15.7"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesRpmHandler))
}

func postProjectsIdPackagesRpmHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesRpmRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesRpm(ctx, req.Id, authorizationHeader))
}

type PostProjectsIdPackagesRpmAuthorizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdPackagesRpmAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesRpmAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_rpm_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 15.7"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesRpmAuthorizeHandler))
}

func postProjectsIdPackagesRpmAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesRpmAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesRpmAuthorize(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesRubygemsFileNameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FileName string `json:"file_name" jsonschema:"description=Spec file name"`
}

func registerGetProjectsIdPackagesRubygemsFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesRubygemsFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesRubygemsFileNameHandler))
}

func getProjectsIdPackagesRubygemsFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesRubygemsFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsFileName(ctx, req.Id, req.FileName, authorizationHeader))
}

type GetProjectsIdPackagesRubygemsQuickMarshal48FileNameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FileName string `json:"file_name" jsonschema:"description=Gemspec file name"`
}

func registerGetProjectsIdPackagesRubygemsQuickMarshal48FileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesRubygemsQuickMarshal48FileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_quick_marshal_4_8_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesRubygemsQuickMarshal48FileNameHandler))
}

func getProjectsIdPackagesRubygemsQuickMarshal48FileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesRubygemsQuickMarshal48FileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName(ctx, req.Id, req.FileName, authorizationHeader))
}

type GetProjectsIdPackagesRubygemsGemsFileNameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	FileName string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetProjectsIdPackagesRubygemsGemsFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesRubygemsGemsFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_gems_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesRubygemsGemsFileNameHandler))
}

func getProjectsIdPackagesRubygemsGemsFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesRubygemsGemsFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsGemsFileName(ctx, req.Id, req.FileName, authorizationHeader))
}

type PostProjectsIdPackagesRubygemsApiV1GemsAuthorizeRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
}

func registerPostProjectsIdPackagesRubygemsApiV1GemsAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdPackagesRubygemsApiV1GemsAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_pkgs_rubygems_api_v1_gems_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdPackagesRubygemsApiV1GemsAuthorizeHandler))
}

func postProjectsIdPackagesRubygemsApiV1GemsAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdPackagesRubygemsApiV1GemsAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize(ctx, req.Id, authorizationHeader))
}

type GetProjectsIdPackagesRubygemsApiV1DependenciesRequest struct {
	Id     int32                                                             `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesRubygemsApiV1Dependencies(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesRubygemsApiV1DependenciesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_rubygems_api_v1_dependencies",
		mcp.WithDescription("This feature was introduced in GitLab 13.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesRubygemsApiV1DependenciesHandler))
}

func getProjectsIdPackagesRubygemsApiV1DependenciesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesRubygemsApiV1DependenciesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdRepositoryTagsRequest struct {
	Id     string                                         `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Params *client.GetApiV4ProjectsIdRepositoryTagsParams `json:"params,omitempty"`
}

func registerGetProjectsIdRepositoryTags(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryTagsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_tags",
		mcp.WithDescription("Get a project repository tags"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryTagsHandler))
}

func getProjectsIdRepositoryTagsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryTagsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryTags(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdRepositoryTagsTagNameRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string `json:"tag_name" jsonschema:"description=The name of the tag"`
}

func registerDeleteProjectsIdRepositoryTagsTagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdRepositoryTagsTagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_repo_tags_tag_name",
		mcp.WithDescription("Delete a repository tag"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdRepositoryTagsTagNameHandler))
}

func deleteProjectsIdRepositoryTagsTagNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdRepositoryTagsTagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdRepositoryTagsTagName(ctx, req.Id, req.TagName, authorizationHeader))
}

type GetProjectsIdRepositoryTagsTagNameRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string `json:"tag_name" jsonschema:"description=The name of the tag"`
}

func registerGetProjectsIdRepositoryTagsTagName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryTagsTagNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_tags_tag_name",
		mcp.WithDescription("Get a single repository tag"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryTagsTagNameHandler))
}

func getProjectsIdRepositoryTagsTagNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryTagsTagNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryTagsTagName(ctx, req.Id, req.TagName, authorizationHeader))
}

type GetProjectsIdRepositoryTagsTagNameSignatureRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	TagName string `json:"tag_name" jsonschema:"description=The name of the tag"`
}

func registerGetProjectsIdRepositoryTagsTagNameSignature(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdRepositoryTagsTagNameSignatureRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_repo_tags_tag_name_signature",
		mcp.WithDescription("Get a tag's signature"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdRepositoryTagsTagNameSignatureHandler))
}

func getProjectsIdRepositoryTagsTagNameSignatureHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdRepositoryTagsTagNameSignatureRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdRepositoryTagsTagNameSignature(ctx, req.Id, req.TagName, authorizationHeader))
}

type GetProjectsIdPackagesTerraformModulesModuleNameModuleSystemRequest struct {
	Id           string                                                                         `json:"id" jsonschema:"description=The ID or full path of a project"`
	ModuleName   string                                                                         `json:"module_name" jsonschema:"description=Module name"`
	ModuleSystem string                                                                         `json:"module_system" jsonschema:"description=Module system"`
	Params       *client.GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemParams `json:"params,omitempty"`
}

func registerGetProjectsIdPackagesTerraformModulesModuleNameModuleSystem(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdPackagesTerraformModulesModuleNameModuleSystemRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_pkgs_terraform_modules_module_name_module_system",
		mcp.WithDescription("This feature was introduced in GitLab 16.7"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdPackagesTerraformModulesModuleNameModuleSystemHandler))
}

func getProjectsIdPackagesTerraformModulesModuleNameModuleSystemHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdPackagesTerraformModulesModuleNameModuleSystemRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem(ctx, req.Id, req.ModuleName, req.ModuleSystem, req.Params, authorizationHeader))
}

type DeleteProjectsIdTerraformStateNameRequest struct {
	Id   string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name int32  `json:"name" jsonschema:"description=null"`
}

func registerDeleteProjectsIdTerraformStateName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdTerraformStateNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_terraform_state_name",
		mcp.WithDescription("Delete a Terraform state of a certain name"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdTerraformStateNameHandler))
}

func deleteProjectsIdTerraformStateNameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdTerraformStateNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdTerraformStateName(ctx, req.Id, req.Name, authorizationHeader))
}

type PostProjectsIdTerraformStateNameRequest struct {
	Id   string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name int32  `json:"name" jsonschema:"description=null"`
}

func registerPostProjectsIdTerraformStateName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdTerraformStateNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_terraform_state_name",
		mcp.WithDescription("Add a new Terraform state or update an existing one"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdTerraformStateNameHandler))
}

func postProjectsIdTerraformStateNameHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdTerraformStateNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdTerraformStateName(ctx, req.Id, req.Name, authorizationHeader))
}

type GetProjectsIdTerraformStateNameRequest struct {
	Id     string                                             `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name   string                                             `json:"name" jsonschema:"description=The name of a Terraform state"`
	Params *client.GetApiV4ProjectsIdTerraformStateNameParams `json:"params,omitempty"`
}

func registerGetProjectsIdTerraformStateName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTerraformStateNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_terraform_state_name",
		mcp.WithDescription("Get a Terraform state by its name"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTerraformStateNameHandler))
}

func getProjectsIdTerraformStateNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTerraformStateNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTerraformStateName(ctx, req.Id, req.Name, req.Params, authorizationHeader))
}

type DeleteProjectsIdTerraformStateNameLockRequest struct {
	Id     string                                                    `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name   int32                                                     `json:"name" jsonschema:"description=null"`
	Params *client.DeleteApiV4ProjectsIdTerraformStateNameLockParams `json:"params,omitempty"`
}

func registerDeleteProjectsIdTerraformStateNameLock(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdTerraformStateNameLockRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_terraform_state_name_lock",
		mcp.WithDescription("Unlock a Terraform state of a certain name"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdTerraformStateNameLockHandler))
}

func deleteProjectsIdTerraformStateNameLockHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdTerraformStateNameLockRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdTerraformStateNameLock(ctx, req.Id, req.Name, req.Params, authorizationHeader))
}

type DeleteProjectsIdTerraformStateNameVersionsSerialRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name   int32  `json:"name" jsonschema:"description=null"`
	Serial int32  `json:"serial" jsonschema:"description=null"`
}

func registerDeleteProjectsIdTerraformStateNameVersionsSerial(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdTerraformStateNameVersionsSerialRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_terraform_state_name_versions_serial",
		mcp.WithDescription("Delete a Terraform state version"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdTerraformStateNameVersionsSerialHandler))
}

func deleteProjectsIdTerraformStateNameVersionsSerialHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdTerraformStateNameVersionsSerialRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx, req.Id, req.Name, req.Serial, authorizationHeader))
}

type GetProjectsIdTerraformStateNameVersionsSerialRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the project"`
	Name   string `json:"name" jsonschema:"description=The name of a Terraform state"`
	Serial int32  `json:"serial" jsonschema:"description=The version number of the state"`
}

func registerGetProjectsIdTerraformStateNameVersionsSerial(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdTerraformStateNameVersionsSerialRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_terraform_state_name_versions_serial",
		mcp.WithDescription("Get a Terraform state version"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdTerraformStateNameVersionsSerialHandler))
}

func getProjectsIdTerraformStateNameVersionsSerialHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdTerraformStateNameVersionsSerialRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx, req.Id, req.Name, req.Serial, authorizationHeader))
}

type GetProjectsIdWikisRequest struct {
	Id     int32                                 `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4ProjectsIdWikisParams `json:"params,omitempty"`
}

func registerGetProjectsIdWikis(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdWikisRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_wikis",
		mcp.WithDescription("Get a list of wiki pages"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdWikisHandler))
}

func getProjectsIdWikisHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdWikisRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdWikis(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdWikisSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The slug of a wiki page"`
}

func registerDeleteProjectsIdWikisSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdWikisSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_wikis_slug",
		mcp.WithDescription("Delete a wiki page"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdWikisSlugHandler))
}

func deleteProjectsIdWikisSlugHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdWikisSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdWikisSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetProjectsIdWikisSlugRequest struct {
	Id     int32                                     `json:"id" jsonschema:"description=null"`
	Slug   string                                    `json:"slug" jsonschema:"description=The slug of a wiki page"`
	Params *client.GetApiV4ProjectsIdWikisSlugParams `json:"params,omitempty"`
}

func registerGetProjectsIdWikisSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdWikisSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_wikis_slug",
		mcp.WithDescription("Get a wiki page"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdWikisSlugHandler))
}

func getProjectsIdWikisSlugHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdWikisSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdWikisSlug(ctx, req.Id, req.Slug, req.Params, authorizationHeader))
}

type PostProjectsIdIssuesRequest struct {
	Id     string                                  `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	Params *client.PostApiV4ProjectsIdIssuesParams `json:"params,omitempty"`
}

func registerPostProjectsIdIssues(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues",
		mcp.WithDescription("New issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesHandler))
}

func postProjectsIdIssuesHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssues(ctx, req.Id, req.Params, authorizationHeader))
}

type GetProjectsIdIssuesRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	Params *client.GetApiV4ProjectsIdIssuesParams `json:"params,omitempty"`
}

func registerGetProjectsIdIssues(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues",
		mcp.WithDescription("List project issues"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesHandler))
}

func getProjectsIdIssuesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssues(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteProjectsIdIssuesIssueIidRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerDeleteProjectsIdIssuesIssueIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdIssuesIssueIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_issues_issue_iid",
		mcp.WithDescription("Delete an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdIssuesIssueIidHandler))
}

func deleteProjectsIdIssuesIssueIidHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdIssuesIssueIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdIssuesIssueIid(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type PutProjectsIdIssuesIssueIidRequest struct {
	Id       string                                         `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int                                            `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
	Params   *client.PutApiV4ProjectsIdIssuesIssueIidParams `json:"params,omitempty"`
}

func registerPutProjectsIdIssuesIssueIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdIssuesIssueIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_issues_issue_iid",
		mcp.WithDescription("Edit an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdIssuesIssueIidHandler))
}

func putProjectsIdIssuesIssueIidHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdIssuesIssueIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdIssuesIssueIid(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerGetProjectsIdIssuesIssueIid(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid",
		mcp.WithDescription("Single project issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidHandler))
}

func getProjectsIdIssuesIssueIidHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIid(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type PutProjectsIdIssuesIssueIidReorderRequest struct {
	Id       string                                                `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int                                                   `json:"issue_iid" jsonschema:"description=The internal ID of the project's issue."`
	Params   *client.PutApiV4ProjectsIdIssuesIssueIidReorderParams `json:"params,omitempty"`
}

func registerPutProjectsIdIssuesIssueIidReorder(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutProjectsIdIssuesIssueIidReorderRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pjs_id_issues_issue_iid_reorder",
		mcp.WithDescription("Reorder an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putProjectsIdIssuesIssueIidReorderHandler))
}

func putProjectsIdIssuesIssueIidReorderHandler(ctx context.Context, request mcp.CallToolRequest, req PutProjectsIdIssuesIssueIidReorderRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4ProjectsIdIssuesIssueIidReorder(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidCloneRequest struct {
	Id       string                                               `json:"id" jsonschema:"description=ID or URL-encoded path of the project."`
	IssueIid int                                                  `json:"issue_iid" jsonschema:"description=Internal ID of a project's issue."`
	Params   *client.PostApiV4ProjectsIdIssuesIssueIidCloneParams `json:"params,omitempty"`
}

func registerPostProjectsIdIssuesIssueIidClone(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidCloneRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_clone",
		mcp.WithDescription("Clone an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidCloneHandler))
}

func postProjectsIdIssuesIssueIidCloneHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidCloneRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidClone(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidSubscribeRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerPostProjectsIdIssuesIssueIidSubscribe(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidSubscribeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_subscribe",
		mcp.WithDescription("Subscribe to an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidSubscribeHandler))
}

func postProjectsIdIssuesIssueIidSubscribeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidSubscribeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidSubscribe(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidUnsubscribeRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerPostProjectsIdIssuesIssueIidUnsubscribe(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidUnsubscribeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_unsubscribe",
		mcp.WithDescription("Unsubscribe from an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidUnsubscribeHandler))
}

func postProjectsIdIssuesIssueIidUnsubscribeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidUnsubscribeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidUnsubscribe(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidTodoRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerPostProjectsIdIssuesIssueIidTodo(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidTodoRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_todo",
		mcp.WithDescription("Create a to-do item"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidTodoHandler))
}

func postProjectsIdIssuesIssueIidTodoHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidTodoRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidTodo(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidNotesRequest struct {
	Id       string                                               `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int                                                  `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
	Params   *client.PostApiV4ProjectsIdIssuesIssueIidNotesParams `json:"params,omitempty"`
}

func registerPostProjectsIdIssuesIssueIidNotes(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidNotesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_notes",
		mcp.WithDescription("Promote an issue to an epic"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidNotesHandler))
}

func postProjectsIdIssuesIssueIidNotesHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidNotesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidNotes(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidTimeEstimateRequest struct {
	Id       string                                                      `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int                                                         `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
	Params   *client.PostApiV4ProjectsIdIssuesIssueIidTimeEstimateParams `json:"params,omitempty"`
}

func registerPostProjectsIdIssuesIssueIidTimeEstimate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidTimeEstimateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_time_estimate",
		mcp.WithDescription("Set a time estimate for an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidTimeEstimateHandler))
}

func postProjectsIdIssuesIssueIidTimeEstimateHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidTimeEstimateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidTimeEstimate(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidResetTimeEstimateRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerPostProjectsIdIssuesIssueIidResetTimeEstimate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidResetTimeEstimateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_reset_time_estimate",
		mcp.WithDescription("Reset the time estimate for an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidResetTimeEstimateHandler))
}

func postProjectsIdIssuesIssueIidResetTimeEstimateHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidResetTimeEstimateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidResetTimeEstimate(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidAddSpentTimeRequest struct {
	Id       string                                                      `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int                                                         `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
	Params   *client.PostApiV4ProjectsIdIssuesIssueIidAddSpentTimeParams `json:"params,omitempty"`
}

func registerPostProjectsIdIssuesIssueIidAddSpentTime(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidAddSpentTimeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_add_spent_time",
		mcp.WithDescription("Add spent time for an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidAddSpentTimeHandler))
}

func postProjectsIdIssuesIssueIidAddSpentTimeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidAddSpentTimeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidAddSpentTime(ctx, req.Id, req.IssueIid, req.Params, authorizationHeader))
}

type PostProjectsIdIssuesIssueIidResetSpentTimeRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerPostProjectsIdIssuesIssueIidResetSpentTime(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostProjectsIdIssuesIssueIidResetSpentTimeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_pjs_id_issues_issue_iid_reset_spent_time",
		mcp.WithDescription("Reset spent time for an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postProjectsIdIssuesIssueIidResetSpentTimeHandler))
}

func postProjectsIdIssuesIssueIidResetSpentTimeHandler(ctx context.Context, request mcp.CallToolRequest, req PostProjectsIdIssuesIssueIidResetSpentTimeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ProjectsIdIssuesIssueIidResetSpentTime(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidTimeStatsRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerGetProjectsIdIssuesIssueIidTimeStats(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidTimeStatsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_time_stats",
		mcp.WithDescription("Get time tracking stats"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidTimeStatsHandler))
}

func getProjectsIdIssuesIssueIidTimeStatsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidTimeStatsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidTimeStats(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidRelatedMergeRequestsRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerGetProjectsIdIssuesIssueIidRelatedMergeRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidRelatedMergeRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_related_mrs",
		mcp.WithDescription("List merge requests related to issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidRelatedMergeRequestsHandler))
}

func getProjectsIdIssuesIssueIidRelatedMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidRelatedMergeRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequests(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidClosedByRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project issue."`
}

func registerGetProjectsIdIssuesIssueIidClosedBy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidClosedByRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_closed_by",
		mcp.WithDescription("List merge requests that close a particular issue on merge"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidClosedByHandler))
}

func getProjectsIdIssuesIssueIidClosedByHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidClosedByRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidClosedBy(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidParticipantsRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerGetProjectsIdIssuesIssueIidParticipants(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidParticipantsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_participants",
		mcp.WithDescription("List participants in an issue"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidParticipantsHandler))
}

func getProjectsIdIssuesIssueIidParticipantsHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidParticipantsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidParticipants(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidUserAgentDetailRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerGetProjectsIdIssuesIssueIidUserAgentDetail(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidUserAgentDetailRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_user_agent_detail",
		mcp.WithDescription("Get user agent details"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidUserAgentDetailHandler))
}

func getProjectsIdIssuesIssueIidUserAgentDetailHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidUserAgentDetailRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidUserAgentDetail(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type GetProjectsIdIssuesIssueIidMetricImagesRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
}

func registerGetProjectsIdIssuesIssueIidMetricImages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetProjectsIdIssuesIssueIidMetricImagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pjs_id_issues_issue_iid_metric_images",
		mcp.WithDescription("List metric images"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getProjectsIdIssuesIssueIidMetricImagesHandler))
}

func getProjectsIdIssuesIssueIidMetricImagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetProjectsIdIssuesIssueIidMetricImagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4ProjectsIdIssuesIssueIidMetricImages(ctx, req.Id, req.IssueIid, authorizationHeader))
}

type DeleteProjectsIdIssuesIssueIidMetricImagesImageIdRequest struct {
	Id       string `json:"id" jsonschema:"description=The global ID or URL-encoded path of the project."`
	IssueIid int    `json:"issue_iid" jsonschema:"description=The internal ID of a project's issue."`
	ImageId  int    `json:"image_id" jsonschema:"description=The ID of the image."`
}

func registerDeleteProjectsIdIssuesIssueIidMetricImagesImageId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteProjectsIdIssuesIssueIidMetricImagesImageIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pjs_id_issues_issue_iid_metric_images_image_id",
		mcp.WithDescription("Delete metric image"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteProjectsIdIssuesIssueIidMetricImagesImageIdHandler))
}

func deleteProjectsIdIssuesIssueIidMetricImagesImageIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteProjectsIdIssuesIssueIidMetricImagesImageIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ProjectsIdIssuesIssueIidMetricImagesImageId(ctx, req.Id, req.IssueIid, req.ImageId, authorizationHeader))
}
