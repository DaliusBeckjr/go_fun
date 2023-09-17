package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	// sendsSoFar = sendsSoFar + sendsToAdd
	sendsSoFar++
}

// Assignment
// fix the bug to accurately track the number of SMS messages sent

// Alter the incrementSends function so that it returns the result after incrementing the sendsSoFar.
// Alter main() to capture the return value from incrementSends() and overwrite the previous sendsSoFar value.
