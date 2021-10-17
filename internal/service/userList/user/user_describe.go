package user

import (
	"fmt"

	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (u *DummyUserService) Describe(User_id uint64) (*userlist.User, error) {
	if len(u.Users) < int(User_id) {
		return nil, fmt.Errorf("no user with such index")
	}

	return &u.Users[User_id-1], nil
}
