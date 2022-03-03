package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type LaunchMenuIterator struct {
	pos int
	s []MenuItem
}

func NewLaunchMenuIterator(s []MenuItem) Iterator {
	return &LaunchMenuIterator{s: s}
}

func (l *LaunchMenuIterator) HasNext() bool {
	return len(l.s) > l.pos
}

func (l *LaunchMenuIterator) Next() interface{} {
	result := l.s[l.pos]
	l.pos++
	return result
}

type DinnerMenuIterator struct {
    pos int
    keys []string
	m map[string]MenuItem
}

func NewDinnerMenuIterator(m map[string]MenuItem) Iterator {
	keys := make([]string,0)

	for key, _ := range m {
		keys = append(keys, key)
	}
	return &DinnerMenuIterator{m: m, keys: keys}
}

func (d DinnerMenuIterator) HasNext() bool {
	return len(d.m) > d.pos
}

func (d *DinnerMenuIterator) Next() interface{} {
	result := d.m[d.keys[d.pos]]
	d.pos++
	return result
}

type Menu interface {
    CreateIterator() Iterator
}

type MenuItem struct {
    Name string
    Description string
    IsVegi bool
    Price float64
}

func NewMenuItem(name string, description string, isVegi bool, price float64) *MenuItem {
	return &MenuItem{Name: name, Description: description, IsVegi: isVegi, Price: price}
}

type LunchMenu struct {
    MenuItems []MenuItem
}

func (l *LunchMenu) CreateIterator() Iterator {
	return NewLaunchMenuIterator(l.MenuItems)
}

type DinnerMenu struct {
    MenuItems map[string]MenuItem
}

func (d DinnerMenu) CreateIterator() Iterator {
	return NewDinnerMenuIterator(d.MenuItems)
}

func main() {
	sliceMenu :=[]MenuItem{
		{
			Name:        "김치찌개",
			Description: "돼지고기 김치찌개",
			IsVegi:      false,
			Price:       8000,
		},
		{
			Name:        "된장찌개",
			Description: "차돌박이 김치찌개",
			IsVegi:      false,
			Price:       10000,
		},
	}

	Launch := LunchMenu{
		MenuItems: sliceMenu,
	}
	LaunchIter := Launch.CreateIterator()
	fmt.Println("[점심메뉴]")
	for LaunchIter.HasNext() {
		fmt.Println(LaunchIter.Next())
	}

	MapMenu := make(map[string]MenuItem)
	MapMenu["menu1"] = MenuItem{
		Name:        "돈까쓰",
		Description: "돼지고기 등심 돈까쓰",
		IsVegi:      false,
		Price:       9500,
	}
	MapMenu["menu2"] = MenuItem{
		Name:        "티본 스테이크",
		Description: "한우 티본 스테이크 900g",
		IsVegi:      false,
		Price:       15000,
	}

	Dinner := DinnerMenu{
		MenuItems: MapMenu,
	}
	DinnerIter := Dinner.CreateIterator()
	fmt.Println("[저녁메뉴]")
	for DinnerIter.HasNext() {
		fmt.Println(DinnerIter.Next())
	}

}
