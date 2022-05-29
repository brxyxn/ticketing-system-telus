package middleware

type UserService interface {
	GetAuthToken(login *Login) error
}

type tokenService struct {
	cacheRepo TokenRepository
}

func NewUserService(cacheRepository TokenRepository) UserService {
	return &tokenService{cacheRepository}
}

func (s *tokenService) GetAuthToken(login *Login) error {
	return s.cacheRepo.GetAuthToken(login)
}
