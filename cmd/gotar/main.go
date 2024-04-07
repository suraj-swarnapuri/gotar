package main

import (
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
)

func main() {
	// ...
	board := fretboard.NewFretboard()

	//sc := scale.NewMajorScale(note.C)

	board.Display(nil, false) // base default board

}
