package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**************************************************************************************
 * Strategy Pattern
 ****************************************************** sanghyuk, 2022/01/06 20:52:18 *
 * 알고리즘군을 정의하고 각각을 캡슐화 하여 바꿔 쓸 수 있게 만든다. 이 패턴을 이용하면 알고리즘을 활용하는 클라이언트와 독립적으로 알고리즘을 변경할 수 있습니다.
 * 객체지향의 디자인 원칙 (참고)
	* 변경이 일어나는 부분을 캡슐화 한다.
	* is...a... 관계 보단 has...a... 관계를 활용 한다.
	* 구현이 아닌 인터페이스에 맞춰서 프로그래밍 한다.
/**************************************************************************************/

type WeaponBehavior interface {
    UseWeapon()
}

type SwordBehavior struct {}

func (s SwordBehavior) UseWeapon() {
	fmt.Println("Use Sword...")
}

type KnifeBehavior struct {}

func (k KnifeBehavior) UseWeapon() {
	fmt.Println("Use Knife...")
}

type BowAndArrowBehavior struct {}

func (b BowAndArrowBehavior) UseWeapon() {
	fmt.Println("Use Bow And Arrow...")
}

type AxeBehavior struct {}

func (a AxeBehavior) UseWeapon() {
	fmt.Println("Use Axe...")
}

type Character struct {
	name string
    weapon WeaponBehavior
}

func (c *Character) SetWeapon(w WeaponBehavior) {
	c.weapon = w
}

func (c *Character) Fight() {
	fmt.Printf("I'm a %s ", c.name)
	c.weapon.UseWeapon()
}

type King struct {
    Character
}

func NewKing() *King {
	k := &King{}
	k.name = "King"
	return k
}

type Knight struct {
    Character
}

func NewKnight() *Knight {
	kn := &Knight{}
	kn.name = "Knight"
	return kn
}

type Queen struct {
    Character
}

func NewQueen() *Queen {
	q := &Queen{}
	q.name = "Queen"
	return q
}

type Troll struct {
    Character
}

func NewTroll() *Troll {
	t := &Troll{}
	t.name = "troll"
	return t
}

func GetRandom(w WeaponBehavior) ( c Character ) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)

	switch  {
	case i > 80:
		c = NewKing().Character
	case i > 60:
		c =  NewKnight().Character
	case i > 40 :
		c =  NewQueen().Character
	case i > 20:
		c =  NewTroll().Character
	default:
		c =  NewTroll().Character
	}

	c.SetWeapon(w)

	return c
}

func main() {
	king := NewKing()
	king.SetWeapon(SwordBehavior{})
	king.Fight()

	knight := NewKnight()
	knight.SetWeapon(KnifeBehavior{})
	knight.Fight()

	queen := NewQueen()
	queen.SetWeapon(BowAndArrowBehavior{})
	queen.Fight()

	troll := NewTroll()
	troll.SetWeapon(AxeBehavior{})
	troll.Fight()

	r := GetRandom(BowAndArrowBehavior{})
	r.Fight()
}
