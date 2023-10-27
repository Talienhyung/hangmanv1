package hangman

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// Displays a given ascii character in x y
func (data *Game) displayAscii(x, y, version int, borderColor termbox.Attribute) {
	var ascii [95][9]string
	switch data.letterFile { //choose the right font for ascii art
	case "shadow.txt":
		ascii = readAscii("Ressources/Ascii_Letter/shadow.txt")
	case "standard.txt":
		ascii = readAscii("Ressources/Ascii_Letter/standard.txt")
	case "thinkertoy.txt":
		ascii = readAscii("Ressources/Ascii_Letter/thinkertoy.txt")
	default:
		ascii = readAscii("Ressources/Ascii_Letter/standard.txt")
	}
	for i := 0; i <= 8; i++ { //displays the correct character
		runes := []rune(ascii[version-32][i])
		for index, j := range runes {
			termbox.SetCell(x+index, y+i, j, borderColor, termbox.ColorDefault)
		}
	}
}

// Displays the last letter entered by a user in the terminal, followed by the final result (win or lose).
func (data *Game) asciiBox(word string) {
	switch word {
	case "win": //display WIN if player win
		data.displayAscii(55+18, 15, 'I', termbox.ColorGreen)
		data.displayAscii(55, 15, 'W', termbox.ColorGreen)
		data.displayAscii(55+16+14, 15, 'N', termbox.ColorGreen)
	case "lose": //display lose if player lose
		data.displayAscii(56, 15, 'L', termbox.ColorRed)
		data.displayAscii(55+11, 15, 'O', termbox.ColorRed)
		data.displayAscii(55+16+5, 15, 'S', termbox.ColorRed)
		data.displayAscii(55+16+14, 15, 'E', termbox.ColorRed)
	default: //displays the first rune of the last input
		runes := []rune(word)
		data.displayAscii(55+16, 15, int(runes[0]), termbox.ColorLightRed)
	}
}

// displayAsciiText displays ASCII art text using the 'letterFile' font.
// The function selects the font based on 'letterFile' and displays the ASCII art text.
func (data *Game) displayAsciiText(words []rune) {
	// Define a 2D array 'ascii' to hold ASCII art characters.
	var ascii [95][9]string

	// Select the font for the ASCII art based on the 'letterFile' field.
	switch data.letterFile {
	case "shadow":
		ascii = readAscii("Ressources/Ascii_Letter/shadow.txt")
	case "standard":
		ascii = readAscii("Ressources/Ascii_Letter/standard.txt")
	case "thinkertoy":
		ascii = readAscii("Ressources/Ascii_Letter/thinkertoy.txt")
	default:
		ascii = readAscii("Ressources/Ascii_Letter/standard.txt")
	}

	// Loop through each line of the ASCII art (9 lines in total).
	for line := 0; line <= 8; line++ {
		// Loop through each letter (rune) in the 'words' slice.
		for _, letter := range words {
			// Print the ASCII character for the current letter on the current line.
			// The ASCII value of the letter is used to index 'ascii' array.
			fmt.Printf(ascii[letter-32][line])
		}
		// After printing a line of text, move to the next line.
		fmt.Println("")
	}
}
