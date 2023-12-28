package main

import (
	"fmt"
)

func main(){
  var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i,v := range myString{
		fmt.Println(i,v)
	}
}

// 0 114
// 1 233
// 3 115
// 4 117
// 5 109
// 6 233

// UTF-8 has variable length encoding.  é uses 2 bytes
// range knows how to decode the string thingies properly

// with runes (better fof indexing/iterating strings):
// 0 114
// 1 233
// 2 115
// 3 117
// 4 109
// 5 233