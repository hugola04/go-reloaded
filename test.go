package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/source"
)

func main() {
	data, _ := os.ReadFile("test.txt")

	texte := string(data)

	texte = source.Niveauun(texte)

	mots := strings.Fields(texte)
	
	motsModifies := source.Niveaudeux(mots)

	var resultat []string
	for _, m := range motsModifies {
		if m != "" {
			resultat = append(resultat, m)
		}
	}
	texte = strings.Join(resultat, " ")

	fmt.Println(texte)
}
