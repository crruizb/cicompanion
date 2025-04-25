package data

import (
	"context"
	"database/sql"
	"log/slog"
	"time"
)

type UsersStore struct {
	DB *sql.DB
}

func NewUsersStore(db *sql.DB) *UsersStore {
	return &UsersStore{
		DB: db,
	}
}

type User struct {
	Id        string `db:"id"`
	Username  string `json:"login" db:"username"`
	GithubPAT string `json:"githubPat" db:"github_pat"`
}

func (m UsersStore) GetUser(username string) (*User, error) {
	query := `
		SELECT id, username, github_pat FROM users
		WHERE username = $1
	`

	args := []any{username}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := &User{}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Username, &user.GithubPAT)
	if err != nil {
		slog.Error("error finding user", "err", err.Error())
		return nil, err
	}

	return user, nil
}

func (m UsersStore) InsertUser(username, githubPat string) (*User, error) {
	query := `
		INSERT INTO users (username, github_pat)
		VALUES ($1, $2)
		RETURNING id
	`

	args := []any{username, githubPat}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := &User{
		Username:  username,
		GithubPAT: githubPat,
	}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		slog.Error("error creating user", "err", err.Error())
		return nil, err
	}

	return user, nil
}
