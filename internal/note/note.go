// Package note contains the Scalel and functions for the atomic unit of measurement of music, a note.
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

type Interval int

const (
	ZERO Interval = iota
	I
	II
	III
	IV
	V
	VI
	VII
)

func (i Interval) String() string {
	names := [...]string{
		"- ",
		"1 ",
		"2 ",
		"3 ",
		"4 ",
		"5 ",
		"6 ",
		"7 ",
	}

	return names[i]
}

func (n Note) String() string {
	names := [...]string{
		"- ",
		"C ",
		"C#",
		"D ",
		"D#",
		"E ",
		"F ",
		"F#",
		"G ",
		"G#",
		"A ",
		"A#",
		"B ",
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
