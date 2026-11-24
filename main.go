package main

import (
	"fmt"
	"os"
	"strings"

	Niveauun "go-reloaded/source"
	Niveaudeux "go-reloaded/source"
	Niveautrois "go-reloaded/source"
	Niveauquatre "go-reloaded/source"
	Niveaucinq "go-reloaded/source"
)

func main() {
	data, _ := os.ReadFile("test.txt")

	texte := string(data)

	texte = Niveauun.Niveauun(texte)

	mots := strings.Fields(texte)
	mots = Niveaudeux.Niveaudeux(mots)
	texte = strings.Join(mots, " ")

	texte = Niveautrois.Niveautrois(texte)
	texte = Niveauquatre.Niveauquatre(texte)
	texte = Niveaucinq.Niveaucinq(texte)

	fmt.Println(texte)
}
