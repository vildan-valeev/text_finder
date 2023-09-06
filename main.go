package main

import (
	"flag"
	"fmt"
	"text_finder/conv"
	"text_finder/simple"
)

func main() {
	f := flag.String("s", "simple", "Run mode")
	flag.Parse()
	s := *f
	fmt.Printf("Run mode is %s.\n", s)

	err := Run(s)
	if err != nil {
		fmt.Println(err)
	}
}

func Run(s string) error {
	if s == "conv" {
		return conv.Conveyor()
	}

	return simple.Simple()
}
