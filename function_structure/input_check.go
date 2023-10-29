package hangman

// Function to check whether the given letter is in ToFind
func (game *HangManData) LetterInWord(oneRune rune) {
	var place []int
	for index, letters := range game.ToFind { // Transforms letters into lower case if they are not already lower case
		if oneRune >= 'A' && oneRune <= 'Z' {
			oneRune = oneRune + 32
		}
		if letters == oneRune || letters == oneRune-32 {
			place = append(place, index) // Saves the index(es) of the position where the letter was found
		}
	}
	if len(place) != 0 { // If any letters have been found then replace the corresponding slots with the letters
		game.LastFail = false
		for _, index := range place {
			game.Word[index] = oneRune
		}
	} else { // If the letter is not found, an attempt is lost
		game.LastFail = true
		game.Attempts--
		game.HangmanPositions++
	}
}

// Function to check whether the given word is ToFind (return true if this is the case)
func (game *HangManData) IsThisTheWord(word string) bool {
	var oneRune rune
	if len(word) != len(game.ToFind) { // Check that words are the same size
		return false
	} else {
		wordRune := []rune(word)
		toFindRune := []rune(game.ToFind)
		for index, runes := range wordRune {
			oneRune = runes
			if runes >= 'A' && runes <= 'Z' { // Puts letters in lower case for easier comparison
				oneRune = oneRune + 32
			}
			if oneRune != toFindRune[index] && oneRune != toFindRune[index]+32 {
				return false
			}
		}
	}
	return true
}
