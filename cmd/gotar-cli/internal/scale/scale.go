package scale

import (
	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

type CommandRun func(cmd *cobra.Command, args []string) error

var Tonic note.Note

var ScaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Print a visualization of a scale on the fretboard.",
	Long:  "Display a visualization of a particular scale on the fretboard.",
}

var fretboardCmd = &cobra.Command{
	Use:   "fretboard",
	Short: "Prints a vizualization of the 12 frets of the guitar.",
	Long:  "Prints 12 frets of the guitar with all the notes.",
	RunE:  displayScale(),
}

func init() {
	ScaleCmd.PersistentFlags().BoolP("show-interval", "i", false, "displays intervals instead of notes")
	ScaleCmd.PersistentFlags().StringP("tonic", "t", "", "the tonic of the scale, aka the first note of the scale.")

	ScaleCmd.AddCommand(fretboardCmd)
	ScaleCmd.AddCommand(MajorScaleCmd)
	ScaleCmd.AddCommand(AeolianScaleCmd)
	ScaleCmd.AddCommand(LydianScaleCmd)
	ScaleCmd.AddCommand(LocrianScaleCmd)
	ScaleCmd.AddCommand(DorianScaleCmd)
	ScaleCmd.AddCommand(PhyrgianScaleCmd)
	ScaleCmd.AddCommand(MixolydianScaleCmd)
}

func getFlags(cmd *cobra.Command) (note.Note, bool) {
	intervalOn, _ := cmd.Flags().GetBool("show-interval")
	root, _ := cmd.Flags().GetString("tonic")
	tonic := note.GetNote(root)
	return tonic, intervalOn
}

func displayScale() CommandRun {
	return func(cms *cobra.Command, args []string) error {
		board := fretboard.NewFretboard()
		board.Display(nil, false)
		return nil
	}
}
