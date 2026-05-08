package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostBroadcastMessagesRequest struct {
	Body client.PostApiV4BroadcastMessagesJSONRequestBody `json:"body"`
}

func registerPostBroadcastMessages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostBroadcastMessagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_broadcast_messages",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postBroadcastMessagesHandler))
}

func postBroadcastMessagesHandler(ctx context.Context, request mcp.CallToolRequest, req PostBroadcastMessagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4BroadcastMessages(ctx, req.Body, authorizationHeader))
}

type GetBroadcastMessagesRequest struct {
	Params *client.GetApiV4BroadcastMessagesParams `json:"params,omitempty"`
}

func registerGetBroadcastMessages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBroadcastMessagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_broadcast_messages",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBroadcastMessagesHandler))
}

func getBroadcastMessagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetBroadcastMessagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BroadcastMessages(ctx, req.Params, authorizationHeader))
}

type DeleteBroadcastMessagesIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=Broadcast message ID"`
}

func registerDeleteBroadcastMessagesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteBroadcastMessagesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_broadcast_messages_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteBroadcastMessagesIdHandler))
}

func deleteBroadcastMessagesIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteBroadcastMessagesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4BroadcastMessagesId(ctx, req.Id, authorizationHeader))
}

type PutBroadcastMessagesIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=Broadcast message ID"`

	Body client.PutApiV4BroadcastMessagesIdJSONRequestBody `json:"body"`
}

func registerPutBroadcastMessagesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutBroadcastMessagesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_broadcast_messages_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putBroadcastMessagesIdHandler))
}

func putBroadcastMessagesIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutBroadcastMessagesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4BroadcastMessagesId(ctx, req.Id, req.Body, authorizationHeader))
}

type GetBroadcastMessagesIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=Broadcast message ID"`
}

func registerGetBroadcastMessagesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBroadcastMessagesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_broadcast_messages_id",
		mcp.WithDescription("This feature was introduced in GitLab 8.12."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBroadcastMessagesIdHandler))
}

func getBroadcastMessagesIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetBroadcastMessagesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BroadcastMessagesId(ctx, req.Id, authorizationHeader))
}
