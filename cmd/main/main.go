package main

import (
	"fmt"
)

func main() {
	fmt.Println("gormexp")
	testBasicUserCrud()
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




