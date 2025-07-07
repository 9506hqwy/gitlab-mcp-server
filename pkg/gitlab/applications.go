package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetApplications(s *server.MCPServer) {
	tool := mcp.NewTool("get_applications",
		mcp.WithDescription("List all registered applications"),
	)

	s.AddTool(tool, getApplicationsHandler)
}

func getApplicationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Applications(ctx, authorizationHeader))
}

func registerDeleteApplicationsId(s *server.MCPServer) {
	tool := mcp.NewTool("delete_applications_id",
		mcp.WithDescription("Delete a specific application"),
		mcp.WithNumber("id",
			mcp.Description("The ID of the application (not the application_id)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, deleteApplicationsIdHandler)
}

func deleteApplicationsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.DeleteApiV4ApplicationsId(ctx, id, authorizationHeader))
}

func registerPostApplicationsIdRenewSecret(s *server.MCPServer) {
	tool := mcp.NewTool("post_applications_id_renew_secret",
		mcp.WithDescription("Renew the secret of a specific application"),
		mcp.WithNumber("id",
			mcp.Description("The ID of the application (not the application_id)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, postApplicationsIdRenewSecretHandler)
}

func postApplicationsIdRenewSecretHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))

	return toResult(c.PostApiV4ApplicationsIdRenewSecret(ctx, id, authorizationHeader))
}
