package usecase

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	users := &usersMock{
		users: []user{
			{
				ID:       "user1",
				LoginID:  "login1",
				Password: "pass1",
				Name:     "name1",
			},
			{
				ID:       "user2",
				LoginID:  "login2",
				Password: "pass2",
				Name:     "name2",
			},
		},
	}

	_, err := Register(
		&RegisterDependencies{
			Users: users,
		},
		&domain.ProvisionalUser{
			LoginID:  "newLogin",
			Password: "newPass",
			Name:     "newName",
		},
	)

	if err != nil {
		t.Fatalf("failed to register user. %s", err.Error())
	}

	if !reflect.DeepEqual(users.users, []user{
		{
			ID:       "user1",
			LoginID:  "login1",
			Password: "pass1",
			Name:     "name1",
		},
		{
			ID:       "user2",
			LoginID:  "login2",
			Password: "pass2",
			Name:     "name2",
		},
		{
			ID:       "newId",
			LoginID:  "newLogin",
			Password: "newPass",
			Name:     "newName",
		},
	}) {
		t.Error("users are not same")
	}
}
