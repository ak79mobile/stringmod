package main

import (
	
	"fmt"
	str "github.com/goDummies1/stringmod/quotes"
	qt "github.com/goDummies1/stringmod/strings"
)

func main() {
	
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)
	
	fmt.Println(qt.GetEmoji("turtle"))
}

