package main

import (
	"fmt"
	"oop/encap"
	"oop/inhe"
)

func main() {
	usr1 := encap.User{}
	usr1.SetNama("Jerry")
	fmt.Println(usr1)
	opr1 := inhe.Operator{}
	opr1.SetNama("Jono")
	fmt.Println(opr1)

	// trs1 := inhe.Transaksi{}
	// // trs1.Pelanggan.Hp =
	// opr2 := inhe.Operator{}
	// // opr2.Hp =

	// tkn1 := inhe.Teknisi{}

	// var fWork poly.FieldWork
	fmt.Print("ini kenalan user ")
	opr1.User.Kenalan()
	fmt.Print("ini kenalan operator ")
	opr1.Kenalan()
	// fWork = opr1
	// fWork = tkn1

	// usr1.
}
