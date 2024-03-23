package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// AeolianScale represents the Aeolian scale
type AeolianScale struct {
	notes []note.Note
	tonic note.Note
}

// GenerateNotes generates the notes of the Aeolian scale
func  NewAeolianScale(n note.Note) AeolianScale {
	as := AeolianScale{}
	as.tonic = n
	notes := make([]note.Note, 7)
	// Aeolian scale is W-H-W-W-H-W-W
	notes[0] = n
	notes[1] = notes[0].StepUp().StepUp()
	notes[2] = notes[1].StepUp()
	notes[3] = notes[2].StepUp().StepUp()
	notes[4] = notes[3].StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp()
	as.notes = notes
	return as
}

func (as *AeolianScale) Notes() []note.Note {
	return as.notes
}

// Name returns the name of the Aeolian scale
func (as *AeolianScale) Name() string {
	return fmt.Sprintf("%s aeolian", as.tonic.String())
}