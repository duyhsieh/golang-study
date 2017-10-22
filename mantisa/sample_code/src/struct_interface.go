package main
import ("fmt")

type Behavior interface {
    Move() int
    Attack() string
}
type Character struct {
    Name string
    //Action Behavior
}

type Ninja struct {
    Character
}

type Mage struct {
    Character
}

func(p *Ninja) Move() int {
    return 4
}

func(p *Ninja) Attack() string {
    return "Shuriken"
}

func(p *Mage) Move() int {
    return 2
}
func(p *Mage) Attack() string {
    return "Fireball"
}

func TakeAction(p Behavior) {
    fmt.Println("Move=", p.Move(), ", Attack=", p.Attack())
}

func main() {
    n := Ninja{Character{Name:"Orochi"}}
    m := Mage{Character{Name:"Sakura"}}
    TakeAction(&n)  // note: & get address can be casted to pointers
    TakeAction(&m) 
}
