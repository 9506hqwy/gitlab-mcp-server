package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type GetGeoProxyRequest struct {
}

func registerGetGeoProxy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGeoProxyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_geo_proxy",
		mcp.WithDescription("Returns a Geo proxy response"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGeoProxyHandler))
}

func getGeoProxyHandler(ctx context.Context, request mcp.CallToolRequest, req GetGeoProxyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GeoProxy(ctx, authorizationHeader))
}

type GetGeoRetrieveReplicableNameReplicableIdRequest struct {
	ReplicableName string `json:"replicable_name" jsonschema:"description=The replicable name of a replicator instance"`
	ReplicableId   int32  `json:"replicable_id" jsonschema:"description=The replicable ID of a replicable instance"`
}

func registerGetGeoRetrieveReplicableNameReplicableId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGeoRetrieveReplicableNameReplicableIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_geo_retrieve_replicable_name_replicable_id",
		mcp.WithDescription("Returns a replicable file from store (via CDN or sendfile)"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGeoRetrieveReplicableNameReplicableIdHandler))
}

func getGeoRetrieveReplicableNameReplicableIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetGeoRetrieveReplicableNameReplicableIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GeoRetrieveReplicableNameReplicableId(ctx, req.ReplicableName, req.ReplicableId, authorizationHeader))
}

type GetGeoRepositoriesGlRepositoryPipelineRefsRequest struct {
	GlRepository string `json:"gl_repository" jsonschema:"description=The repository to check"`
}

func registerGetGeoRepositoriesGlRepositoryPipelineRefs(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGeoRepositoriesGlRepositoryPipelineRefsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_geo_repositories_gl_repo_pipeline_refs",
		mcp.WithDescription("Returns the list of pipeline refs for the project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGeoRepositoriesGlRepositoryPipelineRefsHandler))
}

func getGeoRepositoriesGlRepositoryPipelineRefsHandler(ctx context.Context, request mcp.CallToolRequest, req GetGeoRepositoriesGlRepositoryPipelineRefsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GeoRepositoriesGlRepositoryPipelineRefs(ctx, req.GlRepository, authorizationHeader))
}

type PostGeoNodeProxyIdGraphqlRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of the Geo node"`
}

func registerPostGeoNodeProxyIdGraphql(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGeoNodeProxyIdGraphqlRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_geo_node_proxy_id_graphql",
		mcp.WithDescription("Query the GraphQL endpoint of an existing Geo node"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGeoNodeProxyIdGraphqlHandler))
}

func postGeoNodeProxyIdGraphqlHandler(ctx context.Context, request mcp.CallToolRequest, req PostGeoNodeProxyIdGraphqlRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GeoNodeProxyIdGraphql(ctx, req.Id, authorizationHeader))
}
