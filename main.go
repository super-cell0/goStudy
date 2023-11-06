package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

func main() {

}

// 1106数字类型
func numberType2() {
	//十进制
	var t = 10
	fmt.Printf("%d\n", t) //
	fmt.Printf("%b\n", t) //%b二进制

	//八进制
	var e = 077
	fmt.Printf("%o\n", e)

	//十六进制
	var c = 0xff
	fmt.Printf("%X\n", c)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数类型
	var d complex64
	d = 1 + 2i
	var d2 complex128
	d2 = 2 + 3i
	fmt.Println(d)
	fmt.Println(d2)
}

func numberType() {
	var i8 int8
	var ui8 uint8
	var f32 float32

	fmt.Println(i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	println(ui8, unsafe.Sizeof(ui8), math.MaxUint8)
	println(f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
}

func stringAdd2() {

	var buffer bytes.Buffer
	buffer.WriteString("Tom")
	buffer.WriteString(" what's up")
	fmt.Println(buffer.String())
}

func stringAdd() {
	var name = "beidixiaoxiong"
	var intDemo = "8989"
	var addString = strings.Join([]string{name, intDemo}, ",")
	println(addString)
}

func demoFor() {
	var count = 10
	for i := 0; i < count; i++ {
		println(i)
	}
}

func demoBool() {
	var age = 18
	if age >= 18 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func createArray() {
	var names = []int{1, 2, 3}
	for i, name := range names {
		println(i, name)
	}
}

func demoIota() {
	//iota是一个预定义的标识符，用于创建枚举常量序列。
	//当定义一个枚举类型的常量时，iota会被重置为0，然后每个连续的常量声明都会使iota递增1
	const (
		a1 = iota
		a2 = iota
		a3 = iota
		_
		a4 = iota
	)

	fmt.Println(a1, a2, a3, a4)
}

func constStatement() {
	const PI float64 = 3.1415

	fmt.Println(PI)
}

func cardGame() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Two of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "diamonds"
}

func intCard() int {
	return 45
}

// 变量批量声明
func varStatement() {
	var (
		name   string
		age    int
		gender string
	)

	name = "chen"
	age = 24
	gender = "男"

	fmt.Println(name, age, gender)

	//还有类型推断
	//短变量声明 := 只能放在函数内部
}

func getNameAndAge() (string, int) {
	return "chen", 24
}
