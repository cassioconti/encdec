package encdec

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "encdec",
	Short: "encdec is a encrypt/decrypt CLI tool",
	Long:  `A fast encrypt/decrypt tool allowing secret management.`,
}

func Execute() error {
	return rootCmd.Execute()
}
