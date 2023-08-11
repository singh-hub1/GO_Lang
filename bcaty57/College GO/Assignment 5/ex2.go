package main

import "fmt"

func prod(my chan int){
	for i := 0; i < 10; i++ {
		my<-i
	}
	close(my)
}
func main(){
	ch:=make(chan int)
	go prod(ch)
	for{
		  v,ok:=<-ch
		  if ok==false{
			  break
		  }
		  fmt.Println(v,ok)
	}
}