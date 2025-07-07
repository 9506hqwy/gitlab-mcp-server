package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetBulkImports(s *server.MCPServer) {
	tool := mcp.NewTool("get_bulk_imports",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("sort",
			mcp.Description("Return GitLab Migrations sorted in created by `asc` or `desc` order. (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("status",
			mcp.Description("Return GitLab Migrations with specified status"),

			mcp.Enum("created", "started", "finished", "timeout", "failed", "canceled"),
		),
	)

	s.AddTool(tool, getBulkImportsHandler)
}

func getBulkImportsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetBulkImports(request)
	return toResult(c.GetApiV4BulkImports(ctx, &params, authorizationHeader))
}

func parseGetBulkImports(request mcp.CallToolRequest) client.GetApiV4BulkImportsParams {
	params := client.GetApiV4BulkImportsParams{}

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

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	return params
}

func registerGetBulkImportsEntities(s *server.MCPServer) {
	tool := mcp.NewTool("get_bulk_imports_entities",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("sort",
			mcp.Description("Return GitLab Migrations sorted in created by `asc` or `desc` order. (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithString("status",
			mcp.Description("Return all GitLab Migrations' entities with specified status"),

			mcp.Enum("created", "started", "finished", "timeout", "failed", "canceled"),
		),
	)

	s.AddTool(tool, getBulkImportsEntitiesHandler)
}

func getBulkImportsEntitiesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetBulkImportsEntities(request)
	return toResult(c.GetApiV4BulkImportsEntities(ctx, &params, authorizationHeader))
}

func parseGetBulkImportsEntities(request mcp.CallToolRequest) client.GetApiV4BulkImportsEntitiesParams {
	params := client.GetApiV4BulkImportsEntitiesParams{}

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

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
	}

	return params
}

func registerGetBulkImportsImportId(s *server.MCPServer) {
	tool := mcp.NewTool("get_bulk_imports_import_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithNumber("import_id",
			mcp.Description("The ID of user's GitLab Migration"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getBulkImportsImportIdHandler)
}

func getBulkImportsImportIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	import_id := int32(request.GetInt("import_id", math.MinInt))

	return toResult(c.GetApiV4BulkImportsImportId(ctx, import_id, authorizationHeader))
}

func registerGetBulkImportsImportIdEntities(s *server.MCPServer) {
	tool := mcp.NewTool("get_bulk_imports_import_id_entities",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithNumber("import_id",
			mcp.Description("The ID of user's GitLab Migration"),
			mcp.Required(),
		),
		mcp.WithString("status",
			mcp.Description("Return import entities with specified status"),

			mcp.Enum("created", "started", "finished", "timeout", "failed", "canceled"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getBulkImportsImportIdEntitiesHandler)
}

func getBulkImportsImportIdEntitiesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	import_id := int32(request.GetInt("import_id", math.MinInt))
	params := parseGetBulkImportsImportIdEntities(request)
	return toResult(c.GetApiV4BulkImportsImportIdEntities(ctx, import_id, &params, authorizationHeader))
}

func parseGetBulkImportsImportIdEntities(request mcp.CallToolRequest) client.GetApiV4BulkImportsImportIdEntitiesParams {
	params := client.GetApiV4BulkImportsImportIdEntitiesParams{}

	status := request.GetString("status", "")
	if status != "" {

		params.Status = &status
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

func registerGetBulkImportsImportIdEntitiesEntityId(s *server.MCPServer) {
	tool := mcp.NewTool("get_bulk_imports_import_id_entities_entity_id",
		mcp.WithDescription("This feature was introduced in GitLab 14.1."),
		mcp.WithNumber("import_id",
			mcp.Description("The ID of user's GitLab Migration"),
			mcp.Required(),
		),
		mcp.WithNumber("entity_id",
			mcp.Description("The ID of GitLab Migration entity"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getBulkImportsImportIdEntitiesEntityIdHandler)
}

func getBulkImportsImportIdEntitiesEntityIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	import_id := int32(request.GetInt("import_id", math.MinInt))
	entity_id := int32(request.GetInt("entity_id", math.MinInt))

	return toResult(c.GetApiV4BulkImportsImportIdEntitiesEntityId(ctx, import_id, entity_id, authorizationHeader))
}

func registerGetBulkImportsImportIdEntitiesEntityIdFailures(s *server.MCPServer) {
	tool := mcp.NewTool("get_bulk_imports_import_id_entities_entity_id_failures",
		mcp.WithDescription("This feature was introduced in GitLab 16.6"),
		mcp.WithNumber("import_id",
			mcp.Description("The ID of user's GitLab Migration"),
			mcp.Required(),
		),
		mcp.WithNumber("entity_id",
			mcp.Description("The ID of GitLab Migration entity"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getBulkImportsImportIdEntitiesEntityIdFailuresHandler)
}

func getBulkImportsImportIdEntitiesEntityIdFailuresHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	import_id := int32(request.GetInt("import_id", math.MinInt))
	entity_id := int32(request.GetInt("entity_id", math.MinInt))

	return toResult(c.GetApiV4BulkImportsImportIdEntitiesEntityIdFailures(ctx, import_id, entity_id, authorizationHeader))
}
