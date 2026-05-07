package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostGroupsIdAccessRequestsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
}

func registerPostGroupsIdAccessRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdAccessRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdAccessRequestsHandler))
}

func postGroupsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdAccessRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdAccessRequests(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdAccessRequestsRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
	Params *client.GetApiV4GroupsIdAccessRequestsParams `json:"params,omitempty"`
}

func registerGetGroupsIdAccessRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdAccessRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_access_requests",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdAccessRequestsHandler))
}

func getGroupsIdAccessRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdAccessRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdAccessRequests(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdAccessRequestsUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the access requester"`
}

func registerDeleteGroupsIdAccessRequestsUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdAccessRequestsUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_access_requests_user_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.11."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdAccessRequestsUserIdHandler))
}

func deleteGroupsIdAccessRequestsUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdAccessRequestsUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdAccessRequestsUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type GetGroupsIdEpicsEpicIidAwardEmojiRequest struct {
	Id      string                                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
	EpicIid int32                                                `json:"epic_iid" jsonschema:"description=ID ('iid' for merge requests/issues/epics, 'id' for snippets) of an awardable."`
	Params  *client.GetApiV4GroupsIdEpicsEpicIidAwardEmojiParams `json:"params,omitempty"`
}

func registerGetGroupsIdEpicsEpicIidAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdEpicsEpicIidAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_epics_epic_iid_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdEpicsEpicIidAwardEmojiHandler))
}

func getGroupsIdEpicsEpicIidAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdEpicsEpicIidAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdEpicsEpicIidAwardEmoji(ctx, req.Id, req.EpicIid, req.Params, authorizationHeader))
}

type DeleteGroupsIdEpicsEpicIidAwardEmojiAwardIdRequest struct {
	Id      int32 `json:"id" jsonschema:"description=null"`
	EpicIid int32 `json:"epic_iid" jsonschema:"description=null"`
	AwardId int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteGroupsIdEpicsEpicIidAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdEpicsEpicIidAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_epics_epic_iid_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdEpicsEpicIidAwardEmojiAwardIdHandler))
}

func deleteGroupsIdEpicsEpicIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdEpicsEpicIidAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId(ctx, req.Id, req.EpicIid, req.AwardId, authorizationHeader))
}

type GetGroupsIdEpicsEpicIidAwardEmojiAwardIdRequest struct {
	Id      int32 `json:"id" jsonschema:"description=null"`
	EpicIid int32 `json:"epic_iid" jsonschema:"description=null"`
	AwardId int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetGroupsIdEpicsEpicIidAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdEpicsEpicIidAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_epics_epic_iid_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdEpicsEpicIidAwardEmojiAwardIdHandler))
}

func getGroupsIdEpicsEpicIidAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdEpicsEpicIidAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId(ctx, req.Id, req.EpicIid, req.AwardId, authorizationHeader))
}

type GetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiRequest struct {
	Id      int32                                                           `json:"id" jsonschema:"description=null"`
	EpicIid int32                                                           `json:"epic_iid" jsonschema:"description=null"`
	NoteId  int32                                                           `json:"note_id" jsonschema:"description=null"`
	Params  *client.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiParams `json:"params,omitempty"`
}

func registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_epics_epic_iid_notes_note_id_award_emoji",
		mcp.WithDescription("Get a list of all emoji reactions for a specified awardable. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiHandler))
}

func getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(ctx, req.Id, req.EpicIid, req.NoteId, req.Params, authorizationHeader))
}

type DeleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id      int32 `json:"id" jsonschema:"description=null"`
	EpicIid int32 `json:"epic_iid" jsonschema:"description=null"`
	NoteId  int32 `json:"note_id" jsonschema:"description=null"`
	AwardId int32 `json:"award_id" jsonschema:"description=ID of an emoji reaction."`
}

func registerDeleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_epics_epic_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Only an administrator or the author of the reaction can delete an emoji reaction. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdHandler))
}

func deleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.EpicIid, req.NoteId, req.AwardId, authorizationHeader))
}

type GetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdRequest struct {
	Id      int32 `json:"id" jsonschema:"description=null"`
	EpicIid int32 `json:"epic_iid" jsonschema:"description=null"`
	NoteId  int32 `json:"note_id" jsonschema:"description=null"`
	AwardId int32 `json:"award_id" jsonschema:"description=ID of the emoji reaction."`
}

func registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_epics_epic_iid_notes_note_id_award_emoji_award_id",
		mcp.WithDescription("Get a single emoji reaction from an issue, snippet, or merge request. This feature was introduced in 8.9"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdHandler))
}

func getGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(ctx, req.Id, req.EpicIid, req.NoteId, req.AwardId, authorizationHeader))
}

type GetGroupsIdBadgesRequest struct {
	Id     string                               `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user."`
	Params *client.GetApiV4GroupsIdBadgesParams `json:"params,omitempty"`
}

func registerGetGroupsIdBadges(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdBadgesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_badges",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdBadgesHandler))
}

func getGroupsIdBadgesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdBadgesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdBadges(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdBadgesRenderRequest struct {
	Id     string                                     `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user."`
	Params *client.GetApiV4GroupsIdBadgesRenderParams `json:"params,omitempty"`
}

func registerGetGroupsIdBadgesRender(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdBadgesRenderRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_badges_render",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdBadgesRenderHandler))
}

func getGroupsIdBadgesRenderHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdBadgesRenderRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdBadgesRender(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdBadgesBadgeIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user."`
	BadgeId int32  `json:"badge_id" jsonschema:"description=The badge ID"`
}

func registerDeleteGroupsIdBadgesBadgeId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdBadgesBadgeIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_badges_badge_id",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdBadgesBadgeIdHandler))
}

func deleteGroupsIdBadgesBadgeIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdBadgesBadgeIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdBadgesBadgeId(ctx, req.Id, req.BadgeId, authorizationHeader))
}

type GetGroupsIdBadgesBadgeIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user."`
	BadgeId int32  `json:"badge_id" jsonschema:"description=The badge ID"`
}

