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
	mc := &MajorChord{}
	mc.chordMap[note.I] = sc.GetNote(note.I)
	mc.chordMap[note.III] = sc.GetNote(note.III)
	mc.chordMap[note.V] = sc.GetNote(note.V)

	return *mc
}

func (mc MajorChord) Name() string {
	return fmt.Sprintf("%s Major Chord", mc.chordMap[note.I].String())
}
