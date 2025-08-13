package main

// In Go, defer statements are executed in LIFO (Last In, First Out) order - like a stack.
// This means the last defer statement you write will be the first one to execute when the function

func main() {
	defer println("test 1") // Added to defer stack: [test 1]
	defer println("test 2") // Added to defer stack: [test 2, test 1]
	defer println("test 3") // Added to defer stack: [test 3, test 2, test 1]
	defer println("test 4") // Added to defer stack: [test 4, test 3, test 2, test 1]

	// When function ends, defer stack executes from top to bottom:
}