func registerGetGroupsIdBadgesBadgeId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdBadgesBadgeIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_badges_badge_id",
		mcp.WithDescription("This feature was introduced in GitLab 10.6."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdBadgesBadgeIdHandler))
}

func getGroupsIdBadgesBadgeIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdBadgesBadgeIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdBadgesBadgeId(ctx, req.Id, req.BadgeId, authorizationHeader))
}

type GetGroupsIdCustomAttributesRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetGroupsIdCustomAttributes(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdCustomAttributesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_custom_attributes",
		mcp.WithDescription("Get all custom attributes on a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdCustomAttributesHandler))
}

func getGroupsIdCustomAttributesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdCustomAttributesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdCustomAttributes(ctx, req.Id, authorizationHeader))
}

type DeleteGroupsIdCustomAttributesKeyRequest struct {
	Id  int32  `json:"id" jsonschema:"description=null"`
	Key string `json:"key" jsonschema:"description=The key of the custom attribute"`
}

func registerDeleteGroupsIdCustomAttributesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdCustomAttributesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_custom_attributes_key",
		mcp.WithDescription("Delete a custom attribute on a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdCustomAttributesKeyHandler))
}

func deleteGroupsIdCustomAttributesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdCustomAttributesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdCustomAttributesKey(ctx, req.Id, req.Key, authorizationHeader))
}

type GetGroupsIdCustomAttributesKeyRequest struct {
	Id  int32  `json:"id" jsonschema:"description=null"`
	Key string `json:"key" jsonschema:"description=The key of the custom attribute"`
}

func registerGetGroupsIdCustomAttributesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdCustomAttributesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_custom_attributes_key",
		mcp.WithDescription("Get a custom attribute on a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdCustomAttributesKeyHandler))
}

func getGroupsIdCustomAttributesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdCustomAttributesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdCustomAttributesKey(ctx, req.Id, req.Key, authorizationHeader))
}

type GetGroupsRequest struct {
	Params *client.GetApiV4GroupsParams `json:"params,omitempty"`
}

func registerGetGroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps",
		mcp.WithDescription("Get a groups list"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsHandler))
}

func getGroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Groups(ctx, req.Params, authorizationHeader))
}

type DeleteGroupsIdRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerDeleteGroupsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id",
		mcp.WithDescription("Remove a group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdHandler))
}

func deleteGroupsIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsId(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdRequest struct {
	Id     string                         `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdParams `json:"params,omitempty"`
}

func registerGetGroupsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id",
		mcp.WithDescription("Get a single group, with containing projects."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdHandler))
}

func getGroupsIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsId(ctx, req.Id, req.Params, authorizationHeader))
}

type PostGroupsIdArchiveRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdArchive(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdArchiveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_archive",
		mcp.WithDescription("Archive a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdArchiveHandler))
}

func postGroupsIdArchiveHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdArchiveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdArchive(ctx, req.Id, authorizationHeader))
}

type PostGroupsIdUnarchiveRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdUnarchive(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdUnarchiveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_unarchive",
		mcp.WithDescription("Unarchive a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdUnarchiveHandler))
}

func postGroupsIdUnarchiveHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdUnarchiveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdUnarchive(ctx, req.Id, authorizationHeader))
}

type PostGroupsIdRestoreRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdRestore(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdRestoreRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_restore",
		mcp.WithDescription("Restore a group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdRestoreHandler))
}

func postGroupsIdRestoreHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdRestoreRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdRestore(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdGroupsSharedRequest struct {
	Id     string                                     `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdGroupsSharedParams `json:"params,omitempty"`
}

func registerGetGroupsIdGroupsShared(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdGroupsSharedRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_groups_shared",
		mcp.WithDescription("Get a list of shared groups this group was invited to"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdGroupsSharedHandler))
}

func getGroupsIdGroupsSharedHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdGroupsSharedRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdGroupsShared(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdInvitedGroupsRequest struct {
	Id     string                                      `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdInvitedGroupsParams `json:"params,omitempty"`
}

func registerGetGroupsIdInvitedGroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdInvitedGroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_invited_groups",
		mcp.WithDescription("Get a list of invited groups in this group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdInvitedGroupsHandler))
}

func getGroupsIdInvitedGroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdInvitedGroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdInvitedGroups(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdProjectsRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdProjectsParams `json:"params,omitempty"`
}

func registerGetGroupsIdProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdProjectsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pjs",
		mcp.WithDescription("Get a list of projects in this group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdProjectsHandler))
}

func getGroupsIdProjectsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdProjectsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdProjects(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdProjectsSharedRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdProjectsSharedParams `json:"params,omitempty"`
}

func registerGetGroupsIdProjectsShared(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdProjectsSharedRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pjs_shared",
		mcp.WithDescription("Get a list of shared projects in this group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdProjectsSharedHandler))
}

func getGroupsIdProjectsSharedHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdProjectsSharedRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdProjectsShared(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdSubgroupsRequest struct {
	Id     string                                  `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdSubgroupsParams `json:"params,omitempty"`
}

func registerGetGroupsIdSubgroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdSubgroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_subgroups",
		mcp.WithDescription("Get a list of subgroups in this group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdSubgroupsHandler))
}

func getGroupsIdSubgroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdSubgroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdSubgroups(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdDescendantGroupsRequest struct {
	Id     string                                         `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdDescendantGroupsParams `json:"params,omitempty"`
}

func registerGetGroupsIdDescendantGroups(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdDescendantGroupsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_descendant_groups",
		mcp.WithDescription("Get a list of descendant groups of this group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdDescendantGroupsHandler))
}

func getGroupsIdDescendantGroupsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdDescendantGroupsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdDescendantGroups(ctx, req.Id, req.Params, authorizationHeader))
}

type PostGroupsIdProjectsProjectIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID of a group"`
	ProjectId string `json:"project_id" jsonschema:"description=The ID or path of the project"`
}

func registerPostGroupsIdProjectsProjectId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdProjectsProjectIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_pjs_project_id",
		mcp.WithDescription("Transfer a project to the group namespace. Available only for admin."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdProjectsProjectIdHandler))
}

func postGroupsIdProjectsProjectIdHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdProjectsProjectIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdProjectsProjectId(ctx, req.Id, req.ProjectId, authorizationHeader))
}

type GetGroupsIdTransferLocationsRequest struct {
	Id     string                                          `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdTransferLocationsParams `json:"params,omitempty"`
}

func registerGetGroupsIdTransferLocations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdTransferLocationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_transfer_locations",
		mcp.WithDescription("Get the groups to where the current group can be transferred to"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdTransferLocationsHandler))
}

func getGroupsIdTransferLocationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdTransferLocationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdTransferLocations(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdShareGroupIdRequest struct {
	Id      string `json:"id" jsonschema:"description=The ID of a group"`
	GroupId int32  `json:"group_id" jsonschema:"description=The ID of the shared group"`
}

func registerDeleteGroupsIdShareGroupId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdShareGroupIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_share_group_id",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdShareGroupIdHandler))
}

func deleteGroupsIdShareGroupIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdShareGroupIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdShareGroupId(ctx, req.Id, req.GroupId, authorizationHeader))
}

type PostGroupsIdLdapSyncRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerPostGroupsIdLdapSync(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdLdapSyncRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_ldap_sync",
		mcp.WithDescription("Sync a group with LDAP."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdLdapSyncHandler))
}

func postGroupsIdLdapSyncHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdLdapSyncRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdLdapSync(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdAuditEventsRequest struct {
	Id     int32                                     `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdAuditEventsParams `json:"params,omitempty"`
}

func registerGetGroupsIdAuditEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdAuditEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_audit_events",
		mcp.WithDescription("Get a list of audit events in this group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdAuditEventsHandler))
}

func getGroupsIdAuditEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdAuditEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdAuditEvents(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdAuditEventsAuditEventIdRequest struct {
	Id           int32 `json:"id" jsonschema:"description=null"`
	AuditEventId int32 `json:"audit_event_id" jsonschema:"description=The ID of the audit event"`
}

func registerGetGroupsIdAuditEventsAuditEventId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdAuditEventsAuditEventIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_audit_events_audit_event_id",
		mcp.WithDescription("Get a specific audit event in this group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdAuditEventsAuditEventIdHandler))
}

func getGroupsIdAuditEventsAuditEventIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdAuditEventsAuditEventIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdAuditEventsAuditEventId(ctx, req.Id, req.AuditEventId, authorizationHeader))
}

type GetGroupsIdSamlUsersRequest struct {
	Id     int32                                   `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdSamlUsersParams `json:"params,omitempty"`
}

func registerGetGroupsIdSamlUsers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdSamlUsersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_saml_users",
		mcp.WithDescription("Get a list of SAML users of the group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdSamlUsersHandler))
}

func getGroupsIdSamlUsersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdSamlUsersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdSamlUsers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdProvisionedUsersRequest struct {
	Id     int32                                          `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdProvisionedUsersParams `json:"params,omitempty"`
}

func registerGetGroupsIdProvisionedUsers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdProvisionedUsersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_provisioned_users",
		mcp.WithDescription("Get a list of users provisioned by the group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdProvisionedUsersHandler))
}

func getGroupsIdProvisionedUsersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdProvisionedUsersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdProvisionedUsers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdUsersRequest struct {
	Id     int32                               `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdUsersParams `json:"params,omitempty"`
}

func registerGetGroupsIdUsers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdUsersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_users",
		mcp.WithDescription("Get a list of users for the group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdUsersHandler))
}

func getGroupsIdUsersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdUsersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdUsers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdSshCertificatesRequest struct {
	Id     int32                                         `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdSshCertificatesParams `json:"params,omitempty"`
}

func registerGetGroupsIdSshCertificates(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdSshCertificatesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_ssh_certificates",
		mcp.WithDescription("Get a list of ssh certificates created for a group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdSshCertificatesHandler))
}

func getGroupsIdSshCertificatesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdSshCertificatesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdSshCertificates(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdSshCertificatesSshCertificatesIdRequest struct {
	Id                int32 `json:"id" jsonschema:"description=null"`
	SshCertificatesId int32 `json:"ssh_certificates_id" jsonschema:"description=null"`
}

func registerDeleteGroupsIdSshCertificatesSshCertificatesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdSshCertificatesSshCertificatesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_ssh_certificates_ssh_certificates_id",
		mcp.WithDescription("Removes a Groups::SshCertificate"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdSshCertificatesSshCertificatesIdHandler))
}

func deleteGroupsIdSshCertificatesSshCertificatesIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdSshCertificatesSshCertificatesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdSshCertificatesSshCertificatesId(ctx, req.Id, req.SshCertificatesId, authorizationHeader))
}

