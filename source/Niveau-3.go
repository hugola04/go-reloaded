package source

import "strings"

func Niveautrois(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}

	for _, mot := range mots {
		if estPonctuation(mot) {
			if len(resultat) > 0 {
				resultat[len(resultat)-1] += mot
			}
		} else {
			resultat = append(resultat, mot)
		}
	}

	return strings.Join(resultat, " ")
}

func estPonctuation(mot string) bool {
	for _, caractere := range mot {
		if !strings.ContainsRune(".,!?:;", caractere) {
			return false
		}
	}

	return len(mot) > 0
}