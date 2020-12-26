package traderdb

import "context"

type User struct {
	Id             int    `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
	IsActive       bool   `json:"isActive"`
}

const insertUserQuery = `
INSERT INTO users(email, password, is_active)
VALUES ($1, $2, $3)
`

func InsertNewUser(db PGConnection, user User) error {
	if _, err := db.Exec(
		context.Background(),
		insertUserQuery,
		user.Email,
		user.HashedPassword,
		user.IsActive,
	); err != nil {
		return err
	}
	return nil
}

const selectUserByEmailQuery = `
SELECT id, password, is_active FROM users WHERE email = $1
`

func GetUserByEmail(db PGConnection, email string) (user User, err error) {
	var id int
	var hashedPassword string
	var isActive bool
	row := db.QueryRow(context.Background(), selectUserByEmailQuery, email)
	if err = row.Scan(&id, &hashedPassword, &isActive); err != nil {
		return user, err
	}
	user.Id = id
	user.Email = email
	user.HashedPassword = hashedPassword
	user.IsActive = isActive
	return user, nil
}
