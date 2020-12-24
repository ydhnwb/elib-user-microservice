package authservice

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/elib-user-microservice/application/repository"
	"github.com/ydhnwb/elib-user-microservice/domain/entity"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/security"
)

//AuthService is a contract
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	RegisterUser(u dto.UserRegisterDTO) (entity.User, error)
	FindByEmail(email string) interface{}
	IsEmailDuplicate(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService is creates a new instance of AuthService
func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		userRepository: repo,
	}
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
	res := s.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := security.ComparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (s *authService) RegisterUser(u dto.UserRegisterDTO) (entity.User, error) {
	//check for duplicate
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&u))
	if err != nil {
		log.Fatalf("%s", err)
	}
	return s.userRepository.InsertUser(user)
}

func (s *authService) FindByEmail(email string) interface{} {
	return s.userRepository.FindByEmail(email)
}

func (s *authService) IsEmailDuplicate(email string) bool {
	res := s.userRepository.FindByEmail(email)
	return res != nil
}
