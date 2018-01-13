package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/binatify/whoami/checker"
	"github.com/binatify/whoami/types"
)

var (
	l bool
)

func main() {
	flag.BoolVar(&l, "l", false, "list all basic types description in go")
	flag.Parse()

	fmt.Println()
	fmt.Println()

	if l {
		fmt.Println(">>>> All basic types are:")
		fmt.Println()
		types.List("bool", "",
			"uint8", "uint16", "uint32", "uint64", "",
			"int8", "int16", "int32", "int64", "",
			"float32", "float64", "",
			"complex64", "complex128", "",
			"byte", "rune",
			"uint", "int", "uintptr", "",
			"string")

		fmt.Println()
		fmt.Println()
		return
	}

	var (
		reader = bufio.NewReader(os.Stdin)
		text   = ""
	)

	for {
		fmt.Print("Who am I?: ")
		text, _ = reader.ReadString('\n')

		text = strings.TrimSpace(text)
		if text == "bye" {
			break
		}

		checker.Assert(text)
	}

	fmt.Println()
	fmt.Println("See you next time!")
}
