package encap

import "fmt"

// Encapsulation
// obyek bisa menyimpan informasi.
// obyek HANYA MENAMPILKAN INFORMASI YANG DIBUTUHKAN
type User struct {
	nik    int
	nama   string
	alamat string
	Hp     string
	umur   int
}

func (u *User) UpdateNIK(nikBaru int) {
	u.nik = nikBaru
}

func (u *User) SetNama(nama string) {
	u.nama = nama
}

func (u *User) Kenalan() {
	fmt.Println("Halo nama saya", u.Hp)
}

// func main() {
// 	usr1 := User{1234567, "Jerry", "Pasuruan", "081234", 26}
// 	usr1.nik = 678990
// 	usr1.Umur = 30
// }
