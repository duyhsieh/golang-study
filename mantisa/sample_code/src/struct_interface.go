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


func BadTakeAction1(p *Ninja) {
    fmt.Println("Bad TakeAction1: Ninja: Move=", p.Move(), ", Attack=", p.Attack())
}

func BadTakeAction2(p *Ninja, p1 *Mage ) {
    fmt.Println("Bad TakeAction2: Ninja: Move=", p.Move(), ", Attack=", p.Attack())
    fmt.Println("Bad TakeAction2: Mage: Move=", p1.Move(), ", Attack=", p1.Attack())
}
// do not keep adding data types to new functions .................., so use interface
func GoodTakeAction(p ...Behavior) {
    for _, v:= range p {

        fmt.Println("GoodTakeAction:Move=", v.Move(), ", Attack=", v.Attack())
    }
}

func main() {
    n := Ninja{Character{Name:"Orochi"}}
    m := Mage{Character{Name:"Sakura"}}
    BadTakeAction1(&n)
    BadTakeAction2(&n,&m)
    //TakeAction(&n)  // note: & get address can be casted to pointers
    //TakeAction(&m) 
    GoodTakeAction(&n, &m)
}
