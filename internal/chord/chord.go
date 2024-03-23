// Package chord contains the model and functions for chords, a series of notes played together.
package chord

import "github.com/suraj-swarnapuri/gotar/internal/note"

type Numeral int
const (
	I Numeral = iota
	II
	III
	IV
	V
	VI
	VII
)
type Chord interface {
	Name() string
	Notes() []note.Note
}
