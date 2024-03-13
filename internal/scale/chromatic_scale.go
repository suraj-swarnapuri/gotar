package scale

import "github.com/suraj-swarnapuri/gotar/internal/note"

type ChromaticScale struct {
	notes []note.Note
}

func (cs *ChromaticScale) GenerateNotes(n note.Note) []note.Note {
	notes := make([]note.Note, 12)
	notes[0] = note.C
	for i := 1; i < 12; i++ {
		notes[i] = notes[i-1].StepUp()
	}
	cs.notes = notes
	return notes
}

func (cs *ChromaticScale) Contains(n note.Note) bool {
	for _, note := range cs.notes {
		if note == n {
			return true
		}
	}
	return false
}
