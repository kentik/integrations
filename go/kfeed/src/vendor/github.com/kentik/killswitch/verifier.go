package killswitch

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

type rsaPublicKey struct {
	*rsa.PublicKey
}

func (r *rsaPublicKey) Verify(message []byte, sig []byte) error {
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, d, sig)
}

func buildPublicKeyVerifier() (Verifier, error) {
	return parsePublicKey([]byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt28iO1cxr7f5fn83W46/
TDFAvzEV3X2Qe3tJfMZfatj95jT7g+0L679q1G+IxjBj9/N2Cu6l8YP/L9RZC1Nx
aoMrz2pX7FUJ2eEC8ueXCQcyjonTaY/qt4xMCx6Gksyh4jILQKE9fTbJ7Wo9tKYS
q5q6XD0Wx9eGtDRHA2LF9vzxiIhT/3/44GXUVAHEvK3CsbLRuB+hw60gkrSinm/y
DzOADppwd9bmODnj+QTzGkEDS9H3elQArF3UiGRWfVESVIFGiwo8vBUeRTbCpkix
StbqgrlsW/N9YVwFZpDLLznabeHu2vqvTSUkkTT1kF67ZWH52UqXdwFjI8Yp/ny1
NwIDAQAB
-----END PUBLIC KEY-----`))
}

// parsePublicKey parses a PEM encoded public key.
func parsePublicKey(pemBytes []byte) (Verifier, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsa, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	return newVerifierFromKey(rawkey)
}

func newVerifierFromKey(k interface{}) (Verifier, error) {
	var sshKey Verifier
	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}
