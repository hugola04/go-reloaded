package main

import (
	"fmt"
	"os"

	Niveauun "go-reloaded/source"
	Niveautrois "go-reloaded/source"
)

func main() {
	data, _ := os.ReadFile("test.txt")

	texte := string(data)

	texte = Niveauun.Niveauun(texte)
	//rajoute ton niveaudeux
	texte = Niveautrois.Niveautrois(texte)

	fmt.Println(texte)
}
