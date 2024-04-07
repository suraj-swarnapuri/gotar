// package chord is responsible for logic about creating chords: a series of notes in a scale.
package chord

import (
	"github.com/suraj-swarnapuri/gotar/internal/note"
)

type BaseChord struct {
	chordMap map[note.Interval]note.Note
}

func (bc BaseChord) GetNote(interval note.Interval) note.Note {
	return bc.chordMap[interval]
}

func (bc BaseChord) ListNotes() []note.Note {
	notes := make([]note.Note, 0, len(bc.chordMap))
	for _, n := range bc.chordMap {
		notes = append(notes, n)
	}
	return notes
}

/*
  given a list of notes how to determine the chord
  1. sort the list in ascending order
  2. determine the interval between the notes
  3. based on interval categorize it
  4. find the root
  5. process extensions


*/
