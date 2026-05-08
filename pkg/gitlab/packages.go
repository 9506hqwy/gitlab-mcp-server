package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetPackagesConanV1UsersAuthenticateRequest struct {
}

func registerGetPackagesConanV1UsersAuthenticate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1UsersAuthenticateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_users_authenticate",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1UsersAuthenticateHandler))
}

func getPackagesConanV1UsersAuthenticateHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1UsersAuthenticateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1UsersAuthenticate(ctx, authorizationHeader))
}

type GetPackagesConanV1UsersCheckCredentialsRequest struct {
}

func registerGetPackagesConanV1UsersCheckCredentials(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1UsersCheckCredentialsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_users_check_credentials",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1UsersCheckCredentialsHandler))
}

func getPackagesConanV1UsersCheckCredentialsHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1UsersCheckCredentialsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1UsersCheckCredentials(ctx, authorizationHeader))
}

type GetPackagesConanV1ConansSearchRequest struct {
	Params *client.GetApiV4PackagesConanV1ConansSearchParams `json:"params,omitempty"`
}

func registerGetPackagesConanV1ConansSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_search",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansSearchHandler))
}

func getPackagesConanV1ConansSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansSearch(ctx, req.Params, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetPackagesConanV1PingRequest struct {
}

func registerGetPackagesConanV1Ping(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1PingRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_ping",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1PingHandler))
}

func getPackagesConanV1PingHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1PingRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1Ping(ctx, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceRequest struct {
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type DeletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerDeletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&DeletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("delete_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(deletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler))
}

func deletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler(ctx context.Context, request mcp.CallToolRequest, req DeletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestRequest struct {
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsRequest struct {
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan package ID"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.ConanPackageReference, authorizationHeader))
}

type GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler))
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, authorizationHeader))
}

type PutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	FileName        string `json:"file_name" jsonschema:"description=Package file name"`

	Body client.PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameJSONRequestBody `json:"body"`
}

func registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_export_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler))
}

func putPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req PutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.FileName, req.Body, authorizationHeader))
}

type GetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest struct {
	PackageName     string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion  string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel  string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision  string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	FileName        string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_export_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler))
}

func getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.FileName, authorizationHeader))
}

type PutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest struct {
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan Package ID"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Conan Package Revision"`
	FileName              string `json:"file_name" jsonschema:"description=Package file name"`

	Body client.PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameJSONRequestBody `json:"body"`
}

func registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&PutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("put_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_package_conan_package_reference_package_revision_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(putPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler))
}

func putPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req PutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, req.FileName, req.Body, authorizationHeader))
}

type GetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest struct {
	PackageName           string `json:"package_name" jsonschema:"description=Package name"`
	PackageVersion        string `json:"package_version" jsonschema:"description=Package version"`
	PackageUsername       string `json:"package_username" jsonschema:"description=Package username"`
	PackageChannel        string `json:"package_channel" jsonschema:"description=Package channel"`
	RecipeRevision        string `json:"recipe_revision" jsonschema:"description=Conan Recipe Revision"`
	ConanPackageReference string `json:"conan_package_reference" jsonschema:"description=Conan Package ID"`
	PackageRevision       string `json:"package_revision" jsonschema:"description=Conan Package Revision"`
	FileName              string `json:"file_name" jsonschema:"description=Package file name"`
}

func registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_package_conan_package_reference_package_revision_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler))
}

func getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, req.PackageName, req.PackageVersion, req.PackageUsername, req.PackageChannel, req.RecipeRevision, req.ConanPackageReference, req.PackageRevision, req.FileName, authorizationHeader))
}

type GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsRequest struct {
	ModuleNamespace string `json:"module_namespace" jsonschema:"description=Group's ID or slug"`
	ModuleName      string `json:"module_name" jsonschema:"description="`
	ModuleSystem    string `json:"module_system" jsonschema:"description=null"`
}

func registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_terraform_modules_v1_module_namespace_module_name_module_system_versions",
		mcp.WithDescription("List versions for a module"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsHandler))
}

func getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(ctx, req.ModuleNamespace, req.ModuleName, req.ModuleSystem, authorizationHeader))
}

type GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadRequest struct {
	ModuleNamespace string `json:"module_namespace" jsonschema:"description=Group's ID or slug"`
	ModuleName      string `json:"module_name" jsonschema:"description="`
	ModuleSystem    string `json:"module_system" jsonschema:"description=null"`
}

func registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_terraform_modules_v1_module_namespace_module_name_module_system_download",
		mcp.WithDescription("Download the latest version of a module"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadHandler))
}

func getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(ctx, req.ModuleNamespace, req.ModuleName, req.ModuleSystem, authorizationHeader))
}

type GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemRequest struct {
	ModuleNamespace string `json:"module_namespace" jsonschema:"description=Group's ID or slug"`
	ModuleName      string `json:"module_name" jsonschema:"description="`
	ModuleSystem    string `json:"module_system" jsonschema:"description=null"`
}

func registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_pkgs_terraform_modules_v1_module_namespace_module_name_module_system",
		mcp.WithDescription("Get details about the latest version of a module"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemHandler))
}

func getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemHandler(ctx context.Context, request mcp.CallToolRequest, req GetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(ctx, req.ModuleNamespace, req.ModuleName, req.ModuleSystem, authorizationHeader))
}
