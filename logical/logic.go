package main

import (
	"fmt"
)

func main() {
	//PALINDROME
	var str string = "angin"
	result := isPalindrome(str)
	if result == true {
		fmt.Println("==========")
		fmt.Println("palindrome")
		fmt.Println("==========")
	} else {
		fmt.Println("==============")
		fmt.Println("not palindrome")
		fmt.Println("==============")
	}
	// reverse word
	fmt.Println("=============verse word===========")
	fmt.Printf("%v\n", reverse("I am A Great human"))
	fmt.Println("==================================")

	//FizzBuzz
	fmt.Println("===========FizzBuzz==========")
	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			// Multiple of 3
			fmt.Printf("fizz")
		}
		if i%5 == 0 {
			// Multiple of 5
			fmt.Printf("buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			// Neither, so print the number itself
			fmt.Printf("%d", i)
		}

		// A trailing new line (so both fizz + buzz can be printed on the same line)
		fmt.Printf("\n")

	}
	//Leap Year
	fmt.Println("===========Leap Year==========")
	fmt.Println("Enter Year : ")
	var year int
	fmt.Scanln(&year)

	fmt.Println(checkYear(year))

}

//palindrome
func isPalindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {

		if input[i] != input[len(input)-i-1] {

			return false
		}
	}
	return true
}

// reverse word
func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func checkYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
