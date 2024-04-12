package main

import (
	"github.com/suraj-swarnapuri/gotar/internal/chord"
	"github.com/suraj-swarnapuri/gotar/internal/fretboard"
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

func main() {
	// ...
	board := fretboard.NewFretboard()

	chrd := chord.NewChord(note.C)

	board.Display(chrd.Major().Seven(), true) // base default board

}
