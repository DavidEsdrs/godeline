package position

type Position struct {
	Ln, Col int
	// index holds and absolute index, instead of relative indexes (that is what
	// Ln and Col should be)
	Index int
}

// receives 4 arguments represeting the start and the and of 2 texts segments and returns if they overlaps
func Overlap(p1Start, p1End, p2Start, p2End Position) bool {
	return true
}

func IsInSameLine(p1, p2 Position) bool {
	return p1.Ln == p2.Ln
}

// returns the column and the line of the given index in the given text
func GetPosition(text string, idx int) Position {
	var (
		currentCol  int = 1
		currentLn   int
		currentIdx  int
		currentChar rune
	)

	for currentIdx < idx {
		currentChar = rune(text[currentIdx])

		if currentChar == '\n' {
			currentLn++
			currentCol = 0
		} else {
			currentCol++
		}

		currentIdx++
	}

	return Position{currentLn, currentCol, idx}
}
