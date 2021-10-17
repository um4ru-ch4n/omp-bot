package user

import (
	"fmt"

	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (u *DummyUserService) List(cursor uint64, limit uint64) ([]userlist.User, error) {
	if cursor > uint64(len(u.Users)-1) {
		return nil, fmt.Errorf("not enough users")
	}

	if cursor+limit >= uint64(len(u.Users)) {
		return u.Users[cursor:], nil
	}

	return u.Users[cursor : cursor+limit], nil
}
