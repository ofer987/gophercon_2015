package main

import "fmt"

func main() {
	var num int

	num = 20

	fmt.Printf("Address of num is %v\n", &num)
	fmt.Printf("Value of num is %v\n", num)

	type Something struct {
		Name   string
		Age    int
		Active bool
	}

	var dan Something

	dan = Something{
		Name:   "Dan",
		Active: true,
	}

	fmt.Printf("%s is %v years old and is %s in the community", dan.Name, dan.Age, dan.Active)
}
