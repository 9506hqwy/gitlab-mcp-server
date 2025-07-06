package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetAvatar(s *server.MCPServer) {
	tool := mcp.NewTool("avatar_",
		mcp.WithDescription("Return avatar url for a user"),
		mcp.WithString("email",
			mcp.Description("Public email address of the user"),
			mcp.Required(),
		),
		mcp.WithNumber("size",
			mcp.Description("Single pixel dimension for Gravatar images"),
		),
	)

	s.AddTool(tool, getAvatarHandler)
}

func getAvatarHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetAvatar(request)
	return toResult(c.GetApiV4Avatar(ctx, &params, authorizationHeader))
}

func parseGetAvatar(request mcp.CallToolRequest) client.GetApiV4AvatarParams {
	params := client.GetApiV4AvatarParams{}

	email := request.GetString("email", "")
	if email != "" {

		params.Email = email
	}

	size := request.GetInt("size", math.MinInt)
	if size != math.MinInt {
		size := int32(size)
		params.Size = &size
	}

	return params
}
