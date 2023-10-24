package hangman

func (game *HangManData) UsedLetter(oneRune rune) {
	if oneRune < 'A' || oneRune > 'Z' {
		oneRune = oneRune - 32
	}
	for _, letter := range game.ListUsed {
		if letter == oneRune {
			return
		}
	}
	if oneRune >= 'A' && oneRune <= 'Z' {
		game.ListUsed = append(game.ListUsed, oneRune)
	}
}

func (game *HangManData) UsedWord(word string) {
	runeWord := []rune(word)
	for index, runes := range runeWord { //transforms letters into lower case if they are not already lower case
		if runes < 'a' || runes > 'z' {
			runeWord[index] = runes + 32
		}
	}

	for _, words := range game.ListWord {
		if string(runeWord) == words {
			return
		}
	}
	game.ListWord = append(game.ListWord, string(runeWord))
}
