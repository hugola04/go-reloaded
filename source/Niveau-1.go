package source

import (
	"strconv"
	"strings"
)

func Niveauun(texte string) string {
	mots := strings.Fields(texte)
	resultat := []string{}

	for i := 0; i < len(mots); i++ {
		if mots[i] == "(hex)" && len(resultat) > 0 {
			index := len(resultat) - 1

			nombre, err:= strconv.ParseInt(resultat[index], 16, 64)

			if err == nil {
				resultat[index] = strconv.FormatInt(nombre, 10)
			}

			continue
		}

		if mots[i] == "(bin)" && len(resultat) > 0 {
			index := len(resultat) - 1

			nombre, err := strconv.ParseInt(resultat[index], 2, 64)

			if err == nil {
				resultat[index] = strconv.FormatInt(nombre, 10)
			}

			continue
		}

		resultat = append(resultat, mots[i])
	}

	return strings.Join(resultat, " ")
}