package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/koizr/go-todo-sample/auth/domain"
	"time"
)

const (
	subject  = "sub"
	expire   = "exp"
	issuedAt = "iat"
)

func GenerateToken(secret string, user *domain.User, now *time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		subject:  user.ID,
		issuedAt: now.Unix(),
		expire:   now.Add(time.Minute * 5).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ParseToken(token *jwt.Token, now *time.Time) (string, error) {
	claims := token.Claims.(jwt.MapClaims)

	// 期限切れならエラー
	exp, ok := claims[expire].(int64)
	if !ok {
		return "", errors.New("expire does not exist in claims")
	}
	if exp < now.Unix() {
		return "", errors.New("token is expired")
	}

	// ユーザーID を取れなかったらエラー
	userID, ok := claims[subject].(string)
	if !ok {
		return "", errors.New("UserID does not exist in claims")
	}

	return userID, nil
}
