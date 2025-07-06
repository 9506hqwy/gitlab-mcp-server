package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetKeysId(s *server.MCPServer) {
	tool := mcp.NewTool("keys_id",
		mcp.WithDescription("Get SSH key with user by ID of an SSH key. Note only administrators can lookup SSH key with user by ID of an SSH key"),
		mcp.WithString("id",
			mcp.Description("The ID of an SSH key (example: 2)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getKeysIdHandler)
}

func getKeysIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")

	return toResult(c.GetApiV4KeysId(ctx, id, authorizationHeader))
}

func registerGetKeys(s *server.MCPServer) {
	tool := mcp.NewTool("keys_",
		mcp.WithDescription("You can search for a user that owns a specific SSH key. Note only administrators can lookup SSH key with the fingerprint of an SSH key"),
		mcp.WithString("fingerprint",
			mcp.Description("The fingerprint of an SSH key (example: ba:81:59:68:d7:6c:cd:02:02:bf:6a:9b:55:4e:af:d1)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getKeysHandler)
}

func getKeysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetKeys(request)
	return toResult(c.GetApiV4Keys(ctx, &params, authorizationHeader))
}

func parseGetKeys(request mcp.CallToolRequest) client.GetApiV4KeysParams {
	params := client.GetApiV4KeysParams{}

	fingerprint := request.GetString("fingerprint", "")
	if fingerprint != "" {

		params.Fingerprint = fingerprint
	}

	return params
}
