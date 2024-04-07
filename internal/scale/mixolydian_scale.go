package scale

import (
	"fmt"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// MixolydianScale represents the Mixolydian scale
type MixolydianScale struct {
	BaseScale
	tonic note.Note
}

// GenerateNotes generates the notes of the Mixolydian scale
func NewMixolydianScale(n note.Note) *MixolydianScale {
	ms := &MixolydianScale{}
	ms.tonic = n
	notes := make([]note.Note, 7)
	// Mixolydian scale is W-W-H-W-W-H-W
	notes[0] = n
	notes[1] = notes[0].StepUp().StepUp()
	notes[2] = notes[1].StepUp()
	notes[3] = notes[2].StepUp().StepUp()
	notes[4] = notes[3].StepUp().StepUp()
	notes[5] = notes[4].StepUp()
	notes[6] = notes[5].StepUp()
	ms.notes = notes
	return ms
}

// Name returns the name of the Mixolydian scale
func (ms *MixolydianScale) Name() string {
	return fmt.Sprintf("%s Mixolydian Scale", ms.tonic.String())
}
