package main

import "fmt"

func main() {
	fmt.Println("\nHello World\nThis is a test\n")
	foo()
	fmt.Println("Something else")
	for i := 0; i<50; i++ {
		if i %5 == 0{
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
