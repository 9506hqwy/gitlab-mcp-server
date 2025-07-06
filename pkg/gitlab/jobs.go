package gitlab

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetJobsIdArtifacts(s *server.MCPServer) {
	tool := mcp.NewTool("jobs_id_artifacts",
		mcp.WithDescription("Download the artifacts file for job"),
		mcp.WithNumber("id",
			mcp.Description("Job's ID"),
			mcp.Required(),
		),
		mcp.WithString("token",
			mcp.Description("Job's authentication token"),
		),
		mcp.WithBoolean("direct_download",
			mcp.Description("Perform direct download from remote storage instead of proxying artifacts (default: false)"),
		),
	)

	s.AddTool(tool, getJobsIdArtifactsHandler)
}

func getJobsIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := int32(request.GetInt("id", math.MinInt))
	params := parseGetJobsIdArtifacts(request)
	return toResult(c.GetApiV4JobsIdArtifacts(ctx, id, &params, authorizationHeader))
}

func parseGetJobsIdArtifacts(request mcp.CallToolRequest) client.GetApiV4JobsIdArtifactsParams {
	params := client.GetApiV4JobsIdArtifactsParams{}

	token := request.GetString("token", "")
	if token != "" {

		params.Token = &token
	}

	direct_download := request.GetBool("direct_download", false)
	params.DirectDownload = &direct_download

	return params
}
