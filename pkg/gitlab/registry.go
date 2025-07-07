package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetRegistryRepositoriesId(s *server.MCPServer) {
	tool := mcp.NewTool("get_registry_repositories_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.6."),
		mcp.WithString("id",
			mcp.Description("The ID of the repository"),
			mcp.Required(),
		),
		mcp.WithBoolean("tags",
			mcp.Description("Determines if tags should be included (default: false)"),
		),
		mcp.WithBoolean("tags_count",
			mcp.Description("Determines if the tags count should be included (default: false)"),
		),
		mcp.WithBoolean("size",
			mcp.Description("Determines if the size should be included (default: false)"),
		),
	)

	s.AddTool(tool, getRegistryRepositoriesIdHandler)
}

func getRegistryRepositoriesIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetRegistryRepositoriesId(request)
	return toResult(c.GetApiV4RegistryRepositoriesId(ctx, id, &params, authorizationHeader))
}

func parseGetRegistryRepositoriesId(request mcp.CallToolRequest) client.GetApiV4RegistryRepositoriesIdParams {
	params := client.GetApiV4RegistryRepositoriesIdParams{}

	tags := request.GetBool("tags", false)
	params.Tags = &tags

	tags_count := request.GetBool("tags_count", false)
	params.TagsCount = &tags_count

	size := request.GetBool("size", false)
	params.Size = &size

	return params
}
