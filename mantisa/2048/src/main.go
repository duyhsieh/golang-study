package main
import ("fmt"
        "bufio"
        "os"
        "gamemodule" // this is directory name(path), not package name; package gamecore in this dir are imported
)
/*
const DIM int = 4
const CELLS int = DIM*DIM
type Game struct {
    Grid [DIM][DIM]int
    Merged [DIM][DIM]bool
    GenIndex int 
}

func NewGame() *Game {
    // ========= if variable is referenced after functon finished, it will be allocated on heap, not stack. ============
    g := Game { 
        GenIndex:0,
        //Grid: [DIM][DIM]int{ {0,0,0,2}, {0,0,0,0}, {0,0,2,2}, {1024,1024,2,2}}, 
        // note commna is needed at last 
        Merged: [DIM][DIM]bool{ 
            {false,false,false,false},
            {false,false,false,false},
            {false,false,false,false},
            {false,false,false,false}},
    }
    return &g
}
*/
func main() {
    var g *gamecore.Game = gamecore.NewGame()
    fin := false
    // init wo random 2
    g.GenNumber()
    g.GenNumber()
    g.PrintGrid()

    for {
        if g.IsDead() {
            fmt.Println("GameOver!")
            fin = true
        }

        if fin == true {
            break
        }
        fmt.Println("Move direction: u, d, l, r")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        c := input[0]
        
        var valid_act = false
        switch c {
        case 'u':
            g.ClearMergeState()
            valid_act = g.MoveUp()
        case 'd':
            g.ClearMergeState()
            valid_act = g.MoveDown()
        case 'l':
            g.ClearMergeState()
            valid_act = g.MoveLeft()
        case 'r':
            g.ClearMergeState()
            valid_act = g.MoveRight()
        default:
            fmt.Println("error input")
        }
        if valid_act == false {
            fmt.Println("Action takes no effort!")
        } else {
            if g.IsComplete() == true {
                fin = true
                fmt.Println("You Win!")
            } else {
                g.GenNumber()
            }
        }
        g.PrintGrid()
    }
}


