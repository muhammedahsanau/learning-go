package main

import "fmt"

// implementing some of the data types in this file

func main()  {
	var name string = "ahsan"
	fmt.Println("Hello World name:", name,)

	name = "wali"
	fmt.Println("Hello World name:", name, )

	fmt.Printf("typeof name: %T \t", name)

	const isUserEnabled bool = true
	fmt.Println("\n Hello World isUserEnabled:", isUserEnabled,)
	fmt.Printf("typeof isUserEnabled: %T \t", isUserEnabled)

	var age int = 23
	fmt.Println("\n Hello World age:", age,)
	fmt.Printf("typeof age: %T \t", age)

	var height float32 = 5.6
	fmt.Println("\n Hello World height:", height,)
	fmt.Printf("typeof height: %T \t", height)	

	var bigFloat float64 = 5.6
	fmt.Println("\n Hello World bigFloat:", bigFloat,)
	fmt.Printf("typeof bigFloat: %T \t", bigFloat)


	var smallValue uint8 = 255
	fmt.Println("\n Hello World smallValue:", smallValue,)
	fmt.Printf("typeof smallValue: %T \t", smallValue)

	var bigValue uint16 = 65535
	fmt.Println("\n Hello World bigValue:", bigValue,)
	fmt.Printf("typeof bigValue: %T \t", bigValue)

	myNumber := 5
	fmt.Println("\n Hello World myNumber:", myNumber,)
	fmt.Printf("typeof myNumber: %T \t", myNumber)

	userName := "ahsan"
	fmt.Println("\n Hello World userName:", userName,)
	fmt.Printf("typeof userName: %T \t", userName)

	var myName = "ahsan"
	fmt.Println("\n Hello World myName:", myName,)
	fmt.Printf("typeof myName: %T \t", myName)

	myFloatValue := 5.6
	fmt.Println("\n Hello World myFloatValue:", myFloatValue,)
	fmt.Printf("typeof myFloatValue: %T \t", myFloatValue)
}