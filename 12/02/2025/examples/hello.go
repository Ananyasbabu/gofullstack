package main

import "fmt"

type student struct {
	Name string 
	Rgno float64
	Dept string
}

func main() {
//  ifelseDemo()
//  forThreeVarDemo()
// greater()
// forConditional()
// forPythonStyle()
st := student{Name:"student1", Rgno:12.3 , Dept: "cs"}
fmt.Println("name:",st.Name,"\nrgno:",st.Rgno,"\ndepart:",st.Dept)
}

func ifelseDemo(){
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("Adult")
	}else {
		fmt.Println("Minor")
	}
}

func forThreeVarDemo() {
	sum:=0
	for i :=1; i < 5; i++ {
		sum+=i
	}
	fmt.Println(sum)
}

func greater() {
	var a,b,c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if a>b && a>c {
		fmt.Println(a,"is greatest")
	}else if b>c && b>a {
		fmt.Println(b,"is greatest")
	}else {
		fmt.Println(c,"is greatest")
	}

}

func forConditional() {
	n:=1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}

func forPythonStyle() {
	ar :=[]string{"hello", "world" , "Golang","NIe"}
	for _, s := range ar {
		fmt.Println(s)
	}
	arr :=[]int{1,3,34,3}
	for i , s :=range arr {
		fmt.Println(i,s)
	}
}

