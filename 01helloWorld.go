package main

import (
	"fmt"

	"rsc.io/quote"
)

func printHello() {
	fmt.Println("Holis mundo")
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
}
