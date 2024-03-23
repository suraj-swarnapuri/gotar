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


	scale := scale.NewMajorScale(note.C)

	//chord := chord.NewMinorChord(&scale)

	board.Display(&scale)
}
