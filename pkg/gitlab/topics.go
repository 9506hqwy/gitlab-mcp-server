package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetTopics(s *server.MCPServer) {
	tool := mcp.NewTool("get_topics",
		mcp.WithDescription("This feature was introduced in GitLab 14.5."),
		mcp.WithString("search",
			mcp.Description("Return list of topics matching the search criteria (example: search)"),
		),
		mcp.WithBoolean("without_projects",
			mcp.Description("Return list of topics without assigned projects"),
		),
		mcp.WithNumber("organization_id",
			mcp.Description("The organization id for the topics"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getTopicsHandler)
}

func getTopicsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetTopics(request)
	return toResult(c.GetApiV4Topics(ctx, &params, authorizationHeader))
}

func parseGetTopics(request mcp.CallToolRequest) client.GetApiV4TopicsParams {
	params := client.GetApiV4TopicsParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	without_projects := request.GetBool("without_projects", false)
	params.WithoutProjects = &without_projects

	organization_id := request.GetInt("organization_id", math.MinInt)
	if organization_id != math.MinInt {
		organization_id := int32(organization_id)
		params.OrganizationId = &organization_id
	}

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

func registerDeleteTopicsId(s *server.MCPServer) {
	tool := mcp.NewTool("delete_topics_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.9."),
		mcp.WithNumber("id",
			mcp.Description("ID of project topic"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, deleteTopicsIdHandler)
}

func deleteTopicsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.DeleteApiV4TopicsId(ctx, id, authorizationHeader))
}

func registerGetTopicsId(s *server.MCPServer) {
	tool := mcp.NewTool("get_topics_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.5."),
		mcp.WithNumber("id",
			mcp.Description("ID of project topic"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getTopicsIdHandler)
}

func getTopicsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4TopicsId(ctx, id, authorizationHeader))
}
