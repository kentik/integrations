package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"

	"github.com/kentik/killswitch"
)

var (
	kRS1 = []byte(`*hash<net.hostLookupOrder,string`)
)

// parsePrivateKey parses a PEM encoded private key.
func parsePrivateKey(pemBytes []byte) (killswitch.Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return newSignerFromKey(rawkey)
}

func newSignerFromKey(k interface{}) (killswitch.Signer, error) {
	var sshKey killswitch.Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha256
func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

func showPrivateKey() ([]byte, error) {
	pk, err := encrypt(kRS1, genPrivateKey())
	if err != nil {
		return nil, err
	} else {
		return pk, err
	}
}

// buildPrivateKeySigner returns a signer with our private key
func buildPrivateKeySigner() (killswitch.Signer, error) {
	data, err := base64.StdEncoding.DecodeString(string(genPrivateKeyFull()))
	if err != nil {
		return nil, err
	} else {
		raw, err := decrypt(kRS1, data)
		if err != nil {
			return nil, err
		} else {
			return parsePrivateKey(raw)
		}
	}
}

func genPrivateKey() []byte {
	return []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAt28iO1cxr7f5fn83W46/TDFAvzEV3X2Qe3tJfMZfatj95jT7
g+0L679q1G+IxjBj9/N2Cu6l8YP/L9RZC1NxaoMrz2pX7FUJ2eEC8ueXCQcyjonT
aY/qt4xMCx6Gksyh4jILQKE9fTbJ7Wo9tKYSq5q6XD0Wx9eGtDRHA2LF9vzxiIhT
/3/44GXUVAHEvK3CsbLRuB+hw60gkrSinm/yDzOADppwd9bmODnj+QTzGkEDS9H3
elQArF3UiGRWfVESVIFGiwo8vBUeRTbCpkixStbqgrlsW/N9YVwFZpDLLznabeHu
2vqvTSUkkTT1kF67ZWH52UqXdwFjI8Yp/ny1NwIDAQABAoIBAQCdzSo6uGETFIa6
rsA1sJCbAEf98kEoIMv5nm7yu0j8hx2tO+kfbf5yWWKSzDxymtB1Tuk+sFzrby1J
vTi5CJiqE5vNvGNU+TcXS34Y7ug8qQdyHdlUl61JJ5WHf0Qv54BgMcMCX3OhU0/V
QS9CSBgJrnshvJ/rXVsRjWOF4yQAu9Y/rQDr0DvUysxtJWNgoqiE1Iz5kKeU9I2z
RbmObsvX+IZLUvvPbChqFcmmZrw7EC9CSg6g5prFSljlpVNtbOZT5S6wRTjDVyo3
bBV89w+d8RqyUJmkGVEm42NumpeSwuPqN02izr50o/uI4AqHEK7KuUs+FAsPTKyB
lDsz3etZAoGBANvBajnIbTegYwkYcWPERwaL3U/OwgIr8/Ml3unVvyZJHpi+SM77
XezHqF4OPNQ9XG9CBPNV/Y2tenvi0AfGxS9+o+Ft0zL7X4OTbyTxJu+e4gqaSEgC
KEDdLXAeT73jb3ygSuI7sIggWTT0t0tT3ZNWrW5I6g7u24vukZwOK+bVAoGBANWw
JSIbkc4i7MU/j4XTudaN5JMLCWNk1KHO/TswAg9CN9anlKtHVmN8krkagdrcvuXp
+5TXLlxiKlSE8uh1KwCYmEmWG2BCdKNS2ufEH2yso4+8SpnK8G7CGwU14eDLrVkp
eX/GGOdGy2+Ko4ctWMIvXuKTX2XERbxDHoKmbcnbAoGBAJ0KoVDdzD7+XQe48f8r
2t3wwZ0w0jAfHNxb6esNFubTRgw9n2Za+AonvEhKwGmj2BSiB0ul3eaLXIZ/1c5Y
271PMOn+Q/mg7ebnS3wI4ZxH3J1bF+BtujpwVPJUlwPKLnfPTPRTV5pQE6/mWb55
FlUekh3H+YvmYfqj6GavMexpAoGAJa3dnloGJ7b40P8YK5zd0/tJJrR3f1M0OyPo
extIAbDHb8405M67aOd6Z7FI3HK4JopPljsrLZcRp72Zp2uhnYVKtQ8G3L5bGsFt
YBixAdSfMqUc8mlaY+1OHmkV3zGK64HorqMbmQxeqthjZV6VnAgGTyV1WNh3A29C
Xf9CtKkCgYABiddvX6xStWsXtqyCExtS/Tgx33DQLjx8cgUUzN03jJ1AnlZoZFPd
OkDG1epjqsjury2ZMX7Vzxwi7A7eEFSy+lu4umJARJnMK1YDej1YtxYFpStfPaIY
Ltbh5c+A5bkMRaaGDW7hj18vpCwLWfBhEp61xVUk/JQaifm/aZdkWQ==
-----END RSA PRIVATE KEY-----`)
}

func genPrivateKeyFull() []byte {
	return []byte(`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`)
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
