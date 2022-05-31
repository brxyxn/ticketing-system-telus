package middleware

// redis
type TokenRepository interface {
	GetAuthToken(login *Login) error
}
