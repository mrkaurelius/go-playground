package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type App struct {
	ID string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App
	Org
}

func main() {
	fmt.Println("merhaba yalan dunya")

	// initing struct
	user := User{"1234", "asdfjafawef1234"}
	fmt.Println(user)

	// encoding to json
	userJson, _ := json.Marshal(user)
	fmt.Println(string(userJson))

	// indenting
	indentedJson := &bytes.Buffer{}
	json.Indent(indentedJson, userJson, "", "    ")
	fmt.Println(indentedJson)

	var parsedUser = &User{}
	json.Unmarshal(indentedJson.Bytes(), parsedUser)
	fmt.Println(parsedUser)

	userMap := make(map[string]User)
	userMap["1234"] = user
	fmt.Printf("Usertoken form map %v\n", userMap["1234"].Token)

	var dat map[string]interface{}
	fmt.Println(dat)

	// indented json
	data := []byte(`
		{
			"id": "k34rAT4",
			"name": "My Awesome Org"
		}
	`)

	fmt.Println(string(data))

	var appWithOrg AppWithOrg
	json.Unmarshal(data, &appWithOrg)
	fmt.Println(appWithOrg)

	app := appWithOrg.App
	org := appWithOrg.Org
	fmt.Println(app, org)

	// App type assert
	var i interface{} = App{}
	app1, ok := i.(App)

	fmt.Println(app1)
	fmt.Println(ok)

	fmt.Println(appWithOrg.ID)
	fmt.Println(appWithOrg.Name)
}
