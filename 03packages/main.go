package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"
)

// implementing some of the data types in this file

func main()  {
	complaintId := uuid.New().String()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("wellcome to the bank, What is your name ?")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for banking with us", input, " what is your complaint?")

	userAnswer, _ := reader.ReadString('\n')
	println(userAnswer)
	fmt.Printf(`Thanks for your feedback, we will shortly look into your query \n 
		please note the complain number(%d) for future refference`, complaintId)
}