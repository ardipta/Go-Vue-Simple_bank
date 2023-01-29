package token

import "time"

// Macker is an interface for manage tokens
type Maker interface {
	// CreateToken create a new token for a specific username and password
	CreateToken(username string, duration time.Duration) (string, error)

	// verify token if valid or not
	VerifyToken(token string) (*Payload, error)
}
