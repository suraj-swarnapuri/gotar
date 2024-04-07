package scale

import (
	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

var ScaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Print a visualization of a scale on the fretboard.",
	Long:  "Display a visualization of a particular scale on the fretboard.",
	Run:   displayScale,
}

func init() {
	ScaleCmd.Flags().BoolP("show-interval", "i", false, "displays intervals instead of notes")
	ScaleCmd.Flags().StringP("tonic", "t", "", "the tonic of the scale, aka the first note of the scale.")
	ScaleCmd.Flags().StringP("mode", "m", "", "picks which mode to use eg")
}

func displayScale(cmd *cobra.Command, args []string) {
	board := fretboard.NewFretboard()
	intervalOn, _ := cmd.Flags().GetBool("show-interval")
	tonic, _ := cmd.Flags().GetString("tonic")
	mode, _ := cmd.Flags().GetString("mode")

	scale := getScale(tonic, mode)

	board.Display(scale, intervalOn)
}

func getScale(tonic, mode string) scale.Scale {
	rootNote := note.GetNote(tonic)
	return scale.GetScale(rootNote, mode)
}
