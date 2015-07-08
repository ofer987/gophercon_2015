package main

import "fmt"

type Speaker interface {
	SayHello() string
}

type English struct {
	Name string
}

type Chinese struct {
	Name string
}

func (sp English) SayHello() string {
	return "Hello World " + sp.Name
}

func (sp Chinese) SayHello() string {
	return "你好世界 " + sp.Name
}

func main() {
	dan := English{Name: "Dan"}
	andy := Chinese{Name: "Andy"}

	speakers := []Speaker{dan, andy}

	for _, speaker := range speakers {
		fmt.Printf("%s\n", speaker.SayHello())
	}
}
