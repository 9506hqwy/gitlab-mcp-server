package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
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

type PostGeoStatusRequest struct {
	Body client.PostApiV4GeoStatusJSONRequestBody `json:"body"`
}

func registerPostGeoStatus(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGeoStatusRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_geo_status",
		mcp.WithDescription("Posts the current node status to the primary site"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGeoStatusHandler))
}

func postGeoStatusHandler(ctx context.Context, request mcp.CallToolRequest, req PostGeoStatusRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GeoStatus(ctx, req.Body, authorizationHeader))
}

type PostGeoProxyGitSshInfoRefsUploadPackRequest struct {
	Body client.PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody `json:"body"`
}

func registerPostGeoProxyGitSshInfoRefsUploadPack(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGeoProxyGitSshInfoRefsUploadPackRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_geo_proxy_git_ssh_info_refs_upload_pack",
		mcp.WithDescription("Responsible for making HTTP GET /repo.git/info/refs?service=git-upload-pack request from secondary gitlab-shell to primary"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGeoProxyGitSshInfoRefsUploadPackHandler))
}

func postGeoProxyGitSshInfoRefsUploadPackHandler(ctx context.Context, request mcp.CallToolRequest, req PostGeoProxyGitSshInfoRefsUploadPackRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GeoProxyGitSshInfoRefsUploadPack(ctx, req.Body, authorizationHeader))
}

type PostGeoProxyGitSshUploadPackRequest struct {
	Body client.PostApiV4GeoProxyGitSshUploadPackJSONRequestBody `json:"body"`
}

func registerPostGeoProxyGitSshUploadPack(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGeoProxyGitSshUploadPackRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_geo_proxy_git_ssh_upload_pack",
		mcp.WithDescription("Responsible for making HTTP POST /repo.git/git-upload-pack request from secondary gitlab-shell to primary"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGeoProxyGitSshUploadPackHandler))
}

func postGeoProxyGitSshUploadPackHandler(ctx context.Context, request mcp.CallToolRequest, req PostGeoProxyGitSshUploadPackRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GeoProxyGitSshUploadPack(ctx, req.Body, authorizationHeader))
}

type PostGeoProxyGitSshInfoRefsReceivePackRequest struct {
	Body client.PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody `json:"body"`
}

func registerPostGeoProxyGitSshInfoRefsReceivePack(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGeoProxyGitSshInfoRefsReceivePackRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_geo_proxy_git_ssh_info_refs_receive_pack",
		mcp.WithDescription("Responsible for making HTTP GET /repo.git/info/refs?service=git-receive-pack request from secondary gitlab-shell to primary"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGeoProxyGitSshInfoRefsReceivePackHandler))
}

func postGeoProxyGitSshInfoRefsReceivePackHandler(ctx context.Context, request mcp.CallToolRequest, req PostGeoProxyGitSshInfoRefsReceivePackRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GeoProxyGitSshInfoRefsReceivePack(ctx, req.Body, authorizationHeader))
}

type PostGeoProxyGitSshReceivePackRequest struct {
	Body client.PostApiV4GeoProxyGitSshReceivePackJSONRequestBody `json:"body"`
}

func registerPostGeoProxyGitSshReceivePack(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostGeoProxyGitSshReceivePackRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_geo_proxy_git_ssh_receive_pack",
		mcp.WithDescription("Responsible for making HTTP POST /repo.git/info/refs?service=git-receive-pack request from secondary gitlab-shell to primary"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postGeoProxyGitSshReceivePackHandler))
}

func postGeoProxyGitSshReceivePackHandler(ctx context.Context, request mcp.CallToolRequest, req PostGeoProxyGitSshReceivePackRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4GeoProxyGitSshReceivePack(ctx, req.Body, authorizationHeader))
}
