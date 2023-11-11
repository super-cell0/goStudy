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
	sliceCopy()
}

// sliceCopy
func sliceCopy() {
	var sl1 = []int{23, 45, 67, 44}
	//直接赋值的话sl1也会跟着改变 因为赋值给sl2他们共享一个内存地址
	//var sl2 = sl1
	var sl2 = make([]int, 4)
	copy(sl2, sl1)
	sl2[0] = 100
	fmt.Printf("%v\n", sl2)
	fmt.Printf("%v\n", sl1)
}

// 切片删除
func sliceDelete() {
	var sl = []int{1, 2, 3, 4, 5}
	//删除下标为2的元素
	sl = append(sl[:2], sl[3:]...)
	// a = append(a[:index], al[index+1:]...)
	fmt.Printf("%v\n", sl)
}

func sliceUpdate() {
	var sl = []int{23, 78, 99}
	sl[1] = 2020
	sl[2] = 2023
	fmt.Printf("%v\n", sl)
}

func sliceQuery() {
	var sl = []int{1, 3, 5, 6, 32}
	var q = 5
	for i, i2 := range sl {
		if i2 == q {
			fmt.Printf("下标为: %v\n", i)
		}
	}
}

func sliceAppend() {
	var sl []int
	sl = append(sl, 100)
	sl = append(sl, 23)
	sl = append(sl, 43)

	fmt.Printf("%v\n", sl)
}

// 数组
func arrayDemo6() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	var a2 = a1[:3]
	var a3 = a1[3:]
	var a4 = a1[2:5]
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", a4)
}

// slice
func sliceFor() {
	var sl = []int{2, 3, 5, 8}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%v: %v\n", i, sl[i])
	}

	fmt.Printf("for rang\n")

	for i, i2 := range sl {
		fmt.Printf("%v: %v\n", i, i2)
	}

}

func sliceDemo2() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	var s1To3 = s1[0:3]
	fmt.Printf("%v\n", s1To3)
	var s1To30 = s1[3:]
	fmt.Printf("%v\n", s1To30)
	var s1To5 = s1[2:5]
	fmt.Printf("%v\n", s1To5)
	var s1ToAll = s1[:]
	fmt.Printf("%v\n", s1ToAll)
}

func sliceDemo() {
	var s1 []int
	var s2 []string
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)

	var s3 = make([]int, 3)
	fmt.Printf("%v\n", s3)
	//切片的长度和容量
	fmt.Printf("%v\n", len(s3))
	fmt.Printf("%v\n", cap(s3))
}

// go 数组
func arrayDemo3() {
	var a1 = [3]int{1, 2, 3}
	fmt.Printf("%v\n", len(a1))
}

func arrayDemo4() {
	var array = [3]int{1, 2, 3}
	//for i := 0; i < len(array); i++ {
	//	fmt.Printf("array[%v]: %v\n", i, array[i])
	//}
	for r, a := range array {
		fmt.Printf("%v: %v\n", r, a)
	}
}

func arrayDemo2() {
	var array = [2]int{45, 22}
	fmt.Printf("%v\n", array)

	var array2 = [2]string{"hello", "world"}
	fmt.Printf("%v\n", array2)

	var array3 = [...]int{34, 3, 23}
	fmt.Printf("%v\n", array3)

	//指定索引值的方式来初始化
	var array4 = [...]int{1: 4, 4: 89}
	fmt.Printf("%v\n", array4)
}

func arrayDemo() {
	var array1 [4]int
	fmt.Printf("%v\n", array1)

	var array2 [2]string
	var array3 [2]bool
	fmt.Printf("%v\n", array2)
	fmt.Printf("%v\n", array3)
}

// continue
func forDemo6() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				goto END
			}
			fmt.Printf("%v %v\n", i, j)
		}
	}
END:
	fmt.Println("END")
}

func continueDemo1() {
	for a := 1; a <= 10; a++ {
		if a%2 == 0 {
			fmt.Printf("a: %v\n", a)
		} else {
			continue
		}
	}
}

func test1() {
	var i = 1
	if i >= 2 {
		fmt.Println("2")
	} else {
		goto END
	}
END:
	fmt.Println("END...")
}

func continueDemo2() {
	for i := 0; i <= 10; i++ {
	MY_LABEL:
		for j := 0; j < +10; j++ {
			if i == 2 && j == 2 {
				continue MY_LABEL
			}
			fmt.Printf("%v %v\n", i, j)
		}
	}
}

// for循环
func forDemo() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func forDemo5() {
MyLabel:
	for i := 0; i <= 5; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break MyLabel
		}
	}

	fmt.Println("END")
}

func switchDemo3() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break
		}
	}
}

func forDemo1() {
	//永真循环
	for {
		fmt.Println("hello")
	}
}

func forDemo2() {
	var array = [...]int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func forDemo3() {
	// []什么都不加表示切片
	var sliceDemo = []int{1, 2, 3, 4, 5}
	for _, value := range sliceDemo {
		fmt.Printf("value: %v\n", value)
	}
}

func forDemo4() {
	var m = make(map[string]string, 0)
	m["name"] = "chen"
	m["age"] = "20"
	m["email"] = "zdas@qq.com"

	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}

// switch
func switchDemo2() {
	var a = 100
	switch a {
	case 100:
		fmt.Println("满分")
		//可以可以执行满足条件的下一个case
		fallthrough
	case 80:
		fmt.Println("no")
	}
}

func switchDemo1() {
	var day = 1
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入")
	}
}

func switchDemo() {
	var score = "a"
	switch score {
	case "a":
		fmt.Println("优秀")
	case "b":
		fmt.Println("良好")
	case "c":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

// 算数运算符s
func operatorDemo6() {
	var a, b, c = 1, 2, 3
	if a > b {
		if a > c {
			fmt.Println("a")
		}
	} else {
		if b > c {
			fmt.Println("b")
		} else {
			fmt.Println("c")
		}
	}
}

func operatorDemo5() {
	//Monday Tuesday Wednesday Thursday Friday Saturday Sunday
	var c string
	fmt.Println("请输入一个字符")
	fmt.Scan(&c)

	if c == "S" || c == "s" {
		fmt.Println("请输入第二个字符")
		fmt.Scan(&c)
		if c == "a" || c == "A" {
			fmt.Println("Saturday")
		} else if c == "u" || c == "U" {
			fmt.Println("Sunday")
		} else {
			fmt.Println("输入的是其他单词")
		}
	} else if c == "T" || c == "t" {
		fmt.Println("请输入第二个字符")
		fmt.Scan(&c)
		if c == "h" || c == "H" {
			fmt.Println("Thursday")
		} else if c == "u" || c == "U" {
			fmt.Println("Tuesday")
		} else {
			fmt.Println("输入的是其他单词")
		}
	} else if c == "M" || c == "m" {
		fmt.Println("Monday")
	} else if c == "W" || c == "w" {
		fmt.Println("Wednesday")
	} else if c == "F" || c == "f" {
		fmt.Println("Friday")
	} else {
		fmt.Println("输入的是其他单词")
	}

}

func operatorDemo4() {
	var score = 90
	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score >= 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}

}

func operatorDemo3() {
	var num int
	fmt.Println("请输入一个数字")
	fmt.Scan(&num)

	if num%2 == 0 {
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
	var name = ""
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

// 1110fotmat
func testFormat() {
	var website = Website{Nmae: "chencan"}
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
