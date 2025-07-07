package gitlab

import (
	"context"

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
