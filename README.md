# encdec

CLI to encrypt / decrypt files.

## Install

```sh
go install github.com/cassioconti/encdec
```

## Usage

```sh
SECRET=abcdefghijklmnopqrstuvwxyz123456

# Encrypt - will output the encrypted file to test/data/my-test-file.json.enc
encdec encrypt test/data/my-test-file.json $SECRET

# Decrypt - will output the decrypted file to test/data/my-test-file.json
encdec decrypt test/data/my-test-file.json.enc $SECRET
```
