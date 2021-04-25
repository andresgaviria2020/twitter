package utilidades

import (
	"time"
	"twitter/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("AndrésGaviria2021")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
