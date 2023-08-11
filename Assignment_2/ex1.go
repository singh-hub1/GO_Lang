package main
import "fmt"

func main(){
	var a int

	fmt.Println("Enetr the number")
	fmt.Scanln(&a)

	var res int
	res=pallin(a)
	if(res==a){
		fmt.Printf("%d is pallindrome number",a)
	}else{
		fmt.Printf("%d is not a palllindrome number",a)
	}

}

func pallin(x int)(int){
	var r,sum int=0,0

	for ;x>0;x=x/10{
		r=x%10
		sum=sum*10+r

	}
	return sum
}