package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetAdminBatchedBackgroundMigrationsIdRequest struct {
	Id     int32                                                    `json:"id" jsonschema:"description=The batched background migration id"`
	Params *client.GetApiV4AdminBatchedBackgroundMigrationsIdParams `json:"params,omitempty"`
}

func registerGetAdminBatchedBackgroundMigrationsId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminBatchedBackgroundMigrationsIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_batched_background_migrations_id",
		mcp.WithDescription("Retrieve a batched background migration"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminBatchedBackgroundMigrationsIdHandler))
}

func getAdminBatchedBackgroundMigrationsIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminBatchedBackgroundMigrationsIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminBatchedBackgroundMigrationsId(ctx, req.Id, req.Params, authorizationHeader))
}

type GetAdminBatchedBackgroundMigrationsRequest struct {
	Params *client.GetApiV4AdminBatchedBackgroundMigrationsParams `json:"params,omitempty"`
}

func registerGetAdminBatchedBackgroundMigrations(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminBatchedBackgroundMigrationsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_batched_background_migrations",
		mcp.WithDescription("Get the list of batched background migrations"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminBatchedBackgroundMigrationsHandler))
}

func getAdminBatchedBackgroundMigrationsHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminBatchedBackgroundMigrationsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminBatchedBackgroundMigrations(ctx, req.Params, authorizationHeader))
}

type GetAdminCiVariablesRequest struct {
	Params *client.GetApiV4AdminCiVariablesParams `json:"params,omitempty"`
}

func registerGetAdminCiVariables(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminCiVariablesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_ci_variables",
		mcp.WithDescription("List all instance-level variables"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminCiVariablesHandler))
}

func getAdminCiVariablesHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminCiVariablesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminCiVariables(ctx, req.Params, authorizationHeader))
}

type DeleteAdminCiVariablesKeyRequest struct {
	Key string `json:"key" jsonschema:"description=The key of a variable"`
}

func registerDeleteAdminCiVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteAdminCiVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_admin_ci_variables_key",
		mcp.WithDescription("Delete an existing instance-level variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteAdminCiVariablesKeyHandler))
}

func deleteAdminCiVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteAdminCiVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4AdminCiVariablesKey(ctx, req.Key, authorizationHeader))
}

type GetAdminCiVariablesKeyRequest struct {
	Key string `json:"key" jsonschema:"description=The key of a variable"`
}

func registerGetAdminCiVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminCiVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_ci_variables_key",
		mcp.WithDescription("Get the details of a specific instance-level variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminCiVariablesKeyHandler))
}

func getAdminCiVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminCiVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminCiVariablesKey(ctx, req.Key, authorizationHeader))
}

type GetAdminDatabasesDatabaseNameDictionaryTablesTableNameRequest struct {
	DatabaseName string `json:"database_name" jsonschema:"description=The database name"`
	TableName    string `json:"table_name" jsonschema:"description=The table name"`
}

func registerGetAdminDatabasesDatabaseNameDictionaryTablesTableName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminDatabasesDatabaseNameDictionaryTablesTableNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_databases_database_name_dictionary_tables_table_name",
		mcp.WithDescription("Retrieve dictionary details"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminDatabasesDatabaseNameDictionaryTablesTableNameHandler))
}

func getAdminDatabasesDatabaseNameDictionaryTablesTableNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminDatabasesDatabaseNameDictionaryTablesTableNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName(ctx, req.DatabaseName, req.TableName, authorizationHeader))
}

type GetAdminClustersRequest struct {
}

func registerGetAdminClusters(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminClustersRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Returns a list of instance clusters."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminClustersHandler))
}

func getAdminClustersHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminClustersRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminClusters(ctx, authorizationHeader))
}

type DeleteAdminClustersClusterIdRequest struct {
	ClusterId int32 `json:"cluster_id" jsonschema:"description=The cluster ID"`
}

func registerDeleteAdminClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeleteAdminClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_admin_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Deletes an existing instance cluster. Does not remove existing resources within the connected Kubernetes cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deleteAdminClustersClusterIdHandler))
}

func deleteAdminClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req DeleteAdminClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4AdminClustersClusterId(ctx, req.ClusterId, authorizationHeader))
}

type GetAdminClustersClusterIdRequest struct {
	ClusterId int32 `json:"cluster_id" jsonschema:"description=The cluster ID"`
}

func registerGetAdminClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetAdminClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_admin_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Returns a single instance cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getAdminClustersClusterIdHandler))
}

func getAdminClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req GetAdminClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminClustersClusterId(ctx, req.ClusterId, authorizationHeader))
}
