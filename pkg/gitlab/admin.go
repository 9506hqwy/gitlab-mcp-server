package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetAdminBatchedBackgroundMigrationsId(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_batched_background_migrations_id",
		mcp.WithDescription("Retrieve a batched background migration"),
		mcp.WithString("database",
			mcp.Description("The name of the database (default: main)"),

			mcp.Enum("main", "ci", "sec", "embedding", "geo"),
		),
		mcp.WithNumber("id",
			mcp.Description("The batched background migration id"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getAdminBatchedBackgroundMigrationsIdHandler)
}

func getAdminBatchedBackgroundMigrationsIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetAdminBatchedBackgroundMigrationsId(request)
	return toResult(c.GetApiV4AdminBatchedBackgroundMigrationsId(ctx, id, &params, authorizationHeader))
}

func parseGetAdminBatchedBackgroundMigrationsId(request mcp.CallToolRequest) client.GetApiV4AdminBatchedBackgroundMigrationsIdParams {
	params := client.GetApiV4AdminBatchedBackgroundMigrationsIdParams{}

	database := request.GetString("database", "")
	if database != "" {

		params.Database = &database
	}

	return params
}

func registerGetAdminBatchedBackgroundMigrations(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_batched_background_migrations",
		mcp.WithDescription("Get the list of batched background migrations"),
		mcp.WithString("database",
			mcp.Description("The name of the database, the default `main` (default: main)"),

			mcp.Enum("main", "ci", "sec", "embedding", "geo"),
		),
		mcp.WithString("job_class_name",
			mcp.Description("Filter migrations by job class name."),
		),
	)

	s.AddTool(tool, getAdminBatchedBackgroundMigrationsHandler)
}

func getAdminBatchedBackgroundMigrationsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetAdminBatchedBackgroundMigrations(request)
	return toResult(c.GetApiV4AdminBatchedBackgroundMigrations(ctx, &params, authorizationHeader))
}

func parseGetAdminBatchedBackgroundMigrations(request mcp.CallToolRequest) client.GetApiV4AdminBatchedBackgroundMigrationsParams {
	params := client.GetApiV4AdminBatchedBackgroundMigrationsParams{}

	database := request.GetString("database", "")
	if database != "" {

		params.Database = &database
	}

	job_class_name := request.GetString("job_class_name", "")
	if job_class_name != "" {

		params.JobClassName = &job_class_name
	}

	return params
}

func registerGetAdminCiVariables(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_ci_variables",
		mcp.WithDescription("List all instance-level variables"),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getAdminCiVariablesHandler)
}

func getAdminCiVariablesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetAdminCiVariables(request)
	return toResult(c.GetApiV4AdminCiVariables(ctx, &params, authorizationHeader))
}

func parseGetAdminCiVariables(request mcp.CallToolRequest) client.GetApiV4AdminCiVariablesParams {
	params := client.GetApiV4AdminCiVariablesParams{}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	return params
}

func registerDeleteAdminCiVariablesKey(s *server.MCPServer) {
	tool := mcp.NewTool("delete_admin_ci_variables_key",
		mcp.WithDescription("Delete an existing instance-level variable"),
		mcp.WithString("key",
			mcp.Description("The key of a variable"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, deleteAdminCiVariablesKeyHandler)
}

func deleteAdminCiVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	key := request.GetString("key", "")

	return toResult(c.DeleteApiV4AdminCiVariablesKey(ctx, key, authorizationHeader))
}

func registerGetAdminCiVariablesKey(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_ci_variables_key",
		mcp.WithDescription("Get the details of a specific instance-level variable"),
		mcp.WithString("key",
			mcp.Description("The key of a variable"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getAdminCiVariablesKeyHandler)
}

func getAdminCiVariablesKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	key := request.GetString("key", "")

	return toResult(c.GetApiV4AdminCiVariablesKey(ctx, key, authorizationHeader))
}

func registerGetAdminDatabasesDatabaseNameDictionaryTablesTableName(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_databases_database_name_dictionary_tables_table_name",
		mcp.WithDescription("Retrieve dictionary details"),
		mcp.WithString("database_name",
			mcp.Description("The database name"),
			mcp.Required(),
			mcp.Enum("main", "ci"),
		),
		mcp.WithString("table_name",
			mcp.Description("The table name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getAdminDatabasesDatabaseNameDictionaryTablesTableNameHandler)
}

func getAdminDatabasesDatabaseNameDictionaryTablesTableNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	database_name := request.GetString("database_name", "")
	table_name := request.GetString("table_name", "")

	return toResult(c.GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName(ctx, database_name, table_name, authorizationHeader))
}

func registerGetAdminClusters(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_clusters",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Returns a list of instance clusters."),
	)

	s.AddTool(tool, getAdminClustersHandler)
}

func getAdminClustersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4AdminClusters(ctx, authorizationHeader))
}

func registerDeleteAdminClustersClusterId(s *server.MCPServer) {
	tool := mcp.NewTool("delete_admin_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Deletes an existing instance cluster. Does not remove existing resources within the connected Kubernetes cluster."),
		mcp.WithNumber("cluster_id",
			mcp.Description("The cluster ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, deleteAdminClustersClusterIdHandler)
}

func deleteAdminClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	cluster_id := int32(request.GetInt("cluster_id", math.MinInt))

	return toResult(c.DeleteApiV4AdminClustersClusterId(ctx, cluster_id, authorizationHeader))
}

func registerGetAdminClustersClusterId(s *server.MCPServer) {
	tool := mcp.NewTool("get_admin_clusters_cluster_id",
		mcp.WithDescription("This feature was introduced in GitLab 13.2. Returns a single instance cluster."),
		mcp.WithNumber("cluster_id",
			mcp.Description("The cluster ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getAdminClustersClusterIdHandler)
}

func getAdminClustersClusterIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	cluster_id := int32(request.GetInt("cluster_id", math.MinInt))

	return toResult(c.GetApiV4AdminClustersClusterId(ctx, cluster_id, authorizationHeader))
}
