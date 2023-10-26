package hangman

// Adds the rune passed as a parameter to ListUsed if it's not already there
func (game *HangManData) UsedLetter(oneRune rune) {
	if oneRune < 'A' || oneRune > 'Z' { // Make the letter uppercase if it isn't already
		oneRune = oneRune - 32
	}
	for _, letter := range game.ListUsed { // Search the letter into ListUsed
		if letter == oneRune {
			return
		}
	}
	if oneRune >= 'A' && oneRune <= 'Z' { // Add the letter into ListUsed (only if it's a letter)
		game.ListUsed = append(game.ListUsed, oneRune)
	}
}

// Adds the rune passed as a parameter to ListUsed if it's not already there
func (game *HangManData) UsedWord(word string) {
	runeWord := []rune(word)
	for index, runes := range runeWord { // Transforms letters into lower case if they are not already lower case
		if runes < 'a' || runes > 'z' {
			runeWord[index] = runes + 32
		}
	}

	for _, words := range game.ListWord { // Search the word into ListWord
		if string(runeWord) == words {
			return
		}
	}
	game.ListWord = append(game.ListWord, string(runeWord)) // Add the word into ListUsed
}
