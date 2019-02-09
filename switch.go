package main

import "fmt"

func main() {

	switch env := "dev"; env {
	case "dev":
		fmt.Println("env is dev")
	case "production":
		fmt.Println("AT production")
	default:
		fmt.Println("middle of nowhere")
	}

	switch {
	case isTure():
		fmt.Println("success")
	default:
		fmt.Println("failed")
	}
}

func isTure() bool {
	return true
}
