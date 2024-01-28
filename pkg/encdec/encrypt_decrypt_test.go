package encdec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncDec(t *testing.T) {
	secret := "abcdefghijklmnopqrstuvwxyz123456"
	contentOriginal := []byte("This is an example of the content. Content can be any []byte.")
	encDec := NewEncoderDecoder()
	contentEncrypted, err := encDec.Encrypt(contentOriginal, secret)
	assert.Nil(t, err)
	contentDecrypted, err := encDec.Decrypt(contentEncrypted, secret)
	assert.Nil(t, err)
	assert.Equal(t, contentOriginal, contentDecrypted)
}
