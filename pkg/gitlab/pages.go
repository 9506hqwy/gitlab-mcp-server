package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetPagesDomains(s *server.MCPServer) {
	tool := mcp.NewTool("pages_domains",
		mcp.WithDescription("Get all pages domains"),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getPagesDomainsHandler)
}

func getPagesDomainsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetPagesDomains(request)
	return toResult(c.GetApiV4PagesDomains(ctx, &params, authorizationHeader))
}

func parseGetPagesDomains(request mcp.CallToolRequest) client.GetApiV4PagesDomainsParams {
	params := client.GetApiV4PagesDomainsParams{}

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

	return params
}
