package data

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

type MonitoringPostgresStore struct {
	DB *sql.DB
}

func NewMonitoringPostgresStore(db *sql.DB) *MonitoringPostgresStore {
	return &MonitoringPostgresStore{
		DB: db,
	}
}

type MonitoringURL struct {
	Id     int     `json:"id" db:"id"`
	Name   string  `json:"name" db:"name"`
	URL    string  `json:"url" db:"url"`
	Status *string `json:"status" db:"status"`
}

func (m *MonitoringPostgresStore) AddMonitoringURLs(monitoringURLs []MonitoringURL, userId string) error {
	valueStrings := make([]string, 0, len(monitoringURLs))
	valueArgs := make([]any, 0, len(monitoringURLs)*3)

	for i, monitoringURL := range monitoringURLs {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		valueArgs = append(valueArgs, monitoringURL.Name)
		valueArgs = append(valueArgs, monitoringURL.URL)
		valueArgs = append(valueArgs, userId)
	}

	stmt := fmt.Sprintf("INSERT INTO monitoring_urls (name, url, user_id) VALUES %s", strings.Join(valueStrings, ","))
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, stmt, valueArgs...)
	return err
}

func (m *MonitoringPostgresStore) GetMonitoringURLsByUserId(userId string) ([]MonitoringURL, error) {
	query := `
		SELECT id, name, url, status FROM monitoring_urls
		WHERE user_id = $1;
	`

	args := []any{userId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mURLs := []MonitoringURL{}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		slog.Error("error getting monitoring urls", "err", err.Error())
		return nil, err
	}

	for rows.Next() {
		mURL := &MonitoringURL{}
		err := rows.Scan(&mURL.Id, &mURL.Name, &mURL.URL, &mURL.Status)
		if err != nil {
			slog.Error("error getting monitoring url", "err", err.Error())
			return nil, err
		}

		mURLs = append(mURLs, *mURL)
	}

	return mURLs, nil
}

func (m *MonitoringPostgresStore) GetAllMonitoringURLs() ([]MonitoringURL, error) {
	query := `
		SELECT name, url, status FROM monitoring_urls
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mURLs := []MonitoringURL{}

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		slog.Error("error getting monitoring urls", "err", err.Error())
		return nil, err
	}

	for rows.Next() {
		mURL := &MonitoringURL{}
		err := rows.Scan(&mURL.Name, &mURL.URL, &mURL.Status)
		if err != nil {
			slog.Error("error getting monitoring url", "err", err.Error())
			return nil, err
		}

		mURLs = append(mURLs, *mURL)
	}

	return mURLs, nil
}

func (r *MonitoringPostgresStore) UpdateMonitoringURL(id int, userId string, mURL MonitoringURL) (*MonitoringURL, error) {
	query := `
		UPDATE monitoring_urls SET name = $3, url = $4, status = null
		WHERE id = $1 and user_id = $2
		RETURNING id, name, url, status;
	`

	args := []any{id, userId, mURL.Name, mURL.URL}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	monitoringURL := &MonitoringURL{}

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&monitoringURL.Id, &monitoringURL.Name, &monitoringURL.URL, &monitoringURL.Status)
	if err != nil {
		slog.Error("error updating monitoring", "err", err.Error())
		return nil, err
	}

	return monitoringURL, nil
}

func (r *MonitoringPostgresStore) DeleteMonitoringURL(id int, userId string) error {
	query := `
		DELETE FROM monitoring_urls WHERE id = $1 and user_id = $2;
	`

	args := []any{id, userId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		slog.Error("error deleting monitoring", "err", err.Error())
		return err
	}

	return nil
}
