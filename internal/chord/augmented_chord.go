package chord

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type AugmentedChord struct {
	BaseChord
}

func NewAugmentedChord(sc scale.Scale) AugmentedChord {
	ac := &AugmentedChord{}
	ac.chordMap[note.I] = sc.GetNote(note.I)
	ac.chordMap[note.III] = sc.GetNote(note.III)
	ac.chordMap[note.V] = sc.GetNote(note.V).StepUp()

	return *ac
}

func (ac AugmentedChord) Name() string {
	return fmt.Sprintf("%s Augmented Chord", ac.chordMap[note.I].String())
}
