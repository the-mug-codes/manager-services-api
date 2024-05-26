package messagebird_utils

import (
	"errors"
	"strings"
)

func GetAccessTokenFromBearer(authorization string) (accessToken string, err error) {
	if len(authorization) == 0 {
		return accessToken, errors.New("authorization header not provided")
	}
	accessToken = strings.Split(authorization, "Bearer ")[0]
	if len(accessToken) == 0 {
		return accessToken, errors.New("invalid return accessToken")
	}
	return accessToken, err
}
