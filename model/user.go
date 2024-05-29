package model

import (
	"PaperSubmission/utils"
	"errors"
	"log"
	"sync"
	"time"
)

type User struct {
	ID           int64     `gorm:"primary_key"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	Name         string    `gorm:"column:name"`
	Organization string    `gorm:"column:organization"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type UserModelInterface interface {
	Add(user User) error
	Delete(user User) error
	Edit(user User) error
	Get(user User) (*User, error)
	GetList(request *utils.ListQuery) ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetUserNum() (int64, error)
}

type UserModel struct {
}

var (
	userModel     *UserModel
	userModelOnce sync.Once // 单例模式
)

func NewUserModel() *UserModel { // 暴露给上层用来调用本层方法
	userModelOnce.Do(func() {
		userModel = new(UserModel)
	})
	return userModel
}

func (s UserModel) Add(user User) error {
	if err := GetDB().Create(&user).Error; err != nil {
		log.Println(err.Error())
		return errors.New("创建用户错误")
	}
	return nil
}

func (s UserModel) Delete(user User) error {
	if err := GetDB().Model(&user).Update("is_deleted", true).Error; err != nil {
		log.Println(err.Error())
		return errors.New("删除用户错误")
	}
	return nil
}

func (s UserModel) Edit(user User) error {
	if err := GetDB().Model(&user).Updates(map[string]interface{}{
		"name":         user.Name,
		"organization": user.Organization,
		"password":     user.Password,
		"email":        user.Email,
	}).Error; err != nil {
		log.Println(err.Error())
		return errors.New("编辑用户信息错误")
	}
	return nil
}

func (s UserModel) Get(user User) (*User, error) {
	if err := GetDB().Model(&user).Select("id", "email", "password", "name", "organization", "created_at").Find(&user).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询用户信息错误")
	}
	return &user, nil
}

func (s UserModel) GetList(request *utils.ListQuery) ([]*User, error) {
	var users []*User
	limit, offset := utils.Page(request.PageSize, request.Page) // 分页
	if err := GetDB().Order("id desc").Limit(limit).Offset(offset).Select("id", "email", "password", "name", "organization", "created_at").Find(&users).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询用户信息错误")
	}
	return users, nil
}

func (s UserModel) GetByEmail(email string) (*User, error) {
	var user User
	if err := GetDB().Where("email = ?", email).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if user.ID != 0 && user.IsDeleted == false { // 存在且没有被删除
		return &user, nil
	}
	return nil, nil // 不存在
}

func (s UserModel) GetUserNum() (int64, error) {
	var count int64
	if err := GetDB().Model(&User{}).Count(&count).Error; err != nil {
		log.Println(err.Error())
		return -1, errors.New("查询用户数量错误")
	}
	return count, nil
}
