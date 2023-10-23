package Hangman

import (
	"fmt"
)

var inputs string

// This is the hangman classic game
func ClassicGame(data HangManData) {
	var game HangManData = data
	gameOver := false

	for !gameOver {
		printRune(game.Word)
		fmt.Printf("Attention plus que %d mauvaise reponse accepté !\n", game.Attempts)
		letter := input("Entrez une lettre ou un mot :", inputs)
		if game.meca(letter) {
			gameOver = true
		}

		if game.HangmanPositions >= 0 && game.HangmanPositions <= 9 {
			game.displayHangmanClassic()
		}

		if game.endGame() {
			gameOver = true
		}
	}
	if game.Attempts != 0 {
		fmt.Println("Felicitation ! Vous avez gagné !")
	} else {
		fmt.Println("Le mot était " + game.ToFind + ". Vous ferez mieux la prochaine fois !!")
	}
}

func (game *HangManData) endGame() bool {
	if game.Attempts == 0 {
		return true
	}
	for _, runes := range game.Word {
		if runes == '_' {
			return false
		}
	}
	return true
}

func (hang HangManData) displayHangmanClassic() {
	hangMan := readHang("Ressources/HangMan_Position/hangman.txt")
	fmt.Println("")
	for i := 0; i <= 7; i++ {
		fmt.Println(hangMan[hang.HangmanPositions][i])
	}
}

func printRune(tab []rune) {
	for _, runes := range tab {
		fmt.Print(string(runes))
	}
	fmt.Print("\n")
}

func input(s string, inputs string) string {
	fmt.Println(s)
	fmt.Scanln(&inputs)
	return inputs
}
