package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
