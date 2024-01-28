package encdec

import (
	"log"
	"os"
	"strings"

	"github.com/cassioconti/encdec/pkg/encdec"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Action to decrypt a file",
	Long: `Action to decrypt a file. It takes two arguments:
	- file: The file to be decrypted
	- secret: The secret used to encrypt`,
	Args:    cobra.ExactArgs(2),
	Example: "encdec decrypt <file> <secret>",
	Run: func(cmd *cobra.Command, args []string) {
		contentEnc, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("Failed to load file %s: %v\n", args[0], err)
		}

		secret := strings.TrimSpace(args[1])
		if len(secret) < 32 {
			log.Fatalf("Secret must have at least 32 characters\n")
		}

		content, err := encdec.NewEncoderDecoder().Decrypt(contentEnc, secret)
		if err != nil {
			log.Fatalf("Failed to decrypt file %s: %v\n", args[0], err)
		}

		outputFileName := strings.TrimSuffix(args[0], ".enc")
		if args[0] == outputFileName {
			outputFileName += ".dec"
		}

		err = os.WriteFile(outputFileName, content, 0666)
		if err != nil {
			log.Fatalf("Failed to write file %s: %v\n", outputFileName, err)
		}
	},
}
