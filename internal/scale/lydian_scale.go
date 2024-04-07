package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// LydianScale represents the Lydian scale
type LydianScale struct {
	BaseScale
	tonic note.Note
}

// GenerateNotes generates the notes of the Lydian scale
func NewLydianScale(n note.Note) LydianScale {
	ls := LydianScale{}
	ls.tonic = n
	notes := make([]note.Note, 7)
	// Lydian scale is W-W-W-H-W-W-H
	notes[0] = n
	notes[1] = notes[0].StepUp().StepUp()
	notes[2] = notes[1].StepUp().StepUp()
	notes[3] = notes[2].StepUp()
	notes[4] = notes[3].StepUp().StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp()
	ls.notes = notes
	return ls
}

// Name returns the name of the Lydian scale
func (ls *LydianScale) Name() string {
	return fmt.Sprintf("%s Lydian Scale", ls.tonic.String())
}
