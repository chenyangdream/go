package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() {
	gobook := Book{
		"GoProgram",
		[]string{"AAA", "BBB", "CCC"},
		"ituring.com.cn",
		true,
		9.99,
	}

	fmt.Println(gobook)

	b, err := json.Marshal(gobook)

	if err != nil {
		fmt.Println("json Masrshal error")
		return
	}

	var decodeBook Book

	json.Unmarshal(b, &decodeBook)
	fmt.Println(decodeBook)

	var r interface{}
	err1 := json.Unmarshal(b, &r)

	if err1 != nil {
		fmt.Println("json Unmarshal b error")
	}

	fmt.Println(r)

	newBook, ok := r.(map[string]interface{})

	if ok {
		for k, v := range newBook {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			case float64:
				fmt.Println(k, "is float64", v2)
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}
