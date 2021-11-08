package GIANT

import (
	"ap/GIANT/product"
	"ap/abstract_factory"
)

// Grand Galaxy Park
type GIANT struct{}

func (g *GIANT) CreateBaju() abstract_factory.Baju {
	return &product.BajuKaos{}
}

func (g *GIANT) CreateCelana() abstract_factory.Celana {
	return &product.CelanaLevis{}
}

func (g *GIANT) CreateJaket() abstract_factory.Jaket {
	return &product.JaketLevis{}
}
