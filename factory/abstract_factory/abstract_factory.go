package main

import "fmt"

type Dough interface {
    Dough() string
}

type ThinCrustDough struct {

}

func (t ThinCrustDough) Dough() string {
	return "씬 크러스트"
}

type CrustDough struct {

}

func (c CrustDough) Dough() string {
	return "크러스트"
}

type Sauce interface {
    Sauce() string
}

type MarinaraSauce struct {

}

func (m MarinaraSauce) Sauce() string {
	return "마리나라"
}

type TomatoSauce struct {

}

func (t TomatoSauce) Sauce() string {
	return "토마토"
}

type Cheese interface {
    Cheese() string
}

type ReggianoCheese struct {

}

func (r ReggianoCheese) Cheese() string {
	return "레지아노"
}

type MozzarellaCheese struct {

}

func (m MozzarellaCheese) Cheese() string {
	return "모짜렐라"
}

type Veggie interface {
    Veggie() string
}

type Garlic struct {

}

func (g Garlic) Veggie() string {
	return "마늘"
}

type Onion struct {

}

func (o Onion) Veggie() string {
	return "양파"
}

type Mushroom struct {

}

func (m Mushroom) Veggie() string {
	return "버섯"
}

type RedPepper struct {

}

func (r RedPepper) Veggie() string {
	return "고추"
}

type Pepperoni interface {
    Pepperoni() string
}

type SlicedPepperoni struct {
    
}

func (s SlicedPepperoni) Pepperoni() string {
	return "얇게 자른 페퍼로니"
}

type Clam interface {
    Clam() string
}

type FreshClam struct {
    
}

func (f FreshClam) Clam() string {
	return "신선한 조개"
}

type FrozenClam struct {

}

func (f FrozenClam) Clam() string {
	return "냉동 조개"
}

type PizzaIngredientFactory interface {
    CreateDough() Dough
    CreateSauce() Sauce
    CreateCheese() Cheese
    CreateVeggies() []Veggie
    CreatePeperoni() Pepperoni
    CreateClam() Clam
}

type NYPizzaIngredientFactory struct {
    
}

func (N NYPizzaIngredientFactory) CreateDough() Dough {
	return ThinCrustDough{}
}

func (N NYPizzaIngredientFactory) CreateSauce() Sauce {
	return MarinaraSauce{}
}

func (N NYPizzaIngredientFactory) CreateCheese() Cheese {
	return ReggianoCheese{}
}

func (N NYPizzaIngredientFactory) CreateVeggies() []Veggie {
	return []Veggie{Garlic{}, Onion{}, Mushroom{}, RedPepper{}}
}

func (N NYPizzaIngredientFactory) CreatePeperoni() Pepperoni {
	return SlicedPepperoni{}
}

func (N NYPizzaIngredientFactory) CreateClam() Clam {
	return FreshClam{}
}

type ChicagoPizzaIngredientFactory struct {

}

func (c ChicagoPizzaIngredientFactory) CreateDough() Dough {
	return CrustDough{}
}

func (c ChicagoPizzaIngredientFactory) CreateSauce() Sauce {
	return TomatoSauce{}
}

func (c ChicagoPizzaIngredientFactory) CreateCheese() Cheese {
	return MozzarellaCheese{}
}

func (c ChicagoPizzaIngredientFactory) CreateVeggies() []Veggie {
	return []Veggie{Garlic{}, Onion{}, Mushroom{}}
}

func (c ChicagoPizzaIngredientFactory) CreatePeperoni() Pepperoni {
	return SlicedPepperoni{}
}

func (c ChicagoPizzaIngredientFactory) CreateClam() Clam {
	return FrozenClam{}
}

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type BasePizza struct {
	Name string
	Dough Dough
	Sauce Sauce
	Veggies []Veggie
	Cheese Cheese
	Pepperoni Pepperoni
	Clam Clam
}

func (b BasePizza) Prepare() {
	panic("재정의 하거라.")
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
	PizzaIngredientFactory
}

func NewNYStyleCheesePizza(pizzaIngredientFactory PizzaIngredientFactory) *NYStyleCheesePizza {
	return &NYStyleCheesePizza{PizzaIngredientFactory: pizzaIngredientFactory,BasePizza: BasePizza{Name: "뉴욕치즈피자"}}
}

func (s NYStyleCheesePizza) Prepare()  {
	fmt.Println("["+s.Name+"]" + " 준비중...")
	s.Dough = s.PizzaIngredientFactory.CreateDough()
	s.Sauce = s.PizzaIngredientFactory.CreateSauce()
	s.Cheese = s.PizzaIngredientFactory.CreateCheese()
	fmt.Printf("\t[도우] : %s\n\t[소스] : %s\n\t[치즈] : %s\n", s.Dough.Dough(), s.Sauce.Sauce(), s.Cheese.Cheese())
}

type NYStylePeperoniPizza struct {
	BasePizza
	PizzaIngredientFactory
}

