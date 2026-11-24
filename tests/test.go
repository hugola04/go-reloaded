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

	// Niveau 1 (prend un string, renvoie un string)
	texte = source.Niveauun(texte)

	// Niveau 2 (prend un []string, renvoie un []string)
	mots := strings.Fields(texte)
	mots = source.Niveaudeux(mots)
	texte = strings.Join(mots, " ")

	// Niveau 3 (prend un string, renvoie un string)
	texte = source.Niveautrois(texte)

	fmt.Println(texte)
}
