package main

import "fmt"

type Beverage interface {
    Cost() float64
    GetDescription() string
}

type Espresso struct {

}

func NewEspresso() Beverage {
	return &Espresso{}
}

func (e Espresso) Cost() float64 {
	return 1.99
}

func (e Espresso) GetDescription() string {
	return "에스프레소"
}

type HouseBlend struct {

}

func NewHouseBlend() Beverage {
	return &HouseBlend{}
}

func (h HouseBlend) Cost() float64 {
	return 0.89
}

func (h HouseBlend) GetDescription() string {
	return "하우스 블렌드 커피"
}

type DarkRoast struct {

}

func NewDarkRoast() Beverage {
	return &DarkRoast{}
}

func (d DarkRoast) Cost() float64 {
	return 1.25
}

func (d DarkRoast) GetDescription() string {
	return "다크 로스트 커피"
}

type Decaf struct {

}

func NewDecaf() Beverage {
	return &Decaf{}
}

func (d Decaf) Cost() float64 {
	return 1.18
}

func (d Decaf) GetDescription() string {
	return "디카페인 커피"
}

// 데코레이터
type CondimentDecorator interface {
    Beverage
}

type Mocha struct {
    beverage Beverage
}

func (m Mocha) Cost() float64 {
	return 0.50 + m.beverage.Cost()
}

func (m Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", 모카"
}

func NewMocha(beverage Beverage) CondimentDecorator {
	return &Mocha{beverage: beverage}
}

type Soy struct {
    beverage Beverage
}

func NewSoy(beverage Beverage) CondimentDecorator {
	return &Soy{beverage: beverage}
}

func (s Soy) Cost() float64 {
	return 0.15 + s.beverage.Cost()
}

func (s Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", 두유"
}

type Whip struct {
    beverage Beverage
}

func NewWhip(beverage Beverage) CondimentDecorator {
	return &Whip{beverage: beverage}
}

func (w Whip) Cost() float64 {
	return w.beverage.Cost() + 0.28
}

func (w Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", 휘핑크림"
}

type SteamMilk struct {
    beverage Beverage
}

func NewSteamMilk(beverage Beverage) CondimentDecorator {
	return &SteamMilk{beverage: beverage}
}

func (s SteamMilk) Cost() float64 {
	return s.beverage.Cost() + 0.10
}

func (s SteamMilk) GetDescription() string {
	return s.beverage.GetDescription() + ", 스팀 우유"
}

func main() {
	beverage := NewEspresso()
	fmt.Printf("%s $%f\n", beverage.GetDescription(), beverage.Cost())

	beverage2 := NewDarkRoast()
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewWhip(beverage2)
	beverage2 = NewSteamMilk(beverage2)
	fmt.Printf("%s $%f\n", beverage2.GetDescription(), beverage2.Cost())

	beverage3 := NewHouseBlend()
	beverage3 = NewSoy(beverage3)
	beverage3 = NewMocha(beverage3)
	beverage3 = NewWhip(beverage3)
	fmt.Printf("%s $%f\n", beverage3.GetDescription(), beverage3.Cost())
}
