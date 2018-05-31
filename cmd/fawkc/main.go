package main

import (
	"fmt"
	"os"

	"github.com/fwip/fawk/parse"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Whoops, my anarchy symbol")
		os.Exit(1)
	}
	awkCmd := os.Args[1]
	rootNode := parse.Parse(awkCmd)

	fmt.Fprintln(os.Stderr, rootNode.String())
	fmt.Println(rootNode.ToGo())
}
