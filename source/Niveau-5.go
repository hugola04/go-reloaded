package source 

import (
	"strings"
	"unicode"
)

func Niveaucinq(texte string) string {
	mots := strings.Fields(texte)

	for i := 0; i < len(mots)-1; i++ {

		switch mots[i] {
		case "a", "A", "an", "An":
			suivant := mots[i+1]

			nettoye := strings.TrimLeft(suivant, "'")

			if len(nettoye) == 0 {
				continue
			}

			premiereLettre := []rune(nettoye)[0]

			if !unicode.IsLetter(premiereLettre) {
				continue
			}

			lettreMin := unicode.ToLower(premiereLettre)

			switch lettreMin {
			case 'a', 'e', 'i', 'o', 'u':
				if mots[i] == "A" || mots[i] == "An" {
					mots[i] = "An"
				} else {
					mots[i] = "an"
				}
			default:
				if mots[i] == "A" || mots[i] == "An" {
					mots[i] = "A"
				} else {
					mots[i] = "a"
				}
			}
		}
	}

	return strings.Join(mots, " ")
}
