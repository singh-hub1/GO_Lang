//1. .WAP in Go to create buffered channel, store few values in it and find channel capacity and length. Read values from channel and find modified length of a channel.

package main
import "fmt"
 func main(){
	
	
		m:=make(chan string,3)
		m<-"vishal singh"
		m<-"Aditi sharama"
		m<-"jasu"

	

	fmt.Println("Length of the channel is :: ",len(m))
	fmt.Println("Capacity of the channel is :: ",cap(m))
	

 }
 