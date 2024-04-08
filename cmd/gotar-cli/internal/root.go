package internal

import (
	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/cmd/gotar-cli/internal/scale"
)

var rootCmd = &cobra.Command{
	Use:   "gotar-cli",
	Short: "Gotar-CLI view scales and chords",
	Long:  `Gotar-CLI is an in-depth visualization of chords and scales on the guitar`,
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.PersistentFlags().GetBool("version")
		if version {
			versionCmd.Run(cmd, args)
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "version output")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(scale.ScaleCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
