// Package scale contains the interface and implementations for scales, a series of notes.
package scale

import "github.com/suraj-swarnapuri/gotar/internal/note"

// Scale is the interface to generate a series of notes.
type Scale interface {
	// Generates a series of notes for the scale in a given key.
	GenerateNotes(n note.Note) []note.Note
	// Returns true if the note is in the scale.
	Contains(n note.Note) bool
}
