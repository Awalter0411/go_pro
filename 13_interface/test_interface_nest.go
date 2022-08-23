package main

import "fmt"

type Swimer interface {
	swim()
}

type Flyer interface {
	fly()
}

type FlyFisher interface {
	Swimer
	Flyer
}

type FlyFish struct {
}

func (f FlyFish) swim() {

}

func (f FlyFish) fly() {

}

func main() {
	var f FlyFisher
	f = FlyFish{}
	fmt.Printf("f: %v\n", f)
}
