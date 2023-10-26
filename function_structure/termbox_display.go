package hangman

import "github.com/nsf/termbox-go"

// Draw a box in x, y with size width/height, color borderColor with title in terminal
func drawBox(x, y, width, height int, borderColor termbox.Attribute, title string) {
	// Draw the box frame
	for i := x; i < x+width; i++ {
		termbox.SetCell(i, y, '─', borderColor, termbox.ColorDefault)
		termbox.SetCell(i, y+height-1, '─', borderColor, termbox.ColorDefault)
	}
	for i := y; i < y+height; i++ {
		termbox.SetCell(x, i, '│', borderColor, termbox.ColorDefault)
		termbox.SetCell(x+width-1, i, '│', borderColor, termbox.ColorDefault)
	}

	// Box corners
	termbox.SetCell(x, y, '┌', borderColor, termbox.ColorDefault)
	termbox.SetCell(x+width-1, y, '┐', borderColor, termbox.ColorDefault)
	termbox.SetCell(x, y+height-1, '└', borderColor, termbox.ColorDefault)
	termbox.SetCell(x+width-1, y+height-1, '┘', borderColor, termbox.ColorDefault)

	// Add title
	for i, ch := range title {
		termbox.SetCell(x+i+1, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

// Main display, efficient for all boxes and hangman
func (hang *HangManData) display() {
	// Main box
	drawBox(0, 0, 100, 24, termbox.ColorWhite, "main")

	// First box inside the main box
	drawBox(55, 0, 45, 15, termbox.ColorLightYellow, "Hangman")
	// HangMan in the first box
	if hang.HangmanPositions >= 0 && hang.HangmanPositions <= 9 {
		hang.DisplayHangman(55+18, 4, termbox.ColorBlue)
	}

	// Second box inside the main box
	drawBox(0, 0, 50, 8, termbox.ColorBlue, "Word...")

	// Third box inside the main box
	drawBox(0, 8, 50, 8, termbox.ColorGreen, "Letter")

	// Fourth box inside the main box
	drawBox(0, 16, 50, 8, termbox.ColorLightMagenta, "Used letter/words")
}

// drawText is a function that draws text
// It takes a slice of runes (text), x and y coordinates, a color attribute, and a cursor flag.
func drawText(text []rune, x, y int, color termbox.Attribute, cursor bool) {
	// Initialize variables to track line and space offsets
	ligne := 0
	space := 0

	// Loop through the text and process each character
	for i, ch := range text {
		// Determine the line and space offsets based on the character index
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

		// Set the character at the specified position with the given color
		termbox.SetCell(x+i-space, y+ligne, ch, termbox.ColorDefault, termbox.ColorDefault)
	}

	// If the cursor flag is true, draw a cursor character at a specific position
	if cursor {
		termbox.SetCell(2+len(text)-space, 10+ligne, '_', termbox.ColorDefault, termbox.ColorDefault)
	}
}

// Display HangMan on the right position
func (hang *HangManData) DisplayHangman(x, y int, borderColor termbox.Attribute) {
	hangMan := readHang("Ressources/HangMan_Position/hangman.txt")
	for i := 0; i <= 7; i++ { // Display ligne by ligne
		runes := []rune(hangMan[hang.HangmanPositions][i])
		for index, j := range runes { // Display rune by rune
			termbox.SetCell(x+index, y+i, j, borderColor, termbox.ColorDefault)
		}
	}
}
