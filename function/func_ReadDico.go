package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// The readfile function returns an array of strings containing all the words in a dictionary
func readFile(fichier string) []string {
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

// The listDictio function returns all files in the Dictinonary directory
func listDictio() []string {
	var listDico []string
	entries, err := os.ReadDir("Ressources/Dictionary/")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		listDico = append(listDico, e.Name())
	}
	return listDico
}

// The ReadAllDico function returns an array of strings containing all the words in the various files in the dictionary folder.
func readAllDico() []string {
	listDico := listDictio()
	var dico []string
	for i := 0; i < len(listDico); i++ {
		newDico := readFile("Ressources/Dictionary/" + listDico[i])
		dico = append(dico, newDico...)
	}
	return dico
}

func ReadTheDico(file string) []string {
	listDico := listDictio()
	for _, j := range listDico {
		if file == j {
			Dico := readFile("Ressources/Dictionary/" + file)
			return Dico
		}
	}
	return readAllDico()
}
