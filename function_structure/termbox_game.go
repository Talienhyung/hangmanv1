package hangman

import "github.com/nsf/termbox-go"

// TermBoxGame is a function that handles the main game loop for a Hangman game using the termbox library.
// It takes the HangManData and Game structs as input parameters.
func (HangMan HangManData) TermBoxGame(game Game) {
	// Initialize the termbox library and handle errors
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Initialize user input and cursor
	word := "/"
	userInput := ""
	gameOver := false

	for {
		// Clear the screen and set up user interface
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		game.asciiBox(word)
		game.AsciiCounter(HangMan.Attempts)

		HangMan.display()

		// Display text
		drawText([]rune(userInput), 2, 10, termbox.ColorDefault, true)
		drawText(HangMan.Word, 2, 4, termbox.ColorDefault, false)
		drawText(HangMan.ListLetter, 2, 17, termbox.ColorDefault, false)
		for i := range HangMan.ListWord {
			drawText([]rune(HangMan.ListWord[i]), 2, 18+i, termbox.ColorDefault, false)
		}
		termbox.Flush()

		// Poll for user input events
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				return // Exit the game loop
			} else if ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter {
				if !gameOver {
					if userInput != "" && !HangMan.UsedVerif(userInput) && userInput != "Empty or already proposed!" {
						// Check if the user's input is a valid guess and update the word or game status
						if HangMan.mainMecanics(userInput) {
							word = "win"
							gameOver = true
						} else {
							word = userInput
						}
						userInput = "" // Clear user input
					} else {
						userInput = "Empty or already proposed!"
					}
				} else {
					if userInput == "QUIT" {
						return
					}
				}
			} else if ev.Key == termbox.KeyDelete {
				userInput = "" // Clear user input
			} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				if userInput != "" && userInput != "Empty or already proposed!" {
					userInput = userInput[:len(userInput)-1] // Remove the last character from user input
				}
			} else {
				if userInput == "Empty or already proposed!" {
					userInput = ""
				}
				userInput += string(ev.Ch) // Add the character to user input
			}
		}

		// Check if the game has ended
		if HangMan.endGame() {
			if HangMan.Attempts <= 0 {
				word = "lose"
				HangMan.Word = []rune(HangMan.ToFind)
			} else {
				word = "win"
			}
			gameOver = true
		}
	}
}
