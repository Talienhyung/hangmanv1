package hangman

import (
	"fmt"
)

// This is the hangman Ascii game
func (game HangManData) AsciiGame(data Game) {
	var inputs string
	gameOver := false
	fmt.Println("Good Luck, you have 10 attempts.")
	data.displayAsciiText(game.Word)

	for !gameOver { // Game loop
		// Display input
		letter := input("\nChoose : ", inputs)

		// Verify input
		if !game.UsedVerif(letter) && letter != " " {
			if game.meca(letter) {
				gameOver = true
			}

			if game.LastFail {
				fmt.Printf("Not present in the word, %d attempts remaining\n", game.Attempts)
			}

			// Display AsciiText and HangMan
			data.displayAsciiText(game.Word)
			if game.HangmanPositions >= 0 {
				game.displayHangmanClassic()
			}

		} else {
			fmt.Println("Empty or already proposed!")
		}

		// Verify if it's the end of the game
		if game.endGame() {
			gameOver = true
		}
	}

	// Announcement of results
	if game.Attempts > 0 {
		fmt.Println("Congrats !")
	} else {
		fmt.Println("The word was " + game.ToFind + ". You'll do better next time!!!")
	}
}
