package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type DeleteRunnersRequest struct {
	Params *client.DeleteApiV4RunnersParams `json:"params,omitempty"`
}

func registerDeleteRunners(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteRunnersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_runners",
		mcp.WithDescription("Delete a registered runner"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteRunnersHandler))
}

func deleteRunnersHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteRunnersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4Runners(ctx, req.Params, authorizationHeader))
}

type PostRunnersRequest struct {
	Body client.PostApiV4RunnersJSONRequestBody `json:"body"`
}

func registerPostRunners(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostRunnersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_runners",
		mcp.WithDescription("Register a new runner for the instance"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postRunnersHandler))
}

func postRunnersHandler(ctx context.Context, request mcp.CallToolRequest, req PostRunnersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4Runners(ctx, req.Body, authorizationHeader))
}

type GetRunnersRequest struct {
	Params *client.GetApiV4RunnersParams `json:"params,omitempty"`
}

func registerGetRunners(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetRunnersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_runners",
		mcp.WithDescription("Get runners available for user"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getRunnersHandler))
}

func getRunnersHandler(ctx context.Context, request mcp.CallToolRequest, req GetRunnersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Runners(ctx, req.Params, authorizationHeader))
}

type DeleteRunnersManagersRequest struct {
	Params *client.DeleteApiV4RunnersManagersParams `json:"params,omitempty"`
}

func registerDeleteRunnersManagers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteRunnersManagersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_runners_managers",
		mcp.WithDescription("Delete a registered runner manager"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteRunnersManagersHandler))
}

func deleteRunnersManagersHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteRunnersManagersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4RunnersManagers(ctx, req.Params, authorizationHeader))
}

type PostRunnersVerifyRequest struct {
	Body client.PostApiV4RunnersVerifyJSONRequestBody `json:"body"`
}

func registerPostRunnersVerify(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostRunnersVerifyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_runners_verify",
		mcp.WithDescription("Validate authentication credentials"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postRunnersVerifyHandler))
}

func postRunnersVerifyHandler(ctx context.Context, request mcp.CallToolRequest, req PostRunnersVerifyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4RunnersVerify(ctx, req.Body, authorizationHeader))
}

type PostRunnersResetAuthenticationTokenRequest struct {
	Body client.PostApiV4RunnersResetAuthenticationTokenJSONRequestBody `json:"body"`
}

func registerPostRunnersResetAuthenticationToken(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostRunnersResetAuthenticationTokenRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_runners_reset_authentication_token",
		mcp.WithDescription("Reset runner authentication token with current token"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postRunnersResetAuthenticationTokenHandler))
}

func postRunnersResetAuthenticationTokenHandler(ctx context.Context, request mcp.CallToolRequest, req PostRunnersResetAuthenticationTokenRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4RunnersResetAuthenticationToken(ctx, req.Body, authorizationHeader))
}

type GetRunnersAllRequest struct {
	Params *client.GetApiV4RunnersAllParams `json:"params,omitempty"`
}

func registerGetRunnersAll(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetRunnersAllRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_runners_all",
		mcp.WithDescription("Get a list of all runners in the GitLab instance (shared and project). Access is restricted to users with either administrator access or auditor access."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getRunnersAllHandler))
}

func getRunnersAllHandler(ctx context.Context, request mcp.CallToolRequest, req GetRunnersAllRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4RunnersAll(ctx, req.Params, authorizationHeader))
}

type DeleteRunnersIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a runner"`
}

func registerDeleteRunnersId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteRunnersIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_runners_id",
		mcp.WithDescription("Remove a runner"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteRunnersIdHandler))
}

func deleteRunnersIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteRunnersIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4RunnersId(ctx, req.Id, authorizationHeader))
}

type PutRunnersIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a runner"`

	Body client.PutApiV4RunnersIdJSONRequestBody `json:"body"`
}

func registerPutRunnersId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutRunnersIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_runners_id",
		mcp.WithDescription("Update runner's details"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putRunnersIdHandler))
}

func putRunnersIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutRunnersIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4RunnersId(ctx, req.Id, req.Body, authorizationHeader))
}

type GetRunnersIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a runner"`
}

func registerGetRunnersId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetRunnersIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_runners_id",
		mcp.WithDescription("At least the Maintainer role is required to get runner details at the project and group level. Instance-level runner details via this endpoint are available to all signed in users."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getRunnersIdHandler))
}

func getRunnersIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetRunnersIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4RunnersId(ctx, req.Id, authorizationHeader))
}

type GetRunnersIdManagersRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a runner"`
}

func registerGetRunnersIdManagers(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetRunnersIdManagersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_runners_id_managers",
		mcp.WithDescription("Get a list of all runner's managers"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getRunnersIdManagersHandler))
}

func getRunnersIdManagersHandler(ctx context.Context, request mcp.CallToolRequest, req GetRunnersIdManagersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4RunnersIdManagers(ctx, req.Id, authorizationHeader))
}

type GetRunnersIdJobsRequest struct {
	Id     int32                               `json:"id" jsonschema:"description=The ID of a runner"`
	Params *client.GetApiV4RunnersIdJobsParams `json:"params,omitempty"`
}

func registerGetRunnersIdJobs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetRunnersIdJobsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_runners_id_jobs",
		mcp.WithDescription("List jobs that are being processed or were processed by the specified runner. The list of jobs is limited to projects where the user has at least the Reporter role."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getRunnersIdJobsHandler))
}

func getRunnersIdJobsHandler(ctx context.Context, request mcp.CallToolRequest, req GetRunnersIdJobsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4RunnersIdJobs(ctx, req.Id, req.Params, authorizationHeader))
}
