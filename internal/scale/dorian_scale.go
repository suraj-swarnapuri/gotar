package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// DorianScale represents the Dorian scale
type DorianScale struct {
	BaseScale
	tonic note.Note
}

// NewDorianScale generates the notes of the Dorian scale
func NewDorianScale(n note.Note) *DorianScale {
	ds := &DorianScale{}
	ds.tonic = n
	notes := make([]note.Note, 7)
	// Dorian scale is W-H-W-W-W-H-W
	notes[0] = n
	notes[1] = notes[0].StepUp().StepUp()
	notes[2] = notes[1].StepUp()
	notes[3] = notes[2].StepUp().StepUp()
	notes[4] = notes[3].StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp()
	ds.notes = notes
	return ds
}

// Name returns the name of the Dorian scale
func (ds *DorianScale) Name() string {
	return fmt.Sprintf("%s Dorian Scale", ds.tonic.String())
}
