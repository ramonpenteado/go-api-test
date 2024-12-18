package userDb

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetUser(id int) (User, error) {
	db, err := Connect()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", id)

	var user User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func CreateUser(user User) (User, error) {
	db, err := Connect()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, created_at", user.Username, user.Email)

	var createdUser User
	err = row.Scan(&createdUser.ID, &createdUser.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return createdUser, nil
}

func UpdateUser(user User) (User, error) {
	db, err := Connect()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	row := db.QueryRow("UPDATE users SET username = $1, email = $2 WHERE id = $3 RETURNING id, updated_at", user.Username, user.Email, user.ID)

	var updatedUser User
	err = row.Scan(&updatedUser.ID, &updatedUser.UpdatedAt)
	if err != nil {
		return User{}, err
	}
	return updatedUser, nil
}

func DeleteUser(id int) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
