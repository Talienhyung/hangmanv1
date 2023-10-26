package hangman

import "github.com/nsf/termbox-go"

// Draw is a function that handles the main game loop for a Hangman game using the termbox library.
// It takes the HangManData and Game structs as input parameters.
func Draw(data HangManData, game Game) {
	// Initialize HangManData with the provided data
	var HangMan HangManData = data

	// Initialize the termbox library and handle errors
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Initialize user input and cursor
	word := "/"
	userInput := ""

	for {
		// Clear the screen and set up user interface
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		game.asciiBox(word)
		HangMan.display()

		// Display user input with a blinking cursor
		drawText([]rune(userInput), 2, 10, termbox.ColorDefault, true)
		drawText(HangMan.Word, 2, 5, termbox.ColorDefault, false)
		drawText(HangMan.ListUsed, 2, 17, termbox.ColorDefault, false)
		termbox.Flush()

		// Poll for user input events
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				return // Exit the game loop
			} else if (ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter) && userInput != "" {
				// Check if the user's input is a valid guess and update the word or game status
				if HangMan.meca(userInput) {
					word = "win"
				} else {
					word = userInput
				}
				userInput = "" // Clear user input
			} else if ev.Key == termbox.KeyDelete {
				userInput = "" // Clear user input
			} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				userInput = userInput[:len(userInput)-1] // Remove the last character from user input
			} else {
				userInput += string(ev.Ch) // Add the character to user input
			}
		}

		// Check if the game has ended
		if HangMan.endGame() {
			if HangMan.Attempts <= 0 {
				word = "lose"
			} else {
				word = "win"
			}
		}
	}
}
