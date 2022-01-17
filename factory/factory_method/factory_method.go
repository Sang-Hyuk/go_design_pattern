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
		Dough:    "씬 크러스트",
		Sauce:    "마리나라 소스",
		Toppings:[]string{"잘게 썬 레지아노 치즈", "마늘"},
	}}
}

type NYStylePeperoniPizza struct {
	BasePizza
}

func NewNYStylePeperoniPizza() *NYStylePeperoniPizza {
	return &NYStylePeperoniPizza{BasePizza: BasePizza{
		Name:     "뉴욕페퍼로니",
		Dough:    "중간 두께의 일반",
		Sauce:    "마리라나 소스",
		Toppings:[]string{"페퍼로니", "레지아노 치즈", "버섯", "양파", "고추"},
	}}
}

type NYStyleClamPizza struct {
	BasePizza
}

func NewNYStyleClamPizza() *NYStyleClamPizza {
	return &NYStyleClamPizza{BasePizza: BasePizza{
		Name:     "뉴욕조개피자",
		Dough:    "씬",
		Sauce:    "마리나라 소스",
		Toppings:[]string{"조개", "올리브"},
	}}
}

type NYStyleVeggiePizza struct {
	BasePizza
}

func NewNYStyleVeggiePizza() *NYStyleVeggiePizza {
	return &NYStyleVeggiePizza{BasePizza: BasePizza{
		Name:     "뉴욕야채피자",
		Dough:    "씬 크러스트",
		Sauce:    "토마토 소스",
		Toppings:[]string{"비건 치즈", "바질"},
	}}
}

type ChicagoStyleCheesePizza struct {
    BasePizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{BasePizza: BasePizza{
		Name:     "시카고 스카일 딥디쉬 치즈 피자",
		Dough:    "아주 두꺼운 크러스트",
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
	case "peperoni":
		return NewNYStylePeperoniPizza()
	case "clam":
		return NewNYStyleClamPizza()
	case "veggie":
		return NewNYStyleVeggiePizza()
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
	pizza = n.OrderPizza("peperoni")
	fmt.Printf("세찬이가 주문한 피자 %s\n", pizza.GetName())
	pizza = n.OrderPizza("clam")
	fmt.Printf("혁수가 주문한 피자 %s\n", pizza.GetName())
	pizza = n.OrderPizza("veggie")
	fmt.Printf("재석이가 주문한 피자 %s\n", pizza.GetName())

	pizza = c.OrderPizza("cheese")
	fmt.Printf("상혁이가 주문한 피자 %s\n", pizza.GetName())
}
