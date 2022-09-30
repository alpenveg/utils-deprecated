package main

import (
	"fmt"
	"gopackages/newpackage"

	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println(xstrings.Shuffle(newpackage.City()))
	fmt.Println(newpackage.Digit())
}
