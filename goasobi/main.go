package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/wsuzume/goasobi/pkg"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "hello runs the goasobi HTTP server",
	}

	rootCmd.AddCommand(newEchoCmd())
	rootCmd.AddCommand(newProxyCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func newEchoCmd() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "echo",
		Short: "echo runs the hello/echo HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := fmt.Sprintf(":%d", port)
			r := pkg.NewServer()

			log.Printf("Server running at http://localhost:%d/\n", port)
			return r.Run(addr)
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 8080, "port to listen on")

	return cmd
}

func newProxyCmd() *cobra.Command {
	var port int
	var target string

	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "proxy reverse-proxies all HTTP traffic to the target server",
		RunE: func(cmd *cobra.Command, args []string) error {
			targetURL, err := url.Parse(target)
			if err != nil {
				return fmt.Errorf("invalid target URL %q: %w", target, err)
			}

			addr := fmt.Sprintf(":%d", port)
			r := pkg.NewProxyServer(targetURL)

			log.Printf("Proxy running at http://localhost:%d/ -> %s\n", port, targetURL)
			return r.Run(addr)
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 8080, "port to listen on")
	cmd.Flags().StringVarP(&target, "target", "t", "http://scaler:8080", "URL to proxy all requests to")

	return cmd
}
