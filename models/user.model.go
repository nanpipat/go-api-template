package models

type User struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
}

func (User) TableName() string {
	return "users"
}
