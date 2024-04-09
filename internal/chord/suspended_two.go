package chord

import (
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type Sus2Chord struct {
	BaseChord
}

func NewSus2Chord(sc scale.Scale) Sus2Chord {
	chordMap := make(map[note.Interval]note.Note)
	chordMap[note.I] = sc.GetNote(note.I)
	chordMap[note.III] = sc.GetNote(note.II)
	chordMap[note.V] = sc.GetNote(note.V)

	s2c := Sus2Chord{
		BaseChord{
			chordMap: chordMap,
		},
	}
	return s2c
}
