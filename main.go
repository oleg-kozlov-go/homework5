package main

import (
	"fmt"
	"math"

)

//1.Создать структуру для круга, квадрата, параллелограмма и прямоугольника.
//Для каждой структуры определить методы рассчета периметра и площади.

type circle struct {
	r int
}
type square struct {
	a int
}
type parallelogram struct {
	a int
	b int
}
type rectangle struct {
	a int
	b int
	h int
}

func (c circle) per() int {
	f := float64(c.r)
	p :=2 * f * math.Pi
	fmt.Println("длинна окружности", p, "круг", c)
	return int(p)
}
func (c circle) area() int {
	f := float64(c.r)
	p :=f * f * math.Pi
	fmt.Println("площадь окружности", p, "круг", c)
	return int(f * f * math.Pi)
}
func (c square) per() int {
	return c.a * 4
}
func (c square) area() int {
	return c.a * c.a
}
func (c parallelogram) per() int {
	return (c.a + c.b) * 2
}
func (c parallelogram) area() int {
	return c.a * c.b
}
func (c rectangle) per() int {
	return (c.a + c.b) * 2
}
func (c rectangle) area() int {
	return c.a * c.h
}

//2.Создать структуру коллекция фигур. Определить для нее методы - количество фигур, добавить фигуру,
//сумма площади, сумма периметров содержащихся фигур. Также добавить метод история вызовов,
//который выводит историю всех вызванные ранее методов этой структуры вместе с их аргументами.
//Для доступа к фигурам использовать интерфейс.

type Fig interface {
	per() int
	area() int
}
type FigCol struct {
	Col     []Fig
	perArg  []perArg
	areaArg []areaArg
}
type perArg struct {
	per     int
}
type areaArg struct {
	area     int
}

func (c *FigCol) addFigCol(a Fig) {
	c.Col = append(c.Col, a)
	return
}
func (c *FigCol) quantityFig() {
	fmt.Println(len(c.Col))
}
func (c *FigCol) sumAreaFig() {
	S := 0
	for _, i := range c.Col {
		S += i.area()
	}
	//fmt.Println(S)
	St := areaArg{ S}
	c.areaArg = append(c.areaArg, St)
	return
}
func (c *FigCol) sumPerFig() {
	P := 0
	for _, i := range c.Col {
		P += i.per()
	}

	d := perArg{ P}
	c.perArg = append(c.perArg, d)
}
func (c *FigCol) outputAllMethods() {
	for _, i := range c.perArg {
		fmt.Println( "общий периметр фигур", i.per)
	}
	for _, i := range c.areaArg {
		fmt.Println( "общая площадь фигур", i.area)
	}
}

//3.Создать структуру цилиндр, которая будет включать круг в качестве основания и высоту.
//Методы те же что и у круга, плюс объем. Для доступу к кругу использовать интерфейс.

type cylinder struct {
	Fig
	h int
}

func (c cylinder) areaCylinder() {
	fmt.Println(c.area()*2+c.per()*c.h, "площадь цилиндра")
}
func (c cylinder) volumeCylinder() {
	fmt.Println(c.area()*c.h, "объем цилиндра")
}

//4.У всех структур должна быть функция конструктор.

func newCircle(r int) circle {
	return circle{r}
}
func newSquare(a int) square {
	fmt.Println(a)
	return square{a}
}
func newParallelogram(a int, b int) parallelogram {

	return parallelogram{a, b}
}
func newRectangle(a int, b int, h int) rectangle {
	return rectangle{a, b, h}
}
func newFigCol() FigCol {
	b := make([]Fig, 0, 0)
	d := make([]perArg, 0, 0)
	f := make([]areaArg, 0, 0)
	return FigCol{b, d, f}
}
func newCylinder(a circle) cylinder {
	return cylinder{a, a.r}
}
func main() {
G:=newCircle(5)
g:=newCylinder(G)
	g.areaCylinder()
    g.volumeCylinder()
	var b int
	c := newFigCol()

	for i := 1; i < 5; i++ {
		c.addFigCol(newCircle(i))
		c.addFigCol(newParallelogram(i, i))
		c.addFigCol(newRectangle(i, i, i))
		c.addFigCol(newSquare(i))
		b=i*4
	}
fmt.Println("колличество фигур",b)

	c.sumAreaFig()
	c.sumPerFig()
	c.outputAllMethods()


}

