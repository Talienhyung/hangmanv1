package hangman

import "github.com/nsf/termbox-go"

func (data *Game) DisplayAscii(x, y, version int, borderColor termbox.Attribute) {
	var ascii [95][9]string
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

	for i := 0; i <= 8; i++ {
		runes := []rune(ascii[version-32][i])
		for index, j := range runes {
			termbox.SetCell(x+index, y+i, j, borderColor, termbox.ColorDefault)
		}
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
