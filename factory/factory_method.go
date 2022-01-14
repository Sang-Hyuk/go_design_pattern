package main

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type BasePizza struct {
	Name string
	Dough string
	Sauce string
	Toppings []string
}

func (b BasePizza) Prepare() {
	fmt.Printf("%s 피자 준비중...\n", b.Name)
	fmt.Printf("%s 도우 만드는 중...\n", b.Dough)
	fmt.Printf("%s 소스 넣는 중...\n", b.Sauce)
	fmt.Println("토핑 추가하는 중: ")
	for _, topping := range b.Toppings {
		fmt.Printf("\t%s\n", topping)
	}
}

func (b BasePizza) Bake() {
	fmt.Println("350도에서 25분 동안 굽기...")
}

func (b BasePizza) Cut() {
	fmt.Println("피자를 대각선 8조각으로 자르기...")
}

func (b BasePizza) Box() {
	fmt.Println("피자를 상자에 담기...")
}

func (b BasePizza) GetName() string {
	return b.Name
}

type NYStyleCheesePizza struct {
	BasePizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{BasePizza: BasePizza{
		Name:     "뉴욕치즈피자",
		Dough:    "씬 크러스트 도우",
		Sauce:    "마리나라 소스",
		Toppings:[]string{"잘게 썬 레지아노 치즈"},
	}}
}

type ChicagoStyleCheesePizza struct {
    BasePizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{BasePizza: BasePizza{
		Name:     "시카고 스카일 딥디쉬 치즈 피자",
		Dough:    "아주 두꺼운 크러스트 도우",
		Sauce:    "플럼 토마토 소스",
		Toppings:[]string{"아주 많은 모짜렐라 치즈"},
	}}
}

type PizzaCreator interface {
    CreatePizza(pizzaType string) Pizza
}

type PizzaStore struct {
    PizzaCreator
}

func (ps PizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := ps.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

type NYPizzaStore struct {
}

func (N NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	switch pizzaType {
	case "cheese":
		return NewNYStyleCheesePizza()
	}

	return nil
}

type ChicagoPizzaStore struct {
}

func (c ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	switch pizzaType {
	case "cheese":
		return NewChicagoStyleCheesePizza()
	}

	return nil
}

func main() {
	n := PizzaStore{&NYPizzaStore{}}
	c := PizzaStore{&ChicagoPizzaStore{}}

	pizza := n.OrderPizza("cheese")
	fmt.Printf("상일이가 주문한 피자 %s\n", pizza.GetName())

	pizza = c.OrderPizza("cheese")
	fmt.Printf("상혁이가 주문한 피자 %s\n", pizza.GetName())
}
