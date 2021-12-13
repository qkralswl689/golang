package main

import (
	"fmt"
	"strings"
)


func multiply(a , b int)  int {
	return a * b
}

func lenAndUpper(name string) (int, string){
 return len(name) , strings.ToUpper(name)
}

func repeatMe(words ...string )  {
	fmt.Println(words)
}
// func main() {

// 	fmt.Println("Hello world!")
// 	something.SayHello()
// }

func main(){

	// constant
	// const name string = "mingki"
	// fmt.Println(name)

	// variable
	// way1
	// var name string = "hyun"
	// name = "ming"
	// fmt.Println(name)

	// way2
	// name := "hyun"
 	// name = "ming"
	// fmt.Println(name)

	// 12/13
	// fmt.Println(multiply(2,2))

	// totalLength, upperName := lenAndUpper("nico")

	// fmt.Println(totalLength,upperName)

	repeatMe("nico","hyun","ming","sunny")
}