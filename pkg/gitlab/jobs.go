package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PostJobsRequestRequest struct {
	Body client.PostApiV4JobsRequestJSONRequestBody `json:"body"`
}

func registerPostJobsRequest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostJobsRequestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_jobs_request",
		mcp.WithDescription("Request a job"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postJobsRequestHandler))
}

func postJobsRequestHandler(ctx context.Context, request mcp.CallToolRequest, req PostJobsRequestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4JobsRequest(ctx, req.Body, authorizationHeader))
}

type PutJobsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=Job's ID"`

	Body client.PutApiV4JobsIdJSONRequestBody `json:"body"`
}

func registerPutJobsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutJobsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_jobs_id",
		mcp.WithDescription("Update a job"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putJobsIdHandler))
}

func putJobsIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutJobsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4JobsId(ctx, req.Id, req.Body, authorizationHeader))
}

type PostJobsIdArtifactsAuthorizeRequest struct {
	Id int32 `json:"id" jsonschema:"description=Job's ID"`

	Body client.PostApiV4JobsIdArtifactsAuthorizeJSONRequestBody `json:"body"`
}

func registerPostJobsIdArtifactsAuthorize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostJobsIdArtifactsAuthorizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_jobs_id_artifacts_authorize",
		mcp.WithDescription("Authorize uploading job artifact"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postJobsIdArtifactsAuthorizeHandler))
}

func postJobsIdArtifactsAuthorizeHandler(ctx context.Context, request mcp.CallToolRequest, req PostJobsIdArtifactsAuthorizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4JobsIdArtifactsAuthorize(ctx, req.Id, req.Body, authorizationHeader))
}

type PostJobsIdArtifactsRequest struct {
	Id int32 `json:"id" jsonschema:"description=Job's ID"`

	Body client.PostApiV4JobsIdArtifactsJSONRequestBody `json:"body"`
}

func registerPostJobsIdArtifacts(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostJobsIdArtifactsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_jobs_id_artifacts",
		mcp.WithDescription("Upload a job artifact"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postJobsIdArtifactsHandler))
}

func postJobsIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest, req PostJobsIdArtifactsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4JobsIdArtifacts(ctx, req.Id, req.Body, authorizationHeader))
}

type GetJobsIdArtifactsRequest struct {
	Id     int32                                 `json:"id" jsonschema:"description=Job's ID"`
	Params *client.GetApiV4JobsIdArtifactsParams `json:"params,omitempty"`
}

func registerGetJobsIdArtifacts(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetJobsIdArtifactsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_jobs_id_artifacts",
		mcp.WithDescription("Download the artifacts file for job"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getJobsIdArtifactsHandler))
}

func getJobsIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest, req GetJobsIdArtifactsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4JobsIdArtifacts(ctx, req.Id, req.Params, authorizationHeader))
}
