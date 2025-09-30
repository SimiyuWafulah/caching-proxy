package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "caching-proxy",
    Short: "A caching proxy server",
    Long:  `A caching proxy server that forwards requests and caches responses.`,
    Run: func(cmd *cobra.Command, args []string) {
        
    },
}

func Execute() {
    rootCmd.Execute()
}