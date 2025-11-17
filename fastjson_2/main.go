package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{ "user": { "name": "John Doe", "age": 30 } }`

	jsonValue, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userObj := jsonValue.GetObject("user")
	fmt.Printf("%+v\n", userObj)
	fmt.Printf("user name: %s\n", userObj.Get("name"))
	fmt.Printf("user age: %s\n", userObj.Get("age"))

	userJSON := jsonValue.Get("user").String()

	var user User

	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}

	fmt.Println(user.Name, user.Age)
}
