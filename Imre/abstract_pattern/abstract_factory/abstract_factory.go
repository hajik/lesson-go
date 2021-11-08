package abstract_factory

type KriteriaUmum interface {
	Harga() float64
	Warna() string
}

// baju

type UkuranBaju string

const (
	BajuL  UkuranBaju = "L"
	BajuM  UkuranBaju = "M"
	BajuXL UkuranBaju = "XL"
)

type Baju interface {
	KriteriaUmum
	Ukuran() UkuranBaju
}

//celana

type UkuranCelana string

const (
	CelanaL  UkuranCelana = "L"
	CelanaM  UkuranCelana = "M"
	CelanaXL UkuranCelana = "XL"
)

type Celana interface {
	KriteriaUmum
	Ukuran() UkuranCelana
}

// jaket

type UkuranJaket string

const (
	JaketL  UkuranJaket = "L"
	JaketM  UkuranJaket = "M"
	JaketXL UkuranJaket = "XL"
)

type JenisJaket string

const (
	JenisJaketAmerika JenisJaket = "AMERIKA"
	JenisJaketJepang  JenisJaket = "JEPANG"
)

type Jaket interface {
	KriteriaUmum
	Ukuran() UkuranJaket
	Jenis() JenisJaket
}

// factory
type ClothFactory interface {
	CreateCelana() Celana
	CreateBaju() Baju
	CreateJaket() Jaket
}
