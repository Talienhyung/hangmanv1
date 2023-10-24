package Hangman

import (
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

func (hang *HangManData) display() {
	// Boîte principale
	drawBox(0, 0, 100, 24, termbox.ColorWhite, "main")

	// Première boîte à l'intérieur de la boîte principale
	drawBox(55, 0, 45, 15, termbox.ColorLightYellow, "Hangman")
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

func drawText(text []rune, x, y int, color termbox.Attribute, cursor bool) {
	ligne := 0
	space := 0
	for i, ch := range text {
		switch {
		case i > 45 && i <= 91:
			space = 46
			ligne = 1
		case i > 91 && i <= 137:
			space = 92
			ligne = 2
		case i > 137 && i <= 183:
			space = 138
			ligne = 3
		case i > 183 && i <= 229:
			space = 184
			ligne = 4
		default:
			ligne = 0
			space = 0
		}
		termbox.SetCell(x+i-space, y+ligne, ch, termbox.ColorDefault, termbox.ColorDefault)

	}
	if cursor {
		termbox.SetCell(2+len(text)-space, 10+ligne, '_', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (data *Game) asciiBox(word string) {
	switch word {
	case "win":
		data.DisplayAscii(55+18, 15, 'I', termbox.ColorGreen)
		data.DisplayAscii(55, 15, 'W', termbox.ColorGreen)
		data.DisplayAscii(55+16+14, 15, 'N', termbox.ColorGreen)
	case "lose":
		data.DisplayAscii(56, 15, 'L', termbox.ColorRed)
		data.DisplayAscii(55+11, 15, 'O', termbox.ColorRed)
		data.DisplayAscii(55+16+5, 15, 'S', termbox.ColorRed)
		data.DisplayAscii(55+16+14, 15, 'E', termbox.ColorRed)
	default:
		runes := []rune(word)
		data.DisplayAscii(55+16, 15, int(runes[0]), termbox.ColorLightRed)
	}
}

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
