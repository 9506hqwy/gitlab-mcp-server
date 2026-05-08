package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type PutNamespacesIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`

	Body client.PutApiV4NamespacesIdJSONRequestBody `json:"body"`
}

func registerPutNamespacesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutNamespacesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_namespaces_id",
		mcp.WithDescription("[DEPRECATED] Update a namespace"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putNamespacesIdHandler))
}

func putNamespacesIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutNamespacesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4NamespacesId(ctx, req.Id, req.Body, authorizationHeader))
}

type GetNamespacesIdRequest struct {
	Id string `json:"id" jsonschema:"description=ID or URL-encoded path of the namespace"`
}

func registerGetNamespacesId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetNamespacesIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_namespaces_id",
		mcp.WithDescription("Get a namespace by ID"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getNamespacesIdHandler))
}

func getNamespacesIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetNamespacesIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4NamespacesId(ctx, req.Id, authorizationHeader))
}

type GetNamespacesIdGitlabSubscriptionRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerGetNamespacesIdGitlabSubscription(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetNamespacesIdGitlabSubscriptionRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_namespaces_id_gitlab_subscription",
		mcp.WithDescription("Returns the subscription for the namespace"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getNamespacesIdGitlabSubscriptionHandler))
}

func getNamespacesIdGitlabSubscriptionHandler(ctx context.Context, request mcp.CallToolRequest, req GetNamespacesIdGitlabSubscriptionRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4NamespacesIdGitlabSubscription(ctx, req.Id, authorizationHeader))
}

type DeleteNamespacesIdStorageLimitExclusionRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`
}

func registerDeleteNamespacesIdStorageLimitExclusion(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteNamespacesIdStorageLimitExclusionRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_namespaces_id_storage_limit_exclusion",
		mcp.WithDescription("Removes a Namespaces::Storage::LimitExclusion"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteNamespacesIdStorageLimitExclusionHandler))
}

func deleteNamespacesIdStorageLimitExclusionHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteNamespacesIdStorageLimitExclusionRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4NamespacesIdStorageLimitExclusion(ctx, req.Id, authorizationHeader))
}

type PostNamespacesIdStorageLimitExclusionRequest struct {
	Id int32 `json:"id" jsonschema:"description=null"`

	Body client.PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody `json:"body"`
}

func registerPostNamespacesIdStorageLimitExclusion(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostNamespacesIdStorageLimitExclusionRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_namespaces_id_storage_limit_exclusion",
		mcp.WithDescription("Creates a Namespaces::Storage::LimitExclusion"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postNamespacesIdStorageLimitExclusionHandler))
}

func postNamespacesIdStorageLimitExclusionHandler(ctx context.Context, request mcp.CallToolRequest, req PostNamespacesIdStorageLimitExclusionRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4NamespacesIdStorageLimitExclusion(ctx, req.Id, req.Body, authorizationHeader))
}

type GetNamespacesStorageLimitExclusionsRequest struct {
	Params *client.GetApiV4NamespacesStorageLimitExclusionsParams `json:"params,omitempty"`
}

func registerGetNamespacesStorageLimitExclusions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetNamespacesStorageLimitExclusionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_namespaces_storage_limit_exclusions",
		mcp.WithDescription("Gets all records for namespaces that have been excluded"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getNamespacesStorageLimitExclusionsHandler))
}

func getNamespacesStorageLimitExclusionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetNamespacesStorageLimitExclusionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4NamespacesStorageLimitExclusions(ctx, req.Params, authorizationHeader))
}

type GetNamespacesRequest struct {
	Params *client.GetApiV4NamespacesParams `json:"params,omitempty"`
}

func registerGetNamespaces(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetNamespacesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_namespaces",
		mcp.WithDescription("Get a list of the namespaces of the authenticated user. If the user is an administrator, a list of all namespaces in the GitLab instance is shown."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getNamespacesHandler))
}

func getNamespacesHandler(ctx context.Context, request mcp.CallToolRequest, req GetNamespacesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Namespaces(ctx, req.Params, authorizationHeader))
}

type GetNamespacesIdExistsRequest struct {
	Id     string                                   `json:"id" jsonschema:"description=Namespace’s path"`
	Params *client.GetApiV4NamespacesIdExistsParams `json:"params,omitempty"`
}

func registerGetNamespacesIdExists(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetNamespacesIdExistsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_namespaces_id_exists",
		mcp.WithDescription("Get existence of a namespace by path. Suggests a new namespace path that does not already exist."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getNamespacesIdExistsHandler))
}

func getNamespacesIdExistsHandler(ctx context.Context, request mcp.CallToolRequest, req GetNamespacesIdExistsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4NamespacesIdExists(ctx, req.Id, req.Params, authorizationHeader))
}
