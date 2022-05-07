package main

import "fmt"

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia")
	dwarfs4 := dwarfs1[0:4]
	dwarfs5 := append(dwarfs4, "Test")
	fmt.Printf("%v\n", dwarfs1)
	fmt.Printf("%v\n", dwarfs2)
	fmt.Printf("%v\n", dwarfs3)
	fmt.Printf("%v\n", dwarfs4)
	fmt.Printf("%v\n", dwarfs5)
}