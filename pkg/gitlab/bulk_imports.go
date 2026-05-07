package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetBulkImportsRequest struct {
	Params *client.GetApiV4BulkImportsParams `json:"params,omitempty"`
}

func registerGetBulkImports(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBulkImportsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_bulk_imports",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBulkImportsHandler))
}

func getBulkImportsHandler(ctx context.Context, request mcp.CallToolRequest, req GetBulkImportsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BulkImports(ctx, req.Params, authorizationHeader))
}

type GetBulkImportsEntitiesRequest struct {
	Params *client.GetApiV4BulkImportsEntitiesParams `json:"params,omitempty"`
}

func registerGetBulkImportsEntities(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBulkImportsEntitiesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_bulk_imports_entities",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBulkImportsEntitiesHandler))
}

func getBulkImportsEntitiesHandler(ctx context.Context, request mcp.CallToolRequest, req GetBulkImportsEntitiesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BulkImportsEntities(ctx, req.Params, authorizationHeader))
}

type GetBulkImportsImportIdRequest struct {
	ImportId int32 `json:"import_id" jsonschema:"description=The ID of user's GitLab Migration"`
}

func registerGetBulkImportsImportId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBulkImportsImportIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_bulk_imports_import_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBulkImportsImportIdHandler))
}

func getBulkImportsImportIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetBulkImportsImportIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BulkImportsImportId(ctx, req.ImportId, authorizationHeader))
}

type GetBulkImportsImportIdEntitiesRequest struct {
	ImportId int32                                             `json:"import_id" jsonschema:"description=The ID of user's GitLab Migration"`
	Params   *client.GetApiV4BulkImportsImportIdEntitiesParams `json:"params,omitempty"`
}

func registerGetBulkImportsImportIdEntities(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBulkImportsImportIdEntitiesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_bulk_imports_import_id_entities",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBulkImportsImportIdEntitiesHandler))
}

func getBulkImportsImportIdEntitiesHandler(ctx context.Context, request mcp.CallToolRequest, req GetBulkImportsImportIdEntitiesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BulkImportsImportIdEntities(ctx, req.ImportId, req.Params, authorizationHeader))
}

type GetBulkImportsImportIdEntitiesEntityIdRequest struct {
	ImportId int32 `json:"import_id" jsonschema:"description=The ID of user's GitLab Migration"`
	EntityId int32 `json:"entity_id" jsonschema:"description=The ID of GitLab Migration entity"`
}

func registerGetBulkImportsImportIdEntitiesEntityId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBulkImportsImportIdEntitiesEntityIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_bulk_imports_import_id_entities_entity_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBulkImportsImportIdEntitiesEntityIdHandler))
}

func getBulkImportsImportIdEntitiesEntityIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetBulkImportsImportIdEntitiesEntityIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BulkImportsImportIdEntitiesEntityId(ctx, req.ImportId, req.EntityId, authorizationHeader))
}

type GetBulkImportsImportIdEntitiesEntityIdFailuresRequest struct {
	ImportId int32 `json:"import_id" jsonschema:"description=The ID of user's GitLab Migration"`
	EntityId int32 `json:"entity_id" jsonschema:"description=The ID of GitLab Migration entity"`
}

func registerGetBulkImportsImportIdEntitiesEntityIdFailures(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetBulkImportsImportIdEntitiesEntityIdFailuresRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_bulk_imports_import_id_entities_entity_id_failures",
		mcp.WithDescription("This feature was introduced in GitLab 16.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getBulkImportsImportIdEntitiesEntityIdFailuresHandler))
}

func getBulkImportsImportIdEntitiesEntityIdFailuresHandler(ctx context.Context, request mcp.CallToolRequest, req GetBulkImportsImportIdEntitiesEntityIdFailuresRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4BulkImportsImportIdEntitiesEntityIdFailures(ctx, req.ImportId, req.EntityId, authorizationHeader))
}

type PostBulkImportsImportIdCancelRequest struct {
	ImportId int32 `json:"import_id" jsonschema:"description=The ID of user's GitLab Migration"`
}

func registerPostBulkImportsImportIdCancel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostBulkImportsImportIdCancelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_bulk_imports_import_id_cancel",
		mcp.WithDescription("This feature was introduced in GitLab 17.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postBulkImportsImportIdCancelHandler))
}

func postBulkImportsImportIdCancelHandler(ctx context.Context, request mcp.CallToolRequest, req PostBulkImportsImportIdCancelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4BulkImportsImportIdCancel(ctx, req.ImportId, authorizationHeader))
}
