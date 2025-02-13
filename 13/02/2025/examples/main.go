package main

import (
	"fmt"
)

func f1(){
	fmt.Pritnln("this is begining of f1 func")

	fmt.Pritnln("this is end of f1 func")

}

func f2(){
	fmt.Println("this is begining of f2 func")

	fmt.Println("this is end of f2 func")

}
func f3(){
	fmt.Println("this is begining of f3 func")

	fmt.Pritnln("this is end of f3 func")

}

func main(){
	fmt.Println("Start of main")
	defer f1()
	f2()
	f3()
	fmt.Println("end of main")
}