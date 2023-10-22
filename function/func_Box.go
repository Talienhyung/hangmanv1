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

func display() {
	// Boîte principale
	drawBox(0, 0, 100, 24, termbox.ColorWhite, "main")

	// Première boîte à l'intérieur de la boîte principale
	drawBox(55, 0, 45, 15, termbox.ColorRed, "Hangman")
	DisplayHangman(55+22-9, 2, 7, termbox.ColorBlue)

	// Deuxième boîte à l'intérieur de la boîte principale
	drawBox(0, 0, 50, 8, termbox.ColorBlue, "Word...")

	// Troisieme boîte à l'intérieur de la boîte principale
	drawBox(0, 8, 50, 8, termbox.ColorGreen, "Letter")

	// Quatrieme boîte à l'intérieur de la boîte principale
	drawBox(0, 16, 50, 8, termbox.ColorLightMagenta, "Used letter")
}

func Draw() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Entrée utilisateur et curseur
	userInput := ""
	cursorVisible := true
	termbox.Flush()

	// Configuration du minuteur pour le curseur clignotant
	cursorTimer := time.NewTicker(500 * time.Millisecond)
	defer cursorTimer.Stop()

	for {
		// Afficher l'entrée utilisateur avec le curseur clignotant
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		display()
		for i, ch := range userInput {
			termbox.SetCell(i, 5, ch, termbox.ColorDefault, termbox.ColorDefault)
		}

		// Afficher ou masquer le curseur clignotant
		if cursorVisible {
			termbox.SetCell(len(userInput), 5, '_', termbox.ColorDefault, termbox.ColorDefault)
		}

		termbox.Flush()

		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter {
				userInput = ""
			} else {
				userInput += string(ev.Ch)
			}
		}
	}
}
