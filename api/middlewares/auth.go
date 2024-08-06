package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/AppDeveloperMLLB/todo_app/apperrors"
	"github.com/AppDeveloperMLLB/todo_app/common"
	"google.golang.org/api/idtoken"
)

// AuthMiddleware is a middleware for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// スキップしたいパス
			if req.URL.Path == "/login" || req.URL.Path == "/" || req.URL.Path == "/callback" {
				next.ServeHTTP(w, req)
				return
			}

			authorization := req.Header.Get("Authorization")
			authHeaders := strings.Split(authorization, " ")
			if len(authHeaders) != 2 {
				err := apperrors.InvalidHeader.Wrap(
					errors.New("invalid req header"),
					"invalid header",
				)
				apperrors.ErrorHandler(w, req, err)
				return
			}

			bearer, token := authHeaders[0], authHeaders[1]
			if bearer != "Bearer" || token == "" {
				err := apperrors.InvalidHeader.Wrap(
					errors.New("invalid req header"),
					"invalid header",
				)
				apperrors.ErrorHandler(w, req, err)
				return
			}

			tokenValidator, err := idtoken.NewValidator(context.Background())
			if err != nil {
				err = apperrors.CreateValidatorFailed.Wrap(
					err,
					"internal auth error",
				)
				apperrors.ErrorHandler(w, req, err)
				return
			}

			clientID := os.Getenv("CLIENT_ID")
			payload, err := tokenValidator.Validate(context.Background(), token, clientID)
			if err != nil {
				log.Println("validate failed")
				err = apperrors.Unauthorized.Wrap(err, "invalid token")
				apperrors.ErrorHandler(w, req, err)
				return
			}

			name, ok := payload.Claims["name"]
			if !ok {
				log.Println("name not found")
				err := apperrors.Unauthorized.Wrap(errors.New("invalid token"), "invalid token")
				apperrors.ErrorHandler(w, req, err)
				return
			}
			req = common.SetUserName(req, name.(string))

			id, ok := payload.Claims["sub"]
			if !ok {
				log.Println("id not found")
				err := apperrors.Unauthorized.Wrap(errors.New("invalid token"), "invalid token")
				apperrors.ErrorHandler(w, req, err)
				return
			}
			req = common.SetUserID(req, id.(string))

			next.ServeHTTP(w, req)
		})
}
