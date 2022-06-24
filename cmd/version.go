package cmd

import (
	"fmt"

	"github.com/aaqaishtyaq/git-link/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of rouster",
	Long: `Returns version number of rouster
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.BuildVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
