package chord

import (
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

// Chord represents a series of notes played together.
type MajorChord struct {
	notes []note.Note
	tonic note.Note
}

// NewMajorChord generates the notes of the major chord
func NewMajorChord(s scale.Scale) MajorChord {
	mc := MajorChord{}
	notes := s.Notes()
	mc.tonic = notes[0]
	mc.notes = append(mc.notes, notes[I])
	mc.notes = append(mc.notes, notes[III])
	mc.notes = append(mc.notes, notes[V])
	return mc
}

func (mc *MajorChord) Notes() []note.Note {
	return mc.notes
}

// Name returns the name of the major chord
func (mc *MajorChord) Name() string {
	return mc.tonic.String() + " major"
}