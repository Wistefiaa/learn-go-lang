package main

import (
	"fmt"

	"github.com/Wistefiaa/learn-go-lang/wistefiaa" // after / is the folder name
)

func main() {
	fmt.Println("Hello world!")
	wistefiaa.SayhelloCEDT() //before . is the package name in that folder
	/* there're 4 methods to declare a variable
	   1. var [name] [type]
	   2. var [name] [type] = [initial value]
	   3. var [name] = [initial value]
	   4. [name] := [initial value]
	   for the data type:
	   int32(int), int64, uint32, uint64 (u mean unsigned)
	   float32(float), float64
	   bool
	   string
	*/
	/*
	  if-else condition
	  if score >= 70 {
	  	fmt.Println("PASS")
	  } else{ always behind the if's '}'
	 	fmt.Println("FAIL")
	  }
	*/
	score := 80
	if score >= 70 {
		fmt.Println("PASS")
	} else {
		fmt.Println(("FAIL"))
	}
	/*
		switch case:
			switch [variable]{
			case value1:
				action for value 1
			case value2:
				action for value 2
			default:
				action for no case match
			}

	*/
	for i := 0; i < 10; i++ { //for loop
		fmt.Printf("Now is %d", i)
	}
	i := 1
	for i > 7 { // while loop
		i++
	}
}