func NewNYStylePeperoniPizza(pizzaIngredientFactory PizzaIngredientFactory) *NYStylePeperoniPizza {
	return &NYStylePeperoniPizza{PizzaIngredientFactory: pizzaIngredientFactory,BasePizza: BasePizza{Name: "뉴욕페퍼로니피자"}}
}

func (s *NYStylePeperoniPizza) Prepare() {
	fmt.Println("["+s.Name+"]" + " 준비중...")
	s.Dough = s.PizzaIngredientFactory.CreateDough()
	s.Sauce = s.PizzaIngredientFactory.CreateSauce()
	s.Cheese = s.PizzaIngredientFactory.CreateCheese()
	s.Pepperoni = s.PizzaIngredientFactory.CreatePeperoni()
	fmt.Printf("\t[도우] : %s\n\t[소스] : %s\n\t[치즈] : %s\n\t[페퍼로니] : %s\n", s.Dough.Dough(), s.Sauce.Sauce(), s.Cheese.Cheese(), s.Pepperoni.Pepperoni())
}

type NYStyleClamPizza struct {
	BasePizza
	PizzaIngredientFactory
}

func NewNYStyleClamPizza(pizzaIngredientFactory PizzaIngredientFactory) *NYStyleClamPizza {
	return &NYStyleClamPizza{PizzaIngredientFactory: pizzaIngredientFactory,BasePizza: BasePizza{Name: "뉴욕조개피자"}}
}

func (s *NYStyleClamPizza) Prepare() {
	fmt.Println("["+s.Name+"]" + " 준비중...")
	s.Dough = s.PizzaIngredientFactory.CreateDough()
	s.Sauce = s.PizzaIngredientFactory.CreateSauce()
	s.Cheese = s.PizzaIngredientFactory.CreateCheese()
	s.Clam = s.PizzaIngredientFactory.CreateClam()
	fmt.Printf("\t[도우] : %s\n\t[소스] : %s\n\t[치즈] : %s\n\t[조개] : %s\n", s.Dough.Dough(), s.Sauce.Sauce(), s.Cheese.Cheese(), s.Clam.Clam())
}

type NYStyleVeggiePizza struct {
	BasePizza
	PizzaIngredientFactory
}

func NewNYStyleVeggiePizza(pizzaIngredientFactory PizzaIngredientFactory) *NYStyleVeggiePizza {
	return &NYStyleVeggiePizza{PizzaIngredientFactory: pizzaIngredientFactory,BasePizza: BasePizza{Name: "뉴욕야채피자"}}
}

func (s *NYStyleVeggiePizza) Prepare() {
	fmt.Println("["+s.Name+"]" + " 준비중...")
	s.Dough = s.PizzaIngredientFactory.CreateDough()
	s.Sauce = s.PizzaIngredientFactory.CreateSauce()
	s.Cheese = s.PizzaIngredientFactory.CreateCheese()
	s.Veggies = s.PizzaIngredientFactory.CreateVeggies()
	fmt.Printf("\t[도우] : %s\n\t[소스] : %s\n\t[치즈] : %s\n\t[야채] : %+v\n", s.Dough.Dough(), s.Sauce.Sauce(), s.Cheese.Cheese(), s.Veggies)
}

type ChicagoStyleCheesePizza struct {
	BasePizza
	PizzaIngredientFactory
}

func NewChicagoStyleCheesePizza(pizzaIngredientFactory PizzaIngredientFactory) *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{PizzaIngredientFactory: pizzaIngredientFactory,BasePizza: BasePizza{Name: "시카고치즈피자"}}
}

func (s *ChicagoStyleCheesePizza) Prepare() {
	fmt.Println("["+s.Name+"]" + " 준비중...")
	s.Dough = s.PizzaIngredientFactory.CreateDough()
	s.Sauce = s.PizzaIngredientFactory.CreateSauce()
	s.Cheese = s.PizzaIngredientFactory.CreateCheese()
	fmt.Printf("\t[도우] : %s\n\t[소스] : %s\n\t[치즈] : %s\n", s.Dough.Dough(), s.Sauce.Sauce(), s.Cheese.Cheese())
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
	pizzaIngredientFactory := &NYPizzaIngredientFactory{}

	switch pizzaType {
	case "cheese":
		return NewNYStyleCheesePizza(pizzaIngredientFactory)
	case "peperoni":
		return NewNYStylePeperoniPizza(pizzaIngredientFactory)
	case "clam":
		return NewNYStyleClamPizza(pizzaIngredientFactory)
	case "veggie":
		return NewNYStyleVeggiePizza(pizzaIngredientFactory)
	}

	return nil
}

type ChicagoPizzaStore struct {
}

func (c ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	pizzaIngredientFactory := &ChicagoPizzaIngredientFactory{}

	switch pizzaType {
	case "cheese":
		return NewChicagoStyleCheesePizza(pizzaIngredientFactory)
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
