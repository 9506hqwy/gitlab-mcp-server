package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostTopicsRequest struct {
	Body client.PostApiV4TopicsJSONRequestBody `json:"body"`
}

func registerPostTopics(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostTopicsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_topics",
		mcp.WithDescription("This feature was introduced in GitLab 14.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postTopicsHandler))
}

func postTopicsHandler(ctx context.Context, request mcp.CallToolRequest, req PostTopicsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4Topics(ctx, req.Body, authorizationHeader))
}

type GetTopicsRequest struct {
	Params *client.GetApiV4TopicsParams `json:"params,omitempty"`
}

func registerGetTopics(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetTopicsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_topics",
		mcp.WithDescription("This feature was introduced in GitLab 14.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getTopicsHandler))
}

func getTopicsHandler(ctx context.Context, request mcp.CallToolRequest, req GetTopicsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Topics(ctx, req.Params, authorizationHeader))
}

type DeleteTopicsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=ID of project topic"`
}

func registerDeleteTopicsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteTopicsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_topics_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.9."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteTopicsIdHandler))
}

func deleteTopicsIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteTopicsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4TopicsId(ctx, req.Id, authorizationHeader))
}

type PutTopicsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=ID of project topic"`

	Body client.PutApiV4TopicsIdJSONRequestBody `json:"body"`
}

func registerPutTopicsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutTopicsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_topics_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putTopicsIdHandler))
}

func putTopicsIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutTopicsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4TopicsId(ctx, req.Id, req.Body, authorizationHeader))
}

type GetTopicsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=ID of project topic"`
}

func registerGetTopicsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetTopicsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_topics_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.5."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getTopicsIdHandler))
}

func getTopicsIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetTopicsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4TopicsId(ctx, req.Id, authorizationHeader))
}

type PostTopicsMergeRequest struct {
	Body client.PostApiV4TopicsMergeJSONRequestBody `json:"body"`
}

func registerPostTopicsMerge(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostTopicsMergeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_topics_merge",
		mcp.WithDescription("This feature was introduced in GitLab 15.4."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postTopicsMergeHandler))
}

func postTopicsMergeHandler(ctx context.Context, request mcp.CallToolRequest, req PostTopicsMergeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4TopicsMerge(ctx, req.Body, authorizationHeader))
}
