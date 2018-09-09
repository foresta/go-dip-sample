// user.go

package user

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}
