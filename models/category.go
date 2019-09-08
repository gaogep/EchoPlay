package models

import "time"

type Category struct {
	ID        int       `gorm:"NOT NULL;PRIMARY_KEY;INDEX;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"TYPE:VARCHAR(100);DEFAULT:''" json:"name"`
	CreatedAt time.Time `gorm:"TYPE:DATETIME" json:"created_at"`
	Posts     []Post    `json:"posts"`
}

func GetCategoryList(pageOffset int, pageSize int, maps map[string]interface{}) (list []Category) {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&list)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Category{}).Where(maps).Count(&count)
	return
}

func ExistedCategoryByName(name string) bool {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return true
	}

	return false
}

func ExistedCategoryById(id int) bool {
	var category Category
	db.Select("id").Where("id = ?", id).First(&category)
	if category.ID > 0 {
		return true
	}

	return false
}

func CreateCategory(name string) bool {
	db.Create(&Category{
		Name:      name,
		CreatedAt: time.Now(),
	})

	return true
}

func UpdateCategory(id int, data map[string]interface{}) bool {
	db.Model(&Category{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteCategory(id int) bool {
	db.Where("id = ?", id).Delete(&Category{})
	return true
}
