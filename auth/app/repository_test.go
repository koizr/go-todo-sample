package app

import (
	"github.com/google/uuid"
	"github.com/koizr/go-todo-sample/auth/domain"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewUser(t *testing.T) {
	provisionalUser := &domain.ProvisionalUser{
		LoginID:  "test-login-id",
		Password: "test_password",
		Name:     "Taro",
	}

	user, err := newUser(provisionalUser)
	if err != nil {
		t.Fatalf("failed to create User. %s", err.Error())
	}

	id, err := uuid.Parse(user.ID)
	if err != nil {
		t.Errorf("user.ID is not UUID. user.ID=%s, err=%s", user.ID, err.Error())
	} else if id.Version() != 4 {
		t.Errorf("user.ID's UUID version is not v4. version=%d", id.Version())
	}

	if user.LoginID != provisionalUser.LoginID {
		t.Errorf("user.LoginID has difference from provisionalUser.LoginID. %s not equals %s", user.LoginID, provisionalUser.LoginID)
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(provisionalUser.Password)) != nil {
		t.Errorf("user.Password is not hash of provisionalUser.Password. hash=%s, rawpassword=%s", user.Password, provisionalUser.Password)
	}
	if user.Name != provisionalUser.Name {
		t.Errorf("user.Name has difference from provisionalUser.Name. %s not equals %s", user.Name, provisionalUser.Name)
	}
}
