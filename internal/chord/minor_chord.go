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
	chordMap := make(map[note.Interval]note.Note)
	chordMap[note.I] = s.GetNote(note.I)
	chordMap[note.III] = s.GetNote(note.III).StepDown()
	chordMap[note.V] = s.GetNote(note.V)

	mc := MinorChord{
		BaseChord{
			chordMap: chordMap,
		},
	}

	return mc
}

func (mc MinorChord) Name() string {
	return fmt.Sprintf("%s Minor Chord", mc.chordMap[note.I].String())
}
