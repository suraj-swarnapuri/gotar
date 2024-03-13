// Package note contains the model and functions for the atomic unit of measurement of music, a note.
package note

// Make a enum called Note that holds the 12 notes in the chromatic scale.
type Note int

const (
	Blank Note = iota
	C
	CSharp
	D
	DSharp
	E
	F
	FSharp
	G
	GSharp
	A
	ASharp
	B
)

func (n Note) String() string {
	names := [...]string{
		"-",
		"C",
		"C#",
		"D",
		"D#",
		"E",
		"F",
		"F#",
		"G",
		"G#",
		"A",
		"A#",
		"B",
	}

	return names[n]
}

// StepUp increments the note by a half step.
func (n Note) StepUp() Note {
	// rescue the note from going out of bounds and reset to 0
	if n == B {
		return C
	}

	return n + 1
}

// StepDown decrements the note by a half step.
func (n Note) StepDown() Note {
	// rescue the note from going out of bounds and reset to 11
	if n == C {
		return B
	}

	return n - 1
}
