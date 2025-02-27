package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TableName overrides the default table name with users
func (User) TableName() string {
	return "users"
}
