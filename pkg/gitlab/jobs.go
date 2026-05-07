package gitlab

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

type GetJobsIdArtifactsRequest struct {
	Id     int32                                 `json:"id" jsonschema:"description=Job's ID"`
	Params *client.GetApiV4JobsIdArtifactsParams `json:"params,omitempty"`
}

func registerGetJobsIdArtifacts(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GetJobsIdArtifactsRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("get_jobs_id_artifacts",
		mcp.WithDescription("Download the artifacts file for job"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(getJobsIdArtifactsHandler))
}

func getJobsIdArtifactsHandler(ctx context.Context, request mcp.CallToolRequest, req GetJobsIdArtifactsRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GetApiV4JobsIdArtifacts(ctx, req.Id, req.Params, authorizationHeader))
}
