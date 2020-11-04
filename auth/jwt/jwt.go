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

type Token = jwt.Token

func GenerateToken(secret string, user *domain.User, now *time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		subject:  user.ID,
		issuedAt: now.Unix(),
		expire:   now.Add(time.Minute * 10).Unix(), // TODO: 期限の扱いを変更する。環境変数から取るか、適当にもう少し伸ばすか
	})
	return token.SignedString([]byte(secret))
}

func ParseToken(token *Token) (domain.UserID, error) {
	claims := token.Claims.(jwt.MapClaims)

	// ユーザーID を取れなかったらエラー
	userID, ok := claims[subject].(domain.UserID)
	if !ok {
		return "", errors.New("UserID does not exist in claims")
	}

	return userID, nil
}
