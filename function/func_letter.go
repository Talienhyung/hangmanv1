package Hangman

//################################## POO #######################################

func (game *HangManData) LetterInWord(oneRune rune) {
	var place []int
	for index, letters := range game.ToFind { //transforms letters into lower case if they are not already lower case
		if oneRune >= 'A' && oneRune <= 'Z' {
			oneRune = oneRune + 32
		}
		if letters == oneRune || letters == oneRune-32 {
			place = append(place, index) //saves the index(es) of the position where the letter was found
		}
	}
	if len(place) != 0 { //if any letters have been found then replace the corresponding slots with the letters
		for _, index := range place {
			game.Word[index] = oneRune
		}
	} else { //if the letter is not found, an attempt is lost
		game.Attempts--
		game.HangmanPositions++
	}
}

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

func (game *HangManData) IsThisTheWord(word string) bool {
	var oneRune rune
	if len(word) != len(game.ToFind) { //check that words are the same size
		return false
	} else {
		wordRune := []rune(word)
		toFindRune := []rune(game.ToFind)
		for index, runes := range wordRune {
			oneRune = runes
			if runes >= 'A' && runes <= 'Z' { //puts letters in lower case for easier comparison
				oneRune = oneRune + 32
			}
			if oneRune != toFindRune[index] && oneRune != toFindRune[index]+32 {
				return false
			}
		}
	}
	return true
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
