<<<<<<< HEAD
package source
=======
package source

import (
	"strings"
	"unicode"
)

// Niveauquatre supprime les espaces autour des guillemets associés et conserve les apostrophes.
func Niveauquatre(texte string) string {
	lines := strings.Split(texte, "\n")
	for i, line := range lines {
		lines[i] = quoteLine(line)
	}
	return strings.Join(lines, "\n")
}

func quoteLine(line string) string {
	runes := []rune(line)
	remove := make([]bool, len(runes))
	quoteIndexes := make([]int, 0)

	for i, r := range runes {
		if r != '\'' {
			continue
		}
		if i > 0 && i+1 < len(runes) &&
			isLetter(runes[i-1]) && isLetter(runes[i+1]) {
			continue
		}
		quoteIndexes = append(quoteIndexes, i)
	}

	for pair := 0; pair+1 < len(quoteIndexes); pair += 2 {
		openIndex := quoteIndexes[pair]
		closeIndex := quoteIndexes[pair+1]

		for i := openIndex + 1; i < closeIndex; i++ {
			if runes[i] != ' ' && runes[i] != '\t' {
				break
			}
			remove[i] = true
		}
		for i := closeIndex - 1; i > openIndex; i-- {
			if runes[i] != ' ' && runes[i] != '\t' {
				break
			}
			remove[i] = true
		}
	}

	var result []rune
	for i, r := range runes {
		if !remove[i] {
			result = append(result, r)
		}
	}

	return string(result)
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}
>>>>>>> main