type GetGroupsIdRunnersRequest struct {
	Id     string                                `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdRunnersParams `json:"params,omitempty"`
}

func registerGetGroupsIdRunners(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdRunnersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_runners",
		mcp.WithDescription("List all runners available in the group as well as its ancestor groups, including any allowed shared runners."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdRunnersHandler))
}

func getGroupsIdRunnersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdRunnersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdRunners(ctx, req.Id, req.Params, authorizationHeader))
}

type PostGroupsIdRunnersResetRegistrationTokenRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdRunnersResetRegistrationToken(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdRunnersResetRegistrationTokenRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_runners_reset_registration_token",
		mcp.WithDescription("Reset runner registration token"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdRunnersResetRegistrationTokenHandler))
}

func postGroupsIdRunnersResetRegistrationTokenHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdRunnersResetRegistrationTokenRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdRunnersResetRegistrationToken(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameRequest struct {
	Id             string `json:"id" jsonschema:"description=The group ID or full group path."`
	Distribution   string `json:"distribution" jsonschema:"description=The Debian Codename or Suite"`
	ProjectId      int32  `json:"project_id" jsonschema:"description=The Project Id"`
	Letter         string `json:"letter" jsonschema:"description=The Debian Classification (first-letter or lib-first-letter)"`
	PackageName    string `json:"package_name" jsonschema:"description=The Debian Source Package Name"`
	PackageVersion string `json:"package_version" jsonschema:"description=The Debian Source Package Version"`
	FileName       string `json:"file_name" jsonschema:"description=The Debian File Name"`
}

func registerGetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs_debian_pool_distribution_project_id_letter_package_name_package_version_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 14.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameHandler))
}

func getGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(ctx, req.Id, req.Distribution, req.ProjectId, req.Letter, req.PackageName, req.PackageVersion, req.FileName, authorizationHeader))
}

type DeleteGroupsIdDependencyProxyCacheRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
}

func registerDeleteGroupsIdDependencyProxyCache(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdDependencyProxyCacheRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_dependency_proxy_cache",
		mcp.WithDescription("Schedules for deletion the cached manifests and blobs for a group.This endpoint requires the Owner role for the group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdDependencyProxyCacheHandler))
}

func deleteGroupsIdDependencyProxyCacheHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdDependencyProxyCacheRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdDependencyProxyCache(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdDeployTokensRequest struct {
	Id     int32                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
	Params *client.GetApiV4GroupsIdDeployTokensParams `json:"params,omitempty"`
}

func registerGetGroupsIdDeployTokens(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdDeployTokensRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_deploy_tokens",
		mcp.WithDescription("Get a list of a group's deploy tokens. This feature was introduced in GitLab 12.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdDeployTokensHandler))
}

func getGroupsIdDeployTokensHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdDeployTokensRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdDeployTokens(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdDeployTokensTokenIdRequest struct {
	Id      int32 `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
	TokenId int32 `json:"token_id" jsonschema:"description=The ID of the deploy token"`
}

func registerDeleteGroupsIdDeployTokensTokenId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdDeployTokensTokenIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_deploy_tokens_token_id",
		mcp.WithDescription("Removes a deploy token from the group. This feature was introduced in GitLab 12.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdDeployTokensTokenIdHandler))
}

func deleteGroupsIdDeployTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdDeployTokensTokenIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdDeployTokensTokenId(ctx, req.Id, req.TokenId, authorizationHeader))
}

type GetGroupsIdDeployTokensTokenIdRequest struct {
	Id      int32 `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
	TokenId int32 `json:"token_id" jsonschema:"description=The ID of the deploy token"`
}

func registerGetGroupsIdDeployTokensTokenId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdDeployTokensTokenIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_deploy_tokens_token_id",
		mcp.WithDescription("Get a single group's deploy token by ID. This feature was introduced in GitLab 14.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdDeployTokensTokenIdHandler))
}

func getGroupsIdDeployTokensTokenIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdDeployTokensTokenIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdDeployTokensTokenId(ctx, req.Id, req.TokenId, authorizationHeader))
}

type GetGroupsIdAvatarRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of the group"`
}

func registerGetGroupsIdAvatar(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdAvatarRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_avatar",
		mcp.WithDescription("This feature was introduced in GitLab 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdAvatarHandler))
}

func getGroupsIdAvatarHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdAvatarRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdAvatar(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdClustersRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The ID of the group"`
	Params *client.GetApiV4GroupsIdClustersParams `json:"params,omitempty"`
}

func registerGetGroupsIdClusters(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdClustersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 12.1. Returns a list of group clusters."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdClustersHandler))
}

func getGroupsIdClustersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdClustersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdClusters(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdClustersClusterIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID of the group"`
	ClusterId int32  `json:"cluster_id" jsonschema:"description=The Cluster ID"`
}

func registerDeleteGroupsIdClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 12.1. Deletes an existing group cluster. Does not remove existing resources within the connected Kubernetes cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdClustersClusterIdHandler))
}

func deleteGroupsIdClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdClustersClusterId(ctx, req.Id, req.ClusterId, authorizationHeader))
}

type GetGroupsIdClustersClusterIdRequest struct {
	Id        string `json:"id" jsonschema:"description=The ID of the group"`
	ClusterId int32  `json:"cluster_id" jsonschema:"description=The cluster ID"`
}

func registerGetGroupsIdClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 12.1. Gets a single group cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdClustersClusterIdHandler))
}

func getGroupsIdClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdClustersClusterId(ctx, req.Id, req.ClusterId, authorizationHeader))
}

type GetGroupsIdRegistryRepositoriesRequest struct {
	Id     string                                             `json:"id" jsonschema:"description=The ID or URL-encoded path of the group accessible by the authenticated user"`
	Params *client.GetApiV4GroupsIdRegistryRepositoriesParams `json:"params,omitempty"`
}

func registerGetGroupsIdRegistryRepositories(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdRegistryRepositoriesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_registry_repositories",
		mcp.WithDescription("Get a list of registry repositories in a group. This feature was introduced in GitLab 12.2."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdRegistryRepositoriesHandler))
}

func getGroupsIdRegistryRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdRegistryRepositoriesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdRegistryRepositories(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdDebianDistributionsRequest struct {
	Id     string                                            `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
	Params *client.GetApiV4GroupsIdDebianDistributionsParams `json:"params,omitempty"`
}

func registerGetGroupsIdDebianDistributions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdDebianDistributionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_debian_distributions",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdDebianDistributionsHandler))
}

func getGroupsIdDebianDistributionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdDebianDistributionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdDebianDistributions(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdDebianDistributionsCodenameRequest struct {
	Id       string                                                       `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
	Codename string                                                       `json:"codename" jsonschema:"description=The Debian Codename"`
	Params   *client.DeleteApiV4GroupsIdDebianDistributionsCodenameParams `json:"params,omitempty"`
}

func registerDeleteGroupsIdDebianDistributionsCodename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdDebianDistributionsCodenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_debian_distributions_codename",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdDebianDistributionsCodenameHandler))
}

func deleteGroupsIdDebianDistributionsCodenameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdDebianDistributionsCodenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdDebianDistributionsCodename(ctx, req.Id, req.Codename, req.Params, authorizationHeader))
}

type GetGroupsIdDebianDistributionsCodenameRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
	Codename string `json:"codename" jsonschema:"description=The Debian Codename"`
}

