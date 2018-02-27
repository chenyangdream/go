package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//type User struct {
	//	UserId   int    `json:"user_id"`
	//	UserName string `json:"user_name"`
	//}
	type User struct {
		UserId   int
		UserName string
	}

	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))

	var user User
	json.Unmarshal(j, &user)
	fmt.Printf("%v\n", user)
	fmt.Println("user.UserId = ", user.UserId, "user.UserName = ", user.UserName)
}
