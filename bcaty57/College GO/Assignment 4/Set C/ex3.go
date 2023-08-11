// of the embedding interfaces
package main

import "fmt"

type authordetails interface{
	details()
}

type authorarticles interface{
	articles()
}

type finaldetails interface{
	authorarticles
	authordetails
}

type author struct{
	name string 
	roll int
	college string
	art int

}

func(a author )details(){

	fmt.Printf("author name %s",a.name)
	fmt.Printf("author roll %d",a.roll)
	fmt.Printf("author college %s",a.college)
}

func (a author) articles(){
	fmt.Printf("author articles %d",a.art)
}


func main(){
	val:=author{
		name: "vishal singh",
		roll:13459,
		college: "agc",
		art: 12,
	}

	var f finaldetails=val
	f.details()
	f.articles()

}