package gitlab

import (
	"context"
	"fmt"
	"io"
	"net/http"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
	"github.com/mark3labs/mcp-go/mcp"
)

type UrlKey struct{}
type TokenKey struct{}

func authorizationHeader(ctx context.Context, req *http.Request) error {
	return bearerAuth(ctx, req)
}

func bearerAuth(ctx context.Context, req *http.Request) error {
	token, ok := ctx.Value(TokenKey{}).(string)
	if !ok || token == "" {
		return fmt.Errorf("missing token")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func newClient(ctx context.Context) (*client.ClientWithResponses, error) {
	hc := http.Client{}

	url, ok := ctx.Value(UrlKey{}).(string)
	if !ok || url == "" {
		return nil, fmt.Errorf("missing url")
	}

	return client.NewClientWithResponses(url, client.WithHTTPClient(&hc))
}

func toResult(response *http.Response, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if response.StatusCode < http.StatusOK || http.StatusMultipleChoices <= response.StatusCode {
		return mcp.NewToolResultError(fmt.Sprintf("%s: %s", response.Status, string(body))), nil
	}

	return mcp.NewToolResultText(string(body)), nil
}
