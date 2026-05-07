package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type DeleteHooksHookIdUrlVariablesKeyRequest struct {
	HookId int32  `json:"hook_id" jsonschema:"description=The ID of the hook"`
	Key    string `json:"key" jsonschema:"description=The key of the variable"`
}

func registerDeleteHooksHookIdUrlVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteHooksHookIdUrlVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_hooks_hook_id_url_variables_key",
		mcp.WithDescription("Un-Set a url variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteHooksHookIdUrlVariablesKeyHandler))
}

func deleteHooksHookIdUrlVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteHooksHookIdUrlVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4HooksHookIdUrlVariablesKey(ctx, req.HookId, req.Key, authorizationHeader))
}

type DeleteHooksHookIdCustomHeadersKeyRequest struct {
	HookId int32  `json:"hook_id" jsonschema:"description=The ID of the hook"`
	Key    string `json:"key" jsonschema:"description=The key of the custom header"`
}

func registerDeleteHooksHookIdCustomHeadersKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteHooksHookIdCustomHeadersKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_hooks_hook_id_custom_headers_key",
		mcp.WithDescription("Un-Set a custom header"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteHooksHookIdCustomHeadersKeyHandler))
}

func deleteHooksHookIdCustomHeadersKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteHooksHookIdCustomHeadersKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4HooksHookIdCustomHeadersKey(ctx, req.HookId, req.Key, authorizationHeader))
}

type GetHooksRequest struct {
	Params *client.GetApiV4HooksParams `json:"params,omitempty"`
}

func registerGetHooks(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetHooksRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_hooks",
		mcp.WithDescription("Get a list of all system hooks"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getHooksHandler))
}

func getHooksHandler(ctx context.Context, request mcp.CallToolRequest, req GetHooksRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Hooks(ctx, req.Params, authorizationHeader))
}

type DeleteHooksHookIdRequest struct {
	HookId int32 `json:"hook_id" jsonschema:"description=The ID of the system hook"`
}

func registerDeleteHooksHookId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteHooksHookIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_hooks_hook_id",
		mcp.WithDescription("Deletes a system hook"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteHooksHookIdHandler))
}

func deleteHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteHooksHookIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4HooksHookId(ctx, req.HookId, authorizationHeader))
}

type PostHooksHookIdRequest struct {
	HookId int32 `json:"hook_id" jsonschema:"description=The ID of the hook"`
}

func registerPostHooksHookId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostHooksHookIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_hooks_hook_id",
		mcp.WithDescription("null"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postHooksHookIdHandler))
}

func postHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest, req PostHooksHookIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4HooksHookId(ctx, req.HookId, authorizationHeader))
}

type GetHooksHookIdRequest struct {
	HookId int32 `json:"hook_id" jsonschema:"description=The ID of the system hook"`
}

func registerGetHooksHookId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetHooksHookIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_hooks_hook_id",
		mcp.WithDescription("Get a system hook by its ID. Introduced in GitLab 14.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getHooksHookIdHandler))
}

func getHooksHookIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetHooksHookIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4HooksHookId(ctx, req.HookId, authorizationHeader))
}
