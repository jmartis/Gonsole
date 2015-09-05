package main 

import "fmt"

func main() {
    fmt.Println("hello world")
    
    var age int  = 10;
    var favNum float64 = 1.618058
    
    fmt.Println(age, favNum)
    
    fmt.Println("6 + 4 =", 6 + 4)
    fmt.Println("6 - 4 =", 6 - 4)
    fmt.Println("6 * 4 =", 6 * 4)
    fmt.Println("6 / 4 =", 6 / 4)
    fmt.Println("6 % 4 =", 6 % 4)
    
    const pi float64 = 3.14159265
    
    var username string = "User Name"
	fmt.Println("My" + username + " is unique")	
	fmt.Printf("%.3f \n", pi)
	   
    var numOne = 1.000
    var num99 = .9999
    
    fmt.Println(numOne - num99)
    
}

