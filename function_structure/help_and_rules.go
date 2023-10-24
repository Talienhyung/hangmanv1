package hangman

import (
	"fmt"
	"os"
	"os/exec"
)

func Help() {
	filePath := "Ressources/doc.txt"

	pagerCommand := "less"

	cmd := exec.Command(pagerCommand, filePath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erreur lors de l'exécution de %s : %v\n", pagerCommand, err)
		os.Exit(1)
	}
}

func Rules() {
	filePath := "Ressources/rules.txt"

	pagerCommand := "less"

	cmd := exec.Command(pagerCommand, filePath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erreur lors de l'exécution de %s : %v\n", pagerCommand, err)
		os.Exit(1)
	}
}
