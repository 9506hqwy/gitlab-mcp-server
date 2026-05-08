package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type DeletePersonalAccessTokensSelfRequest struct {
}

func registerDeletePersonalAccessTokensSelf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeletePersonalAccessTokensSelfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_personal_access_tokens_self",
		mcp.WithDescription("Revoke a personal access token by passing it to the API in a header"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deletePersonalAccessTokensSelfHandler))
}

func deletePersonalAccessTokensSelfHandler(ctx context.Context, request mcp.CallToolRequest, req DeletePersonalAccessTokensSelfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4PersonalAccessTokensSelf(ctx, authorizationHeader))
}

type GetPersonalAccessTokensSelfRequest struct {
}

func registerGetPersonalAccessTokensSelf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPersonalAccessTokensSelfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_personal_access_tokens_self",
		mcp.WithDescription("Get the details of a personal access token by passing it to the API in a header"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPersonalAccessTokensSelfHandler))
}

func getPersonalAccessTokensSelfHandler(ctx context.Context, request mcp.CallToolRequest, req GetPersonalAccessTokensSelfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PersonalAccessTokensSelf(ctx, authorizationHeader))
}

type GetPersonalAccessTokensSelfAssociationsRequest struct {
	Params *client.GetApiV4PersonalAccessTokensSelfAssociationsParams `json:"params,omitempty"`
}

func registerGetPersonalAccessTokensSelfAssociations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPersonalAccessTokensSelfAssociationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_personal_access_tokens_self_associations",
		mcp.WithDescription("Get groups and projects this personal access token can access by passing it to the API in a header"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPersonalAccessTokensSelfAssociationsHandler))
}

func getPersonalAccessTokensSelfAssociationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetPersonalAccessTokensSelfAssociationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PersonalAccessTokensSelfAssociations(ctx, req.Params, authorizationHeader))
}

type PostPersonalAccessTokensSelfRotateRequest struct {
	Body client.PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody `json:"body"`
}

func registerPostPersonalAccessTokensSelfRotate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostPersonalAccessTokensSelfRotateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_personal_access_tokens_self_rotate",
		mcp.WithDescription("Rotates a personal access token by passing it to the API in a header"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postPersonalAccessTokensSelfRotateHandler))
}

func postPersonalAccessTokensSelfRotateHandler(ctx context.Context, request mcp.CallToolRequest, req PostPersonalAccessTokensSelfRotateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4PersonalAccessTokensSelfRotate(ctx, req.Body, authorizationHeader))
}

type GetPersonalAccessTokensRequest struct {
	Params *client.GetApiV4PersonalAccessTokensParams `json:"params,omitempty"`
}

func registerGetPersonalAccessTokens(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPersonalAccessTokensRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_personal_access_tokens",
		mcp.WithDescription("Get all personal access tokens the authenticated user has access to."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPersonalAccessTokensHandler))
}

func getPersonalAccessTokensHandler(ctx context.Context, request mcp.CallToolRequest, req GetPersonalAccessTokensRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PersonalAccessTokens(ctx, req.Params, authorizationHeader))
}

type DeletePersonalAccessTokensIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerDeletePersonalAccessTokensId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeletePersonalAccessTokensIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_personal_access_tokens_id",
		mcp.WithDescription("Revoke a personal access token by using the ID of the personal access token."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deletePersonalAccessTokensIdHandler))
}

func deletePersonalAccessTokensIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeletePersonalAccessTokensIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4PersonalAccessTokensId(ctx, req.Id, authorizationHeader))
}

type GetPersonalAccessTokensIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetPersonalAccessTokensId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPersonalAccessTokensIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_personal_access_tokens_id",
		mcp.WithDescription("Get a personal access token by using the ID of the personal access token."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPersonalAccessTokensIdHandler))
}

func getPersonalAccessTokensIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetPersonalAccessTokensIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PersonalAccessTokensId(ctx, req.Id, authorizationHeader))
}

type PostPersonalAccessTokensIdRotateRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`

	Body client.PostApiV4PersonalAccessTokensIdRotateJSONRequestBody `json:"body"`
}

func registerPostPersonalAccessTokensIdRotate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostPersonalAccessTokensIdRotateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_personal_access_tokens_id_rotate",
		mcp.WithDescription("Rotates a personal access token."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postPersonalAccessTokensIdRotateHandler))
}

func postPersonalAccessTokensIdRotateHandler(ctx context.Context, request mcp.CallToolRequest, req PostPersonalAccessTokensIdRotateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4PersonalAccessTokensIdRotate(ctx, req.Id, req.Body, authorizationHeader))
}
