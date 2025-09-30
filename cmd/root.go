package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var (
    port     int
    origin   string
    clearCache bool
)

var rootCmd = &cobra.Command{
    Use:   "caching-proxy",
    Short: "Caching proxy server",
    Long:  `Caching proxy server simulation to forward requests and caches responses.`,
    Run: func(cmd *cobra.Command, args []string) {
        if clearCache {
            // TODO :Implement clear cache
            return
        }

        if port == 0 || origin == "" {
            cmd.Help()
            os.Exit(1)
        }

        // Start  proxy server
        fmt.Printf("Starting proxy server on port %d for origin %s\n", port, origin)
    },
}

func init() {
    rootCmd.Flags().IntVarP(&port, "port", "p", 0, "Port for the proxy server")
    rootCmd.Flags().StringVarP(&origin, "origin", "o", "", "Origin server URL")
    rootCmd.Flags().BoolVarP(&clearCache, "clear-cache", "c", false, "Clear the cache")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}