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

	var myMap map[string]uint8 = make(map[string]uint8) // keys are string, values are unsigned 8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":32,"Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	// fmt.Println(myMap["jason"])
	var age, ok = myMap2["Jason"]
	// delete(myMap2, "Adam")
	if ok {
    fmt.Printf("Tha age is %v\n", age)
  } else {
    fmt.Println("Invalid Name!!!!")
  }

	for name := range myMap2 { // order is not guaranteed
		fmt.Printf("Name: %v\n",name)
	}

	for i,v := range intArr { // index, value
		fmt.Printf("Index: %v, Value: %v\n",i,v)
	}

	var i int = 0
	for i <10{
		fmt.Println(i)
    i++
	}

	i = 0
	for{
		if i >= 10{
			break
		}
		fmt.Println(i)
    i++
	}

	for j:=0; j<10; j++ {
		fmt.Println((j))
	}
}
