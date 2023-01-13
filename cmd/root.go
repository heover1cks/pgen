package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"idpassgen/pkg/version"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "pgen",
	Short:   "password & passphrase generator",
	Long:    "notyet",
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
