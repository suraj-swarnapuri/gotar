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
	chordMap := make(map[note.Interval]note.Note)
	chordMap[note.I] = sc.GetNote(note.I)
	chordMap[note.III] = sc.GetNote(note.III)
	chordMap[note.V] = sc.GetNote(note.V).StepUp()

	ac := AugmentedChord{
		BaseChord{
			chordMap: chordMap,
		},
	}
	return ac
}

func (ac AugmentedChord) Name() string {
	return fmt.Sprintf("%s Augmented Chord", ac.chordMap[note.I].String())
}
