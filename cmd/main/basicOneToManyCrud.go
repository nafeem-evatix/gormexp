package main

import "gorm.io/gorm"
// model
type Hero struct{
	Id    uint `gorm:"primarykey"`
	Name string
	Weapons []Weapon `gorm:"foreignkey:hero_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (it *Hero) TableName() string {
	return "Hero"
}

type Weapon struct {
	Id    uint `gorm:"primarykey"`
	Name string
	HeroId uint
	Hero *Hero
}

func (it *Weapon) TableName() string {
	return "Weapon"
}
// service

type basicOneToMany struct {
	db *gorm.DB
}

func (it *basicOneToMany) migrate() error {
	return it.db.AutoMigrate(&Hero{},&Weapon{})
}

func (it *basicOneToMany) createHero(hero *Hero) error {
	return it.db.Create(hero).Error
}

func newBasicOneToMany() *basicOneToMany {
	db := getSQLiteGormDBMust(dbName)
	return &basicOneToMany{
		db: db,
	}
}