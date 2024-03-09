package pakexample

import "fmt"

// exported function name starts with capital letter
func SayHello(name string) {
	fmt.Printf(getGreetingString(), name)
}

// not exported function
func getGreetingString() string {
	return "Hello %v!\n"
}
