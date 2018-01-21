package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["11111"] = PersonInfo{"11111", "Jack", "Beijing"}
	personDB["22222"] = PersonInfo{"22222", "John", "Shanghai"}

	person, ok := personDB["11111"]
	if ok == true {
		fmt.Println(person)
	} else {
		fmt.Println("person not found")
	}
	var i = 100

	if i > 0 {
		fmt.Println("i>0")
	} else if i < 0 {
		fmt.Println("i<0")
	} else {
		fmt.Println("i==0")
	}

	var num = 1

	switch {
	case 0 <= num && num <= 3:
		fmt.Printf("0-3\n")
	case 4 <= num && num <= 6:
		fmt.Printf("4-6\n")
	case 7 <= num && num <= 9:
		fmt.Printf("7-9\n")
	}
	myfunc(1, 2, 3, 4, 4, 5, 6)

}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
