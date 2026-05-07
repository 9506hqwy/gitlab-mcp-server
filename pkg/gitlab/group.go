package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type GetGroupIdPackagesComposerPackagesRequest struct {
	Id string `json:"id" jsonschema:"description=The ID or URL-encoded path of a group"`
}

func registerGetGroupIdPackagesComposerPackages(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupIdPackagesComposerPackagesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_group_id_pkgs_composer_packages",
		mcp.WithDescription("This feature was introduced in GitLab 13.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupIdPackagesComposerPackagesHandler))
}

func getGroupIdPackagesComposerPackagesHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupIdPackagesComposerPackagesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupIdPackagesComposerPackages(ctx, req.Id, authorizationHeader))
}

type GetGroupIdPackagesComposerPShaRequest struct {
	Id  string `json:"id" jsonschema:"description=The ID or URL-encoded path of a group"`
	Sha string `json:"sha" jsonschema:"description=Shasum of current json"`
}

func registerGetGroupIdPackagesComposerPSha(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetGroupIdPackagesComposerPShaRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_group_id_pkgs_composer_p_sha",
		mcp.WithDescription("This feature was introduced in GitLab 13.1"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getGroupIdPackagesComposerPShaHandler))
}

func getGroupIdPackagesComposerPShaHandler(ctx context.Context, request mcp.CallToolRequest, req GetGroupIdPackagesComposerPShaRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4GroupIdPackagesComposerPSha(ctx, req.Id, req.Sha, authorizationHeader))
}
