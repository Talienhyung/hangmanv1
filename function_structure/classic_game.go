package hangman

import "fmt"

// This is the hangman classic game
func (game HangManData) ClassicGame() {
	var inputs string
	gameOver := false

	for !gameOver { // Game loop
		// Display word and attempts
		printRune(game.Word)
		fmt.Printf("Attention plus que %d mauvaise(s) réponse(s) acceptée(s) !\n", game.Attempts)
		letter := input("Entrez une lettre ou un mot :", inputs)

		// Verify input
		if game.meca(letter) {
			gameOver = true
		}

		// Display hangman
		if game.HangmanPositions >= 0 && game.HangmanPositions <= 9 {
			game.displayHangmanClassic()
		}

		// Verify if it's the end of the game
		if game.endGame() {
			gameOver = true
		}
	}

	// Announcement of results
	if game.Attempts > 0 {
		fmt.Println("Felicitation ! Vous avez gagné !")
	} else {
		fmt.Println("Le mot était " + game.ToFind + ". Vous ferez mieux la prochaine fois !!")
	}
}
