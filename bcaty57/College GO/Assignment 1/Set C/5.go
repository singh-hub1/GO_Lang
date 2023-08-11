//5. WAP in go language to check whether first string is substring of another string or not.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("1st String : ")
	fmt.Scan(&str1)

	fmt.Print("2nd String : ")
	fmt.Scan(&str2)

	result := strings.Contains(str2, str1)

	if result {
		fmt.Println("\n1st string is substring of 2nd string")
	} else {
		fmt.Println("\n1st string is NOT substring of 2nd string")
	}

}
