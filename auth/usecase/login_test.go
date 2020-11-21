package usecase

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	form := &LoginForm{
		LoginID:  "test2",
		Password: "xyz",
	}
	dep := &depMock{
		users: &usersMock{
			users: []user{
				{
					ID:       "1b8eb725-b5ce-4cf9-8902-bdc747cca17e",
					LoginID:  "test1",
					Password: "abc",
				},
				{
					ID:       "a3773913-b0ea-47b8-b2d8-a85807dfdb33",
					LoginID:  "test2",
					Password: "xyz",
				},
				{
					ID:       "f0004276-476f-412b-acdc-ce1df60b2de4",
					LoginID:  "test3",
					Password: "1234",
				},
			},
		},
		secret:  "test",
		now:     time.Date(2020, 10, 1, 12, 20, 30, 5, time.UTC),
		authExp: time.Minute * 5,
	}

	token, err := Login(dep, form)
	if err != nil {
		t.Fatalf("login is failed. %s", err.Error())
	}
	//goland:noinspection SpellCheckingInspection
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDE1NTUxMzAsImlhdCI6MTYwMTU1NDgzMCwic3ViIjoiYTM3NzM5MTMtYjBlYS00N2I4LWIyZDgtYTg1ODA3ZGZkYjMzIn0.PjL0SuK0nUQWq2mt5ezAgOPH8rkjl13P1nO_u8IEINg" {
		t.Errorf("got unexpected token. token=%s", token)
	}
}

func TestLoginUserNotFound(t *testing.T) {
	dep := &depMock{
		users: &usersMock{
			users: []user{
				{
					ID:       "1b8eb725-b5ce-4cf9-8902-bdc747cca17e",
					LoginID:  "test1",
					Password: "abc",
				},
				{
					ID:       "a3773913-b0ea-47b8-b2d8-a85807dfdb33",
					LoginID:  "test2",
					Password: "xyz",
				},
				{
					ID:       "f0004276-476f-412b-acdc-ce1df60b2de4",
					LoginID:  "test3",
					Password: "1234",
				},
			},
		},
		secret:  "test",
		now:     time.Date(2020, 10, 1, 12, 20, 30, 5, time.UTC),
		authExp: time.Minute * 5,
	}

	_, err := Login(dep, &LoginForm{
		LoginID:  "not exists",
		Password: "xyz",
	})
	if err == nil {
		t.Errorf("expect failed to Login with LoginID that doesn't exist. but it suceeded.")
	}

	_, err = Login(dep, &LoginForm{
		LoginID:  "test2",
		Password: "not exists",
	})
	if err == nil {
		t.Errorf("expect failed to Login with Password that doesn't exist. but it suceeded.")
	}

	_, err = Login(dep, &LoginForm{
		LoginID:  "not exists",
		Password: "not exists",
	})
	if err == nil {
		t.Errorf("expect failed to Login with LoginID and Password that don't exist. but it suceeded.")
	}
}

type depMock struct {
	users   domain.Users
	secret  string
	now     time.Time
	authExp time.Duration
}

func (d *depMock) Users() domain.Users {
	return d.users
}

func (d depMock) Secret() string {
	return d.secret
}

func (d depMock) Now() *time.Time {
	return &d.now
}

func (d depMock) AuthenticationExpire() time.Duration {
	return d.authExp
}
