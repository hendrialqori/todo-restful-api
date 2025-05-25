package domain

type User struct {
	Id        int
	Email     string
	UserName  string
	RoleId    int
	Role      Role
	CreatedAt string
	// UpdatedAt with pointer so when data is null, Scan still able to convert to string. or using sql.NullString
	UpdatedAt *string
}
