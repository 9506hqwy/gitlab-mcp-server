package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetFeatures(s *server.MCPServer) {
	tool := mcp.NewTool("get_features",
		mcp.WithDescription("Get a list of all persisted features, with its gate values."),
	)

	s.AddTool(tool, getFeaturesHandler)
}

func getFeaturesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Features(ctx, authorizationHeader))
}

func registerGetFeaturesDefinitions(s *server.MCPServer) {
	tool := mcp.NewTool("get_features_definitions",
		mcp.WithDescription("Get a list of all feature definitions."),
	)

	s.AddTool(tool, getFeaturesDefinitionsHandler)
}

func getFeaturesDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4FeaturesDefinitions(ctx, authorizationHeader))
}

func registerDeleteFeaturesName(s *server.MCPServer) {
	tool := mcp.NewTool("delete_features_name",
		mcp.WithDescription("Removes a feature gate. Response is equal when the gate exists, or doesn't."),
		mcp.WithNumber("name",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, deleteFeaturesNameHandler)
}

func deleteFeaturesNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	name := int32(request.GetInt("name", math.MinInt))

	return toResult(c.DeleteApiV4FeaturesName(ctx, name, authorizationHeader))
}
