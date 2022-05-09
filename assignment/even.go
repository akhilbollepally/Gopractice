package main

import "fmt"

func main() {
	numberslist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("One way of using for loop:")
// one way using for loop
	for _, element := range numberslist {
		if (element%2==0) {
			fmt.Println(element,"is even number"	)

		}else{
			fmt.Println(element,"is odd number")
		}

	}
	fmt.Println("Another way of using for loop:")
	// Another way of using for loop
	for i:=0;i<len(numberslist);i++{
		if(numberslist[i]%2==0){
			fmt.Println(numberslist[i],"is even number")
		}else{
			fmt.Println(numberslist[i],"is odd number")
		}
	}
}