package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// AeolianScale represents the Aeolian scale
type AeolianScale struct {
	BaseScale
	tonic note.Note
}

// GenerateNotes generates the notes of the Aeolian scale
func NewAeolianScale(n note.Note) *AeolianScale {
	as := &AeolianScale{}
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

// Name returns the name of the Aeolian scale
func (as *AeolianScale) Name() string {
	return fmt.Sprintf("%s Aeolian Scale", as.tonic.String())
}
