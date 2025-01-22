package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ver "github.com/waynezhang/tskks/internal/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ver.VersionString())
	},
}
