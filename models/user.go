package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"NOT NULL;PRIMARY_KEY;INDEX;AUTO_INCREMENT" json:"id"`
	NickName  string    `gorm:"TYPE:VARCHAR(24);DEFAULT:'';UNIQUE_INDEX" json:"nick_name"`
	Mobile    string    `gorm:"TYPE:VARCHAR(100);DEFAULT:''" json:"mobile"`
	Email     string    `gorm:"TYPE:VARCHAR(30);DEFAULT:''" json:"email"`
	PassWd    string    `gorm:"TYPE:VARCHAR(200);DEFAULT:''"`
	CreatedAt time.Time `gorm:"TYPE:DATETIME"`
	Posts     []Post    `json:"posts"`
}

func ExistedUserById(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func ValidateUserByName(name string, passwd string) bool {
	var user User
	db.Select("pass_wd").Where("nick_name = ?", name).First(&user)
	if user.PassWd == passwd {
		return true
	}

	return false
}
