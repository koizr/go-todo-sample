package domain

type UserID = string

type User struct {
	ID UserID
}

type Users interface {
	Find(loginID string, password string) (*User, error)
}

type UserNotFoundError struct {
}

func (u UserNotFoundError) Error() string {
	return "user not found"
}
