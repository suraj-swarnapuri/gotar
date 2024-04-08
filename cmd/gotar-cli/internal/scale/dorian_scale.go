package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var DorianScaleCmd = &cobra.Command{
	Use:   "dorian",
	Short: "Print a visualization of the dorian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Dorian scale on the fretboard in a given key",
	RunE:  displayDorianScale(),
}

func displayDorianScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		dorianScale := sc.NewDorianScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(dorianScale, intervalOn)
		return nil
	}
}
