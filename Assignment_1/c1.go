//WAP in go language to concatenate two strings using pointers
package main

import (
	"fmt"
)

func main() {
	str1 := "vishal"
	str2 := "singh"

	var a *string = &str1
	var b *string = &str2

	fmt.Println(" Two strings are :: ", (*a + " " + *b))

}
