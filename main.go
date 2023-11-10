package main

import (
	"bytes"
	"fmt"
	"godaily/user"
	"math"
	"strings"
	"unsafe"
)

type Website struct {
	Nmae string
}

func main() {
	operatorDemo3()
}

//算数运算符
func operatorDemo3() {
	var num int
	fmt.Println("请输入一个数字")
	fmt.Scan(&num)

	if num %2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

func userDemo() {
	var user = user.Hello()
	fmt.Printf("user: %v\n", user)
}

func operatorDemo2() {
	var name string = ""
	var age int
	var email string
	fmt.Println("please enter name、age、 email")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
}

func operatorDemo() {
	var a = 100
	a++
	fmt.Printf("a: %v\n", a)
	a--
	fmt.Printf("a: %v\n", a)
}

//1110fotmat
func testFormat()  {
	var website = Website{ Nmae: "chencan"}
	fmt.Printf("website: %v\n", website)
	fmt.Printf("website: %#v\n", website)
	
	var number = 10
	fmt.Printf("number: %T\n", number)
}

// 1108字符串函数
func stringFunc() {
	var str = "HELLO WORLD"
	var str2 = "chenqingsong"
	fmt.Printf("%v\n", len(str))
	fmt.Printf("%v\n", strings.Split(str, " "))
	fmt.Printf("%v\n", strings.Contains(str, "hello"))
	fmt.Printf("strings.ToLower(str): %v\n", strings.ToLower(str))
	fmt.Printf("strings.ToUpper(str2): %v\n", strings.ToUpper(str2))
	fmt.Printf("strings.HasPrefix(str, \"HELLO\"): %v\n", strings.HasPrefix(str, "HELLO"))
	fmt.Printf("strings.HasSuffix(str, \"hello\"): %v\n", strings.HasSuffix(str, "hello"))
	fmt.Printf("strings.Index(str, \"O\"): %v\n", strings.Index(str, "O"))
	fmt.Printf("strings.LastIndex(str, \"O\"): %v\n", strings.LastIndex(str, "O"))
}

func demoStr02() {
	var str = "hello"
	print(str + "\n")

	var str2 = "chenqingsong"
	fmt.Println(str2[4:])
}

func demoStr() {
	var s = "hello world"
	var s1 = "chen"
	s2 := "story"
	var s4 = `
	line1
	line2 
	lin3
	`
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s4)
}

func stringSplicing() {
	//字符串拼接
	var name = "chen"
	var age = "23"
	var message = fmt.Sprintf("%s, %s", name, age)
	fmt.Printf("%s\n", message)

	var message2 = strings.Join([]string{name, age}, " ")
	fmt.Printf("%s\n", message2)

	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString("world")
	fmt.Printf("%v\n", buffer.String())

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
		A1 = iota
		A2 = iota
		A3 = iota
		_
		A4 = iota
	)

	fmt.Println(A1, A2, A3, A4)
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
