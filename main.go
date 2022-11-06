package main

import (
	"fmt"

	"github.com/alpenveg/utils/newpackage"
	"github.com/alpenveg/utils/wordz"

	"github.com/alpenveg/utils/utils"
	"github.com/huandu/xstrings"
)

func main() {

	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}

	fmt.Println(xstrings.Shuffle(newpackage.City()))
	fmt.Println(newpackage.Digit())
}
