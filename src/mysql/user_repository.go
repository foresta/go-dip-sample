package mysql

import "github.com/foresta/go-dip-sample/src/user"

type UserRepository struct {
}

func (repo *UserRepsitory) Store(u *user.User) {
	// TODO::
}

func (repo *UserRepository) FindAll() []*user.User {
	// TODO::
}
