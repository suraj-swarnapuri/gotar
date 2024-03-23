package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

type MajorScale struct {
	notes []note.Note
	tonic note.Note
}

func NewMajorScale(n note.Note) MajorScale {
	ms := MajorScale{}
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

func (ms *MajorScale) Name() string {
	return fmt.Sprintf("%s major", ms.tonic.String())
}

func (ms *MajorScale) Notes() []note.Note {
	return ms.notes
}
