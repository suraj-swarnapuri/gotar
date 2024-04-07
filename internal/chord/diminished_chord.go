package chord

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type NewDiminishedChord struct {
	BaseChord
}

func NewNewDiminishedChord(sc scale.Scale) NewDiminishedChord {
	mc := &NewDiminishedChord{}
	mc.chordMap[note.I] = sc.GetNote(note.I)
	mc.chordMap[note.III] = sc.GetNote(note.III).StepDown()
	mc.chordMap[note.V] = sc.GetNote(note.V).StepDown()

	return *mc
}

func (mc NewDiminishedChord) Name() string {
	return fmt.Sprintf("%s Diminished Chord", mc.chordMap[note.I].String())
}
