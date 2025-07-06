package gitlab

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetPackagesConanV1UsersAuthenticate(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_users_authenticate",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
	)

	s.AddTool(tool, getPackagesConanV1UsersAuthenticateHandler)
}

func getPackagesConanV1UsersAuthenticateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1UsersAuthenticate(ctx, authorizationHeader))
}

func registerGetPackagesConanV1UsersCheckCredentials(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_users_check_credentials",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
	)

	s.AddTool(tool, getPackagesConanV1UsersCheckCredentialsHandler)
}

func getPackagesConanV1UsersCheckCredentialsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1UsersCheckCredentials(ctx, authorizationHeader))
}

func registerGetPackagesConanV1ConansSearch(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_search",
		mcp.WithDescription("This feature was introduced in GitLab 12.4"),
		mcp.WithString("q",
			mcp.Description("Search query (example: Hello*)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansSearchHandler)
}

func getPackagesConanV1ConansSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetPackagesConanV1ConansSearch(request)
	return toResult(c.GetApiV4PackagesConanV1ConansSearch(ctx, &params, authorizationHeader))
}

func parseGetPackagesConanV1ConansSearch(request mcp.CallToolRequest) client.GetApiV4PackagesConanV1ConansSearchParams {
	params := client.GetApiV4PackagesConanV1ConansSearchParams{}

	q := request.GetString("q", "")
	if q != "" {

		params.Q = q
	}

	return params
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_search",
		mcp.WithDescription("This feature was introduced in GitLab 18.0"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetPackagesConanV1Ping(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_ping",
		mcp.WithDescription("This feature was introduced in GitLab 12.2"),
	)

	s.AddTool(tool, getPackagesConanV1PingHandler)
}

func getPackagesConanV1PingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4PackagesConanV1Ping(ctx, authorizationHeader))
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx, package_name, package_version, package_username, package_channel, conan_package_reference, authorizationHeader))
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx, package_name, package_version, package_username, package_channel, conan_package_reference, authorizationHeader))
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_digest",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_packages_conan_package_reference_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	conan_package_reference := request.GetString("conan_package_reference", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx, package_name, package_version, package_username, package_channel, conan_package_reference, authorizationHeader))
}

func registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_conans_package_name_package_version_package_username_package_channel_download_urls",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler)
}

func getPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")

	return toResult(c.GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx, package_name, package_version, package_username, package_channel, authorizationHeader))
}

func registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_export_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.6"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Conan Recipe Revision (example: 0)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name (example: conanfile.py)"),
			mcp.Required(),
			mcp.Enum("conanfile.py", "conanmanifest.txt", "conan_sources.tgz", "conan_export.tgz", "conaninfo.txt", "conan_package.tgz"),
		),
	)

	s.AddTool(tool, getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler)
}

func getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx, package_name, package_version, package_username, package_channel, recipe_revision, file_name, authorizationHeader))
}

func registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_conan_v1_files_package_name_package_version_package_username_package_channel_recipe_revision_package_conan_package_reference_package_revision_file_name",
		mcp.WithDescription("This feature was introduced in GitLab 12.5"),
		mcp.WithString("package_name",
			mcp.Description("Package name (example: my-package)"),
			mcp.Required(),
		),
		mcp.WithString("package_version",
			mcp.Description("Package version (example: 1.0)"),
			mcp.Required(),
		),
		mcp.WithString("package_username",
			mcp.Description("Package username (example: my-group+my-project)"),
			mcp.Required(),
		),
		mcp.WithString("package_channel",
			mcp.Description("Package channel (example: stable)"),
			mcp.Required(),
		),
		mcp.WithString("recipe_revision",
			mcp.Description("Conan Recipe Revision (example: 0)"),
			mcp.Required(),
		),
		mcp.WithString("conan_package_reference",
			mcp.Description("Conan Package ID (example: 103f6067a947f366ef91fc1b7da351c588d1827f)"),
			mcp.Required(),
		),
		mcp.WithString("package_revision",
			mcp.Description("Conan Package Revision (example: 0)"),
			mcp.Required(),
		),
		mcp.WithString("file_name",
			mcp.Description("Package file name (example: conaninfo.txt)"),
			mcp.Required(),
			mcp.Enum("conanfile.py", "conanmanifest.txt", "conan_sources.tgz", "conan_export.tgz", "conaninfo.txt", "conan_package.tgz"),
		),
	)

	s.AddTool(tool, getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler)
}

func getPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	package_name := request.GetString("package_name", "")
	package_version := request.GetString("package_version", "")
	package_username := request.GetString("package_username", "")
	package_channel := request.GetString("package_channel", "")
	recipe_revision := request.GetString("recipe_revision", "")
	conan_package_reference := request.GetString("conan_package_reference", "")
	package_revision := request.GetString("package_revision", "")
	file_name := request.GetString("file_name", "")

	return toResult(c.GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx, package_name, package_version, package_username, package_channel, recipe_revision, conan_package_reference, package_revision, file_name, authorizationHeader))
}

func registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_terraform_modules_v1_module_namespace_module_name_module_system_versions",
		mcp.WithDescription("List versions for a module"),
		mcp.WithString("module_namespace",
			mcp.Description("Group's ID or slug"),
			mcp.Required(),
		),
		mcp.WithString("module_name",
			mcp.Description(""),
			mcp.Required(),
		),
		mcp.WithString("module_system",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsHandler)
}

func getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	module_namespace := request.GetString("module_namespace", "")
	module_name := request.GetString("module_name", "")
	module_system := request.GetString("module_system", "")

	return toResult(c.GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(ctx, module_namespace, module_name, module_system, authorizationHeader))
}

func registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_terraform_modules_v1_module_namespace_module_name_module_system_download",
		mcp.WithDescription("Download the latest version of a module"),
		mcp.WithString("module_namespace",
			mcp.Description("Group's ID or slug"),
			mcp.Required(),
		),
		mcp.WithString("module_name",
			mcp.Description(""),
			mcp.Required(),
		),
		mcp.WithString("module_system",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadHandler)
}

func getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	module_namespace := request.GetString("module_namespace", "")
	module_name := request.GetString("module_name", "")
	module_system := request.GetString("module_system", "")

	return toResult(c.GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(ctx, module_namespace, module_name, module_system, authorizationHeader))
}

func registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(s *server.MCPServer) {
	tool := mcp.NewTool("pkgs_terraform_modules_v1_module_namespace_module_name_module_system",
		mcp.WithDescription("Get details about the latest version of a module"),
		mcp.WithString("module_namespace",
			mcp.Description("Group's ID or slug"),
			mcp.Required(),
		),
		mcp.WithString("module_name",
			mcp.Description(""),
			mcp.Required(),
		),
		mcp.WithString("module_system",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemHandler)
}

func getPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	module_namespace := request.GetString("module_namespace", "")
	module_name := request.GetString("module_name", "")
	module_system := request.GetString("module_system", "")

	return toResult(c.GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(ctx, module_namespace, module_name, module_system, authorizationHeader))
}
