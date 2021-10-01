package main

import "gorm.io/gorm"

// Model
type Person struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	NID   *NID `gorm:"foreignkey:person_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (it *Person) TableName() string {
	return "Person"
}

type NID struct {
	Id     uint   `gorm:"primarykey"`
	Number string `gorm:"uniqueIndex"`
	PersonId uint
	Person *Person
}

func (it *NID) TableName() string {
	return "NID"
}

// Service

type basicOneToOneCrud struct {
	db *gorm.DB
}

func (it *basicOneToOneCrud) migrate() error {
	return it.db.AutoMigrate(&Person{},&NID{})
}

func (it *basicOneToOneCrud) createPerson(person *Person) error {
	return it.db.Create(person).Error
}

func (it *basicOneToOneCrud) getPerson(personId uint) *Person {
	person := &Person{}
	it.db.Where("id",personId).First(person)

	return person
}

func (it *basicOneToOneCrud) getPersonWithNIDPreload(personId uint) *Person {
	person := &Person{}
	nid := &NID{}
	it.db.
		Preload(nid.TableName()).
		Where("id",personId).First(person)

	return person
}

func (it *basicOneToOneCrud) deletePerson(personId uint) error {
	return it.db.Where("id",personId).Delete(&Person{}).Error
}

func (it *basicOneToOneCrud) getNIDWithPersonPreload(nidId uint) *NID {
	person := &Person{}
	nid := &NID{}
	it.db.
		Preload(person.TableName()).
		Where("id",nidId).First(nid)

	return nid
}

func newBasicOneToOneCrud() *basicOneToOneCrud {
	db := getSQLiteGormDBMust(dbName)
	return &basicOneToOneCrud{
		db: db,
	}
}
