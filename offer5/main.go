package main

import "fmt"

type USER struct {
	age  int
	name string
	u    Name
}
type Name struct {
	a string
}
type USER2 struct {
	age  int
	name string
	u    Name
}

func main() {
	user := USER{
		1,
		"d",
		Name{""},
	}
	user3 := USER2{
		1,
		"d",
		Name{""},
	}
	user2 := USER(user3)
	fmt.Println(user2 == user)
}
