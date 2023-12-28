package main

import "fmt"

func main(){
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{1,2,3}
	fmt.Printf("The length is %v with capaticy %v.\n",len(intSlice),cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Printf("The length is %v with capaticy %v.\n",len(intSlice),cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length is %v with capaticy %v.\n",len(intSlice),cap(intSlice))
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // 3 is length, 8 is capacity
	fmt.Printf("The length is %v with capaticy %v.\n",len(intSlice3),cap(intSlice3))
	fmt.Println(intSlice3)
}
