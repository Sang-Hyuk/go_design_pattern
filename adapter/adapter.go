package main

import "fmt"

type Duck interface {
    Quack()
    Fly()
}

type MallardDuck struct {

}

func (m MallardDuck) Quack() {
	fmt.Println("Quack...")
}

func (m MallardDuck) Fly() {
	fmt.Println("fly...")
}

type Turkey interface {
    Gobble()
    Fly()
}

type WildTurkey struct {

}

func (w WildTurkey) Gobble() {
	fmt.Println("Gobble gobble...")
}

func (w WildTurkey) Fly() {
	fmt.Println("flying short distance...")
}

type TurkeyAdapter struct {
    turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (t TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t TurkeyAdapter) Fly() {
	t.turkey.Fly()
}

func TestDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}

func main() {
	duck := MallardDuck{}

	turkey := WildTurkey{}
	turkeyAdapter := NewTurkeyAdapter(turkey)

	fmt.Println("칠면조가 말하길....")
	turkey.Gobble()
	turkey.Fly()

	fmt.Println("오리가 말하길....")
	TestDuck(duck)

	fmt.Println("칠면조어댑터가 말하길....")
	TestDuck(turkeyAdapter)
}
