package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetMergeRequestsRequest struct {
	Params *client.GetApiV4MergeRequestsParams `json:"params,omitempty"`
}

func registerGetMergeRequests(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetMergeRequestsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_mrs",
		mcp.WithDescription("Get all merge requests the authenticated user has access to. By default it returns only merge requests created by the current user. To get all merge requests, use parameter 'scope=all'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getMergeRequestsHandler))
}

func getMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest, req GetMergeRequestsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4MergeRequests(ctx, req.Params, authorizationHeader))
}
