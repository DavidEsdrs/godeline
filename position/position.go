package position

type Position struct {
	Ln, Col int
	// index holds and absolute index, instead of relative indexes (that is what
	// Ln and Col should be)
	Index int
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
