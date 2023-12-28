package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var intNum int16 = 32767
	intNum++
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2

	var result float32 = floatNum32 / float32(intNum32)
	fmt.Println(result)

	var intnum1 int = 3
	var intnum2 int = 2
	fmt.Println(intnum1 / intnum2)
	fmt.Println(intnum1 % intnum2)

	var myString string = "Hello \nWorld"
	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	var myBoolean bool
	fmt.Println(myBoolean)

	myvar := "text"
	fmt.Println(myvar)

	var1, var2 := 10, 20
	fmt.Println(var1, var2)

	const myConst string = "Hello World"
	fmt.Println(myConst)

	const pi float32 = 3.14159
	fmt.Println(pi)
}