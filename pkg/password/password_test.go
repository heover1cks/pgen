package password

import "testing"

func TestGeneratePass(t *testing.T) {
	opts := NewOptions()
	opts.WithPassLength(16)
	opts.WithUpperCase(true)
	opts.WithMaxSpecialCharacters(2)
	println(opts.GeneratePassword())
}
