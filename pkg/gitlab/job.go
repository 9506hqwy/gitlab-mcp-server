package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGetJob(s *server.MCPServer) {
	tool := mcp.NewTool("job_",
		mcp.WithDescription("Get current job using job token"),
	)

	s.AddTool(tool, getJobHandler)
}

func getJobHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Job(ctx, authorizationHeader))
}

func registerGetJobAllowedAgents(s *server.MCPServer) {
	tool := mcp.NewTool("job_allowed_agents",
		mcp.WithDescription("Retrieves a list of agents for the given job token"),
	)

	s.AddTool(tool, getJobAllowedAgentsHandler)
}

func getJobAllowedAgentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4JobAllowedAgents(ctx, authorizationHeader))
}
