package main

import (
	"fmt"
	"os"

	Niveauun "go-reloaded/source"
)

func main() {
	data,_ := os.ReadFile("test.txt")

	texte := Niveauun.Niveauun(string(data))

	fmt.Println(texte)
}