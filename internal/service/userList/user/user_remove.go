package user

import "fmt"

func (u *DummyUserService) Remove(User_id uint64) (bool, error) {
	if len(u.Users) < int(User_id) {
		return false, fmt.Errorf("no user with such index")
	}

	if User_id == uint64(len(u.Users)) {
		u.Users = u.Users[:User_id-1]
		return true, nil
	}

	u.Users = append(u.Users[:User_id-1], u.Users[User_id:]...)

	return true, nil
}
