package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// PhrygianScale represents the Phrygian scale
type PhrygianScale struct {
	BaseScale
	tonic note.Note
}

// GenerateNotes generates the notes of the Phrygian scale
func NewPhyrgianScale(n note.Note) PhrygianScale {
	ps := PhrygianScale{}
	ps.tonic = n
	notes := make([]note.Note, 7)
	// Phrygian scale is H-W-W-W-H-W-W
	notes[0] = n
	notes[1] = notes[0].StepUp()
	notes[2] = notes[1].StepUp().StepUp()
	notes[3] = notes[2].StepUp()
	notes[4] = notes[3].StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp()
	ps.notes = notes
	return ps
}

// Name returns the name of the Phrygian scale
func (ps *PhrygianScale) Name() string {
	return fmt.Sprintf("%s Phrygian Scale", ps.tonic.String())
}
