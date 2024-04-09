package chord

import (
	"github.com/spf13/cobra"
	chrd "github.com/suraj-swarnapuri/gotar/internal/chord"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	sc "github.com/suraj-swarnapuri/gotar/internal/scale"
)

var AugmentedChordCmd = &cobra.Command{
	Use:   "augmented",
	Short: "Print a visualization of a augmented scale of a given key",
	RunE:  displayAugmentedChord(),
}

func displayAugmentedChord() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)

		majorScale := sc.NewMajorScale(tonic)
		cd := chrd.NewAugmentedChord(majorScale)

		board := fretboard.NewFretboard()
		board.Display(cd, intervalOn)
		return nil
	}
}
