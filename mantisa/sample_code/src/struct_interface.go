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

type Assassin struct {
    Behavior
    Character
}

// Actually you can have multiple anonymous fields. It's just implicit composition. 
// So if you embed two interfaces having same signature, compilie will fail
// Since only structs can be receiver to function (causing a new implementation), you must declare method decoration with struct
// So struct is like an object with methods. Interface is a reference to some object having those APIs.
type CriticalAttack struct { 
    Behavior  // used to initialize initial implementation from whatever passed in , and then override
}

// this can actually implement decoration pattern
func(c* CriticalAttack) Attack() string {
    //return "Critical !"
    return "Critical Attack! " + c.Behavior.Attack() // c.Behavior.Attack() can make magic that it behaves as whatever you passed in...
}

// just like decorate it with Critical Attack
func MakeCriticalAttack(src Behavior) Behavior {
    return &CriticalAttack{src}  // return a CriticalAttack object inheriting implementations from src, while overriding some of the them
}

// only override Attack
func(p *Assassin) Attack() string {
    return  "Assassin dagger"
}

func(p *Ninja) Move() int {
    return 9
}

func(p *Ninja) Attack() string {
    return "Ninja Shuriken"
}

func(p *Mage) Move() int {
    return 2
}
func(p *Mage) Attack() string {
    return "Mage Fireball"
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
    a := Assassin{&Ninja{Character{Name:"Not_Important_Only_To_Take_Implementation"}}, Character{Name:"Taki"} }
    var x = MakeCriticalAttack(&Ninja{} ) // use NInja as default implementation, but override Attack (inspiration from GO reverse)

    BadTakeAction1(&n)
    BadTakeAction2(&n,&m)
    GoodTakeAction(&n, &m, &a)
    fmt.Println("--- demo changing different interface implementations on the fly ----")
    GoodTakeAction(&Ninja{},  &Mage{}, x )

}
