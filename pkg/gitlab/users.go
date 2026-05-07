package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetUsersIdEventsRequest struct {
	Id     string                              `json:"id" jsonschema:"description=The ID or username of the user"`
	Params *client.GetApiV4UsersIdEventsParams `json:"params,omitempty"`
}

func registerGetUsersIdEvents(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsersIdEventsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_users_id_events",
		mcp.WithDescription("This feature was introduced in GitLab 8.13."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsersIdEventsHandler))
}

func getUsersIdEventsHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsersIdEventsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsersIdEvents(ctx, req.Id, req.Params, authorizationHeader))
}

type GetUsersUserIdProjectsRequest struct {
	UserId string                                    `json:"user_id" jsonschema:"description=The ID or username of the user"`
	Params *client.GetApiV4UsersUserIdProjectsParams `json:"params,omitempty"`
}

func registerGetUsersUserIdProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsersUserIdProjectsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_users_user_id_pjs",
		mcp.WithDescription("Get a user projects"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsersUserIdProjectsHandler))
}

func getUsersUserIdProjectsHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsersUserIdProjectsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsersUserIdProjects(ctx, req.UserId, req.Params, authorizationHeader))
}

type GetUsersUserIdContributedProjectsRequest struct {
	UserId string                                               `json:"user_id" jsonschema:"description=The ID or username of the user"`
	Params *client.GetApiV4UsersUserIdContributedProjectsParams `json:"params,omitempty"`
}

func registerGetUsersUserIdContributedProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsersUserIdContributedProjectsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_users_user_id_contributed_pjs",
		mcp.WithDescription("Get projects that a user has contributed to"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsersUserIdContributedProjectsHandler))
}

func getUsersUserIdContributedProjectsHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsersUserIdContributedProjectsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsersUserIdContributedProjects(ctx, req.UserId, req.Params, authorizationHeader))
}

type GetUsersUserIdStarredProjectsRequest struct {
	UserId string                                           `json:"user_id" jsonschema:"description=The ID or username of the user"`
	Params *client.GetApiV4UsersUserIdStarredProjectsParams `json:"params,omitempty"`
}

func registerGetUsersUserIdStarredProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetUsersUserIdStarredProjectsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_users_user_id_starred_pjs",
		mcp.WithDescription("Get projects starred by a user"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getUsersUserIdStarredProjectsHandler))
}

func getUsersUserIdStarredProjectsHandler(ctx context.Context, request mcp.CallToolRequest, req GetUsersUserIdStarredProjectsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4UsersUserIdStarredProjects(ctx, req.UserId, req.Params, authorizationHeader))
}
