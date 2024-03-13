package scale

import "github.com/suraj-swarnapuri/gotar/internal/note"

type MajorScale struct {
	notes []note.Note
}

func (ms *MajorScale) GenerateNotes(n note.Note) []note.Note {
	notes := make([]note.Note, 7)
	notes[0] = n
	notes[1] = notes[0].StepUp().StepUp()
	notes[2] = notes[1].StepUp().StepUp()
	notes[3] = notes[2].StepUp()
	notes[4] = notes[3].StepUp().StepUp()
	notes[5] = notes[4].StepUp().StepUp()
	notes[6] = notes[5].StepUp().StepUp()
	ms.notes = notes
	return notes
}

func (ms *MajorScale) Contains(n note.Note) bool {
	for _, note := range ms.notes {
		if note == n {
			return true
		}
	}
	return false
}
