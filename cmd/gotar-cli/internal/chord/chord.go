package chord

import (
	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

type CommandRun func(cmd *cobra.Command, args []string) error

var ChordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Print a visualization of a chord on the fretboard.",
}

func init() {
	ChordCmd.PersistentFlags().BoolP("show-interval", "i", false, "displays intervals instead of notes")
	ChordCmd.PersistentFlags().StringP("tonic", "t", "", "the tonic of the scale, aka the first note of the scale.")
	ChordCmd.MarkPersistentFlagRequired("tonic")

	ChordCmd.AddCommand(MajorChordCmd)
	ChordCmd.AddCommand(MinorChordCmd)
	ChordCmd.AddCommand(AugmentedChordCmd)
	ChordCmd.AddCommand(DiminishedChordCmd)
}

func getFlags(cmd *cobra.Command) (note.Note, bool) {
	intervalOn, _ := cmd.Flags().GetBool("show-interval")
	root, _ := cmd.Flags().GetString("tonic")
	tonic := note.GetNote(root)
	return tonic, intervalOn
}
