package scale

import (
	"errors"

	sc "github.com/suraj-swarnapuri/gotar/internal/scale"

	"github.com/spf13/cobra"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

var PhyrgianScaleCmd = &cobra.Command{
	Use:   "phrygian",
	Short: "Print a visualization of the Phrygian scale on the fretboard in a given key.",
	Long:  "Display a visualization of the Phrygian scale on the fretboard in a given key",
	RunE:  displayPhrygianScale(),
}

func displayPhrygianScale() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)
		if tonic == note.Blank {
			return errors.New("the tonic / root note needs to be provided")
		}

		phrygianScale := sc.NewPhyrgianScale(tonic)

		board := fretboard.NewFretboard()
		board.Display(phrygianScale, intervalOn)
		return nil
	}
}
