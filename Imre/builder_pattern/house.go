package main

import (
	"fmt"
	"log"
)

type Rumah struct {
	jumlahJendela int
	jumlahPintu   int
	jumlahKamar   int
	dapur         bool
}

type RumahBuilder struct {
	Rumah
}

func (r *RumahBuilder) BuildJendela() *RumahBuilder {
	r.Rumah.jumlahJendela++

	return r
}

func (r *RumahBuilder) BuildPintu() *RumahBuilder {
	r.Rumah.jumlahPintu++

	return r
}

func (r *RumahBuilder) BuildKamar() *RumahBuilder {
	r.Rumah.jumlahKamar++

	return r
}

func (r *RumahBuilder) BuildDapur() *RumahBuilder {

	r.Rumah.dapur = true

	return r

}

func (r *RumahBuilder) BuildRumah() (*Rumah, error) {

	if !r.Rumah.dapur {
		return nil, fmt.Errorf("Dapur kosong! Dapur harus dibangun!")
	}

	return &r.Rumah, nil

}

func main() {

	house := RumahBuilder{}

	res, err := house.BuildJendela().BuildDapur().BuildPintu().BuildPintu().BuildKamar().BuildRumah()

	if err != nil {
		log.Println("Failed:", err)
	}

	fmt.Println("house : ", res)

}
