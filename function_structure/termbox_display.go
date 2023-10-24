package hangman

import "github.com/nsf/termbox-go"

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

func (hang *HangManData) DisplayHangman(x, y int, borderColor termbox.Attribute) {
	hangMan := readHang("Ressources/HangMan_Position/hangman.txt")
	for i := 0; i <= 7; i++ {
		runes := []rune(hangMan[hang.HangmanPositions][i])
		for index, j := range runes {
			termbox.SetCell(x+index, y+i, j, borderColor, termbox.ColorDefault)
		}
	}
}
