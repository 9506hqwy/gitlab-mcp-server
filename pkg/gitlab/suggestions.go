package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PutSuggestionsIdApplyRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of the suggestion"`

	Body client.PutApiV4SuggestionsIdApplyJSONRequestBody `json:"body"`
}

func registerPutSuggestionsIdApply(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutSuggestionsIdApplyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_suggestions_id_apply",
		mcp.WithDescription("Apply suggestion patch in the Merge Request it was created"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putSuggestionsIdApplyHandler))
}

func putSuggestionsIdApplyHandler(ctx context.Context, request mcp.CallToolRequest, req PutSuggestionsIdApplyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4SuggestionsIdApply(ctx, req.Id, req.Body, authorizationHeader))
}

type PutSuggestionsBatchApplyRequest struct {
	Body client.PutApiV4SuggestionsBatchApplyJSONRequestBody `json:"body"`
}

func registerPutSuggestionsBatchApply(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutSuggestionsBatchApplyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_suggestions_batch_apply",
		mcp.WithDescription("Apply multiple suggestion patches in the Merge Request where they were created"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putSuggestionsBatchApplyHandler))
}

func putSuggestionsBatchApplyHandler(ctx context.Context, request mcp.CallToolRequest, req PutSuggestionsBatchApplyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4SuggestionsBatchApply(ctx, req.Body, authorizationHeader))
}
