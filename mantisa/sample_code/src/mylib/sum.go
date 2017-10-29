package mylib

func Sum(x int, y int) int {
    return sum(x,y)
}

func sum(x int, y int) int { // accessible within package
    return x+y
}

