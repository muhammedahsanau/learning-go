// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main()  {
// 	reader := bufio.NewReader(os.Stdin)

//		fmt.Print("Enter a number: ")
//		userInput, _ := reader.ReadString('a')
//		fmt.Println(userInput)
//	}
package main

import (
	"fmt"
)

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Println(i, "", j, "") 
  fmt.Print( i, "", j, "")
  fmt.Printf( i, "", j, "")
  fmt.Printf("%#v", i)
  fmt.Printf("%#v", j)
  fmt.Printf("%#v", i)
  fmt.Printf("%#v", j)
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}