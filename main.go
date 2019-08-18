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
		fmt.Print("INSERT COMMAND::: ")

		text, _ := reader.ReadString('\n')

		fmt.Println(input.Msg(text))
	}

}
