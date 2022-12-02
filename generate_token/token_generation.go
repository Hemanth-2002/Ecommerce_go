package generatetoken

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secretcode")

func GenerateJWt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "KHK"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenstring, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("error : %s", err.Error())
	}
	fmt.Printf("token is : %s", tokenstring)
	fmt.Println()
	return tokenstring, nil
}
