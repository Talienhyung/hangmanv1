package hangman

import "fmt"

// This is the hangman classic game
func ClassicGame(data HangManData) {
	var inputs string
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
	if game.Attempts > 0 {
		fmt.Println("Felicitation ! Vous avez gagné !")
	} else {
		fmt.Println("Le mot était " + game.ToFind + ". Vous ferez mieux la prochaine fois !!")
	}
}
