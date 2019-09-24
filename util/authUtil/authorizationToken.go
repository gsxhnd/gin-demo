package authUtil

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.

// token info contain of user type, user id, expireTime
type TokenInfo struct {
	Type       uint64 // user type 1:app user 2:
	Username   string
	UserID     uint64 `json:"user_id"`
	ExpireTime int64  `json:"exp"`
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*TokenInfo, error) {
	ctx := &TokenInfo{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Type = uint64(claims["type"].(float64))
		ctx.Username = claims["username"].(string)
		ctx.ExpireTime = int64(uint64(claims["exp"].(float64)))
		ctx.UserID = uint64(claims["user_id"].(float64))
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(token string) (*TokenInfo, error) {
	// Load the jwt secret from config
	secret := viper.GetString("JWTSecret")

	if len(token) == 0 {
		return &TokenInfo{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	_, _ = fmt.Sscanf(token, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.
func Sign(c TokenInfo, secret string, expTime int64) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("JWTSecret")
	}
	if c.Type == 1 {
		// The token content.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type":     c.Type,
			"username": c.Username,
			"user_id":  c.UserID,
			"exp":      expTime,
			"nbf":      time.Now().Unix(),
			"iat":      time.Now().Unix(),
		})
		// Sign the token with the specified secret.
		tokenString, err = token.SignedString([]byte(secret))
		return
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type":     c.Type,
			"username": c.Username,
			"nbf":      time.Now().Unix(),
			"iat":      time.Now().Unix(),
		})
		// Sign the token with the specified secret.
		tokenString, err = token.SignedString([]byte(secret))
		return
	}
}
