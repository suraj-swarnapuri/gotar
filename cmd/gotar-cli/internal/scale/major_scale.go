package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var MajorScaleCmd = &cobra.Command{
	Use:   "ionian",
	Short: "Print a visualization of the Ionian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Ionian scale on the fretboard in a given key",
	RunE:  displayMajorScale(),
}

func displayMajorScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		majorScale := sc.NewMajorScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(majorScale, intervalOn)
		return nil
	}
}
