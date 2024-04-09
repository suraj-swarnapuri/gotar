package chord

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type DiminishedChord struct {
	BaseChord
}

func NewDiminishedChord(sc scale.Scale) DiminishedChord {
	chordMap := make(map[note.Interval]note.Note)
	chordMap[note.I] = sc.GetNote(note.I)
	chordMap[note.III] = sc.GetNote(note.III).StepDown()
	chordMap[note.V] = sc.GetNote(note.V).StepDown()

	mc := DiminishedChord{
		BaseChord{
			chordMap: chordMap,
		},
	}
	return mc
}

func (mc DiminishedChord) Name() string {
	return fmt.Sprintf("%s Diminished Chord", mc.chordMap[note.I].String())
}
