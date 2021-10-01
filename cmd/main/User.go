package main

type User struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}

func (user *User) TableName() string {
	return "User"
}
