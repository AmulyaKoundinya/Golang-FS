package main
import "fmt"
type student struct{
	Name string
	RollNo int
	Marks float64
}
func main(){
	st := student{Name: "Anjana" , RollNo: 123 , Marks: 98.2}
	fmt.Println("Name:",st.Name, "\nRollNo:", st.RollNo, "\nMarks:",st.Marks)
}
/*func forPythonStyle(){
	strings := []string{"hi", "hey" , "bye" , "NIE"}
	for _, s := range strings {
		fmt.Println(s)
	}
	}
	

	//var num1,num2 int
	//fmt.Println("Enter two numbers(Do not enter zero as the second number ):")

	//fmt.Scanln(&num1,&num2)
	//sum := num1+num2
	 fmt.Println("Sum is:",sum)
	diff := num1-num2
	fmt.Println("Diff is:",diff)
	prod := num1*num2
	fmt.Println("Product is:",prod)
	div := num1/num2
	fmt.Println("Quotient is:",div)*/

	
	

