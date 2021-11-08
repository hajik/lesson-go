package product

import (
	"ap/abstract_factory/"
)

// jaket

type JaketLevis struct{}

func (j *JaketLevis) Harga() float64 {
	return 60000
}

func (j *JaketLevis) Warna() string {
	return "PUTIH BIRU"
}

func (j *JaketLevis) Ukuran() abstract_factory.UkuranJaket {
	return abstract_factory.JaketXL
}

func (J *JaketLevis) Jenis() abstract_factory.JenisJaket {
	return abstract_factory.JenisJaketAmerika
}
