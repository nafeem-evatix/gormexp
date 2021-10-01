package main

import (
	"fmt"
)

func main() {
	fmt.Println("gormexp")
	// testBasicUserCrud()
	testOneToOneCrud()
}

func testOneToOneCrud() {
	basicOneToOne := newBasicOneToOneCrud()
	handleError(basicOneToOne.migrate())
	person := &Person{
		Name:  "ABC",
		NID:   &NID{
			Number: "11122223333",
		},
	}

	handleError(basicOneToOne.createPerson(person))
	fmt.Println(basicOneToOne.getPerson(person.Id))
	fmt.Println(basicOneToOne.getPersonWithNIDPreload(person.Id).NID)

	fmt.Println(basicOneToOne.getNIDWithPersonPreload(1).Person)
	handleError(basicOneToOne.deletePerson(person.Id))
}

func testBasicUserCrud() {
	basicUserCrud := newBasicUserCrud()
	handleError(basicUserCrud.migrate())
	user := &User{
		Name:  "user1",
		Email: "user1@email.com",
	}

	handleError(basicUserCrud.create(user))
	fmt.Println(basicUserCrud.get(user.Id))

	user.Name = "user1ChangedName"

	handleError(basicUserCrud.update(user.Id,user))
	fmt.Println(basicUserCrud.get(user.Id))

	handleError(basicUserCrud.delete(user.Id))
	fmt.Println(basicUserCrud.get(user.Id))
}




