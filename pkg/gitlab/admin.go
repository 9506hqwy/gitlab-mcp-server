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

type PutAdminBatchedBackgroundMigrationsIdResumeRequest struct {
	Id int32 `json:"id" jsonschema:"description=The batched background migration id"`

	Body client.PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody `json:"body"`
}

func registerPutAdminBatchedBackgroundMigrationsIdResume(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutAdminBatchedBackgroundMigrationsIdResumeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_admin_batched_background_migrations_id_resume",
		mcp.WithDescription("Resume a batched background migration"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putAdminBatchedBackgroundMigrationsIdResumeHandler))
}

func putAdminBatchedBackgroundMigrationsIdResumeHandler(ctx context.Context, request mcp.CallToolRequest, req PutAdminBatchedBackgroundMigrationsIdResumeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4AdminBatchedBackgroundMigrationsIdResume(ctx, req.Id, req.Body, authorizationHeader))
}

type PutAdminBatchedBackgroundMigrationsIdPauseRequest struct {
	Id int32 `json:"id" jsonschema:"description=The batched background migration id"`

	Body client.PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody `json:"body"`
}

func registerPutAdminBatchedBackgroundMigrationsIdPause(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutAdminBatchedBackgroundMigrationsIdPauseRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_admin_batched_background_migrations_id_pause",
		mcp.WithDescription("Pause a batched background migration"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putAdminBatchedBackgroundMigrationsIdPauseHandler))
}

func putAdminBatchedBackgroundMigrationsIdPauseHandler(ctx context.Context, request mcp.CallToolRequest, req PutAdminBatchedBackgroundMigrationsIdPauseRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4AdminBatchedBackgroundMigrationsIdPause(ctx, req.Id, req.Body, authorizationHeader))
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

type PostAdminCiVariablesRequest struct {
	Body client.PostApiV4AdminCiVariablesJSONRequestBody `json:"body"`
}

func registerPostAdminCiVariables(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostAdminCiVariablesRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_admin_ci_variables",
		mcp.WithDescription("Create a new instance-level variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postAdminCiVariablesHandler))
}

func postAdminCiVariablesHandler(ctx context.Context, request mcp.CallToolRequest, req PostAdminCiVariablesRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4AdminCiVariables(ctx, req.Body, authorizationHeader))
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

type PutAdminCiVariablesKeyRequest struct {
	Key string `json:"key" jsonschema:"description=The key of a variable"`

	Body client.PutApiV4AdminCiVariablesKeyJSONRequestBody `json:"body"`
}

func registerPutAdminCiVariablesKey(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutAdminCiVariablesKeyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_admin_ci_variables_key",
		mcp.WithDescription("Update an instance-level variable"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putAdminCiVariablesKeyHandler))
}

func putAdminCiVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest, req PutAdminCiVariablesKeyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4AdminCiVariablesKey(ctx, req.Key, req.Body, authorizationHeader))
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

type PutAdminClustersClusterIdRequest struct {
	ClusterId int32 `json:"cluster_id" jsonschema:"description=The cluster ID"`

	Body client.PutApiV4AdminClustersClusterIdJSONRequestBody `json:"body"`
}

func registerPutAdminClustersClusterId(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutAdminClustersClusterIdRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_admin_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Updates an existing instance cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putAdminClustersClusterIdHandler))
}

func putAdminClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest, req PutAdminClustersClusterIdRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4AdminClustersClusterId(ctx, req.ClusterId, req.Body, authorizationHeader))
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

type PostAdminClustersAddRequest struct {
	Body client.PostApiV4AdminClustersAddJSONRequestBody `json:"body"`
}

func registerPostAdminClustersAdd(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostAdminClustersAddRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_admin_clusters_add",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Adds an existing Kubernetes instance cluster."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postAdminClustersAddHandler))
}

func postAdminClustersAddHandler(ctx context.Context, request mcp.CallToolRequest, req PostAdminClustersAddRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4AdminClustersAdd(ctx, req.Body, authorizationHeader))
}

type PostAdminMigrationsTimestampMarkRequest struct {
	Timestamp int32 `json:"timestamp" jsonschema:"description=The migration version timestamp"`

	Body client.PostApiV4AdminMigrationsTimestampMarkJSONRequestBody `json:"body"`
}

func registerPostAdminMigrationsTimestampMark(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PostAdminMigrationsTimestampMarkRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("post_admin_migrations_timestamp_mark",
		mcp.WithDescription("Mark the migration as successfully executed"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(postAdminMigrationsTimestampMarkHandler))
}

func postAdminMigrationsTimestampMarkHandler(ctx context.Context, request mcp.CallToolRequest, req PostAdminMigrationsTimestampMarkRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PostApiV4AdminMigrationsTimestampMark(ctx, req.Timestamp, req.Body, authorizationHeader))
}
