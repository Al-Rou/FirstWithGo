package main

import "fmt"

//This function checks if an entered word is a palindrome or not
func main() {
	var input string
	//Prompt the user to enter a word
	fmt.Print("Enter your word: ")
	//Takes the word and stores in a variable of type string
	fmt.Scan(&input)
	//Compares the symmetry in the word
	for i := 0; i < len(input)/2; i++ {
		//If one character is not the same as its symmetric associated one, the program stops and
		//outputs an appropriate message
		if (input[i] != input[len(input)-1-i]) && (input[i] != input[len(input)-1-i]-32) && (input[i]-32 != input[len(input)-1-i]) {
			fmt.Print("The entered word is NOT palindrome!")
			return
		}
	}
	//If the program comes out of the loop, it means there is a symmetry in the word. Hence a palindrome!
	fmt.Print("It is palindrome!")
}
