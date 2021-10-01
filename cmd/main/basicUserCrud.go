package main

import "gorm.io/gorm"

type basicUserCrud struct {
	db *gorm.DB
}

func (it *basicUserCrud) migrate() error {
	return it.db.AutoMigrate(&User{})
}

func (it *basicUserCrud) create(user *User) error {
	return it.db.Create(user).Error
}

func (it *basicUserCrud) update(id uint,user *User) error {
	return it.db.Where("id",id).Updates(user).Error
}

func (it *basicUserCrud) get(id uint) *User {
	user := &User{}
	it.db.Where("id",id).First(user)

	return user
}

func (it *basicUserCrud) delete(id uint) error {
	return it.db.Where("id",id).Delete(&User{}).Error
}

func newBasicUserCrud() *basicUserCrud {
	db := getSQLiteGormDBMust("user.db")
	return &basicUserCrud{
		db: db,
	}
}



