package chord

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type MinorChord struct {
	BaseChord
}

func NewMinorChord(s scale.Scale) MinorChord {
	mc := &MinorChord{}
	mc.chordMap[note.I] = s.GetNote(note.I)
	mc.chordMap[note.III] = s.GetNote(note.III).StepDown()
	mc.chordMap[note.V] = s.GetNote(note.V)

	return *mc
}

func (mc MinorChord) Name() string {
	return fmt.Sprintf("%s Minor Chord", mc.chordMap[note.I].String())
}
