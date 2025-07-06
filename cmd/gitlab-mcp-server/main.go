package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/9506hqwy/gitlab-mcp-server/pkg/gitlab"
)

var version = "<version>"
var commit = "<commit>"

func fromArgument(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, gitlab.UrlKey{}, viper.GetString("url"))
	ctx = context.WithValue(ctx, gitlab.TokenKey{}, viper.GetString("token"))
	return ctx
}

var rootCmd = &cobra.Command{
	Use:     "gitlab-mcp-server",
	Short:   "GitLab MCP Server",
	Long:    "GitLab MCP Server",
	Version: fmt.Sprintf("%s\nCommit: %s", version, commit),
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewMCPServer(
			"GitLab MCP Server",
			"0.1.0",
			server.WithToolCapabilities(false),
		)

		gitlab.RegisterTools(s, viper.GetBool("readonly"))

		if err := server.ServeStdio(s, server.WithStdioContextFunc(fromArgument)); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Fatalf("Server error: %v", err)
			}
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("url", "https://127.0.0.1", "GitLab server URL.")
	rootCmd.PersistentFlags().String("token", "", "GitLab server token.")
	rootCmd.PersistentFlags().Bool("readonly", true, "HTTP GET method only.")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("readonly", rootCmd.PersistentFlags().Lookup("readonly"))
}

func initConfig() {
	viper.SetEnvPrefix("gitlab")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
