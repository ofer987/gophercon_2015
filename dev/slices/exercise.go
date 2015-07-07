package main

import "fmt"

func main() {
	part_a()
	part_b()
}

func part_a() {
	nil_slice := []int{}

	for i := 0; i < 10; i++ {
		nil_slice = append(nil_slice, i)
	}

	for i, item := range nil_slice {
		fmt.Printf("[%v] and the value is %v\n", i, item)
	}
}

func part_b() {
	persons := []string{"Dan", "Edith", "Alexander", "Ron", "Keren", "Effi"}
	print_names(persons)

	siblings := persons[3:5:5]
	print_names(siblings)
}

func print_names(items []string) {
	for i, item := range items {
		fmt.Printf("[%v] My name is %s\n", i, item)
	}
}
