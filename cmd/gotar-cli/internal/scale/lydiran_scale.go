package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var LydianScaleCmd = &cobra.Command{
	Use:   "lydian",
	Short: "Print a visualization of the Lydian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Lydian scale on the fretboard in a given key",
	RunE:  displayLydianScale(),
}

func displayLydianScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		lydianScale := sc.NewLydianScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(lydianScale, intervalOn)
		return nil
	}
}
