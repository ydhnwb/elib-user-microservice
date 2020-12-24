package userservice

import (
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/elib-user-microservice/application/repository"
	"github.com/ydhnwb/elib-user-microservice/domain/entity"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/dto"
)

//UserService is a contract
type UserService interface {
	UpdateProfile(u dto.UserUpdateDTO) (entity.User, error)
	GetOwnProfile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (s *userService) UpdateProfile(u dto.UserUpdateDTO) (entity.User, error) {
	//repository is almost only needs entity shape, not the DTO
	//so after the DTO validated by controller/handler
	//i converts this dto to a shape that needed by repository
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&u))
	if err != nil {
		return user, err
	}
	res, e := s.userRepository.UpdateUser(user)
	if e != nil {
		return user, e
	}
	return res, nil
}

func (s *userService) GetOwnProfile(userID string) entity.User {
	//I am optimists here  about userID
	//First, userID is extracted from jwt token
	//and jwt token must be valid, because there is a middleware to check it
	return s.userRepository.Profile(userID)
}
