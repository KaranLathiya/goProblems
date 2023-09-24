package main 
import  (
	"fmt"
	"time"
)

func main(){
	var arr []int=[]int{1, 23, 29, 17, 4, 9, 11, 13}
	ch := make(chan int, 1)
	go func(){
		for i:=0;i<8;i++{
		fmt.Println(time.Now(),"sending")	
		ch<-arr[i]
		fmt.Println(time.Now(),"sent")	
		}
		close(ch)
	}()

	for {
		select {
		case v, open := <-ch:
			if !open {
				fmt.Print(time.Now(),"done")
				return
			}
			if v%2==0{
				fmt.Println(v," is even")
			} else{
				fmt.Println(v," is odd")
			}
			for i:=2;i<(v/2+1);i++{
				if v%i==0{
					fmt.Println(v,"is non-prime")
					break
				}	else if i==v/2{
					fmt.Println(v,"is prime")
					break
				}
			}
			fmt.Println(time.Now(), "received", v)
		}
	}
}