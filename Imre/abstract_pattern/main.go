package main

import (
	"ap/GGP"
	"ap/abstract_factory"
	"fmt"
)

func main() {

	var cloth abstract_factory.ClothFactory

	cloth = &GGP.GGP{}

	baju := cloth.CreateBaju()

	fmt.Println(baju.Harga())
}
