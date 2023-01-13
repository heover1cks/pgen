package words

import (
	"log"
	"os"
	"strings"
)

func RetrieveWordsFromOSDict() []string {
	wordBytes, err := os.ReadFile("/usr/share/dict/words")

	// TODO: use embed or options to use custom & default words list
	if err != nil {
		log.Fatalf("cannot open /usr/share/dict/words: %v", err)
	}

	return strings.Split(string(wordBytes), "\n")
}
