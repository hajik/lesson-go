package product

import (
	"ap/abstract_factory"
)

// celana

type CelanaLevis struct{}

func (c *CelanaLevis) Harga() float64 {
	return 200000
}

func (c *CelanaLevis) Warna() string {
	return "BIRU TUA"
}

func (c *CelanaLevis) Ukuran() abstract_factory.UkuranCelana {
	return abstract_factory.CelanaXL
}
