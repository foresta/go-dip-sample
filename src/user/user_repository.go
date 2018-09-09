package user

type Repository interface {
	Store(u *User) error
	FindAll() []*User
}
