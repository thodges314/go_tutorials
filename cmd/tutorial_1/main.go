package main

import (
	"errors"
	"fmt"
)

func main(){
	// var printValue string = "Hello, World!"
	// printMe(printValue)
	var num int = 11
	var denom int = 4
	var result, remainder, err = intDivision(num, denom)
	// if(err!= nil){
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0{
	// 	fmt.Printf("The result of integer division is %v.", result)
	// } else{
	// 	fmt.Printf("The result of integer division is %v and the remainder is %v.", result, remainder)
	// }
	switch{
		case err!= nil:
      fmt.Printf(err.Error())
    case remainder == 0:
      fmt.Printf("The result of integer division is %v.", result)
    default:
      fmt.Printf("The result of integer division is %v and the remainder is %v.", result, remainder)
	}
	// in switch, break is implied
	switch remainder{
	case 0:
		fmt.Printf("The division was exact.")
	case 1,2:
		fmt.Printf("The division was close.")
	default:
		fmt.Printf("The division was not exact.")
	} 
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(num int, denom int) (int, int, error){
	var err error
	if denom == 0{
    err = errors.New("Denominator cannot be 0")
  } 
	var result int = num / denom
	var remainder int = num % denom
	return result, remainder, err
}