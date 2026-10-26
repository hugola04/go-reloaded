package source

import (
	"strconv"
	"strings"
	"unicode"
)

func capitalize(word string) string {
	runes := []rune(word)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}
	return string(runes)
}

func Niveaudeux(word []string) []string {
	for i := 0; i < len(word); i++ {
		
		if (word[i] == "(up," || word[i] == "(low," || word[i] == "(cap,") && i+1 < len(word) {
			action := word[i][1 : len(word[i])-1]
			
			nettoye := strings.TrimRight(word[i+1], ".,!?:;)' ") 
			count, err := strconv.Atoi(nettoye)
			
			if err == nil {
				word[i] = "" 
				word[i+1] = ""

				modifies := 0
				for j := i - 1; j >= 0 && modifies < count; j-- {
					if word[j] == "" || word[j] == "," || word[j] == "." || word[j] == "!" || word[j] == "?" || word[j] == "'" {
						continue 
					}
					
					if action == "up" {
						word[j] = strings.ToUpper(word[j])
					} else if action == "low" {
						word[j] = strings.ToLower(word[j])
					} else if action == "cap" {
						word[j] = capitalize(word[j])
					}
					modifies++ 
				}
			}
		}

		if i > 0 {
			if word[i] == "(up)" {
				word[i-1] = strings.ToUpper(word[i-1])
				word[i] = ""
			} else if word[i] == "(low)" {
				word[i-1] = strings.ToLower(word[i-1])
				word[i] = ""
			} else if word[i] == "(cap)" {
				word[i-1] = capitalize(word[i-1])
				word[i] = ""
			}
		}
	}
	return word
}