package postgre

const (
	createUserQuery = `insert into users (username) values ($1);`
)
