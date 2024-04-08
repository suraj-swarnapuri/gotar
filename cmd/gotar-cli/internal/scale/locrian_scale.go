package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var LocrianScaleCmd = &cobra.Command{
	Use:   "locrian",
	Short: "Print a visualization of the Locrian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Locrian scale on the fretboard in a given key",
	RunE:  displayLocrianScale(),
}

func displayLocrianScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		locrianScale := sc.NewLocrianScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(locrianScale, intervalOn)
		return nil
	}
}
