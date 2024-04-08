// Package scale contains the interface and implementations for scales, a series of notes.
package scale

import (
	"strings"

	"github.com/suraj-swarnapuri/gotar/internal/note"
)

// Scale is the interface to generate a series of notes.
type Scale interface {
	// Returns series of notes for the scale in a given key.
	ListNotes() []note.Note
	// Returns the specific note in the scale
	GetNote(num note.Interval) note.Note
	// Returns the name of the scale.
	Name() string
}

type BaseScale struct {
	notes []note.Note
}

func (bs BaseScale) ListNotes() []note.Note {
	return bs.notes
}

func (bs BaseScale) GetNote(num note.Interval) note.Note {
	return bs.notes[num]
}

func GetScale(tonic note.Note, mode string) Scale {
	if tonic == note.Blank {
		return nil
	}
	mode = strings.ToLower(mode)
	switch mode {
	case "ionian":
		return NewMajorScale(tonic)
	case "dorian":
		return NewDorianScale(tonic)
	case "locrian":
		return NewLocrianScale(tonic)
	case "lydian":
		return NewLydianScale(tonic)
	case "mixolydian":
		return NewMixolydianScale(tonic)
	case "aeolian":
		return NewAeolianScale(tonic)
	case "phrygian":
		return NewPhyrgianScale(tonic)
	default:
		return nil
	}
}
