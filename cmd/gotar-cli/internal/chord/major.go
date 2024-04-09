package chord

import (
	"github.com/spf13/cobra"
	chrd "github.com/suraj-swarnapuri/gotar/internal/chord"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	sc "github.com/suraj-swarnapuri/gotar/internal/scale"
)

var MajorChordCmd = &cobra.Command{
	Use:   "major",
	Short: "Print a visualization of a major scale of a given key",
	RunE:  displayMajorChord(),
}

func displayMajorChord() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)

		majorScale := sc.NewMajorScale(tonic)
		cd := chrd.NewMajorChord(majorScale)

		board := fretboard.NewFretboard()
		board.Display(cd, intervalOn)
		return nil
	}
}
