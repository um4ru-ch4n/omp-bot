package user

import (
	"fmt"

	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (u *DummyUserService) Update(User_id uint64, user userlist.User) error {
	if len(u.Users) < int(User_id) {
		return fmt.Errorf("no user with such index")
	}

	u.Users[User_id-1] = user

	return nil
}
