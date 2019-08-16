package main

import (
	"fmt"

	"github.com/ggalihpp/galactic-market/roman"
)

func main() {

	x, err := roman.Rtoi("MCMXLIV")
	if err != nil {
		fmt.Println("ERRORTJCUSKS:: ", err.Error())
	}

	fmt.Println("HASILNYA:: ", x)
}
