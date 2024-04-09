package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var AeolianScaleCmd = &cobra.Command{
	Use:   "aeolian",
	Short: "Print a visualization of the Aeolian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Aeolian scale on the fretboard in a given key",
	RunE:  displayAeolianScale(),
}

func displayAeolianScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		aeolianScale := sc.NewAeolianScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(aeolianScale, intervalOn)
		return nil
	}
}
