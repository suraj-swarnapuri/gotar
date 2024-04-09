package chord

import (
	"github.com/spf13/cobra"
	chrd "github.com/suraj-swarnapuri/gotar/internal/chord"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	sc "github.com/suraj-swarnapuri/gotar/internal/scale"
)

var DiminishedChordCmd = &cobra.Command{
	Use:   "diminished",
	Short: "Print a visualization of a diminished scale of a given key",
	RunE:  displayDiminishedChord(),
}

func displayDiminishedChord() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)

		majorscale := sc.NewMajorScale(tonic)
		cd := chrd.NewDiminishedChord(majorscale)

		board := fretboard.NewFretboard()
		board.Display(cd, intervalOn)
		return nil
	}
}
