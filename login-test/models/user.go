package models

import (
	"errors"
	"login-test/utils/token"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	// Username string `gorm:"size:255;not null;unique" json:"username"`
	// Password string `gorm:"size:255;not null;" json:"password"`
	Username string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null;"`
}

func (u User) Save() (User, error) {
	if DB == nil {
		return User{}, errors.New("DB is not exist.")
	}

	err := DB.Create(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// BeforeSaveというメソッド名にしておくとGinがフックしてくれるらしい
// 参考(https://go.genzouw.com/2023/04/12/post-358/)
func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Username = strings.ToLower(u.Username)

	return nil
}

func (u User) PrepareOutput() User {
	u.Password = ""
	return u
}

func GenerateToken(username string, password string) (string, error) {
	var user User

	err := DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
