package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name string
}

func (h Human) Eat() {
	fmt.Println(h.Name, "OmnonNom!")
}

func (h Human) Sleep() {
	fmt.Println(h.Name, "Zzz zZZzz zZZZ...")
}

func (h Human) Repeat() {
	fmt.Println(h.Name, "Says: And this is a new day!")
}

type Action struct {
	humanIsAlive bool
	Human
}

func (a Action) Validate() error {
	if !a.humanIsAlive {
		return fmt.Errorf("%s is not alive", a.Name)
	}
	if strings.TrimSpace(a.Name) == "" {
		return fmt.Errorf("field name is not defined or contains only spaces")
	}
	return nil
}

func main() {
	a := Action{
		Human:        Human{Name: "Ugly"},
		humanIsAlive: true,
	}

	if err := a.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	a.Eat()
	a.Sleep()
	a.Repeat()
}
