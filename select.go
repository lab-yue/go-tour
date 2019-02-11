package main

import (
	"fmt"
	"time"
)

func living(study, work, sleep chan string) {
	task := []string{"meeting", "reading"}

	for {
		select {
		case work <- task[0]:
			fmt.Println("work: " + task[0])
		case study <- task[1]:
			fmt.Println("study: ", task[1])
		case <-sleep:
			fmt.Println("end of day")
			return
		default:
			fmt.Print(".")
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func main() {
	study := make(chan string)
	work := make(chan string)
	sleep := make(chan string)

	go func() {
		fmt.Println(<-work)
		fmt.Println(<-study)
		sleep <- "rest"
	}()

	living(study, work, sleep)
}
