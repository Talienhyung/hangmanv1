package Hangman

import (
	"time"

	"github.com/nsf/termbox-go"
)

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

func display(hang int, letter rune) {
	// Boîte principale
	DisplayAscii(55+16, 15, int(letter), termbox.ColorLightRed, "standard")
	drawBox(0, 0, 100, 24, termbox.ColorWhite, "main")

	// Première boîte à l'intérieur de la boîte principale
	drawBox(55, 0, 45, 15, termbox.ColorRed, "Hangman")
	DisplayHangman(55+18, 4, hang, termbox.ColorBlue)

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

// FONCTION DRAW EN PERIODE DE TEST, AJOUE LA PARTUE VERIFICATION DES LETTRES? EN COUR : AJOUE D'UNE FONCTION VERIFICATION DE MOT !
func Draw() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Entrée utilisateur et curseur

	var HangMan HangManData

	userInput := ""
	HangMan.ToFind := "ecriture"
	HangMan.Word := []rune("________")
	HangMan.Attempts := 0
	var letter rune
	var listWordUsed []string
	var listUsed []rune
	cursorVisible := true
	termbox.Flush()

	// Configuration du minuteur pour le curseur clignotant
	cursorTimer := time.NewTicker(500 * time.Millisecond)
	defer cursorTimer.Stop()

	for {
		// Afficher l'entrée utilisateur avec le curseur clignotant
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		display(HangMan.Attempts, 'A')

		drawText([]rune(userInput), 2, 10, termbox.ColorDefault)
		drawText(HangMan.Word, 2, 5, termbox.ColorDefault)
		drawText(listUsed, 2, 17, termbox.ColorDefault)

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
				if moreThanOneLetter(userInput) {
					HangMan.Word, hang, listWordUsed = moreThanOneLetter(userInput, HangMan.ToFind, HangMan.Word, listWordUsed, HangMan.Attempts)
				}
				HangMan.Word, hang, listUsed = oneOrmoreLetter(userInput, HangMan.ToFind, HangMan.Word, listUsed, HangMan.Attempts)
				userInput = ""
			} else {
				userInput += string(ev.Ch)
			}
		}
	}
}

func moreThanOneLetter(userInput string) bool {
	runes := []rune(userInput)
	if len(runes) > 1 {
		return true
	}
	return false
}

func oneLetter(userInput, ToFind string, Word, listUsed []rune, hang int) ([]rune, int, []rune) {
	runes := []rune(userInput)
	a, b := LetterInWord(runes[0], ToFind, Word, hang)
	return a, b, UsedLetter(runes[0], listUsed)
}
