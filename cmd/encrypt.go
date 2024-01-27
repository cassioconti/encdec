package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strings"

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

		contentEnc, err := encrypt([]byte(secret), content)
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

func encrypt(secret, content []byte) ([]byte, error) {
	myCipher, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	contentB64 := base64.StdEncoding.EncodeToString(content)
	contentEncrypted := make([]byte, aes.BlockSize+len(contentB64))
	iv := contentEncrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfbEnc := cipher.NewCFBEncrypter(myCipher, iv)
	cfbEnc.XORKeyStream(contentEncrypted[aes.BlockSize:], []byte(contentB64))
	return contentEncrypted, nil
}
