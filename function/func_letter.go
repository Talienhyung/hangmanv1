package Hangman

// fonction qui cherche si une lettre donné ce trouve dans le mot donné puis renvoie le mot qui sera affichier selon si une lettre a été trouvée ou non
func LetterInWord(oneRune rune, word string, displayWord []rune, hang int) ([]rune, int) {
	var place []int
	for index, letters := range word {
		if oneRune >= 'A' && oneRune <= 'Z' {
			oneRune = oneRune + 32
		}
		if letters == oneRune || letters == oneRune-32 {
			place = append(place, index)
		}
	}
	if len(place) != 0 {
		for _, index := range place {
			displayWord[index] = oneRune
		}
	} else {
		hang++
	}

	return displayWord, hang
}

func UsedLetter(oneRune rune, listUsed []rune) []rune {
	if oneRune < 'A' || oneRune > 'Z' {
		oneRune = oneRune - 32
	}
	for _, letter := range listUsed {
		if letter == oneRune {
			return listUsed
		}
	}
	if oneRune >= 'A' && oneRune <= 'Z' {
		listUsed = append(listUsed, oneRune)
	}
	return listUsed
}
