// model.go

package models

import (
	"errors"


	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	// Password       string   `json:"-"`
	FirstName      string   `json:"firstName,omitempty"`
	LastName       string   `json:"lastName,omitempty"`
	City           string   `json:"city,omitempty"`
	State          string   `json:"state,omitempty"`
	Country        string   `json:"country,omitempty"`
	Street         string   `json:"street,omitempty"`
	Gender         string   `json:"gender,omitempty"`
	Name           string   `json:"name,omitempty"`
	Dob            string   `json:"dob,omitempty"`
	Phone          string   `json:"phone,omitempty"`
	ProfilePicture string   `json:"profilePicture,omitempty"`
	FollowersCount int64    `gorm:"-" json:"followersCount"`
	FollowingCount int64    `gorm:"-" json:"followingCount"`
	ListingsCount  int64    `gorm:"-" json:"listingsCount"`
	ImageLinks     []string `gorm:"-" json:"imageLinks"`
	// Dob            time.Time `json:"dob"`
}

func (user *User) SignIn(db *gorm.DB) error {
	// chester.burke@example.com          | 7777
	// fmt.Println("email:", user.Email, "", "\npassword:", user.Password)

	// err := db.First("id,email,profile_picture,username,first_name,last_name", "password = ? and email = ?", user.Password, user.Email).Error
	err := db.Select("id,email,profile_picture,username,first_name,last_name").Where("password = ? and email = ? ", user.Password, user.Email).Find(&user).Error
	if user.ID <= 0 {
		return errors.New("User Not Found")
	}
	return err
}