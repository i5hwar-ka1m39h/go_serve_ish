package usecase

import (
	"context"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
)

type userUsecase struct {
	userRepository model.UserRepository
	contextTime    time.Duration
}

func NewUserUsecase(userRepository model.UserRepository, time time.Duration) model.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTime:    time,
	}
}

func (userUc *userUsecase) CreateUser(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, userUc.contextTime)
	defer cancel()

	err := userUc.userRepository.CreateSingle(ctx, user)

	return err
}

func (userUc *userUsecase) UpdateUser(c context.Context, userId string, user map[string]any) error {
	ctx, cancel := context.WithTimeout(c, userUc.contextTime)
	defer cancel()

	err := userUc.userRepository.UpdateSingle(ctx, userId, user)

	return err
}

func (userUc *userUsecase) GetUserById(c context.Context, userId string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, userUc.contextTime)
	defer cancel()

	user, err := userUc.userRepository.GetSingleById(ctx, userId)
	return user, err

}

func (userUc *userUsecase) SearchUser(c context.Context, someStr string) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(c, userUc.contextTime)
	defer cancel()

	users, err := userUc.userRepository.GetByString(ctx, someStr)
	return users, err
}
