package chord

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type MajorChord struct {
	BaseChord
}

func NewMajorChord(sc scale.Scale) MajorChord {
	chordMap := make(map[note.Interval]note.Note)
	chordMap[note.I] = sc.GetNote(note.I)
	chordMap[note.III] = sc.GetNote(note.III)
	chordMap[note.V] = sc.GetNote(note.V)

	mc := MajorChord{
		BaseChord{
			chordMap: chordMap,
		},
	}

	return mc
}

func (mc MajorChord) Name() string {
	return fmt.Sprintf("%s Major Chord", mc.chordMap[note.I].String())
}
