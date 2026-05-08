package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostSnippetsRequest struct {
	Body client.PostApiV4SnippetsJSONRequestBody `json:"body"`
}

func registerPostSnippets(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostSnippetsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_snippets",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postSnippetsHandler))
}

func postSnippetsHandler(ctx context.Context, request mcp.CallToolRequest, req PostSnippetsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4Snippets(ctx, req.Body, authorizationHeader))
}

type GetSnippetsRequest struct {
	Params *client.GetApiV4SnippetsParams `json:"params,omitempty"`
}

func registerGetSnippets(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsHandler))
}

func getSnippetsHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Snippets(ctx, req.Params, authorizationHeader))
}

type GetSnippetsPublicRequest struct {
	Params *client.GetApiV4SnippetsPublicParams `json:"params,omitempty"`
}

func registerGetSnippetsPublic(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsPublicRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets_public",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsPublicHandler))
}

func getSnippetsPublicHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsPublicRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4SnippetsPublic(ctx, req.Params, authorizationHeader))
}

type GetSnippetsAllRequest struct {
	Params *client.GetApiV4SnippetsAllParams `json:"params,omitempty"`
}

func registerGetSnippetsAll(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsAllRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets_all",
		mcp.WithDescription("This feature was introduced in GitLab 16.3."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsAllHandler))
}

func getSnippetsAllHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsAllRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4SnippetsAll(ctx, req.Params, authorizationHeader))
}

type DeleteSnippetsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a snippet"`
}

func registerDeleteSnippetsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteSnippetsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_snippets_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteSnippetsIdHandler))
}

func deleteSnippetsIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteSnippetsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4SnippetsId(ctx, req.Id, authorizationHeader))
}

type PutSnippetsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a snippet"`

	Body client.PutApiV4SnippetsIdJSONRequestBody `json:"body"`
}

func registerPutSnippetsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutSnippetsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_snippets_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putSnippetsIdHandler))
}

func putSnippetsIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutSnippetsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4SnippetsId(ctx, req.Id, req.Body, authorizationHeader))
}

type GetSnippetsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a snippet"`
}

func registerGetSnippetsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsIdHandler))
}

func getSnippetsIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4SnippetsId(ctx, req.Id, authorizationHeader))
}

type GetSnippetsIdRawRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a snippet"`
}

func registerGetSnippetsIdRaw(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsIdRawRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets_id_raw",
		mcp.WithDescription("This feature was introduced in GitLab 8.15."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsIdRawHandler))
}

func getSnippetsIdRawHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsIdRawRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4SnippetsIdRaw(ctx, req.Id, authorizationHeader))
}

type GetSnippetsIdFilesRefFilePathRawRequest struct {
	Id       int32  `json:"id" jsonschema:"description=null"`
	Ref      string `json:"ref" jsonschema:"description=The name of branch, tag or commit"`
	FilePath string `json:"file_path" jsonschema:"description=The url encoded path to the file, e.g. lib%2Fclass%2Erb"`
}

func registerGetSnippetsIdFilesRefFilePathRaw(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsIdFilesRefFilePathRawRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets_id_files_ref_file_path_raw",
		mcp.WithDescription("Get raw snippet file contents from the repository"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsIdFilesRefFilePathRawHandler))
}

func getSnippetsIdFilesRefFilePathRawHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsIdFilesRefFilePathRawRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4SnippetsIdFilesRefFilePathRaw(ctx, req.Id, req.Ref, req.FilePath, authorizationHeader))
}

type GetSnippetsIdUserAgentDetailRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of a snippet"`
}

func registerGetSnippetsIdUserAgentDetail(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetSnippetsIdUserAgentDetailRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_snippets_id_user_agent_detail",
		mcp.WithDescription("Get the user agent details for a snippet"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getSnippetsIdUserAgentDetailHandler))
}

func getSnippetsIdUserAgentDetailHandler(ctx context.Context, request mcp.CallToolRequest, req GetSnippetsIdUserAgentDetailRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4SnippetsIdUserAgentDetail(ctx, req.Id, authorizationHeader))
}
