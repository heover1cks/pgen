package passphrase

import "testing"

func TestGenerateID(t *testing.T) {
	opts := NewOptions()
	opts.WithSeparator("-")
	opts.WithWordsCount(2)
	println(opts.GeneratePassPhrase())
}
