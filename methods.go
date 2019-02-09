package main

import (
	"fmt"
	"strings"
)


type human struct {
	name   string
	health int
}

type worker interface {
	work()
	say(string)
	info()
	sleep()
	didNotSleep()
}

func (character human) say(text string) {
	b:= []int{1,2}
	a:=string(b)
	fmt.Println(character.name + " says: " + text)
}

func (character *human) sleep() {
	character.health += 99
}

func (character human) didNotSleep() {
	character.health += 99
}

func (character human) work() {
	fmt.Println(character.name, "is working")
}
func (character *human) info() {
	fmt.Println("name:", character.name)
	fmt.Println("health: ", character.health)
}
func main() {
	var peter worker = &human{"Peter", 1}
	npc := peter
	peter.say("hello")
	peter.info()
	peter.didNotSleep()
	peter.info()
	peter.sleep()
	peter.info()
	peter.work()
	npc.say("Guten Morgen!")

}
