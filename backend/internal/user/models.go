package user

import (
	"context"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func FromRequest(r *http.Request) (*User, error) {
	// Fetch the key set from Better Auth
	keySet, err := jwk.Fetch(context.Background(), "http://localhost:5173/api/auth/jwks")
	if err != nil {
		return nil, err
	}

	// Parse the JWT token from the request
	token, err := jwt.ParseRequest(r, jwt.WithKeySet(keySet))
	if err != nil {
		return nil, err
	}

	// Get user ID from subject
	userID := token.Subject()

	// Extract email and name from claims
	var email, name string
	v, ok := token.Get("email")
	if ok {
		email = v.(string)
	}

	v, ok = token.Get("name")
	if ok {
		name = v.(string)
	}

	return &User{
		ID:    userID,
		Email: email,
		Name:  name,
	}, nil
}
