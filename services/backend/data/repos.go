package data

import (
	"context"
	"database/sql"
	"log/slog"
	"time"
)

type ReposPostgresStore struct {
	DB *sql.DB
}

func NewReposPostgresStore(db *sql.DB) *ReposPostgresStore {
	return &ReposPostgresStore{
		DB: db,
	}
}

type Repo struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (r *ReposPostgresStore) AddRepo(repo Repo, userId string) error {
	query := `
		INSERT INTO repos (id, name, user_id)
		VALUES ($1, $2, $3)
	`

	args := []any{repo.Id, repo.Name, userId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Err()
	if err != nil {
		slog.Error("error creating repo", "err", err.Error())
		return err
	}

	return nil
}

func (r *ReposPostgresStore) GetRepos(userId string) ([]Repo, error) {
	query := `
		SELECT id, name
		FROM repos
		WHERE user_id = $1
	`

	args := []any{userId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	repos := []Repo{}

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		slog.Error("error getting repos", "err", err.Error())
		return nil, err
	}

	for rows.Next() {
		repo := &Repo{}
		err := rows.Scan(&repo.Id, &repo.Name)
		if err != nil {
			slog.Error("error getting repo", "err", err.Error())
			return nil, err
		}

		repos = append(repos, *repo)

	}

	return repos, nil
}

func (r *ReposPostgresStore) Deleterepo(repoId int, userId string) error {
	query := `
		DELETE FROM repos WHERE id = $1 and user_id = $2;
	`

	args := []any{repoId, userId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		slog.Error("error deleting repo", "err", err.Error())
		return err
	}

	return nil
}
