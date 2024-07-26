package titlecase

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToTitleCase converts a string to title case using the English language. It will
// lowercase small words like "a", "an", "and", "as", "at", "but", "by", "en", "for", etc.
func ToTitleCase(input string) string {
	words := strings.Fields(input)
	caser := cases.Title(language.English)

	smallwords := " a an and as at but by en for if in of on or the to v v. via vs vs. "

	for i, word := range words {
		if strings.Contains(smallwords, " "+strings.ToLower(word)+" ") && i != 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = caser.String(word)
		}
	}
	return strings.Join(words, " ")
}
