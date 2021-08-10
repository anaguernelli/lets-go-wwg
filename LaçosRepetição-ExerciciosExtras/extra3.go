package main

import (
	"fmt"
	"strconv"
)

//Package 'strconv' implements conversions to and from string
//representations of basic data types.

func main() {
	var oi string
	for i := 0; i < 10; i++ {
		oi += strconv.Itoa(i + 1)
		fmt.Printf("%d\n", oi)
	}

}
//??
//  +=   ->   oi = oi + x 