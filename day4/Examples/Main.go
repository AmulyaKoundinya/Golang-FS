package main
import(
	"fmt"
)
 func f1(){
	fmt.Println("This is beg of func f1")
	fmt.Println("This is end of func f1")

}
func f2(){
	fmt.Println("This is beg of func f2")
	fmt.Println("This is end of func f2")

}
func f3(){
	fmt.Println("This is beg of func f3")
	fmt.Println("This is end of func f3")

}
func main(){
	 fmt.Println("This is beg of func main")
	f1()
	f2()
	f3()
	fmt.Println("This is end of func main")
}
