// model.go

package models

import (
	"errors"
	"net/http"
	"strconv"

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
	Cart           Cart     //added cuz of foreign keys
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

func (user *User) GetUserByID(db *gorm.DB) error {
	var follows Follows
	var listing Listing

	err := db.First(&user, "id = ?", user.ID).Error
	if err != nil {

		return err
	}
	user.FollowersCount = follows.GetFollowersCount(user, db)
	user.FollowingCount = follows.GetFollowingCount(user, db)
	user.ListingsCount = listing.GetListingsCount(user, db)

	return err

}

func (user *User) GetUser(db *gorm.DB) error {
	var follows Follows
	var listing Listing

	err := db.First(&user, "username = ?", user.Username).Error
	if err != nil {

		return err
	}
	user.FollowersCount = follows.GetFollowersCount(user, db)
	user.FollowingCount = follows.GetFollowingCount(user, db)
	user.ListingsCount = listing.GetListingsCount(user, db)

	return err

}

func (user *User) GetAllUserImages(db *gorm.DB) error {
	var images Images
	links, err := images.GetUsersImagesFromBucket(user.Username)

	user.ImageLinks = links
	// fmt.Println(user.ImageLinks)

	return err

}
func (user *User) UpdateUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("UPDATE users SET first_name='%s', dob=%s WHERE username=%s", user.FirstName, user.Dob, user.Username)
	// _, err := db.Exec(statement)

	return db.Save(&user).Error
	// return db.Model(&User{}).Where("username = ?", user.Username).Update("first_name", "yikes").Error
}

func (user *User) DeleteUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("DELETE FROM users WHERE username=%s", user.Username)
	// _, err := db.Exec(statement)

	return db.Where("username = ?", user.Username).Delete(&user).Error
}

func (user *User) CreateUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("INSERT INTO users(first_name, dob) VALUES('%s', %s)", user.FirstName, user.Dob)
	// _, err := db.Exec(statement)

	// if err != nil {
	// 	return err
	// }

	// err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.Username)

	// if err != nil {
	// 	return err
	// }

	return db.Select("Name", "Age", "CreatedAt").Create(&user).Error
}

func (user *User) GetUsersWithSearch(db *gorm.DB) ([]User, error) {

	// the beautiful beautiful query
	// SELECT * FROM users WHERE concat(first_name," ",last_name)  LIKE '%Joel Amo%'

	// }// Get all records
	users := []User{}

	// err := db.Where(`concat(username, " ",first_name," ",last_name) LIKE "%@name%"`, map[string]interface{}{"name": user.Username}).Find(&users).Error
	trx := db.Where(`instr(concat(username, " ",first_name," ",last_name),?)`, user.Name).Find(&users)
	err := trx.Error

	return users, err
}

func (user *User) GetUsers(db *gorm.DB, start, count int) ([]User, error) {
	// statement := fmt.Sprintf("SELECT username, first_name, dob FROM users LIMIT %d OFFSET %d", count, start)
	// rows, err := db.Query(statement)

	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	// users := []User{}

	// for rows.Next() {
	// 	var user User
	// 	if err := rows.Scan(&user.Username, &user.FirstName, &user.Dob); err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, user)
	// }

	// }// Get all records
	users := []User{}

	// SELECT * FROM users;

	return users, db.Find(&users).Error
}

func (user *User) GetUsersByPage(r *http.Request, db *gorm.DB) ([]User, error) {

	// }// Get all records
	users := []User{}

	// SELECT * FROM users;

	return users, db.Scopes(usersPaginate(r)).Find(&users).Error
}

func usersPaginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
