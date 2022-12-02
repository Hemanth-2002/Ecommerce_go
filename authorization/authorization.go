package authorization

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secretcode")

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error occured")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprint(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not authorized")
		}
	})
}
