package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wsuzume/goasobi/pkg"
)

func main() {
	var port int

	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "hello runs the goasobi HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := fmt.Sprintf(":%d", port)
			r := pkg.NewServer()

			log.Printf("Server running at http://localhost:%d/\n", port)
			return r.Run(addr)
		},
	}

	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "port to listen on")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
