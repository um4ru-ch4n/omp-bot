package user

import (
	"fmt"

	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (u *DummyUserService) Describe(userID uint64) (*userlist.User, error) {
	if len(u.Users) < int(userID) {
		return nil, fmt.Errorf("no user with such index")
	}

	return &u.Users[userID-1], nil
}
