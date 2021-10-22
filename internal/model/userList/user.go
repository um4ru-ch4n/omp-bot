package entitylist

import "fmt"

type User struct {
	Username string
}

func (u *User) String() string {
	return fmt.Sprintf("Username: %s", u.Username)
}
