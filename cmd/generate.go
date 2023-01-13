package cmd

import (
	"github.com/spf13/cobra"
	"idpassgen/pkg/passphrase"
	"idpassgen/pkg/password"
)

var includeUpperCase, includeLowerCase, includeNumbers, includeSpecialCharacters bool
var maxNumbers, maxSpecialCharacters, passwordLength, wordsCount int
var separator string

func init() {
	rootCmd.AddCommand(passwordGenCommand)
	rootCmd.AddCommand(passphraseGenCommand)

	passwordGenCommand.Flags().BoolVarP(&includeUpperCase, "upper", "u", true, "whether include uppercase or not")
	passwordGenCommand.Flags().BoolVarP(&includeLowerCase, "lower", "l", true, "whether include lowercase or not")
	passwordGenCommand.Flags().BoolVarP(&includeNumbers, "num", "n", true, "whether include numbers or not")
	passwordGenCommand.Flags().BoolVarP(&includeSpecialCharacters, "special", "s", true, "whether include special characters or not")
	passwordGenCommand.Flags().IntVarP(&maxNumbers, "num.max", "N", 3, "maximum numbers")
	passwordGenCommand.Flags().IntVarP(&maxSpecialCharacters, "special.max", "S", 3, "maximum special characters")
	passwordGenCommand.Flags().IntVarP(&passwordLength, "length", "L", 6, "password length")

	passphraseGenCommand.Flags().BoolVarP(&includeNumbers, "num", "n", false, "whether include numbers or not")
	passphraseGenCommand.Flags().StringVarP(&separator, "separator", "s", "-", "passphrase separator")
	passphraseGenCommand.Flags().IntVarP(&wordsCount, "word.count", "c", 3, "words count")

}

var passwordGenCommand = &cobra.Command{
	Use:   "password",
	Short: "generate password",
	Long:  "generate password",
	Run: func(cmd *cobra.Command, args []string) {
		opts := password.NewOptions()
		opts.WithPassLength(passwordLength)
		opts.WithUpperCase(includeUpperCase)
		opts.WithLowerCase(includeLowerCase)
		opts.WithNumbers(includeNumbers)
		opts.WithSpecialCharacters(includeSpecialCharacters)
		opts.WithMaxNumbers(maxNumbers)
		opts.WithMaxSpecialCharacters(maxSpecialCharacters)
		println(opts.GeneratePassword())
	},
}

var passphraseGenCommand = &cobra.Command{
	Use:   "passphrase",
	Short: "generate passphrase",
	Long:  "generate passphrase",
	Run: func(cmd *cobra.Command, args []string) {
		opts := passphrase.NewOptions()
		opts.WithWordsCount(wordsCount)
		opts.WithSeparator(separator)
		opts.WithNumber(includeNumbers)
		println(opts.GeneratePassPhrase())
	},
}
