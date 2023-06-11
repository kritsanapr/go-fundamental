package mapandstruct

import "fmt"

func Learn() {
	fmt.Println("Learn Map and Struct")

	// Map
	// Create map with make
	var mmap = make(map[string]int)
	mmap["value1"] = 1
	mmap["value2"] = 2

	m := map[string]int{"token": 1, "token2": 2}
	fmt.Println("map", m)
	fmt.Println("map token", m["token"])

	// assign map value
	m["token3"] = 3
	m["token4"] = 4
	fmt.Println("map token3", m["token3"])
	fmt.Println("map token4", m["token4"])

	// delete map value
	delete(m, "token3")
	fmt.Println("map token3", m["token3"])

	// loop map
	for key, value := range m {
		fmt.Println("loop map", key, value)
	}

	// Struct
	type User struct {
		Firstname string
		Lastname  string
		Age       int
	}

	var user User
	user.Firstname = "Kritsana"
	user.Lastname = "Prasit"
	user.Age = 25
	fmt.Println("struct", user)

	// assign struct value
	user2 := User{"Kritsana", "Prasit", 25}
	fmt.Println("struct", user2)
}
