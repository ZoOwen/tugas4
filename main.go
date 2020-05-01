package main

import "fmt"

type Sepeda struct {
	warna string
	ban   int
	gigi  int
}

func (s *Sepeda) waktuTempuh(jarak float32) float32 {

	waktutempuh := jarak * 2.5

	return waktutempuh
}

func main() {
	arraySepeda := make([]Sepeda, 5, 5)
	arraySepeda[0] = Sepeda{warna: "biru", ban: 2, gigi: 5}
	arraySepeda[1] = Sepeda{warna: "hitam", ban: 2, gigi: 6}
	arraySepeda[2] = Sepeda{warna: "putih", ban: 2, gigi: 1}
	arraySepeda[3] = Sepeda{warna: "ungu", ban: 3, gigi: 1}
	arraySepeda[4] = Sepeda{warna: "merah", ban: 2, gigi: 3}
	fmt.Println(arraySepeda)

	for _, value := range arraySepeda {

		fmt.Println("Warna Sepeda = " + value.warna)
		fmt.Println("jumlah ban   = ", value.ban)
		fmt.Println("jumlah gigi  = ", value.gigi)
		fmt.Println("waktu tempuh sejauh 20 km adalah ", value.waktuTempuh(20), "menit ")
	}
}
