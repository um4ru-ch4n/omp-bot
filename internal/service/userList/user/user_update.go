package user

import (
	"fmt"

	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (u *DummyUserService) Update(userID uint64, user userlist.User) error {
	if len(u.Users) < int(userID) {
		return fmt.Errorf("no user with such index")
	}

	u.Users[userID-1] = user

	return nil
}
