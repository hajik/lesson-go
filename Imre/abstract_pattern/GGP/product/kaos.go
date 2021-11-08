package product

import (
	"ap/abstract_factory"
)

// kaos

type BajuKaos struct{}

func (k *BajuKaos) Harga() float64 {
	return 50000
}

func (k *BajuKaos) Warna() string {
	return "MERAH"
}

func (k *BajuKaos) Ukuran() abstract_factory.UkuranBaju {
	return abstract_factory.BajuL
}
