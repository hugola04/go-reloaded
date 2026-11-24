package source

import (
	"strconv"
	"strings"
	"unicode"
)

func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}
	return string(runes)
}

func estUnMot(s string) bool {
	if s == "" || s == "(up)" || s == "(low)" || s == "(cap)" || s == "(hex)" || s == "(bin)" {
		return false
	}
	for _, r := range s {
		if !strings.ContainsRune(".,!?:;'", r) {
			return true
		}
	}
	return false
}

func Niveaudeux(word []string) []string {
	for i := 0; i < len(word); i++ {

		// Cas (up, N), (low, N), (cap, N)
		if (word[i] == "(up," || word[i] == "(low," || word[i] == "(cap,") && i+1 < len(word) {
			action := word[i][1 : len(word[i])-1]

			if strings.HasSuffix(word[i+1], ")") {
				nettoye := strings.TrimSuffix(word[i+1], ")")
				count, err := strconv.Atoi(nettoye)

				if err == nil && count >= 0 {
					word[i] = ""
					word[i+1] = ""

					modifies := 0
					for j := i - 1; j >= 0 && modifies < count; j-- {
						if !estUnMot(word[j]) {
							continue
						}
						switch action {
						case "up":
							word[j] = strings.ToUpper(word[j])
						case "low":
							word[j] = strings.ToLower(word[j])
						case "cap":
							word[j] = capitalize(word[j])
						}
						modifies++
					}
				}
			}
		}

		// Cas (up), (low), (cap)
		if word[i] == "(up)" || word[i] == "(low)" || word[i] == "(cap)" {
			action := word[i][1 : len(word[i])-1]
			word[i] = ""

			for j := i - 1; j >= 0; j-- {
				if estUnMot(word[j]) {
					switch action {
					case "up":
						word[j] = strings.ToUpper(word[j])
					case "low":
						word[j] = strings.ToLower(word[j])
					case "cap":
						word[j] = capitalize(word[j])
					}
					break
				}
			}
		}
	}

	var resultat []string
	for _, w := range word {
		if w != "" {
			resultat = append(resultat, w)
		}
	}
	return resultat
}
