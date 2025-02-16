package models

//import "gorm.io/gorm"

type User struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TableName overrides the default table name
func (User) TableName() string {
	return "users"
}
