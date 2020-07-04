package controller

import (
	"fmt"
	"go_app/consts"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var secretkey = consts.Consts.SecretKey

// generae JWT when user login
func GenerateJWT(w http.ResponseWriter, userId string) *JSONError {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}
	tokenString, err := token.SignedString([]byte(secretkey))
	// ログイン照会後の処理なので500エラーとする
	if err != nil {
		return &JSONError{err, "認証できませんでした", http.StatusInternalServerError}
	}

	cookieId := &http.Cookie{
		Name:  "isOnSession",
		Value: tokenString,
	}

	http.SetCookie(w, cookieId)
	return nil
}

// return JWT -> map or error
func AuthByJwt(r *http.Request) (map[string]interface{}, *JSONError) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretkey)
		return b, nil
	})
	if err != nil {
		return nil, &JSONError{err, fmt.Sprintf("認証されておりません。%v", err), http.StatusUnauthorized}
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