func registerGetGroupsIdDebianDistributionsCodename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdDebianDistributionsCodenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_debian_distributions_codename",
		mcp.WithDescription("This feature was introduced in 14.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdDebianDistributionsCodenameHandler))
}

func getGroupsIdDebianDistributionsCodenameHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdDebianDistributionsCodenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdDebianDistributionsCodename(ctx, req.Id, req.Codename, authorizationHeader))
}

type GetGroupsIdDebianDistributionsCodenameKeyAscRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
	Codename string `json:"codename" jsonschema:"description=The Debian Codename"`
}

func registerGetGroupsIdDebianDistributionsCodenameKeyAsc(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdDebianDistributionsCodenameKeyAscRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_debian_distributions_codename_key_asc",
		mcp.WithDescription("This feature was introduced in 14.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdDebianDistributionsCodenameKeyAscHandler))
}

func getGroupsIdDebianDistributionsCodenameKeyAscHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdDebianDistributionsCodenameKeyAscRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc(ctx, req.Id, req.Codename, authorizationHeader))
}

type GetGroupsIdExportDownloadRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerGetGroupsIdExportDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdExportDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_export_download",
		mcp.WithDescription("This feature was introduced in GitLab 12.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdExportDownloadHandler))
}

func getGroupsIdExportDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdExportDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdExportDownload(ctx, req.Id, authorizationHeader))
}

type PostGroupsIdExportRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdExport(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdExportRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_export",
		mcp.WithDescription("This feature was introduced in GitLab 12.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdExportHandler))
}

func postGroupsIdExportHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdExportRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdExport(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdExportRelationsDownloadRequest struct {
	Id     string                                                `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdExportRelationsDownloadParams `json:"params,omitempty"`
}

func registerGetGroupsIdExportRelationsDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdExportRelationsDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_export_relations_download",
		mcp.WithDescription("This feature was introduced in GitLab 13.12"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdExportRelationsDownloadHandler))
}

func getGroupsIdExportRelationsDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdExportRelationsDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdExportRelationsDownload(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdExportRelationsStatusRequest struct {
	Id     string                                              `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdExportRelationsStatusParams `json:"params,omitempty"`
}

func registerGetGroupsIdExportRelationsStatus(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdExportRelationsStatusRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_export_relations_status",
		mcp.WithDescription("This feature was introduced in GitLab 13.12"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdExportRelationsStatusHandler))
}

func getGroupsIdExportRelationsStatusHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdExportRelationsStatusRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdExportRelationsStatus(ctx, req.Id, req.Params, authorizationHeader))
}

type PostGroupsImportAuthorizeRequest struct {
}

func registerPostGroupsImportAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsImportAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_import_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 12.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsImportAuthorizeHandler))
}

func postGroupsImportAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsImportAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsImportAuthorize(ctx, authorizationHeader))
}

type GetGroupsIdPackagesRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=ID or URL-encoded path of the group"`
	Params *client.GetApiV4GroupsIdPackagesParams `json:"params,omitempty"`
}

func registerGetGroupsIdPackages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs",
		mcp.WithDescription("Get a list of project packages at the group level. This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesHandler))
}

func getGroupsIdPackagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackages(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdPlaceholderReassignmentsRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerGetGroupsIdPlaceholderReassignments(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPlaceholderReassignmentsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_placeholder_reassignments",
		mcp.WithDescription("This feature was added in GitLab 17.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPlaceholderReassignmentsHandler))
}

func getGroupsIdPlaceholderReassignmentsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPlaceholderReassignmentsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPlaceholderReassignments(ctx, req.Id, authorizationHeader))
}

type PostGroupsIdPlaceholderReassignmentsAuthorizeRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdPlaceholderReassignmentsAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdPlaceholderReassignmentsAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_placeholder_reassignments_authorize",
		mcp.WithDescription("This feature was introduced in GitLab 17.10"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdPlaceholderReassignmentsAuthorizeHandler))
}

func postGroupsIdPlaceholderReassignmentsAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdPlaceholderReassignmentsAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdPlaceholderReassignmentsAuthorize(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdVariablesRequest struct {
	Id     string                                  `json:"id" jsonschema:"description=The ID of a group or URL-encoded path of the group owned by the authenticated user"`
	Params *client.GetApiV4GroupsIdVariablesParams `json:"params,omitempty"`
}

func registerGetGroupsIdVariables(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdVariablesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_variables",
		mcp.WithDescription("Get a list of group-level variables"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdVariablesHandler))
}

func getGroupsIdVariablesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdVariablesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdVariables(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdVariablesKeyRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID of a group or URL-encoded path of the group owned by the authenticated user"`
	Key string `json:"key" jsonschema:"description=The key of a variable"`
}

func registerDeleteGroupsIdVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_variables_key",
		mcp.WithDescription("Delete an existing variable from a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdVariablesKeyHandler))
}

func deleteGroupsIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdVariablesKey(ctx, req.Id, req.Key, authorizationHeader))
}

type GetGroupsIdVariablesKeyRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID of a group or URL-encoded path of the group owned by the authenticated user"`
	Key string `json:"key" jsonschema:"description=The key of the variable"`
}

func registerGetGroupsIdVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_variables_key",
		mcp.WithDescription("Get the details of a group’s specific variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdVariablesKeyHandler))
}

func getGroupsIdVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdVariablesKey(ctx, req.Id, req.Key, authorizationHeader))
}

type GetGroupsIdIntegrationsRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetGroupsIdIntegrations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdIntegrationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_integrations",
		mcp.WithDescription("Get a list of all active integrations."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdIntegrationsHandler))
}

func getGroupsIdIntegrationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdIntegrationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdIntegrations(ctx, req.Id, authorizationHeader))
}

type DeleteGroupsIdIntegrationsSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The name of the integration"`
}

func registerDeleteGroupsIdIntegrationsSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdIntegrationsSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_integrations_slug",
		mcp.WithDescription("Disable the integration. Integration settings are preserved."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdIntegrationsSlugHandler))
}

func deleteGroupsIdIntegrationsSlugHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdIntegrationsSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdIntegrationsSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetGroupsIdIntegrationsSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The name of the integration"`
}

func registerGetGroupsIdIntegrationsSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdIntegrationsSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_integrations_slug",
		mcp.WithDescription("Get the integration settings."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdIntegrationsSlugHandler))
}

func getGroupsIdIntegrationsSlugHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdIntegrationsSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdIntegrationsSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetGroupsIdInvitationsRequest struct {
	Id     string                                    `json:"id" jsonschema:"description=The group ID"`
	Params *client.GetApiV4GroupsIdInvitationsParams `json:"params,omitempty"`
}

func registerGetGroupsIdInvitations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdInvitationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_invitations",
		mcp.WithDescription("This feature was introduced in GitLab 13.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdInvitationsHandler))
}

func getGroupsIdInvitationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdInvitationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdInvitations(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdInvitationsEmailRequest struct {
	Id    string `json:"id" jsonschema:"description=The group ID"`
	Email string `json:"email" jsonschema:"description=The email address of the invitation"`
}

func registerDeleteGroupsIdInvitationsEmail(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdInvitationsEmailRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_invitations_email",
		mcp.WithDescription("Removes an invitation from a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdInvitationsEmailHandler))
}

func deleteGroupsIdInvitationsEmailHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdInvitationsEmailRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdInvitationsEmail(ctx, req.Id, req.Email, authorizationHeader))
}

type GetGroupsIdUploadsRequest struct {
	Id     int32                                 `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdUploadsParams `json:"params,omitempty"`
}

func registerGetGroupsIdUploads(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdUploadsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_uploads",
		mcp.WithDescription("Get the list of uploads of a group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdUploadsHandler))
}

func getGroupsIdUploadsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdUploadsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdUploads(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdUploadsUploadIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	UploadId int32 `json:"upload_id" jsonschema:"description=The ID of a group upload"`
}

func registerDeleteGroupsIdUploadsUploadId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdUploadsUploadIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_uploads_upload_id",
		mcp.WithDescription("Delete a single group upload"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdUploadsUploadIdHandler))
}

func deleteGroupsIdUploadsUploadIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdUploadsUploadIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdUploadsUploadId(ctx, req.Id, req.UploadId, authorizationHeader))
}

type GetGroupsIdUploadsUploadIdRequest struct {
	Id       int32 `json:"id" jsonschema:"description=null"`
	UploadId int32 `json:"upload_id" jsonschema:"description=The ID of a group upload"`
}

func registerGetGroupsIdUploadsUploadId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdUploadsUploadIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_uploads_upload_id",
		mcp.WithDescription("Download a single group upload by ID"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdUploadsUploadIdHandler))
}

func getGroupsIdUploadsUploadIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdUploadsUploadIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdUploadsUploadId(ctx, req.Id, req.UploadId, authorizationHeader))
}

type DeleteGroupsIdUploadsSecretFilenameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=null"`
	Secret   string `json:"secret" jsonschema:"description=The 32-character secret of a group upload"`
	Filename string `json:"filename" jsonschema:"description=The filename of a group upload"`
}

func registerDeleteGroupsIdUploadsSecretFilename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdUploadsSecretFilenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_uploads_secret_filename",
		mcp.WithDescription("Delete a single group upload by secret and filename"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdUploadsSecretFilenameHandler))
}

func deleteGroupsIdUploadsSecretFilenameHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdUploadsSecretFilenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdUploadsSecretFilename(ctx, req.Id, req.Secret, req.Filename, authorizationHeader))
}

type GetGroupsIdUploadsSecretFilenameRequest struct {
	Id       int32  `json:"id" jsonschema:"description=null"`
	Secret   string `json:"secret" jsonschema:"description=The 32-character secret of a group upload"`
	Filename string `json:"filename" jsonschema:"description=The filename of a group upload"`
}

func registerGetGroupsIdUploadsSecretFilename(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdUploadsSecretFilenameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_uploads_secret_filename",
		mcp.WithDescription("Download a single project upload by secret and filename"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdUploadsSecretFilenameHandler))
}

func getGroupsIdUploadsSecretFilenameHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdUploadsSecretFilenameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdUploadsSecretFilename(ctx, req.Id, req.Secret, req.Filename, authorizationHeader))
}

type GetGroupsIdMembersRequest struct {
	Id     string                                `json:"id" jsonschema:"description=The group ID"`
	Params *client.GetApiV4GroupsIdMembersParams `json:"params,omitempty"`
}

func registerGetGroupsIdMembers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdMembersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_members",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdMembersHandler))
}

func getGroupsIdMembersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdMembersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdMembers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdMembersAllRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=The group ID"`
	Params *client.GetApiV4GroupsIdMembersAllParams `json:"params,omitempty"`
}

func registerGetGroupsIdMembersAll(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdMembersAllRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_members_all",
		mcp.WithDescription("Gets a list of group or project members viewable by the authenticated user, including those who gained membership through ancestor group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdMembersAllHandler))
}

func getGroupsIdMembersAllHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdMembersAllRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdMembersAll(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdMembersUserIdRequest struct {
	Id     string                                         `json:"id" jsonschema:"description=The group ID"`
	UserId int32                                          `json:"user_id" jsonschema:"description=The user ID of the member"`
	Params *client.DeleteApiV4GroupsIdMembersUserIdParams `json:"params,omitempty"`
}

func registerDeleteGroupsIdMembersUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdMembersUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_members_user_id",
		mcp.WithDescription("Removes a user from a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdMembersUserIdHandler))
}

func deleteGroupsIdMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdMembersUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdMembersUserId(ctx, req.Id, req.UserId, req.Params, authorizationHeader))
}

type GetGroupsIdMembersUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The group ID"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerGetGroupsIdMembersUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdMembersUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_members_user_id",
		mcp.WithDescription("Gets a member of a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdMembersUserIdHandler))
}

func getGroupsIdMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdMembersUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdMembersUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type GetGroupsIdMembersAllUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The group ID"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerGetGroupsIdMembersAllUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdMembersAllUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_members_all_user_id",
		mcp.WithDescription("Gets a member of a group or project, including those who gained membership through ancestor group"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdMembersAllUserIdHandler))
}

func getGroupsIdMembersAllUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdMembersAllUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdMembersAllUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type DeleteGroupsIdMembersUserIdOverrideRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID of a group"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerDeleteGroupsIdMembersUserIdOverride(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdMembersUserIdOverrideRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_members_user_id_override",
		mcp.WithDescription("Remove an LDAP group member access level override."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdMembersUserIdOverrideHandler))
}

func deleteGroupsIdMembersUserIdOverrideHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdMembersUserIdOverrideRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdMembersUserIdOverride(ctx, req.Id, req.UserId, authorizationHeader))
}

type PostGroupsIdMembersUserIdOverrideRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID of a group"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerPostGroupsIdMembersUserIdOverride(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdMembersUserIdOverrideRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_members_user_id_override",
		mcp.WithDescription("Overrides the access level of an LDAP group member."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdMembersUserIdOverrideHandler))
}

func postGroupsIdMembersUserIdOverrideHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdMembersUserIdOverrideRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdMembersUserIdOverride(ctx, req.Id, req.UserId, authorizationHeader))
}

type PutGroupsIdMembersMemberIdApproveRequest struct {
	Id       string `json:"id" jsonschema:"description=The ID of a group"`
	MemberId int32  `json:"member_id" jsonschema:"description=The ID of the member requiring approval"`
}

func registerPutGroupsIdMembersMemberIdApprove(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutGroupsIdMembersMemberIdApproveRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_grps_id_members_member_id_approve",
		mcp.WithDescription("Approves a pending member"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putGroupsIdMembersMemberIdApproveHandler))
}

func putGroupsIdMembersMemberIdApproveHandler(ctx context.Context, request mcp.CallToolRequest, req PutGroupsIdMembersMemberIdApproveRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4GroupsIdMembersMemberIdApprove(ctx, req.Id, req.MemberId, authorizationHeader))
}

type PostGroupsIdMembersApproveAllRequest struct {
	Id string `json:"id" jsonschema:"description=The ID of a group"`
}

func registerPostGroupsIdMembersApproveAll(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdMembersApproveAllRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_members_approve_all",
		mcp.WithDescription("Approves all pending members"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdMembersApproveAllHandler))
}

func postGroupsIdMembersApproveAllHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdMembersApproveAllRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdMembersApproveAll(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdPendingMembersRequest struct {
	Id     string                                       `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdPendingMembersParams `json:"params,omitempty"`
}

func registerGetGroupsIdPendingMembers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPendingMembersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pending_members",
		mcp.WithDescription("Lists all pending members for a group including invited users"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPendingMembersHandler))
}

func getGroupsIdPendingMembersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPendingMembersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPendingMembers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdBillableMembersRequest struct {
	Id     string                                        `json:"id" jsonschema:"description=The ID of a group"`
	Params *client.GetApiV4GroupsIdBillableMembersParams `json:"params,omitempty"`
}

func registerGetGroupsIdBillableMembers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdBillableMembersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_billable_members",
		mcp.WithDescription("Gets a list of billable users of top-level group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdBillableMembersHandler))
}

func getGroupsIdBillableMembersHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdBillableMembersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdBillableMembers(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdBillableMembersUserIdMembershipsRequest struct {
	Id     string                                                         `json:"id" jsonschema:"description=The ID of a group"`
	UserId int32                                                          `json:"user_id" jsonschema:"description=The user ID of the member"`
	Params *client.GetApiV4GroupsIdBillableMembersUserIdMembershipsParams `json:"params,omitempty"`
}

func registerGetGroupsIdBillableMembersUserIdMemberships(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdBillableMembersUserIdMembershipsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_billable_members_user_id_memberships",
		mcp.WithDescription("Get the direct memberships of a billable user of a top-level group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdBillableMembersUserIdMembershipsHandler))
}

func getGroupsIdBillableMembersUserIdMembershipsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdBillableMembersUserIdMembershipsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdBillableMembersUserIdMemberships(ctx, req.Id, req.UserId, req.Params, authorizationHeader))
}

type GetGroupsIdBillableMembersUserIdIndirectRequest struct {
	Id     string                                                      `json:"id" jsonschema:"description=The ID of a group"`
	UserId int32                                                       `json:"user_id" jsonschema:"description=The user ID of the member"`
	Params *client.GetApiV4GroupsIdBillableMembersUserIdIndirectParams `json:"params,omitempty"`
}

func registerGetGroupsIdBillableMembersUserIdIndirect(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdBillableMembersUserIdIndirectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_billable_members_user_id_indirect",
		mcp.WithDescription("Get the indirect memberships of a billable user of a top-level group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdBillableMembersUserIdIndirectHandler))
}

func getGroupsIdBillableMembersUserIdIndirectHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdBillableMembersUserIdIndirectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdBillableMembersUserIdIndirect(ctx, req.Id, req.UserId, req.Params, authorizationHeader))
}

type DeleteGroupsIdBillableMembersUserIdRequest struct {
	Id     string `json:"id" jsonschema:"description=The ID of a group"`
	UserId int32  `json:"user_id" jsonschema:"description=The user ID of the member"`
}

func registerDeleteGroupsIdBillableMembersUserId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdBillableMembersUserIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_billable_members_user_id",
		mcp.WithDescription("Removes a billable member from a group or project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdBillableMembersUserIdHandler))
}

func deleteGroupsIdBillableMembersUserIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdBillableMembersUserIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdBillableMembersUserId(ctx, req.Id, req.UserId, authorizationHeader))
}

type GetGroupsIdMergeRequestsRequest struct {
	Id     string                                      `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user."`
	Params *client.GetApiV4GroupsIdMergeRequestsParams `json:"params,omitempty"`
}

func registerGetGroupsIdMergeRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdMergeRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_mrs",
		mcp.WithDescription("Get all merge requests for this group and its subgroups."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdMergeRequestsHandler))
}

func getGroupsIdMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdMergeRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdMergeRequests(ctx, req.Id, req.Params, authorizationHeader))
}

type PostGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
}

func registerPostGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_pkgs_npm_npm_v1_security_advisories_bulk",
		mcp.WithDescription("This feature was introduced in GitLab 15.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkHandler))
}

func postGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx, req.Id, authorizationHeader))
}

type PostGroupsIdPackagesNpmNpmV1SecurityAuditsQuickRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of the group"`
}

func registerPostGroupsIdPackagesNpmNpmV1SecurityAuditsQuick(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGroupsIdPackagesNpmNpmV1SecurityAuditsQuickRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_grps_id_pkgs_npm_npm_v1_security_audits_quick",
		mcp.WithDescription("This feature was introduced in GitLab 15.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGroupsIdPackagesNpmNpmV1SecurityAuditsQuickHandler))
}

func postGroupsIdPackagesNpmNpmV1SecurityAuditsQuickHandler(ctx context.Context, request mcp.CallToolRequest, req PostGroupsIdPackagesNpmNpmV1SecurityAuditsQuickRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdPackagesNugetIndexRequest struct {
	Id int32 `json:"id" jsonschema:"description=The group ID or full group path."`
}

func registerGetGroupsIdPackagesNugetIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesNugetIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs_nuget_index",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesNugetIndexHandler))
}

func getGroupsIdPackagesNugetIndexHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesNugetIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackagesNugetIndex(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdPackagesNugetV2Request struct {
	Id int32 `json:"id" jsonschema:"description=The group ID or full group path."`
}

func registerGetGroupsIdPackagesNugetV2(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesNugetV2Request{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs_nuget_v2",
		mcp.WithDescription("This feature was introduced in GitLab 16.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesNugetV2Handler))
}

func getGroupsIdPackagesNugetV2Handler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesNugetV2Request) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackagesNugetV2(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdPackagesNugetV2MetadataRequest struct {
	Id int32 `json:"id" jsonschema:"description=The group ID or full group path."`
}

func registerGetGroupsIdPackagesNugetV2Metadata(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesNugetV2MetadataRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs_nuget_v2_metadata",
		mcp.WithDescription("This feature was introduced in GitLab 16.3"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesNugetV2MetadataHandler))
}

func getGroupsIdPackagesNugetV2MetadataHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesNugetV2MetadataRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackagesNugetV2Metadata(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdPackagesNugetQueryRequest struct {
	Id     int32                                            `json:"id" jsonschema:"description=The group ID or full group path."`
	Params *client.GetApiV4GroupsIdPackagesNugetQueryParams `json:"params,omitempty"`
}

func registerGetGroupsIdPackagesNugetQuery(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesNugetQueryRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs_nuget_query",
		mcp.WithDescription("This feature was introduced in GitLab 12.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesNugetQueryHandler))
}

func getGroupsIdPackagesNugetQueryHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesNugetQueryRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackagesNugetQuery(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdPackagesPypiSimpleRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID or full path of the group."`
}

func registerGetGroupsIdPackagesPypiSimple(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdPackagesPypiSimpleRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_pkgs_pypi_simple",
		mcp.WithDescription("This feature was introduced in GitLab 15.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdPackagesPypiSimpleHandler))
}

func getGroupsIdPackagesPypiSimpleHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdPackagesPypiSimpleRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdPackagesPypiSimple(ctx, req.Id, authorizationHeader))
}

type GetGroupsIdReleasesRequest struct {
	Id     string                                 `json:"id" jsonschema:"description=The ID or URL-encoded path of the group owned by the authenticated user"`
	Params *client.GetApiV4GroupsIdReleasesParams `json:"params,omitempty"`
}

func registerGetGroupsIdReleases(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdReleasesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_releases",
		mcp.WithDescription("Returns a list of group releases."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdReleasesHandler))
}

func getGroupsIdReleasesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdReleasesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdReleases(ctx, req.Id, req.Params, authorizationHeader))
}

type GetGroupsIdWikisRequest struct {
	Id     int32                               `json:"id" jsonschema:"description=null"`
	Params *client.GetApiV4GroupsIdWikisParams `json:"params,omitempty"`
}

func registerGetGroupsIdWikis(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdWikisRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_wikis",
		mcp.WithDescription("Get a list of wiki pages"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdWikisHandler))
}

func getGroupsIdWikisHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdWikisRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdWikis(ctx, req.Id, req.Params, authorizationHeader))
}

type DeleteGroupsIdWikisSlugRequest struct {
	Id   int32  `json:"id" jsonschema:"description=null"`
	Slug string `json:"slug" jsonschema:"description=The slug of a wiki page"`
}

func registerDeleteGroupsIdWikisSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteGroupsIdWikisSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_grps_id_wikis_slug",
		mcp.WithDescription("Delete a wiki page"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteGroupsIdWikisSlugHandler))
}

func deleteGroupsIdWikisSlugHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteGroupsIdWikisSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4GroupsIdWikisSlug(ctx, req.Id, req.Slug, authorizationHeader))
}

type GetGroupsIdWikisSlugRequest struct {
	Id     int32                                   `json:"id" jsonschema:"description=null"`
	Slug   string                                  `json:"slug" jsonschema:"description=The slug of a wiki page"`
	Params *client.GetApiV4GroupsIdWikisSlugParams `json:"params,omitempty"`
}

func registerGetGroupsIdWikisSlug(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdWikisSlugRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_wikis_slug",
		mcp.WithDescription("Get a wiki page"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdWikisSlugHandler))
}

func getGroupsIdWikisSlugHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdWikisSlugRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdWikisSlug(ctx, req.Id, req.Slug, req.Params, authorizationHeader))
}

type GetGroupsIdIssuesRequest struct {
	Id     string                               `json:"id" jsonschema:"description=The global ID or URL-encoded path of the group."`
	Params *client.GetApiV4GroupsIdIssuesParams `json:"params,omitempty"`
}

func registerGetGroupsIdIssues(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupsIdIssuesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_grps_id_issues",
		mcp.WithDescription("List group issues"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupsIdIssuesHandler))
}

func getGroupsIdIssuesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupsIdIssuesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupsIdIssues(ctx, req.Id, req.Params, authorizationHeader))
}
