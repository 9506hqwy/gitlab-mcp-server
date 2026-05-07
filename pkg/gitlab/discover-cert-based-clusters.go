package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetDiscoverCertBasedClustersRequest struct {
	Params *client.GetApiV4DiscoverCertBasedClustersParams `json:"params,omitempty"`
}

func registerGetDiscoverCertBasedClusters(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetDiscoverCertBasedClustersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_discover_cert_based_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 17.9. It will be removed in 18.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getDiscoverCertBasedClustersHandler))
}

func getDiscoverCertBasedClustersHandler(ctx context.Context, request mcp.CallToolRequest, req GetDiscoverCertBasedClustersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4DiscoverCertBasedClusters(ctx, req.Params, authorizationHeader))
}
