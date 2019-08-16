package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ggalihpp/galactic-market/input"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nINSERT COMMAND::: ")

		text, _ := reader.ReadString('\n')
		// r, e := roman.Rtoi(text)
		// if e != nil {
		// 	fmt.Println("ERROR!!!:: ", e.Error())
		// } else {
		// 	fmt.Println(r)
		// }

		input.Msg(text)
	}

}
