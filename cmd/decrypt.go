package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strings"

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

		content, err := decrypt([]byte(secret), contentEnc)
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

func decrypt(secret, content []byte) ([]byte, error) {
	myCipher, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	if len(content) < aes.BlockSize {
		return nil, errors.New("the content to be decrypted is too short")
	}

	iv := content[:aes.BlockSize]
	content = content[aes.BlockSize:]
	cfbDec := cipher.NewCFBDecrypter(myCipher, iv)
	cfbDec.XORKeyStream(content, content)
	contentDecrypted, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		return nil, err
	}

	return contentDecrypted, nil
}
