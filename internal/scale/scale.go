// Package scale contains the interface and implementations for scales, a series of notes.
package scale

import "github.com/suraj-swarnapuri/gotar/internal/note"

// Scale is the interface to generate a series of notes.
type Scale interface {
	// Returns series of notes for the scale in a given key.
	Notes() []note.Note
	// Returns the name of the scale.
	Name() string
}
