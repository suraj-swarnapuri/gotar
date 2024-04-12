// package chord is responsible for logic about creating chords: a series of notes in a scale.
package chord

import (
	"github.com/suraj-swarnapuri/gotar/internal/note"
	"github.com/suraj-swarnapuri/gotar/internal/scale"
)

/*
  given a list of notes how to determine the chord
  1. sort the list in ascending order
  2. determine the interval between the notes
  3. based on interval categorize it
  4. find the root
  5. process extensions
*/

/*
Chord:
 - tonic
 - major scale
 - Major()chord
 - Minor()chord
 - Augmented()chord
 - Diminished()chord
 - SuspendTwo()chord
 - SuspendFour()chord
 -

*/

type Chord struct {
	ms      scale.IonianScale
	noteMap map[note.Interval]note.Note
	name    string
}

func NewChord(tonic note.Note) Chord {
	return Chord{
		ms:      *scale.NewMajorScale(tonic),
		noteMap: make(map[note.Interval]note.Note),
		name:    tonic.String(),
	}
}

func (c Chord) Name() string {
	return c.name + " Chord"
}

func (c Chord) GetNote(interval note.Interval) note.Note {
	return c.noteMap[interval]
}

func (c Chord) ListNotes() []note.Note {
	notes := make([]note.Note, 0, 12)
	for k, v := range c.noteMap {
		notes[k-1] = v
	}

	return notes
}

func (c Chord) Major() Chord {
	c.noteMap[note.I] = c.ms.GetNote(note.I)
	c.noteMap[note.III] = c.ms.GetNote(note.III)
	c.noteMap[note.V] = c.ms.GetNote(note.V)
	c.name += " Major"
	return c
}

func (c Chord) Minor() Chord {
	c.noteMap[note.I] = c.ms.GetNote(note.I)
	c.noteMap[note.III] = c.ms.GetNote(note.III).StepDown()
	c.noteMap[note.V] = c.ms.GetNote(note.V)
	c.name += " Minor"
	return c
}

func (c Chord) Diminished() Chord {
	c.noteMap[note.I] = c.ms.GetNote(note.I)
	c.noteMap[note.III] = c.ms.GetNote(note.III).StepDown()
	c.noteMap[note.V] = c.ms.GetNote(note.V).StepDown()
	c.name += " Diminished"
	return c
}

func (c Chord) Augmented() Chord {
	c.noteMap[note.I] = c.ms.GetNote(note.I)
	c.noteMap[note.III] = c.ms.GetNote(note.III)
	c.noteMap[note.V] = c.ms.GetNote(note.V).StepUp()
	c.name += " Augmented"
	return c
}

func (c Chord) SuspendedTwo() Chord {
	c.noteMap[note.I] = c.ms.GetNote(note.I)
	c.noteMap[note.II] = c.ms.GetNote(note.II)
	c.noteMap[note.V] = c.ms.GetNote(note.V)
	c.name += " Sus2"
	return c
}

func (c Chord) SuspendedFour() Chord {
	c.noteMap[note.I] = c.ms.GetNote(note.I)
	c.noteMap[note.IV] = c.ms.GetNote(note.IV)
	c.noteMap[note.V] = c.ms.GetNote(note.V)
	c.name += " Sus4"
	return c
}

func (c Chord) Seven() Chord {
	c.name += "7"
	// min3 and min5 double flag 7
	if c.ms.GetNote(note.III).StepDown() == c.GetNote(note.III) && c.ms.GetNote(note.V).StepDown() == c.GetNote(note.V) {
		c.noteMap[note.VII] = c.ms.GetNote(note.VII).StepDown().StepDown()
		return c
	}
	// min3 min 7
	if c.ms.GetNote(note.III).StepDown() == c.GetNote(note.III) {
		c.noteMap[note.VII] = c.ms.GetNote(note.VII).StepDown()
		return c
	}
	// maj3 maj7
	c.noteMap[note.VII] = c.ms.GetNote(note.VII)
	return c
}
