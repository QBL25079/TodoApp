package users_postgres_repo

type UserModel struct {
	ID          int
	Version    int
	FullName    string
	PhoneNumber *string
}
