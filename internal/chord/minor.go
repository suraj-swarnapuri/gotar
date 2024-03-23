package chord

import (
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

type MinorChord struct {
	notes []note.Note
	tonic note.Note
}

func NewMinorChord(s scale.Scale) MinorChord {
	mc := MinorChord{}
	notes := s.Notes()
	mc.tonic = notes[0]
	mc.notes = append(mc.notes, notes[I])
	mc.notes = append(mc.notes, notes[III].StepDown())
	mc.notes = append(mc.notes, notes[V])
	return mc
}

func (mc *MinorChord) Notes() []note.Note {
	return mc.notes
}

func (mc *MinorChord) Name() string {
	return mc.tonic.String() + " minor"
}