// With the var keyword:
// Syntax
// var array_name = [length]datatype{values} // here length is defined

// or

// var array_name = [...]datatype{values} // here length is inferred

// 2. With the := sign:
// Syntax
// array_name := [length]datatype{values} // here length is defined

// or

// array_name := [...]datatype{values} // here length is inferred

package main

import (
	"fmt"
)

func main() {
  var arr1 = [3]int{}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)

  
  arr4 := [5]int{} //not initialized
  arr5 := [5]int{1,2} //partially initialized
  arr6 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr4)
  fmt.Println(arr5)
  fmt.Println(arr6)
}