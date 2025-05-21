package domain

type User struct {
	Id        int
	Email     string
	UserName  string
	RoleId    int
	Role      Role
	CreatedAt string
	UpdatedAt string
}
