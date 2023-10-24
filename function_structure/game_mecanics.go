package hangman

import "os"

func (hang *HangManData) meca(input string) bool {
	if len(input) > 1 {
		if hang.IsThisTheWord(input) {
			return true
		} else if input == "STOP" {
			hang.Save("Ressources/Save/save.txt")
		} else if input == "QUIT" {
			os.Exit(0)
		} else {
			hang.Attempts -= 2
			hang.HangmanPositions += 2
		}
	} else {
		oneRune := []rune(input)
		hang.LetterInWord(oneRune[0])
		hang.UsedLetter(oneRune[0])
	}
	return false
}

func (game *HangManData) endGame() bool {
	if game.Attempts <= 0 {
		return true
	}
	for _, runes := range game.Word {
		if runes == '_' {
			return false
		}
	}
	return true
}
