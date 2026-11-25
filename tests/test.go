package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/source"
)

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	texte := string(data)

	texte = source.Niveauun(texte)

	mots := strings.Fields(texte)
	mots = source.Niveaudeux(mots)
	texte = strings.Join(mots, " ")

	texte = source.Niveautrois(texte)
	texte = source.Niveauquatre(texte)
	texte = source.Niveaucinq(texte)

	expected, err := os.ReadFile("finis.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if strings.Join(strings.Fields(texte), " ") != strings.Join(strings.Fields(string(expected)), " ") {
		fmt.Fprintln(os.Stderr, "Échec : le résultat des cinq niveaux ne correspond pas à finis.txt.")
		fmt.Fprintf(os.Stderr, "Attendu : %s\nObtenu : %s\n", strings.Join(strings.Fields(string(expected)), " "), strings.Join(strings.Fields(texte), " "))
		os.Exit(1)
	}

	fmt.Println("Test réussi : test.txt après les cinq niveaux correspond à finis.txt.")
}
