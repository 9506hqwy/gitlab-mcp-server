package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetGroupIdPackagesComposerPackages(s *server.MCPServer) {
	tool := mcp.NewTool("group_id___pkgs_composer_packages",
		mcp.WithDescription("This feature was introduced in GitLab 13.1"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of a group"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupIdPackagesComposerPackagesHandler)
}

func getGroupIdPackagesComposerPackagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4GroupIdPackagesComposerPackages(ctx, id, authorizationHeader))
}

func registerGetGroupIdPackagesComposerPSha(s *server.MCPServer) {
	tool := mcp.NewTool("group_id___pkgs_composer_p_sha",
		mcp.WithDescription("This feature was introduced in GitLab 13.1"),
		mcp.WithString("id",
			mcp.Description("The ID or URL-encoded path of a group"),
			mcp.Required(),
		),
		mcp.WithString("sha",
			mcp.Description("Shasum of current json (example: 673594f85a55fe3c0eb45df7bd2fa9d95a1601ab)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGroupIdPackagesComposerPShaHandler)
}

func getGroupIdPackagesComposerPShaHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	sha := request.GetString("sha", "")

	return toResult(c.GetApiV4GroupIdPackagesComposerPSha(ctx, id, sha, authorizationHeader))
}
