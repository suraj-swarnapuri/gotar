package chord

import (
	"github.com/spf13/cobra"
	chrd "github.com/suraj-swarnapuri/gotar/internal/chord"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	sc "github.com/suraj-swarnapuri/gotar/internal/scale"
)

var MinorChordCmd = &cobra.Command{
	Use:   "minor",
	Short: "Print a visualization of a minor scale of a given key",
	RunE:  displayMinorChord(),
}

func displayMinorChord() CommandRun {
	return func(cmd *cobra.Command, args []string) error {
		tonic, intervalOn := getFlags(cmd)

		minorScale := sc.NewMajorScale(tonic)
		cd := chrd.NewMinorChord(minorScale)

		board := fretboard.NewFretboard()
		board.Display(cd, intervalOn)
		return nil
	}
}
