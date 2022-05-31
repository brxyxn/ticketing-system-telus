package middleware

type Login struct {
	Email string
	Token string
	IP    string
}

/* Repository */
// type TokenRepository interface {
// 	GetAuthToken(login *Login) error
// }

// type tokenService struct {
// 	tokenRepository TokenRepository
// }

// func NewTokenService(databaseRepository TokenRepository) TokenRepository {
// 	return &tokenService{databaseRepository}
// }

/* Service */
// type AuthenticationService interface {
// 	GetAuthToken(login *Login) error
// }

// type AuthenticationService interface {
// 	Authenticate(c *fiber.Ctx) error
// }

// type userHandler struct {
// 	service AuthenticationService
// }

// func NewTokenHandler(service AuthenticationService) AuthenticationService {
// 	return &userHandler{service}
// }
