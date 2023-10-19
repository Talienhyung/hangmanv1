package Hangman

import (
	"bufio"
	"fmt"
	"os"
)

func Read(fichier string) []string {
	var dictio []string

	readFile, err := os.Open(fichier)
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile) // Creates a scanner to read the file.

	fileScanner.Split(bufio.ScanLines) // Divides the file into lines.

	// Browse each line of the file.
	for fileScanner.Scan() {
		// Adds the text of the current line to the FindWord slice.
		dictio = append(dictio, fileScanner.Text())
	}

	readFile.Close()

	return dictio
}
