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

type CriticalAttack struct {
    Behavior  // used to initialize initial implementation from whatever passed in , and then override
}

func(c* CriticalAttack) Attack() string {
    //return "Critical !"
    return "Critical Attack! " + c.Behavior.Attack() // c.Behavior.Attack() can make magic that it behaves as whatever you passed in...
}
func SetCriticalAttack(src Behavior) Behavior {
    return &CriticalAttack{src}
}

// only override Attack
func(p *Assassin) Attack() string {
    return "Assassin dagger"
}

func(p *Ninja) Move() int {
    return 4
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
    var x = SetCriticalAttack(&Ninja{} ) // use NInja as default implementation, but override Attack (inspiration from GO reverse)

    BadTakeAction1(&n)
    BadTakeAction2(&n,&m)
    GoodTakeAction(&n, &m, &a)
    fmt.Println("--- demo changing different interface implementations on the fly ----")
    GoodTakeAction(&Ninja{},  &Mage{}, x )

}
