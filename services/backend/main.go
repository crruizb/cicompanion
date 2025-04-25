package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/cicompanion/api"
	"github.com/cicompanion/api/middleware"
	"github.com/cicompanion/data"
	"github.com/cicompanion/githubapi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type config struct {
	oauthRedirectURL   string
	githubClientId     string
	githubClientSecret string
	dbUsername         string
	dbPassword         string
	dbHost             string
	dbName             string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config := config{}
	flag.StringVar(&config.oauthRedirectURL, "redirect-url", os.Getenv("REDIRECT_URL"), "Oauth redirect url")
	flag.StringVar(&config.githubClientId, "github-clientid", os.Getenv("GITHUB_CLIENT_ID"), "Github client id")
	flag.StringVar(&config.githubClientSecret, "github-client-secret", os.Getenv("GITHUB_CLIENT_SECRET"), "Github client secret")
	flag.StringVar(&config.dbUsername, "db-user", os.Getenv("DB_USERNAME"), "Db username")
	flag.StringVar(&config.dbPassword, "db-password", os.Getenv("DB_PASSWORD"), "Db password")
	flag.StringVar(&config.dbHost, "db-host", os.Getenv("DB_HOST"), "Db host")
	flag.StringVar(&config.dbName, "db-name", os.Getenv("DB_NAME"), "Db name")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", config.dbUsername, config.dbPassword, config.dbHost, config.dbName)
	db, err := openDB(dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var githubOauthConfig = &oauth2.Config{
		RedirectURL:  config.oauthRedirectURL,
		ClientID:     config.githubClientId,
		ClientSecret: config.githubClientSecret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
	us := data.NewUsersStore(db)
	gc := githubapi.NewGithubHttpClient()

	rt := api.NewRouter(githubOauthConfig, gc)

	excludePrefix := []string{"/auth/github/login", "/auth/callback"}
	mw := middleware.With(
		middleware.CorsMiddleware(),
		middleware.Auth(us, excludePrefix),
	)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mw(rt),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	slog.Info("starting ci companion service on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
