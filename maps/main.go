package main

import "fmt"

func main() {

	var m1 map[string]string

    fmt.Println("uninitialized map m1", m1)

    //m1["test"] = "ok" - cant isert into nil map

	var m2 map[string]map[string]string

	fmt.Println("uninitialized map 2nd order m2", m2)

	var m3 map[string]string = map[string]string{}

	m3["test"] = "ok"

	fmt.Println("m3", m3)

	m4 := map[string]string{}

	fmt.Println("m4", m4)

	m5 := make(map[string]string)

	m5["firstName"] = "Dmytro"

	fmt.Println("m5", m5)

	firstName := m5["firstName"]

	fmt.Println("firstName", firstName)

	lastName := m5["lastName"]

	fmt.Println("lastName", lastName)

	fullName, exists := m5["fullName"]

	fmt.Println("fullName", fullName, "exists", exists)

	_, exists1 := m5["fullName"]

	fmt.Println("exists1", exists1)

	key := "firstName"
	delete(m5, key)

	_, exists2 := m5[key]
	fmt.Println("exists2", exists2)

	m6 := m5
	m6[key] = "Mykola"
	fmt.Println("m5", m5, "m6", m6)


}
