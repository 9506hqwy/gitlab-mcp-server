package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetDiscoverCertBasedClusters(s *server.MCPServer) {
	tool := mcp.NewTool("discover_cert_based_clusters_",
		mcp.WithDescription("This feature was introduced in GitLab 17.9. It will be removed in 18.0."),
		mcp.WithNumber("group_id",
			mcp.Description("The group ID to find all certificate-based clusters in the hierarchy"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getDiscoverCertBasedClustersHandler)
}

func getDiscoverCertBasedClustersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetDiscoverCertBasedClusters(request)
	return toResult(c.GetApiV4DiscoverCertBasedClusters(ctx, &params, authorizationHeader))
}

func parseGetDiscoverCertBasedClusters(request mcp.CallToolRequest) client.GetApiV4DiscoverCertBasedClustersParams {
	params := client.GetApiV4DiscoverCertBasedClustersParams{}

	group_id := request.GetInt("group_id", math.MinInt)
	if group_id != math.MinInt {
		group_id := int32(group_id)
		params.GroupId = group_id
	}

	return params
}
