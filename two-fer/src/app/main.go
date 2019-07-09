package main

import (
	"fmt"
	"twofer"
)

func main() {
	x := twofer.ShareWith("BOB")
	fmt.Println(x)
}
