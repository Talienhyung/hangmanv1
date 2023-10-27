package hangman

import (
	"os"
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

// Main mecanic of the game which gathers several functions, return true if the game is finished, otherwise false.
func (hang *HangManData) meca(input string) bool {
	if utf8.RuneCountInString(input) > 1 { // If it's a word
		if hang.IsThisTheWord(input) {
			hang.Word = []rune(hang.ToFind)
			return true
		} else if input == "STOP" {
			hang.Save("Ressources/Save/save.txt")
		} else if input == "QUIT" {
			termbox.Close()
			os.Exit(0)
		} else {
			hang.UsedWord(input)
			hang.Attempts -= 2
			hang.HangmanPositions += 2
		}
	} else { // If it's a letter
		oneRune := []rune(input)
		hang.LetterInWord(oneRune[0])
		hang.UsedLetter(oneRune[0])
	}
	return false
}

// Check if the game is finished or not
func (game *HangManData) endGame() bool {
	if game.Attempts <= 0 { // No more attempts
		return true
	}
	for _, runes := range game.Word {
		if runes == '_' {
			return false
		}
	}
	return true // Words found
}
