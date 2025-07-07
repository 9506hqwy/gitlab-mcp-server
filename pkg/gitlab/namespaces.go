package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetNamespacesId(s *server.MCPServer) {
	tool := mcp.NewTool("get_namespaces_id",
		mcp.WithDescription("Get a namespace by ID"),
		mcp.WithString("id",
			mcp.Description("ID or URL-encoded path of the namespace"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getNamespacesIdHandler)
}

func getNamespacesIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4NamespacesId(ctx, id, authorizationHeader))
}

func registerGetNamespacesIdGitlabSubscription(s *server.MCPServer) {
	tool := mcp.NewTool("get_namespaces_id_gitlab_subscription",
		mcp.WithDescription("Returns the subscription for the namespace"),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getNamespacesIdGitlabSubscriptionHandler)
}

func getNamespacesIdGitlabSubscriptionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.GetApiV4NamespacesIdGitlabSubscription(ctx, id, authorizationHeader))
}

func registerDeleteNamespacesIdStorageLimitExclusion(s *server.MCPServer) {
	tool := mcp.NewTool("delete_namespaces_id_storage_limit_exclusion",
		mcp.WithDescription("Removes a Namespaces::Storage::LimitExclusion"),
		mcp.WithNumber("id",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, deleteNamespacesIdStorageLimitExclusionHandler)
}

func deleteNamespacesIdStorageLimitExclusionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.DeleteApiV4NamespacesIdStorageLimitExclusion(ctx, id, authorizationHeader))
}

func registerGetNamespacesStorageLimitExclusions(s *server.MCPServer) {
	tool := mcp.NewTool("get_namespaces_storage_limit_exclusions",
		mcp.WithDescription("Gets all records for namespaces that have been excluded"),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getNamespacesStorageLimitExclusionsHandler)
}

func getNamespacesStorageLimitExclusionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetNamespacesStorageLimitExclusions(request)
	return toResult(c.GetApiV4NamespacesStorageLimitExclusions(ctx, &params, authorizationHeader))
}

func parseGetNamespacesStorageLimitExclusions(request mcp.CallToolRequest) client.GetApiV4NamespacesStorageLimitExclusionsParams {
	params := client.GetApiV4NamespacesStorageLimitExclusionsParams{}

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

func registerGetNamespaces(s *server.MCPServer) {
	tool := mcp.NewTool("get_namespaces",
		mcp.WithDescription("Get a list of the namespaces of the authenticated user. If the user is an administrator, a list of all namespaces in the GitLab instance is shown."),
		mcp.WithString("search",
			mcp.Description("Returns a list of namespaces the user is authorized to view based on the search criteria"),
		),
		mcp.WithBoolean("owned_only",
			mcp.Description("In GitLab 14.2 and later, returns a list of owned namespaces only"),
		),
		mcp.WithBoolean("top_level_only",
			mcp.Description("Only include top level namespaces (default: false)"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
		mcp.WithString("requested_hosted_plan",
			mcp.Description("Name of the hosted plan requested by the customer"),
		),
	)

	s.AddTool(tool, getNamespacesHandler)
}

func getNamespacesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetNamespaces(request)
	return toResult(c.GetApiV4Namespaces(ctx, &params, authorizationHeader))
}

func parseGetNamespaces(request mcp.CallToolRequest) client.GetApiV4NamespacesParams {
	params := client.GetApiV4NamespacesParams{}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	owned_only := request.GetBool("owned_only", false)
	params.OwnedOnly = &owned_only

	top_level_only := request.GetBool("top_level_only", false)
	params.TopLevelOnly = &top_level_only

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

	requested_hosted_plan := request.GetString("requested_hosted_plan", "")
	if requested_hosted_plan != "" {

		params.RequestedHostedPlan = &requested_hosted_plan
	}

	return params
}

func registerGetNamespacesIdExists(s *server.MCPServer) {
	tool := mcp.NewTool("get_namespaces_id_exists",
		mcp.WithDescription("Get existence of a namespace by path. Suggests a new namespace path that does not already exist."),
		mcp.WithString("id",
			mcp.Description("Namespace’s path"),
			mcp.Required(),
		),
		mcp.WithNumber("parent_id",
			mcp.Description("The ID of the parent namespace. If no ID is specified, only top-level namespaces are considered."),
		),
	)

	s.AddTool(tool, getNamespacesIdExistsHandler)
}

func getNamespacesIdExistsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseGetNamespacesIdExists(request)
	return toResult(c.GetApiV4NamespacesIdExists(ctx, id, &params, authorizationHeader))
}

func parseGetNamespacesIdExists(request mcp.CallToolRequest) client.GetApiV4NamespacesIdExistsParams {
	params := client.GetApiV4NamespacesIdExistsParams{}

	parent_id := request.GetInt("parent_id", math.MinInt)
	if parent_id != math.MinInt {
		parent_id := int32(parent_id)
		params.ParentId = &parent_id
	}

	return params
}
