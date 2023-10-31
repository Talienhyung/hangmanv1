package hangman

import "fmt"

// This is the hangman classic game
func (game HangManData) ClassicGame() {
	var inputs string
	gameOver := false
	fmt.Printf("Good Luck, you have %d attempts.\n", game.Attempts)
	printRune(game.Word)

	for !gameOver { // Game loop
		// Display word and attempts
		letter := input("\nChoose : ", inputs)

		// Verify input
		if letter != "" && !game.UsedVerif(letter) {
			if game.meca(letter) {
				gameOver = true
			}

			if game.LastFail {
				fmt.Printf("Not present in the word, %d attempts remaining\n", game.Attempts)
			}

			// Display words and HangMan
			printRune(game.Word)
			if game.HangmanPositions >= 0 {
				game.displayHangmanClassic()
			}

			// Verify if it's the end of the game
			if game.endGame() {
				gameOver = true
			}
		} else {
			fmt.Println("Empty or already proposed!")
		}
	}

	// Announcement of results
	if game.Attempts > 0 {
		fmt.Println("Congrats !")
	} else {
		fmt.Println("The word was " + game.ToFind + ". You'll do better next time!!!")
	}
}
