package lib

func Sum(x int, y int) int {
	a:=0
	return sum(x,y)
}

func sum(x int, y int) int { // accessible within package
    return x+y
}

