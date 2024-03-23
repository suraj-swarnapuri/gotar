package scale

import "github.com/suraj-swarnapuri/gotar/internal/note"

type ChromaticScale struct {
	notes []note.Note
}

func NewChromaticScale() ChromaticScale{
	cs := ChromaticScale{}
	notes := make([]note.Note, 12)
	notes[0] = note.C
	for i := 1; i < 12; i++ {
		notes[i] = notes[i-1].StepUp()
	}
	cs.notes = notes
	return cs
}

func (cs *ChromaticScale) Notes() []note.Note {
	return cs.notes
}

func (cs *ChromaticScale) Name() string {
	return "chromatic"
}

