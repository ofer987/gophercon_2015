func main() {
	dan := English{Name: "Dan"}
	andy := Chinese{Name: "Andy"}

	speakers := []Speaker{dan, andy}

	for _, speaker := range speakers {
		fmt.Printf("%s\n", speaker.SayHello())
	}
}
