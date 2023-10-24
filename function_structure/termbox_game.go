package hangman

import "github.com/nsf/termbox-go"

// FONCTION DRAW EN PERIODE DE TEST, AJOUE LA PARTIE VERIFICATION DES LETTRES? EN COUR : AJOUE D'UNE FONCTION VERIFICATION DE MOT !
func Draw(data HangManData, game Game) {
	var HangMan HangManData = data

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Entrée utilisateur et curseur
	word := "/"
	userInput := ""

	for {
		// Afficher l'entrée utilisateur avec le curseur clignotant
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		game.asciiBox(word)
		HangMan.display()

		drawText([]rune(userInput), 2, 10, termbox.ColorDefault, true)
		drawText(HangMan.Word, 2, 5, termbox.ColorDefault, false)
		drawText(HangMan.ListUsed, 2, 17, termbox.ColorDefault, false)
		termbox.Flush()

		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				return
			} else if (ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter) && userInput != "" {
				if HangMan.meca(userInput) {
					word = "win"
				} else {
					word = userInput
				}
				userInput = ""
			} else if ev.Key == termbox.KeyDelete {
				userInput = ""
			} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				userInput = userInput[:len(userInput)-1]
			} else {
				userInput += string(ev.Ch)
			}
		}
		if HangMan.endGame() {
			if HangMan.Attempts <= 0 {
				word = "lose"
			} else {
				word = "win"
			}
		}
	}
}
