package main

import "fmt"

func main() {

	a := true

	if a {
		fmt.Println("hello a")
	}

	fmt.Println("---------")
	b := 1
	if b == 1 {
		fmt.Println("hello b")
	}

	fmt.Println("---------")
	m1 := map[string]string{"firstN": "Dmytro", "lastN": "Bortnichuk"}

	if firstN, exist := m1["firstN"]; exist && firstN == "Dmytro" {
		fmt.Println("this is", firstN)
	} else if exist {
		fmt.Println("firstN", firstN)
	} else {
		fmt.Println("no firstN")
	}

	fmt.Println("---------")
	c := 2

	if (a || b > 0) && (c == 2 || b < 0) {

	}

	for {
		fmt.Println("while(true)")
		break
	}

	fmt.Println("---------")
	sl1 := []int{0, 10, 20, 30, 40, 50}
	value := 0
	idx := 0
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl1[idx]
		idx++
		println("while - style loop,", "idx", idx, "value", value)
	}

	fmt.Println("---------")
	for i := 0; i < len(sl1)-2; i++ {
		fmt.Println("for - style loop,", "idx", i, "value", sl1[i])
	}

	fmt.Println("---------")
	for _, val := range sl1 {
		fmt.Println("foreach - style loop, no idx,", "value", val)
	}

	fmt.Println("---------")
	for idx, val := range sl1 {
		fmt.Println("foreach - style loop,", "idx", idx, "value", val)
	}

	fmt.Println("---------")
	//m1 := map[string]string{"a":"1", "b":"2", "c":"3"}

	for key := range m1 {
		fmt.Println("foreach - style loop,", "idx", key)
	}

	fmt.Println("---------")
	for key, val := range m1 {
		fmt.Println("foreach - style loop,", "idx", key, "value", val)
	}

	fmt.Println("---------")
	for _, val := range m1 {
		fmt.Println("foreach - style loop,", "value", val)
	}

	m1["flag"] = "ok"

	fmt.Println("switch---------")
	switch m1["firstName"] {
	case "Dmytro", "Stepan":
		fmt.Println("identified") // no fallthrough by default
	case "Mariasia":
		if m1["flag"] == "ok" {
			break; // exit switch
		}
		fmt.Println("this is Mariasia")
		fallthrough
	default:
		fmt.Println("default")
	}

	fmt.Println("switch---------")
	switch  {
	case m1["firstName"] == "Dmytro", m1["firstName"] == "Stepan":
		fmt.Println("Dmytro or Stepan")
	case m1["firstName"] == "Mariasia":
		fmt.Println("Mariasia")
	}

	fmt.Println("---------")
	MyLabel:
	for key, val := range m1 {
		fmt.Println("foreach - style loop,", "idx", key, "value", val)
		switch {
		case m1["firstName"] == "Dmytro":
			fmt.Println("Dmytro found, exiting the loop")
			break MyLabel
		}
	}




}
