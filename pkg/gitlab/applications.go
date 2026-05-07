package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type GetApplicationsRequest struct {
}

func registerGetApplications(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetApplicationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_applications",
		mcp.WithDescription("List all registered applications"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getApplicationsHandler))
}

func getApplicationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetApplicationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4Applications(ctx, authorizationHeader))
}

type DeleteApplicationsIdRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of the application (not the application_id)"`
}

func registerDeleteApplicationsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteApplicationsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_applications_id",
		mcp.WithDescription("Delete a specific application"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteApplicationsIdHandler))
}

func deleteApplicationsIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteApplicationsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4ApplicationsId(ctx, req.Id, authorizationHeader))
}

type PostApplicationsIdRenewSecretRequest struct {
	Id int32 `json:"id" jsonschema:"description=The ID of the application (not the application_id)"`
}

func registerPostApplicationsIdRenewSecret(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostApplicationsIdRenewSecretRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_applications_id_renew_secret",
		mcp.WithDescription("Renew the secret of a specific application"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postApplicationsIdRenewSecretHandler))
}

func postApplicationsIdRenewSecretHandler(ctx context.Context, request mcp.CallToolRequest, req PostApplicationsIdRenewSecretRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4ApplicationsIdRenewSecret(ctx, req.Id, authorizationHeader))
}
