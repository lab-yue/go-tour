package main

import "fmt"

func downloadSteam(c chan string) {
	c <- "downloading steam, "
}
func playAOE2(c chan string) {
	c <- "playing Age of Empires II"
}

func main() {
	channel()
	buffer()
	year()
}

func channel() {
	c := make(chan string)
	go playAOE2(c)
	go downloadSteam(c)
	fmt.Println(<-c, <-c)
}

func buffer() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)
}

func year() {
	y := make(chan int)
	go sendYear(y)
	fmt.Println("start sending")
	for year := range y {
		fmt.Println(year)
	}

}

func sendYear(y chan int) {

	for year := 2019; year > 1999; year-- {
		y <- year
	}
	close(y)
}
