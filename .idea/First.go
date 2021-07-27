package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter your word: ")
	fmt.Scan(&input)
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i]{
			fmt.Print("The entered word is NOT palindrome!")
			return
		}
	}
	fmt.Print("It is palindrome!")
}
