package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

// User struct
type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// response, err := r.Table("user").Insert(user).RunWrite(session)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// user := User{
	// 	Name: "Nikita Muraviov",
	// }
	// response, _ := r.
	// 	Table("user").
	// 	Get("247a611a-43c2-490a-9726-39910314ecc9").
	// 	Delete().
	// 	RunWrite(session)
	// fmt.Printf("%#v\n", response)

	cursor, _ := r.Table("user").Changes(r.ChangesOpts{IncludeInitial: true}).Run(session)
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		fmt.Printf("%#v\n", changeResponse)
	}
}
