package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type encoderDecoder struct{}

func NewEncoderDecoder() EncryptDecrypt {
	return &encoderDecoder{}
}

// Encode implements EncoderDecoder.
func (*encoderDecoder) Encrypt(content []byte, secret string) ([]byte, error) {
	myCipher, err := aes.NewCipher([]byte(secret))
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

// Decode implements EncoderDecoder.
func (*encoderDecoder) Decrypt(encryptedContent []byte, secret string) ([]byte, error) {
	myCipher, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	if len(encryptedContent) < aes.BlockSize {
		return nil, errors.New("the content to be decrypted is too short")
	}

	iv := encryptedContent[:aes.BlockSize]
	encryptedContent = encryptedContent[aes.BlockSize:]
	cfbDec := cipher.NewCFBDecrypter(myCipher, iv)
	cfbDec.XORKeyStream(encryptedContent, encryptedContent)
	contentDecrypted, err := base64.StdEncoding.DecodeString(string(encryptedContent))
	if err != nil {
		return nil, err
	}

	return contentDecrypted, nil
}
