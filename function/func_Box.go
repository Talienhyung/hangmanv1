package Hangman

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (hang *HangManData) meca(input string) bool {
	if len(input) > 1 {
		if hang.IsThisTheWord(input) {
			return true
		} else if input == "STOP" {
			hang.Save("save.txt")
		} else {
			hang.Attempts -= 2
			hang.HangmanPositions += 2
		}
	} else {
		oneRune := []rune(input)
		hang.LetterInWord(oneRune[0])
		hang.UsedLetter(oneRune[0])
	}
	return false
}

func drawBox(x, y, width, height int, borderColor termbox.Attribute, title string) {
	// Dessine le cadre de la boîte
	for i := x; i < x+width; i++ {
		termbox.SetCell(i, y, '─', borderColor, termbox.ColorDefault)
		termbox.SetCell(i, y+height-1, '─', borderColor, termbox.ColorDefault)
	}
	for i := y; i < y+height; i++ {
		termbox.SetCell(x, i, '│', borderColor, termbox.ColorDefault)
		termbox.SetCell(x+width-1, i, '│', borderColor, termbox.ColorDefault)
	}

	// Coins de la boîte
	termbox.SetCell(x, y, '┌', borderColor, termbox.ColorDefault)
	termbox.SetCell(x+width-1, y, '┐', borderColor, termbox.ColorDefault)
	termbox.SetCell(x, y+height-1, '└', borderColor, termbox.ColorDefault)
	termbox.SetCell(x+width-1, y+height-1, '┘', borderColor, termbox.ColorDefault)

	// Ajouter le titre
	for i, ch := range title {
		termbox.SetCell(x+i+1, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (hang *HangManData) display(letter rune) {
	// Boîte principale
	drawBox(0, 0, 100, 24, termbox.ColorWhite, "main")

	// Première boîte à l'intérieur de la boîte principale
	drawBox(55, 0, 45, 15, termbox.ColorRed, "Hangman")
	if hang.HangmanPositions >= 0 && hang.HangmanPositions <= 9 {
		hang.DisplayHangman(55+18, 4, termbox.ColorBlue)
	}

	// Deuxième boîte à l'intérieur de la boîte principale
	drawBox(0, 0, 50, 8, termbox.ColorBlue, "Word...")

	// Troisieme boîte à l'intérieur de la boîte principale
	drawBox(0, 8, 50, 8, termbox.ColorGreen, "Letter")

	// Quatrieme boîte à l'intérieur de la boîte principale
	drawBox(0, 16, 50, 8, termbox.ColorLightMagenta, "Used letter")
}

func drawText(text []rune, x, y int, color termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (data *Game) asciiBox(word string) {
	switch word {
	case "win":
		data.DisplayAscii(55+18, 15, 'I', termbox.ColorGreen)
		data.DisplayAscii(55, 15, 'W', termbox.ColorGreen)
		data.DisplayAscii(55+16+14, 15, 'N', termbox.ColorGreen)
	case "lose":
		data.DisplayAscii(55, 15, 'L', termbox.ColorGreen)
		data.DisplayAscii(55+16, 15, 'O', termbox.ColorGreen)
		data.DisplayAscii(55+16+12, 15, 'S', termbox.ColorGreen)
		data.DisplayAscii(55+16+24, 15, 'E', termbox.ColorGreen)
	}
	// data.DisplayAscii(55+16, 15, int(letter), termbox.ColorLightRed)
}

// FONCTION DRAW EN PERIODE DE TEST, AJOUE LA PARTIE VERIFICATION DES LETTRES? EN COUR : AJOUE D'UNE FONCTION VERIFICATION DE MOT !
func Draw(data HangManData, game Game) {
	var HangMan HangManData = data

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Entrée utilisateur et curseur
	letter := '/'
	word := ""
	game.asciiBox(word)
	userInput := ""
	cursorVisible := true
	termbox.Flush()

	// Configuration du minuteur pour le curseur clignotant
	cursorTimer := time.NewTicker(500 * time.Millisecond)
	defer cursorTimer.Stop()

	for {
		// Afficher l'entrée utilisateur avec le curseur clignotant
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		HangMan.display(letter)
		drawText([]rune(userInput), 2, 10, termbox.ColorDefault)
		drawText(HangMan.Word, 2, 5, termbox.ColorDefault)
		drawText(HangMan.ListUsed, 2, 17, termbox.ColorDefault)

		// Afficher ou masquer le curseur clignotant
		if cursorVisible {
			termbox.SetCell(2+len(userInput), 10, '_', termbox.ColorDefault, termbox.ColorDefault)
		}

		termbox.Flush()

		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				return
			} else if (ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter) && userInput != "" {
				runes := []rune(userInput)
				letter = runes[0]
				if HangMan.meca(userInput) {
					word = "win"
				}
				userInput = ""
			} else {
				userInput += string(ev.Ch)
			}
		}
		if HangMan.endGame() {
			word = "win"
		}
	}
}
