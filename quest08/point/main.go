package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	xS := ""
	yS := ""

	x1 := points.x
	y1 := points.y

	for x1 > 0 {
		xS += string(rune(x1%10) + '0')
		x1 /= 10
	}

	for y1 > 0 {
		yS += string(rune(y1%10) + '0')
		y1 /= 10
	}
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := len(xS) - 1; i >= 0; i-- {
		z01.PrintRune(rune(xS[i]))
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := len(yS) - 1; i >= 0; i-- {
		z01.PrintRune(rune(yS[i]))
	}
	z01.PrintRune('\n')
}
