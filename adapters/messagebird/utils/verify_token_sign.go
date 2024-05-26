package messagebird_utils

import "github.com/golang-jwt/jwt"

func VerifyTokenSign(accessToken string, publicKey string) (parsedToken *jwt.Token, err error) {
	parsedToken, err = jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return ParseKeycloakRSAPublicKey(publicKey)
	})
	return parsedToken, err
}
