package domain

type UserID = string

type User struct {
	ID UserID
}

type Users interface {
	Find(loginID string, password string) (*User, error)
	Add(user *ProvisionalUser) (*User, error)
}

type UserNotFoundError struct {
}

func (u UserNotFoundError) Error() string {
	return "user not found"
}

// 登録前のユーザー
type ProvisionalUser struct {
	LoginID  string
	Password string
	Name     string
}

func Register(user *ProvisionalUser, users Users) (*User, error) {
	return users.Add(user)
}
