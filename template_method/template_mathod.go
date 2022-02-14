package main

import "fmt"

type MakeBeverageTemplate interface {
	PrepareRecipe()
}

type MakeBeverage interface {
    BoilWater()
    PourInCup()
    Brew()
    AddCondiments()
}

type caffeineBeverage struct {
	MakeBeverage
}

func (c caffeineBeverage) BoilWater()  {
	fmt.Println("물 끓이는 중...")
}

func (c caffeineBeverage) PourInCup()  {
	fmt.Println("컵에 따르는 중...")
}

func (c *caffeineBeverage) PrepareRecipe()  {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

type Coffee struct {
	MakeBeverage
}

func (c Coffee) Brew() {
	fmt.Println("필터로 커피를 우려내는 중...")
}

func (c Coffee) AddCondiments() {
	fmt.Println("우유와 설탕을 추가하는 중...")
}

type Tea struct {
	MakeBeverage
}

func (t Tea) Brew() {
	fmt.Println("차 준비중...")
}

func (t Tea) AddCondiments() {
	fmt.Println("커피 준비중...")
}

func main() {
	coffee := caffeineBeverage{&Coffee{}}
	tea := caffeineBeverage{&Tea{}}

	coffee.PrepareRecipe()
	tea.PrepareRecipe()
}
