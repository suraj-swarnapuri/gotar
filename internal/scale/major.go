package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

type IonianScale struct {
	BaseScale
	tonic note.Note
}

func NewMajorScale(n note.Note) *IonianScale {
	ms := &IonianScale{}
	ms.tonic = n
	notes := make([]note.Note, 7)
	// Major scale is W-W-H-W-W-W-H

	notes[0] = n
	notes[1] = notes[0].StepUp().StepUp()
	notes[2] = notes[1].StepUp().StepUp()
	notes[3] = notes[2].StepUp()
	notes[4] = notes[3].StepUp().StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp().StepUp()
	ms.notes = notes
	return ms
}

func (ms *IonianScale) Name() string {
	return fmt.Sprintf("%s Ionian Scale", ms.tonic.String())
}
