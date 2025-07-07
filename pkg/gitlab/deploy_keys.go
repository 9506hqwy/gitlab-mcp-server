package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetDeployKeys(s *server.MCPServer) {
	tool := mcp.NewTool("get_deploy_keys",
		mcp.WithDescription("Get a list of all deploy keys across all projects of the GitLab instance. This endpoint requires administrator access and is not available on GitLab.com."),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("public",
			mcp.Description("Only return deploy keys that are public (default: false)"),
		),
	)

	s.AddTool(tool, getDeployKeysHandler)
}

func getDeployKeysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetDeployKeys(request)
	return toResult(c.GetApiV4DeployKeys(ctx, &params, authorizationHeader))
}

func parseGetDeployKeys(request mcp.CallToolRequest) client.GetApiV4DeployKeysParams {
	params := client.GetApiV4DeployKeysParams{}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	public := request.GetBool("public", false)
	params.Public = &public

	return params
}
