package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

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
