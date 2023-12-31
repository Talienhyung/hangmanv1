package hangman

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

// Main mecanic of the game which gathers several functions, return true if the game is finished, otherwise false.
func (hang *HangManData) mainMecanics(input string) bool {
	if utf8.RuneCountInString(input) > 1 { // If it's a word
		if hang.IsThisTheWord(input) {
			hang.Word = []rune(hang.ToFind)
			hang.LastFail = false
			return true
		} else if input == "STOP" { // If the input is STOP, save the game
			err := hang.Save("Ressources/Save/save.txt")
			if err != nil {
				termbox.Close()
				fmt.Println("Game save failed :", err)
				os.Exit(2)
			}
			termbox.Close()
			fmt.Println("Game save in save.txt")
			os.Exit(0)
		} else if input == "QUIT" { // If the input is QUIT, quit the game
			termbox.Close()
			os.Exit(0)
		} else {
			hang.UsedWord(input)
			hang.Attempts -= 2
			hang.HangmanPositions += 2
			hang.LastFail = true
			if hang.HangmanPositions > 9 { // Avoid out of range
				hang.HangmanPositions = 9
				hang.Attempts = 0
			}
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
