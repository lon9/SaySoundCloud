package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/api/option"
)

const valName = "FIREBASE_ID_TOKEN"

// FirebaseAuthMiddleware contains methods verifying JWT token
type FirebaseAuthMiddleware struct {
	fbase   *firebase.App
	skipper middleware.Skipper
}

// NewFireBaseAuthMiddleware is middleware authentication with firebase
func NewFireBaseAuthMiddleware(credFilePath string, skipper middleware.Skipper) (*FirebaseAuthMiddleware, error) {
	opt := option.WithCredentialsFile(credFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	if skipper == nil {
		skipper = middleware.DefaultSkipper
	}
	return &FirebaseAuthMiddleware{
		fbase:   app,
		skipper: skipper,
	}, nil
}

// Verify verifies token
func (f *FirebaseAuthMiddleware) Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if f.skipper(c) {
			return next(c)
		}

		r := c.Request()
		token := strings.Replace(r.Header.Get(echo.HeaderAuthorization), "Bearer ", "", 1)
		if token == "" {
			return c.String(http.StatusUnauthorized, "Bad token")
		}

		client, err := f.fbase.Auth(context.Background())
		if err != nil {
			log.Println(err)
			return c.String(http.StatusUnauthorized, "Bad token")
		}

		authToken, err := client.VerifyIDToken(context.Background(), token)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusUnauthorized, "Bad token")
		}
		c.Set(valName, authToken)
		return next(c)
	}
}

// ExtractClaims extracts claims
func ExtractClaims(c echo.Context) *auth.Token {
	idToken := c.Get(valName)
	if idToken == nil {
		return new(auth.Token)
	}
	return idToken.(*auth.Token)
}
