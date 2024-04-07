// Pacakge fretboard holds the logic for displaying notes on a standard guitar fretboard.
package fretboard

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type Fretboard struct {
	Board [][]note.Note
}

func NewFretboard() Fretboard {
	// set up a standard tuning fretboard
	fb := &Fretboard{}
	fb.Board = make([][]note.Note, 6)
	for i := range fb.Board {
		fb.Board[i] = make([]note.Note, 12)
	}

	openStrings := []note.Note{note.E, note.A, note.D, note.G, note.B, note.E}
	for i, n := range openStrings {
		currNote := n
		for j := 0; j < 12; j++ {
			fb.Board[i][j] = currNote
			currNote = currNote.StepUp()
		}
	}

	return *fb
}

func (fb Fretboard) print(noteMap map[note.Note]note.Interval, intervalOn bool) {
	for i := 0; i <= 12; i++ {
		if i >= 10 {
			fmt.Printf("%v| ", i)
			continue
		}
		fmt.Printf("%v | ", i)
	}
	fmt.Println("\n--------------------------------------------------")
	for _, ns := range fb.Board {
		for _, n := range ns {
			if !intervalOn {
				fmt.Print(n.String() + "| ")
			} else {
				fmt.Print(noteMap[n].String() + "| ")
			}

		}
		if !intervalOn {
			fmt.Print(ns[0].String() + "\n")
		} else {
			fmt.Print(noteMap[ns[0]].String() + "\n")
		}

	}
}

func (fb Fretboard) Display(sc scale.Scale, intervalOn bool) {
	if sc == nil {
		fb.print(nil, false)
		return
	}

	fmt.Println(sc.Name())

	notes := sc.ListNotes()
	noteMap := make(map[note.Note]note.Interval, len(notes))
	interval := note.I
	for _, n := range notes {
		noteMap[n] = interval
		interval++
	}
	noteMap[note.Blank] = note.ZERO

	for i := range fb.Board {
		for j := range fb.Board[i] {
			if _, ok := noteMap[fb.Board[i][j]]; !ok {
				fb.Board[i][j] = note.Blank
			}
		}
	}
	fb.print(noteMap, intervalOn)
}
