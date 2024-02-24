package repo

const (
	createUserQuery = `insert into users (id, email, username, otp_secret, password, created_at, updated_at)
									values ($1, $2, $3, $4, $5, $6, $7);`
	getUserListQuery = `select
							id, email, username, otp_secret, password, created_at, updated_at deleted_at
						from users limit $1 offset $2;`
	getByLogin = `select
    						id, email, username, otp_secret, password, created_at, updated_at deleted_at
						from users where email = $1 or username = $1;`
	checkUserByLogin = `select count(*) from users where email = $1 or username = $1;`
)
