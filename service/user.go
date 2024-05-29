package service

import (
	"PaperSubmission/model"
	"PaperSubmission/utils"
	"errors"
	"sync"
)

type UserServiceInterface interface {
	Add(user model.User) error
	Delete(user model.User) error
	Edit(user model.User) error
	Get(user model.User) (*model.User, error)
	GetList(request *utils.ListQuery) ([]*model.User, error)
	VerifyPassword(email, password string) (int64, string, error)
	GetUserNum() (int64, error)
}

type UserService struct {
}

var (
	userService     *UserService
	userServiceOnce sync.Once
)

func NewUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (s UserService) Add(user model.User) error {
	email := user.Email
	u, err := model.NewUserModel().GetByEmail(email)
	if err != nil {
		return err
	}
	if u != nil {
		return errors.New("邮箱已被注册")
	}
	return model.NewUserModel().Add(user)
}

func (s UserService) Delete(user model.User) error {
	return model.NewUserModel().Delete(user)
}

func (s UserService) Edit(user model.User) error {
	return model.NewUserModel().Edit(user)
}

func (s UserService) Get(user model.User) (*model.User, error) {
	return model.NewUserModel().Get(user)
}

func (s UserService) GetList(request *utils.ListQuery) ([]*model.User, error) {
	return model.NewUserModel().GetList(request)
}

func (s UserService) VerifyPassword(email, password string) (int64, string, error) {
	user, err := model.NewUserModel().GetByEmail(email)
	if err != nil {
		return 0, "", err
	}
	if user == nil { // 用户不存在
		return 0, "", errors.New("用户不存在")
	}
	if user.Password != password {
		return 0, "", nil
	}
	return user.ID, user.Name, nil
}

func (s UserService) GetUserNum() (int64, error) {
	return model.NewUserModel().GetUserNum()
}
