package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetGeoProxy(s *server.MCPServer) {
	tool := mcp.NewTool("geo_proxy",
		mcp.WithDescription("Returns a Geo proxy response"),
	)

	s.AddTool(tool, getGeoProxyHandler)
}

func getGeoProxyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GeoProxy(ctx, authorizationHeader))
}

func registerGetGeoRetrieveReplicableNameReplicableId(s *server.MCPServer) {
	tool := mcp.NewTool("geo_retrieve_replicable_name_replicable_id",
		mcp.WithDescription("Returns a replicable file from store (via CDN or sendfile)"),
		mcp.WithString("replicable_name",
			mcp.Description("The replicable name of a replicator instance (example: package_file)"),
			mcp.Required(),
		),
		mcp.WithNumber("replicable_id",
			mcp.Description("The replicable ID of a replicable instance"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGeoRetrieveReplicableNameReplicableIdHandler)
}

func getGeoRetrieveReplicableNameReplicableIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	replicable_name := request.GetString("replicable_name", "")
	replicable_id := int32(request.GetInt("replicable_id", math.MinInt))

	return toResult(c.GetApiV4GeoRetrieveReplicableNameReplicableId(ctx, replicable_name, replicable_id, authorizationHeader))
}

func registerGetGeoRepositoriesGlRepositoryPipelineRefs(s *server.MCPServer) {
	tool := mcp.NewTool("geo_repositories_gl_repo_pipeline_refs",
		mcp.WithDescription("Returns the list of pipeline refs for the project"),
		mcp.WithString("gl_repository",
			mcp.Description("The repository to check"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getGeoRepositoriesGlRepositoryPipelineRefsHandler)
}

func getGeoRepositoriesGlRepositoryPipelineRefsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	gl_repository := request.GetString("gl_repository", "")

	return toResult(c.GetApiV4GeoRepositoriesGlRepositoryPipelineRefs(ctx, gl_repository, authorizationHeader))
}
