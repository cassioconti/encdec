package encdec

import (
	"log"
	"os"
	"strings"

	"github.com/cassioconti/encdec/pkg/encdec"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Action to encrypt a file",
	Long: `Action to encrypt a file. It takes two arguments:
	- file: The file to be encrypted
	- secret: The secret (minimum 32 characters long)`,
	Args:    cobra.ExactArgs(2),
	Example: "encdec encrypt <file> <secret>",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("Failed to load file %s: %v\n", args[0], err)
		}

		secret := strings.TrimSpace(args[1])
		if len(secret) < 32 {
			log.Fatalf("Secret must have at least 32 characters\n")
		}

		contentEnc, err := encdec.NewEncoderDecoder().Encrypt(content, secret)
		if err != nil {
			log.Fatalf("Failed to encrypt file %s: %v\n", args[0], err)
		}

		outputFileName := args[0] + ".enc"
		err = os.WriteFile(outputFileName, contentEnc, 0666)
		if err != nil {
			log.Fatalf("Failed to write file %s: %v\n", outputFileName, err)
		}
	},
}
