package product

import (
	"ap/abstract_factory"
)

// jaket

type JaketLevis struct{}

func (j *JaketLevis) Harga() float64 {
	return 500000
}

func (j *JaketLevis) Warna() string {
	return "BIRU DONGKER"
}

func (j *JaketLevis) Ukuran() abstract_factory.UkuranJaket {
	return abstract_factory.JaketL
}

func (J *JaketLevis) Jenis() abstract_factory.JenisJaket {
	return abstract_factory.JenisJaketJepang
}
