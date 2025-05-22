package domain

type Todo struct {
	Id        int
	Title     string
	Status    string
	UserId    int
	User      User
	CreatedAt string
	UpdatedAt string
}
