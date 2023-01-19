package words

import (
	"log"
	"os"
	"strings"

	_ "embed"
)

//go:embed words
var embeddedWords string

func RetrieveWordsFromOSDict() []string {
	wordBytes, err := os.ReadFile("/usr/share/dict/words")

	if err != nil {
		log.Printf("cannot open /usr/share/dict/words: %v", err)
		return RetrieveWordsFromEmbedWords()
	}

	return strings.Split(string(wordBytes), "\n")
}

func RetrieveWordsFromEmbedWords() []string {
	return strings.Split(embeddedWords, "\n")
}
