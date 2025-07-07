package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetDeployTokens(s *server.MCPServer) {
	tool := mcp.NewTool("get_deploy_tokens",
		mcp.WithDescription("Get a list of all deploy tokens across the GitLab instance. This endpoint requires administrator access. This feature was introduced in GitLab 12.9."),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithBoolean("active",
			mcp.Description("Limit by active status"),
		),
	)

	s.AddTool(tool, getDeployTokensHandler)
}

func getDeployTokensHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetDeployTokens(request)
	return toResult(c.GetApiV4DeployTokens(ctx, &params, authorizationHeader))
}

func parseGetDeployTokens(request mcp.CallToolRequest) client.GetApiV4DeployTokensParams {
	params := client.GetApiV4DeployTokensParams{}

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

	active := request.GetBool("active", false)
	params.Active = &active

	return params
}
