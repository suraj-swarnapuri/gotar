package main

import (
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

func main() {
	// ...
	board := fretboard.InitBoard()
	// scale := &scale.ChromaticScale{}
	// scale.GenerateNotes(note.Blank)

	scale := &scale.MajorScale{}
	scale.GenerateNotes(note.C)

	board.Display(scale)
}
