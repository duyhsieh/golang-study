package gamecore

import ("fmt"
        "math/rand"
)

const DIM int = 4
const CELLS int = DIM*DIM
type Game struct {
    Grid [DIM][DIM]int
    Merged [DIM][DIM]bool
    GenIndex int 
}

// public methods: starts with upper case. other package can call it.

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

func(g* Game) IsComplete() bool {
    for i:=0; i<DIM; i++ {
        for j:=0; j<DIM; j++ {
            if g.Grid[i][j] == 2048 {
                return true
            }
        }
    }
    return false
}

func(g* Game) ClearMergeState() {
    for i:=0; i<DIM; i++ {
        for j:=0; j<DIM; j++ {
            g.Merged[i][j] = false
        }
    }
}

func(g* Game) PrintGrid() {
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

func(g* Game) GenNumber() {
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

func(g* Game) IsDead() bool{
    for i:=0; i < DIM; i++ {
        for j:=0; j < DIM; j++ {
            if g.Grid[i][j] == 0 {
                return false
            }
        }
    }
    return true
}

func(g* Game) MoveLeft() bool{
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

func(g* Game) MoveRight() bool{
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
func(g* Game) MoveUp() bool{
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
func(g* Game) MoveDown() bool{
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

// private methods: do not open for calling outside this package

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

