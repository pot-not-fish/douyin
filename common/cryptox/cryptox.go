package cryptox

import (
	"encoding/base64"

	"github.com/forgoer/openssl"
)

var (
	secretKey = []byte("123456")
)

func EncryptoPassword(password string) (string, error) {
	res, err := openssl.DesCBCEncrypt([]byte(password), secretKey, secretKey, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(res), nil
}
