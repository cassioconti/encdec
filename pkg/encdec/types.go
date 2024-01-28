package encdec

type EncryptDecrypt interface {
	Encrypt(content []byte, secret string) ([]byte, error)
	Decrypt(encryptedContent []byte, secret string) ([]byte, error)
}
