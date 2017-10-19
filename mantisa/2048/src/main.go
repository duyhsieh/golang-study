package main
import ("fmt"
        "bufio"
        "os"
        "math/rand"
)

const DIM int = 4
const CELLS int = DIM*DIM

type Game struct {
    Grid [4][4]int
    Merged [4][4]bool
    GenIndex int 

}

func main() {
    fmt.Println("2048 Console Verson...")
    var g Game

    fin := false
    // init wo random 2
    g.genNumber()
    g.genNumber()

    for {
        g.printGrid()
        if g.isDead() {
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
            g.clearMergeState()
            valid_act = g.moveUp()
        case 'd':
            g.clearMergeState()
            valid_act = g.moveDown()
        case 'l':
            g.clearMergeState()
            valid_act = g.moveLeft()
        case 'r':
            g.clearMergeState()
            valid_act = g.moveRight()
        default:
            fmt.Println("error input")
        }
        if valid_act == false {
            fmt.Println("Action takes no effort!")
        } else {
            if g.isComplete() == true {
                fin = true
                fmt.Println("You Win!")
            } else {
                g.genNumber()
            }
        }
    }
}


func(g* Game) isComplete() bool {
    for i:=0; i<DIM; i++ {
        for j:=0; j<DIM; j++ {
            if g.Grid[i][j] == 2048 {
                return true
            }
        }
    }
    return false
}

func(g* Game) clearMergeState() {
    for i:=0; i<DIM; i++ {
        for j:=0; j<DIM; j++ {
            g.Merged[i][j] = false
        }
    }
}

func(g* Game) printGrid() {
    for i:=0; i < DIM; i++ {
        for j:=0; j < DIM; j++ {
            if (i*DIM + j) == g.GenIndex {
                fmt.Printf("%4d*", g.Grid[i][j])
            } else {
                fmt.Printf("%4d ", g.Grid[i][j])
            }
        }
        fmt.Printf("\n")
    }
}

func(g* Game) genNumber() {
    n := rand.Intn(DIM*DIM)
    for i:=0; i < CELLS; i++ {
        idx := (i + n) % CELLS
        column := idx/DIM
        row := idx % DIM
        if g.Grid[column][row] == 0 {
            g.Grid[column][row] = 2
            g.GenIndex = DIM*column + row
            fmt.Printf("New Random 2 Idx:[%d,%d]\n", column, row)
            return
        }
    }
}

func(g* Game) isDead() bool{
    for i:=0; i < DIM; i++ {
        for j:=0; j < DIM; j++ {
            if g.Grid[i][j] == 0 {
                return false
            }
        }
    }
    return true
}

func(g* Game) moveLeft() bool{
    dirty := false
    for row:=0; row < DIM; row++ {
        for col:=0; col < DIM; col++ {
            r := g.rowMoveLeft(row, col, false)
            if dirty == false && r == true {
                dirty = true
            }
        }
    }
    return dirty
}

func(g* Game) moveRight() bool{
    dirty := false
    for row:=0; row < DIM; row++ {
        for col:=DIM-1; col >= 0; col-- {
            r := g.rowMoveRight(row, col, false)
            if dirty == false && r == true {
                dirty = true
            }
        }
    }
    return dirty
}
func(g* Game) moveUp() bool{
    dirty := false
    for col:=0; col < DIM; col++ {
        for row:=0; row < DIM; row++ {
            r := g.colMoveUp(row, col, false)
            if dirty == false && r == true {
                dirty = true
            }
        }
    }
    return dirty
}
func(g* Game) moveDown() bool{
    dirty := false
    for col:=0; col < DIM; col++ {
        for row:=DIM-1; row >= 0; row-- {
            r := g.colMoveDown(row, col, false)
            if dirty == false && r == true {
                dirty = true
            }
        }
    }
    return dirty
}

func(g* Game) rowMoveLeft(row int, col int, dirty bool) bool {
    if col == 0 || g.Grid[row][col] == 0 {
        return dirty
    } else {
        if g.Grid[row][col-1] == 0 {
            g.Grid[row][col-1] = g.Grid[row][col]
            g.Grid[row][col] = 0
            //fmt.Printf("moving %v,%v to zero left:row=%v\n", row, col, g.Grid[row])
            return g.rowMoveLeft(row, col-1, true)
        }else if g.Grid[row][col-1] == g.Grid[row][col] {
            if g.Merged[row][col-1] == false {
                g.Merged[row][col-1] = true
                g.Grid[row][col-1]*=2
                g.Grid[row][col]=0
                //fmt.Printf("merging %v,%v to left:row=%v\n",row, col, g.Grid[row])
                return true
            }
        }
        return dirty
    }
}

func(g* Game) rowMoveRight(row int, col int, dirty bool) bool {
    if col == DIM-1 || g.Grid[row][col] == 0 {
        return dirty
    } else {
        if g.Grid[row][col+1] == 0 {
            g.Grid[row][col+1] = g.Grid[row][col]
            g.Grid[row][col] = 0
            //fmt.Printf("moving %v,%v to zero right:row=%v\n", row, col, g.Grid[row])
            return g.rowMoveRight(row, col+1, true)
        }else if g.Grid[row][col+1] == g.Grid[row][col] {
            if g.Merged[row][col+1] == false {
                g.Merged[row][col+1]=true
                g.Grid[row][col+1]*=2
                g.Grid[row][col]=0
                //fmt.Printf("merging %v,%v to right:row=%v\n",row, col, g.Grid[row])
                return true
            }
        }
        return dirty
    }
}

func(g* Game) colMoveUp(row int, col int, dirty bool) bool {
    if row == 0 || g.Grid[row][col] == 0 {
        return dirty
    } else {
        if g.Grid[row-1][col] == 0 {
            g.Grid[row-1][col] = g.Grid[row][col]
            g.Grid[row][col] = 0
            //fmt.Printf("moving %v,%v to zero up:row=%v\n", row, col, g.Grid[row])
            return g.colMoveUp(row-1, col, true)
        }else if g.Grid[row-1][col] == g.Grid[row][col] {
            if g.Merged[row-1][col] == false {
                g.Merged[row-1][col] = true
                g.Grid[row-1][col]*=2
                g.Grid[row][col]=0
                //fmt.Printf("merging %v,%v to up:row=%v\n",row, col, g.Grid[row])
                return true
            }
        }
        return dirty
    }
}

func(g* Game) colMoveDown(row int, col int, dirty bool) bool {
    if row == DIM-1 || g.Grid[row][col] == 0 {
        return dirty
    } else {
        if g.Grid[row+1][col] == 0 {
            g.Grid[row+1][col] = g.Grid[row][col]
            g.Grid[row][col] = 0
            //fmt.Printf("moving %v,%v to zero down:row=%v\n", row, col, g.Grid[row])
            return g.colMoveDown(row+1, col, true)
        }else if g.Grid[row+1][col] == g.Grid[row][col] {
            if g.Merged[row+1][col] == false {
                g.Merged[row+1][col] = true 
                g.Grid[row+1][col]*=2
                g.Grid[row][col]=0
                //fmt.Printf("merging %v,%v to down:row=%v\n",row, col, g.Grid[row])
                return true
        }
        }
        return dirty
    }
}

