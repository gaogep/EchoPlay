package models

import "time"

type Post struct {
	ID         int       `gorm:"NOT NULL;PRIMARY_KEY;INDEX;AUTO_INCREMENT" json:"pid"`
	Title      string    `gorm:"TYPE:VARCHAR(30);DEFAULT:''" json:"title"`
	Content    string    `gorm:"TYPE:TEXT;DEFAULT:''" json:"content"`
	CreatedAt  time.Time `gorm:"TYPE:DATETIME"`
	User       User      `json:"user"`
	UserID     int       `json:"uid"`
	Category   Category  `json:"category"`
	CategoryID int       `json:"cid"`
}

func CreatePost(data map[string]interface{}) bool {
	db.Create(&Post{
		Title:      data["title"].(string),
		Content:    data["content"].(string),
		UserID:     data["user_id"].(int),
		CategoryID: data["category_id"].(int),
	})

	return true
}

func ExistedPostById(id int) bool {
	var post Post
	db.Select("id").Where("id = ?", id).First(&post)
	if post.ID > 0 {
		return true
	}

	return false
}

func GetPost(id int) Post {
	var post Post
	db.Where("id = ?", id).First(&post)
	db.Model(&post).Related(&post.User)
	return post
}

func UpdatePost(id int, data map[string]interface{}) bool {
	db.Model(&Post{}).Where("id = ?", id).Updates(data)
	return true
}

func DeletePost(id int) bool {
	db.Where("id = ?", id).Delete(&Post{})
	return true
}
