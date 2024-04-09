package chord

import (
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type Sus4Chord struct {
	BaseChord
}

func NewSus4Chord(sc scale.Scale) Sus4Chord {
	chordMap := make(map[note.Interval]note.Note)
	chordMap[note.I] = sc.GetNote(note.I)
	chordMap[note.III] = sc.GetNote(note.IV)
	chordMap[note.V] = sc.GetNote(note.V)

	s4c := Sus4Chord{
		BaseChord{
			chordMap: chordMap,
		},
	}
	return s4c
}
