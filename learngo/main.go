package main

import (
	"fmt"
	"strings"
)


func multiply(a , b int)  int {
	return a * b
}

// 12/13
// func lenAndUpper(name string) (int, string){
//  return len(name) , strings.ToUpper(name)
// }

func repeatMe(words ...string )  {
	fmt.Println(words)
}

// 12/15
func lenAndUpper(name string) (lenght int, uppercase string){

	defer  fmt.Println("I'm done")
	
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
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

	// repeatMe("nico","hyun","ming","sunny")

	totalLength, up := lenAndUpper("nico")

	fmt.Println(totalLength, up)
}