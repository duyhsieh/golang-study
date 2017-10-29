package main
import ("fmt")

type Character struct {
    Name string
    //Action Behavior
}

type Ninja struct {
    Character
}

func(c *Character) Talk() {
    fmt.Println("Character Talk: My name is ", c.Name)
}

func(c *Ninja) Talk() {
    fmt.Println("Ninja Talk: My name is ", c.Name)
}

func main() {
    n := Ninja{Character{Name:"Orochi"}}
    n.Character.Talk()
    n.Talk()
}
