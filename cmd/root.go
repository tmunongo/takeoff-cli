package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "takeoff",
	Short: "takeoff is a CLI tool for bootstrapping docker-based development environments",
	Long: `A fast a friendly CLI tool for bootstrapping your development projects
			with Docker. No need to understand docker-compose or docker, takeoff
			handles it all for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("takeoff ")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
