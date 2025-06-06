package middleware

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/cicompanion/api"
	"github.com/cicompanion/data"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

const (
	authorizationHeader     = "authorization"
	cookieAuthorizationName = "ci_access_token"
	authorizationPrefix     = "bearer"
)

var (
	errInvalidToken = "invalid token"
	errExpiredToken = "expired token"
)

type UserStore interface {
	GetUser(username string) (*data.User, error)
	InsertUser(username, githubPat string) (*data.User, error)
	UpdateUserToken(username, githubPat string) (*data.User, error)
}

func CorsMiddleware(frontendURL string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if origin == "http://127.0.0.1:5173" || origin == frontendURL {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Auth is a middleware that checks for OAuth2 authentication.
func Auth(us UserStore, excludedPaths []string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, path := range excludedPaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fields := []string{}
			c := r.CookiesNamed(cookieAuthorizationName)
			if len(c) == 0 {
				fields = strings.Fields(r.Header.Get(authorizationHeader))
				if len(fields) < 2 {
					api.ForbiddenResponse(w, r, errInvalidToken)
					return
				}
			} else {
				fields = append(fields, "bearer")
				fields = append(fields, c[0].Value)
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationPrefix {
				api.ForbiddenResponse(w, r, errInvalidToken)
				return
			}

			token := &oauth2.Token{
				AccessToken: fields[1],
			}

			// validate token
			oauthUser, err := api.GetUserData(token)
			if err != nil {
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					slog.Info("token is expired")
					api.ForbiddenResponse(w, r, errExpiredToken)
					return
				default:
					api.ForbiddenResponse(w, r, errExpiredToken)
					return
				}
			}

			user, err := us.GetUser(oauthUser.Username)
			if err != nil {
				switch {
				case errors.Is(err, sql.ErrNoRows):
					user, err = us.InsertUser(oauthUser.Username, token.AccessToken)
					if err != nil {
						api.ServerErrorResponse(w, r, err)
						return
					}
				default:
					api.ForbiddenResponse(w, r, errExpiredToken)
					return
				}
			} else {
				if user.GithubPAT != token.AccessToken {
					user, err = us.UpdateUserToken(oauthUser.Username, token.AccessToken)
					if err != nil {
						api.ServerErrorResponse(w, r, err)
						return
					}
				}
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), api.ContextUser, user)))
		})
	}
}
