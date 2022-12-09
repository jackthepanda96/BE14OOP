package inhe

import (
	"fmt"
	"oop/encap"
)

// inheritance -> pewarisan
type Operator struct {
	encap.User
	Shift int
}

func (o Operator) Begadang() {
	fmt.Println("Begadang bro")
}
func (o Operator) Piket(hari string) {
	fmt.Println("Aku piket hari", hari)
}
func (o *Operator) Kenalan() {
	fmt.Println("Halo nama o", o.Hp)
}

type Teknisi struct {
	encap.User
}

func (o Teknisi) Begadang() {
	fmt.Println("Begadang bro")
}
func (o Teknisi) Piket(hari string) {
	fmt.Println("Aku piket hari", hari)
}

type Supervisor struct {
	encap.User
	IsInspeksi bool
}

type Transaksi struct {
	Pelanggan encap.User
	Harga     int
}
