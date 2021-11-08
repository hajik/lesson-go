package GGP

import (
	"ap/GGP/product"
	"ap/abstract_factory"
)

// Grand Galaxy Park
type GGP struct{}

func (g *GGP) CreateBaju() abstract_factory.Baju {
	return &product.BajuKaos{}
}

func (g *GGP) CreateCelana() abstract_factory.Celana {
	return &product.CelanaLevis{}
}

func (g *GGP) CreateJaket() abstract_factory.Jaket {
	return &product.JaketLevis{}
}
