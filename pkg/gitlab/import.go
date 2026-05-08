package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostImportBitbucketRequest struct {
	Body client.PostApiV4ImportBitbucketJSONRequestBody `json:"body"`
}

func registerPostImportBitbucket(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostImportBitbucketRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_import_bitbucket",
		mcp.WithDescription("This feature was introduced in GitLab 17.0."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postImportBitbucketHandler))
}

func postImportBitbucketHandler(ctx context.Context, request mcp.CallToolRequest, req PostImportBitbucketRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ImportBitbucket(ctx, req.Body, authorizationHeader))
}

type PostImportBitbucketServerRequest struct {
	Body client.PostApiV4ImportBitbucketServerJSONRequestBody `json:"body"`
}

func registerPostImportBitbucketServer(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostImportBitbucketServerRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_import_bitbucket_server",
		mcp.WithDescription("This feature was introduced in GitLab 13.2."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postImportBitbucketServerHandler))
}

func postImportBitbucketServerHandler(ctx context.Context, request mcp.CallToolRequest, req PostImportBitbucketServerRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ImportBitbucketServer(ctx, req.Body, authorizationHeader))
}

type PostImportGithubRequest struct {
	Body client.PostApiV4ImportGithubJSONRequestBody `json:"body"`
}

func registerPostImportGithub(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostImportGithubRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_import_github",
		mcp.WithDescription("This feature was introduced in GitLab 11.3.4."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postImportGithubHandler))
}

func postImportGithubHandler(ctx context.Context, request mcp.CallToolRequest, req PostImportGithubRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ImportGithub(ctx, req.Body, authorizationHeader))
}

type PostImportGithubCancelRequest struct {
	Body client.PostApiV4ImportGithubCancelJSONRequestBody `json:"body"`
}

func registerPostImportGithubCancel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostImportGithubCancelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_import_github_cancel",
		mcp.WithDescription("This feature was introduced in GitLab 15.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postImportGithubCancelHandler))
}

func postImportGithubCancelHandler(ctx context.Context, request mcp.CallToolRequest, req PostImportGithubCancelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ImportGithubCancel(ctx, req.Body, authorizationHeader))
}

type PostImportGithubGistsRequest struct {
	Body client.PostApiV4ImportGithubGistsJSONRequestBody `json:"body"`
}

func registerPostImportGithubGists(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostImportGithubGistsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_import_github_gists",
		mcp.WithDescription("This feature was introduced in GitLab 15.8"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postImportGithubGistsHandler))
}

func postImportGithubGistsHandler(ctx context.Context, request mcp.CallToolRequest, req PostImportGithubGistsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ImportGithubGists(ctx, req.Body, authorizationHeader))
}
