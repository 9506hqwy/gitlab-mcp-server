# GitLab MCP Server

This repository provides model context protocol server for GitLab server.

## API Version

* v18.1

## Build

Build binary.

```sh
go build -o bin/gitlab-mcp-server ./cmd/gitlab-mcp-server/main.go
```

Or build container image.

```sh
docker build -t gitlab-mcp-server .
```

Add `Z` option at bind mount operation in *Dockerfile* if using podman with SELinux.

## Usage

Run application.

Specify token.

```sh
$ ./bin/gitlab-mcp-server -h
GitLab MCP Server

Usage:
  gitlab-mcp-server [flags]

Flags:
  -h, --help           help for gitlab-mcp-server
      --readonly       HTTP GET method only. (default true)
      --token string   GitLab server token.
      --url string     GitLab server URL. (default "https://127.0.0.1")
  -v, --version        version for gitlab-mcp-server
```

Set environment variable instead of arguments.

| Argument   | Environment Variable |
| :--------- | :------------------- |
| --url      | GITLAB_URL           |
| --token    | GITLAB_TOKEN         |
| --readonly | GITLAB_READONLY      |

Or run container.

```sh
docker run --rm -i -e GITLAB_URL=<URL> -e GITLAB_TOKEN=<TOKEN> gitlab-mcp-server
```

### Usage with VS code

Add `gitlab-mcp-server` binary to `PATH` environment variable and configure VS code.

```json
{
    "servers": {
        "gitlab": {
            "type": "stdio",
            "command": "gitlab-mcp-server",
            "env": {
                "GITLAB_URL": "${input:gitlab_url}",
                "GITLAB_TOKEN": "${input:gitlab_token}",
            }
        }
    },
    "inputs": [
        {
            "type": "promptString",
            "id": "gitlab_url",
            "description": "GitLab Server URL",
            "password": false
        },
        {
            "type": "promptString",
            "id": "gitlab_token",
            "description": "GitLab Server TOKEN",
            "password": true
        }
    ]
}
```

## Tools

TODO

## Testing

TODO

## Notes

* Check off unused tools at [tool icon] in GitHub Copilot Chat panel bacause vscode limits max 128 tools.
