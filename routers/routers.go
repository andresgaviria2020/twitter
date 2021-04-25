package routers

import (
	"errors"
	"strings"
	"twitter/bd"
	"twitter/models"

	"github.com/dgrijalva/jwt-go"
)

var Email string

var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("AndrésGaviria2021")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Error de token")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims,
		func(token *jwt.Token) (interface{}, error) {
			return miClave, nil
		})

	if err != nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.Id
		}
		return claims, encontrado, ID, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}

	return claims, false, "", err
}
