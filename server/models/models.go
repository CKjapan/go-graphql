package models

// Todo is
type Todo struct {
	ID     string `gorm:"column:id;primary_key"`
	Text   string `gorm:"column:text"`
	Done   bool   `gorm:"column:done"`
	UserID string `gorm:"column:user_id"`
}

// User is
type User struct {
	ID   string `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}
