package main

import (
	"fmt"
	"github.com/zond/godip"
	"github.com/zond/godip/variants"
)

func main() {
	variant := variants.Variants["Classical"]
	s, _ := variant.Start()
	s.Next()
	s.Next()
	// prov := godip.Province("lvn")

	// d := s.Graph().Edges(prov, true)
	d, _, _ := s.Unit(godip.Province("stp"))
	fmt.Println(d)
}