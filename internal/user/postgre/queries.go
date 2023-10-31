package postgre

const (
	createUserQuery  = `insert into users (id, family_name, given_name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7);`
	getUserListQuery = `select id, family_name, given_name, email, password, created_at, updated_at, deleted_at from users limit $1 offset $2;`
)
