# encdec

CLI and packages to encrypt / decrypt files.

## CLI usage

### Install

```sh
go install github.com/cassioconti/encdec/cmd/encdec@latest
```

### Consume

```sh
SECRET=abcdefghijklmnopqrstuvwxyz123456

# Encrypt - will output the encrypted file to test/data/my-test-file.json.enc
encdec encrypt test/data/my-test-file.json $SECRET

# Decrypt - will output the decrypted file to test/data/my-test-file.json
encdec decrypt test/data/my-test-file.json.enc $SECRET
```

## Package usage

### Import

```sh
go get -u github.com/cassioconti/encdec/pkg/encdec
```

### Consume

```golang
import "github.com/cassioconti/encdec/pkg/encdec"

func myFunc() {
    secret := "abcdefghijklmnopqrstuvwxyz123456"
	contentOriginal := []byte("This is an example of the content. Content can be any []byte.")
	encDec := encdec.NewEncoderDecoder()
	contentEncrypted, err := encDec.Encrypt(contentOriginal, secret)
	if err != nil {
		// handle error
	}

	contentDecrypted, err := encDec.Decrypt(contentEncrypted, secret)
	if err != nil {
		// handle error
	}

	fmt.Printf("%s\n", contentDecrypted)
}
```
