package main

import (
	"fmt"
	"strconv"
)

//Says hello
func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

//Introduce yourself
func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

//Guessing your age
func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7. (use spaces between numbers)")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

//Counts while n
func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

//Little Quiz Game
func startQuiz() {
	var ans, right int

	fmt.Println("Let's test your programming knowledge.?")
	fmt.Println("Why do we use methods?")
	fmt.Println("1. To repeat a statement multiple times.")
	fmt.Println("2. To decompose a program into several small subroutines.")
	fmt.Println("3. To determine the execution time of a program.")
	fmt.Println("4. To interrupt the execution of a program.")

	right = 2

	for {
		fmt.Scan(&ans)
		if ans == right {
			break
		} else {
			fmt.Println("Please, try again.")
		}
	}
	fmt.Println("Correct!")
}

//Bye Bye
func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Aid", "2022") // change it as you need
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}
