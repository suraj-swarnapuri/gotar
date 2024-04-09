package main

import (
	"github.com/suraj-swarnapuri/gotar/internal/chord"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

func main() {
	// ...
	board := fretboard.NewFretboard()

	sc := scale.NewMajorScale(note.C)
	chrd := chord.NewDiminishedChord(sc)

	board.Display(chrd, false) // base default board

}
