// Pacakge fretboard holds the logic for displaying notes on a standard guitar fretboard.
package fretboard

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type Fretboard struct {
	board map[string][]note.Note
}

func InitBoard() Fretboard {
	fb := Fretboard{}
	fb.board = make(map[string][]note.Note, 6)
	fb.board["E"] = addString(note.E)
	fb.board["A"] = addString(note.A)
	fb.board["D"] = addString(note.D)
	fb.board["G"] = addString(note.G)
	fb.board["B"] = addString(note.B)
	fb.board["e"] = addString(note.E)

	return fb
}

func (f Fretboard) Display(s scale.Scale) {
	displayString(f.board["E"], s)
	fmt.Println()
	displayString(f.board["A"], s)
	fmt.Println()
	displayString(f.board["D"], s)
	fmt.Println("                   o             o             o             o                    :   ")
	displayString(f.board["G"], s)
	fmt.Println()
	displayString(f.board["B"], s)
	fmt.Println()
	displayString(f.board["e"], s)
}

func addString(n note.Note) []note.Note {
	guitarString := make([]note.Note, 12)
	for i := 0; i < 12; i++ {
		guitarString[i] = n
		n = n.StepUp()
	}
	return guitarString
}

func displayString(notes []note.Note, s scale.Scale) {
	fmt.Print(notes[0])
	fmt.Print(" |--")
	for i, n := range notes {
		n = n.StepUp()
		var tmpN note.Note
		if s.Contains(n) {
			tmpN = n
		} else {
			tmpN = note.Blank
		}

		if i == len(notes)-1 {
			fmt.Printf("%s--|", tmpN.String())
			break
		}

		if len(tmpN.String()) == 2 {
			fmt.Printf("%s--|--", tmpN.String())
		} else {
			fmt.Printf("%s---|--", tmpN.String())
		}

	}
	fmt.Println()
}
