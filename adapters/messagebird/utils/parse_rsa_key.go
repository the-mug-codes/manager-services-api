package messagebird_utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func ParseKeycloakRSAPublicKey(base64String string) (publicKey *rsa.PublicKey, err error) {
	buffer, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return publicKey, err
	}
	parsedKey, err := x509.ParsePKIXPublicKey(buffer)
	if err != nil {
		return publicKey, err
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, err
	}
	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}
