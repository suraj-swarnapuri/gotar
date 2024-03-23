package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// LocrianScale represents the Locrian scale
type LocrianScale struct {
	notes []note.Note
	tonic note.Note
}

// GenerateNotes generates the notes of the Locrian scale
func NewLocrianScale(n note.Note) LocrianScale {
	ls := LocrianScale{}
	ls.tonic = n
	notes := make([]note.Note, 7)
	// Locrian scale is H-W-W-H-W-W-W
	notes[0] = n
	notes[1] = notes[0].StepUp()
	notes[2] = notes[1].StepUp().StepUp()
	notes[3] = notes[2].StepUp()
	notes[4] = notes[3].StepUp().StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp()
	ls.notes = notes
	return ls
}

// Name returns the name of the Locrian scale
func (ls *LocrianScale) Name() string {
	return fmt.Sprintf("%s locrian", ls.tonic.String())
}

// Notes returns the notes of the Locrian scale
func (ls *LocrianScale) Notes() []note.Note {
	return ls.notes
}

