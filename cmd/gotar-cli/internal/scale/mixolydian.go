package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var MixolydianScaleCmd = &cobra.Command{
	Use:   "mixolydian",
	Short: "Print a visualization of the Mixolydian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Mixolydian scale on the fretboard in a given key",
	RunE:  displayMixolydianScale(),
}

func displayMixolydianScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		mixolydianScale := sc.NewMixolydianScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(mixolydianScale, intervalOn)
		return nil
	}
}
