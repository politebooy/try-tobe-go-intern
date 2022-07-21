package main

import (
	"fmt"
	"log"
	"net/http"
)

func eval(expr string) int {
	var result int
	for _, token := range expr {
		switch token {
		case '+':
			result += 1
		case '-':
			result -= 1
		case '*':
			result *= 1
		case '/':
			result /= 1
		default:
			result = int(token - '0')
		}
	}
	return result
}

func main() {
	fmt.Println(eval("1+2*3+4"))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
