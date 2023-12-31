package hangman

import "unicode/utf8"

// Adds the rune passed as a parameter to ListLetter if it's not already there
func (game *HangManData) UsedLetter(oneRune rune) {
	if oneRune < 'A' || oneRune > 'Z' { // Make the letter uppercase if it isn't already
		oneRune = oneRune - 32
	}
	if oneRune >= 'A' && oneRune <= 'Z' { // Add the letter into ListLetter (only if it's a letter)
		game.ListLetter = append(game.ListLetter, oneRune)
	}
}

// Adds the rune passed as a parameter to ListLetter if it's not already there
func (game *HangManData) UsedWord(word string) {
	runeWord := []rune(word)
	for index, runes := range runeWord { // Transforms letters into lower case if they are not already lower case
		if runes < 'a' || runes > 'z' {
			runeWord[index] = runes + 32
		}
	}
	game.ListWord = append(game.ListWord, string(runeWord)) // Add the word into ListLetter
}

// Checks if the input is already in one of the ListWord or ListLetter (return true if it's the case)
func (game *HangManData) UsedVerif(intput string) bool {
	runeWord := []rune(intput)
	oneRune := runeWord[0]
	if utf8.RuneCountInString(intput) > 1 {
		for index, runes := range runeWord { // Transforms letters into lower case if they are not already lower case
			if runes < 'a' || runes > 'z' {
				runeWord[index] = runes + 32
			}
		}
		for _, words := range game.ListWord { // Search the word into ListWord
			if string(runeWord) == words {
				return true
			}
		}
	} else {
		if oneRune < 'A' || oneRune > 'Z' { // Make the letter uppercase if it isn't already
			oneRune = oneRune - 32
		}
		for _, letter := range game.ListLetter { // Search the letter into ListLetter
			if letter == oneRune {
				return true
			}
		}
	}
	return false
}
