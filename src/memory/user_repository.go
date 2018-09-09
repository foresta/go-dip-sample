package memory

import (
	"fmt"
	"sync"

	"github.com/foresta/go-dip-sample/src/user"
)

type userRepository struct {
	mtx   sync.RWMutex
	users map[int]*user.User
}

func NewUserRepository() user.Repository {
	return &userRepository{
		users: make(map[int]*user.User),
	}
}

func (repo *userRepository) Store(u *user.User) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	// id
	id := len(repo.users) + 1
	u.ID = id
	repo.users[u.ID] = u

	fmt.Printf("user is created: %d", u.ID)

	return nil
}

func (repo *userRepository) FindAll() []*user.User {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	c := make([]*user.User, 0, len(repo.users))
	for _, val := range repo.users {
		c = append(c, val)
	}
	return c
}
