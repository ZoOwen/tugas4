package main

import "fmt"

type Sepeda struct {
	nama string
	ban  int
	gigi int
}

func (s *Sepeda) waktuTempuh(jarak float32) float32 {

	waktutempuh := jarak * 2.5

	return waktutempuh
}

func main() {
	arraySepeda := make([]Sepeda, 5, 5)
	arraySepeda[0] = Sepeda{nama: "bmx", ban: 2, gigi: 5}
	arraySepeda[1] = Sepeda{nama: "sepeda gunung", ban: 2, gigi: 6}
	arraySepeda[2] = Sepeda{nama: "ontel", ban: 2, gigi: 1}
	arraySepeda[3] = Sepeda{nama: "sepeda anak", ban: 3, gigi: 1}
	arraySepeda[4] = Sepeda{nama: "Sepeda santai", ban: 2, gigi: 3}
	fmt.Println(arraySepeda)

	for _, value := range arraySepeda {

		fmt.Println("jenis Sepeda = " + value.nama)
		fmt.Println("jumlah ban   = ", value.ban)
		fmt.Println("jumlah gigi  = ", value.gigi)
		fmt.Println("waktu tempuh sepeda adalah sekitar ", value.waktuTempuh(20), "perjam")
	}
}
