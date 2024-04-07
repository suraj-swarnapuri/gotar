package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gotar-CLI",
	Long:  `All software has versions. This is Gotar's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gotar-CLI: v0.0.1")
	},
}
