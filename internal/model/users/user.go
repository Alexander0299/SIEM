package users

type User struct {
	id       int
	username string
	role     string
}

func NewUser(id int, username, role string) *User {
	return &User{id, username, role}
}

func (u *User) GetRole() string {
	return u.role
}

func (u *User) SetRole(newRole string) {
	u.role = newRole
}
