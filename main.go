package main

import (
	"bytes"
	"fmt"
	"godaily/user"
	"math"
	"strings"
	"unsafe"
)

func main() {
	testPerson05()
}

// Person05 构造方法
type Person05 struct {
	name string
	age  int
}

func newPerson05(name string, age int) (*Person05, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age cannot be 0")
	}
	return &Person05{name: name, age: age}, nil
}

func testPerson05() {
	var person, myError = newPerson05("Tom", 34)
	if myError == nil {
		fmt.Printf("%v\n", *person)
	} else {
		fmt.Printf("%v\n", myError)
	}
}

// Animal
// golang本质上没有oop的概念，也没有继承的概念，但是可以通过结构体嵌套实现这个特性。
type Animal struct {
	name string
	age  int
}

func (animal Animal) eat01() {
	fmt.Println("eat...")
}

func (animal Animal) sleep01() {
	fmt.Println("sleep...")
}

type Dog01 struct {
	animal Animal
	color  string
}

type Cat01 struct {
	animal Animal
	breed  string
}

func testAnimal() {
	var dog = Dog01{
		animal: Animal{
			name: "pp",
			age:  34,
		},
		color: "gray",
	}
	dog.animal.eat01()
	dog.animal.sleep01()
	fmt.Printf("%v\n", dog.animal.name)
	fmt.Printf("%v\n", dog.animal.age)
}

func testAnimal01() {
	var cat = Cat01{
		animal: Animal{
			name: "xiaomao",
			age:  3,
		},
		breed: "lan",
	}
	cat.animal.eat01()
	cat.animal.sleep01()
}

// Person04 oop
type Person04 struct {
	name string
	age  int
}

func (person Person04) eat() {
	fmt.Printf("eat...\n")
}

func (person Person04) sleep() {
	fmt.Printf("sleep...\n")
}

func (person Person04) work() {
	fmt.Printf("work...\n")
}

func testPerson04() {
	var person = Person04{
		name: "chen",
		age:  20,
	}
	fmt.Printf("%v\n", person)
	person.eat()
	person.sleep()
	person.work()
}

// PetNew opp
// golang没有面向对象的概念，也没有封装的概念，但是可以通过结构体struct和函数绑定来实现0OP的属性和方
// 法等特性。援收者receiver方法。
// 定义一个宠物接口
type PetNew interface {
	eat()
	sleep()
}

// Pig 定义pig结构体
type Pig struct {
}

// Cat 定义cat结构体
type Cat struct {
}

// pig实现PetNew接口方法
func (pig Pig) eat() {
	fmt.Println("pig eat...\n")
}

func (pig Pig) sleep() {
	fmt.Println("pig sleep...\n")
}

// cat实现PetNew接口方法
func (cat Cat) eat() {
	fmt.Printf("cat eat...\n")
}

func (cat Cat) sleep() {
	fmt.Printf("cat sleep...\n")
}

// Person03 定义一个person结构体
type Person03 struct {
}

// 为person添加一个养宠物的方法
func (person Person03) care(pet PetNew) {
	pet.eat()
	pet.sleep()
}

func testPetNew() {
	var pig = Pig{}
	var cat = Cat{}
	var person = Person03{}
	person.care(pig)
	person.care(cat)
}

// Fly 接口嵌套
type Fly interface {
	fly()
}

type Swim interface {
	swim()
}

// FlyFish 接口的组合
type FlyFish interface {
	Fly
	Swim
}

type Fish struct{}

func (f Fish) fly() {
	fmt.Println("fly...\n")
}

func (f Fish) swim() {
	fmt.Println("swim...\n")
}

func testFish() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}

// Music 一个类型可以实现多个接口
type Music interface {
	playerMusic()
}

type Video interface {
	playerVideo()
}

type Mobile02 struct {
	id int
}

func (m Mobile02) playerMusic() {
	fmt.Printf(": playerMusic\n")
}

func (m Mobile02) playerVideo() {
	fmt.Printf(": playerVideo\n")
}

func testMobile02() {
	var m = Mobile02{
		id: 110,
	}

	m.playerMusic()
	m.playerVideo()
}

// Pet02 多个类型可以实现同一个接口（多态）
type Pet02 interface {
	eat02()
}

type Dog02 struct {
}

type Cat02 struct {
}

func (d Dog02) eat02() {
	fmt.Printf("吃东西\n")
}

func (c Cat02) eat02() {
	fmt.Printf("猫吃东西\n")
}

func testPet02() {
	var pet Pet02
	pet = Dog02{}
	pet = Cat02{}
	pet.eat02()
	pet.eat02()
}

// Pet 本质上和方法的值类型接收者和指针类型接收者，的思考方法是一样的，值接收者是一个拷贝，是一个
// 副本，而指针接收者，传递的是指针。
type Pet interface {
	eat(name string) string
}

type Dog struct {
	name string
}

func (d *Dog) eat(name string) string {
	d.name = "1 花花..."
	fmt.Printf("%v\n", name)
	return "OK"
}

func testEat() {
	var dog = &Dog{
		name: "草草",
	}
	var newEat = dog.eat("新的花花")
	fmt.Printf("2 %v\n", newEat)
	fmt.Printf("3 %v\n", &dog)
}

// go语言的接口，是一种新的类型走义，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方
// 法就是实现了这个接口。
// 语法格式和方法非常类似。
type interfaceDemo1 interface {
	read()
	write()
}

type Computer struct {
	id int
}

type Mobile struct {
	model string
}

func (c Computer) read() {
	fmt.Printf("%v\n", c.id)
	fmt.Printf("read...\n")
}

func (c Computer) write() {
	fmt.Printf("%v\n", c.id)
	fmt.Printf("write...\n")
}

func testInterface() {
	var c = Computer{
		id: 1001,
	}

	c.read()
	c.write()
}

// Person02 结构体实例，有值类型和指针类型，那么方法的接收者是结构体，那么也有值类型和指针类型。区别就是接收者是
// 否复制结构体副本。值类型复制，指针类型不复制
type Person02 struct {
	name string
}

func methodDemo1() {
	var p1 = Person02{
		name: "chen",
	}
	var p2 = &Person02{
		name: "chen",
	}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)
}

func showPerson3(person Person02) {
	person.name = "liu值类型"
}

func showPerson4(person *Person02) {
	//自动解引用
	person.name = "chen指针类型"
}

func showPersonTest() {
	var p1 = Person02{
		name: "值类型",
	}
	var p2 = &Person02{
		name: "指针类型",
	}
	showPerson3(p1)
	fmt.Printf("%v\n", p1)
	showPerson4(p2)
	fmt.Printf("%v\n", *p2)

}

func (per Person02) showPerson5() {
	per.name = "liu值类型"
}

func (per *Person02) showPerson6() {
	per.name = "chen指针类型"
}

func testShowPerson2() {
	var p1 = Person02{
		name: "这个值类型不能改变",
	}
	var p2 = &Person02{
		name: "这是指针类型可以改变",
	}

	p1.showPerson5()
	fmt.Printf("%v\n", p1)
	p2.showPerson6()
	fmt.Printf("%v\n", *p2)
}

// 方法的receivertype并非一定要是struct类型，type定义的类型别名、slice、map、channel、func类型等都可以
// struct结合它的方法就等价于面向对象中的类。只不过struct可以和它的方法分开，并非一定要属于同一个文件，但必须属于同一个包
// 方法有两种接收类型：（TType）和（T*Type），它们之间有区别
// 方法就是函数，所以Go中没有方法重载（overload）的说法，也就是说同一个类型中的所有方法名必须都唯一
// 如果receiver是一个指针类型，则会自动解除引用
// 方法和type是分开的，意味着实例的行为（behavior）和数据存储（field）是分开的，但是它们通过receiver建立起关联关系
func loginDemo() {
	var customer = Customer{
		password: "1234",
	}
	var cus = customer.login("chen", "1234")
	fmt.Printf("%v\n", cus)
}

type Customer struct {
	password string
}

func (customer Customer) login(name string, password string) bool {
	fmt.Printf("%v\n", customer.password)
	if name == "chen" && password == "1234" {
		return true
	} else {
		return false
	}
}

func eatDemo() {
	var per = Person01{
		id: 89,
	}
	per.eat()
}

type Person01 struct {
	id int
}

// receiver
func (per Person01) eat() {
	fmt.Printf("%v\n", per.id)
}

func structNest() {
	type Dog struct {
		name  string
		color string
	}

	type Person struct {
		name string
		id   int
		dog  Dog
	}

	var myDog = Dog{
		name:  "tt",
		color: "black",
	}

	var per = Person{
		id:   1001,
		name: "chen",
		dog:  myDog,
	}
	fmt.Printf("%v\n", per.dog.name)
	fmt.Printf("%v\n", per)
}

func showDemo2() {
	var tom = Person{
		id:    3333,
		name:  "new",
		age:   33,
		email: "new@qq.com",
	}
	var per = &tom
	fmt.Printf("tom: %v\n", per)
	fmt.Printf("------------------\n")
	showPerson2(per)
	fmt.Printf("showPerson2(per): %v\n", per)
}

// 传递结构体指轩，这时在函数内部，能够改变外部结构体内容
func showPerson2(per *Person) {
	per.id = 2222
	per.name = "old"
	per.age = 22
	per.email = "old@qq.com"
	fmt.Printf("%v\n", per)
}

func showDemo1() {
	var tom = Person{
		id:    2222,
		name:  "new",
		age:   22,
		email: "new@qq.com",
	}
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("-------------------\n")
	showPerson(tom)
	fmt.Printf("showPerson(tom): %v\n", tom)
}

// 直接传递结构体，这是是一个副本（拷贝），在函数内部不会改变外面结构内容
func showPerson(person Person) {
	person.id = 1111
	person.name = "old"
	person.email = "old@qq.com"
	person.age = 11
	fmt.Printf("%v\n", person)
}

// 用new创建结构体指针
func newStruct() {
	type Person struct {
		id   int
		name string
	}

	var tom = new(Person)
	tom.id = 1101
	tom.name = "tom"

	fmt.Printf("%p\n", tom)
	fmt.Printf("%v\n", *tom)
}

func structPoint1() {
	type Person struct {
		id   int
		name string
		age  int
	}

	var tom = Person{
		id:   10201,
		name: "tom",
		age:  34,
	}

	var point_person *Person
	point_person = &tom

	fmt.Printf("%v\n", tom)
	fmt.Printf("%p\n", point_person)
	fmt.Printf("%v\n", *point_person)

}

// 结构体指针
func pointStruct() {
	var name string
	name = "tom"
	var point_name *string
	point_name = &name
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", point_name)
	fmt.Printf("%v\n", *point_name)

}

type Website struct {
	Nmae string
}

type Myf func(int, int) int

// go定义结构体
type Person struct {
	id    int
	name  string
	age   int
	email string
}

func structDemo1() {
	var tom Person
	tom.id = 1010
	tom.name = "beidixiaoxiong"
	tom.age = 43
	tom.email = "2334@qq.com"
	fmt.Printf("%v\n", tom)
}

// 匿名结构体
func structDemo2() {
	var tom struct {
		id   int
		name string
	}
	tom.id = 3043
	tom.name = "beidixiaoxiong"
	fmt.Printf("%v\n", tom)
}

// 类型定义&类型别名
func typeDemo1() {
	//类型定义
	type MyInt int
	var age MyInt
	age = 89
	fmt.Printf("%T %v\n", age, age)

	//类型别名
	type MyString = string
	var name MyString
	name = "beidixiaoxiong"
	fmt.Printf("%T %v\n", name, name)
}

// Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据。
// 类型指针不能进行偏移和运算。
// Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）積*（根据地址取值）
func pointDemo1() {
	var ip *int
	fmt.Printf("%v\n", ip)

	var i = 100
	ip = &i
	//取地址
	fmt.Printf("%v\n", ip)
	//取值
	fmt.Printf("%v\n", *ip)

	var strP *string
	var str = "beidiixaoxiong"
	strP = &str
	fmt.Printf("%v\n", strP)
	fmt.Printf("%v\n", *strP)
}

func pointDemo2() {
	var array = [...]int{1, 2, 4}
	var arrayPoint [3]*int
	fmt.Printf("%v\n", arrayPoint)

	for i := 0; i < len(array); i++ {
		arrayPoint[i] = &array[i]
	}
	fmt.Printf("%v\n", arrayPoint)

	for i := 0; i < len(arrayPoint); i++ {
		fmt.Printf("%v\n", *arrayPoint[i])
	}
}

// golang有一个特殊的函数init函数，先于main函数执行，实现包级别的一些初始化操作。
// init函数先于main函数自动执行，不能被其他函数调用；
// init函数没有输入参数、返回值；
// 每个包可以有多个init函数；
// 包的每个源文件也可以有多个init函数，这点比较特殊；
// 同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序。
// 不同包的init函数按照包导入的依赖关系决定执行顺序。
//func init() {
//	fmt.Printf("init...\n")
//}
//
////var demoInitVar = initVar()
//
//func initVar() int {
//	fmt.Printf("init var...\n")
//	return 100
//}

// defer
// go语言中的defer语句会将其后面跟随的语句进行延迟处理。
// 在defer归属的函数即将返回时，将延迟处理的语 句按defer定义的逆序进行执行
// ，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执 行。stack
func deferDemo() {
	fmt.Printf("start\n")
	defer fmt.Printf("start 01\n")
	defer fmt.Printf("start 02\n")
	defer fmt.Printf("start 03\n")
	fmt.Printf("END\n")
}

func fibonacci(n int) []int {
	var fib = make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// 斐波那契数列-递归
func fib(f int) int {
	if f == 2 || f == 1 {
		return 1
	}
	return fib(f-1) + fib(f-2)
}

// go 递归recursion
func recursionDemo2(a int) int {
	if a == 1 {
		//结束条件
		return 1
	} else {
		return a * recursionDemo2(a-1)
	}
}

// 阶乘
func recursionDemo1() {
	var a = 1
	for i := 1; i <= 10; i++ {
		a *= i
	}
	fmt.Printf("%v\n", a)
}

// go闭包closure
func closureAdd() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func closureDemo() {
	var demo = closureAdd()
	var demo1 = demo(10)
	fmt.Printf("%v\n", demo1)
	demo1 = demo(20)
	fmt.Printf("%v\n", demo1)
}

func closureDemo2(base int) (func(a int) int, func(a int) int) {
	var add = func(a int) int {
		base += a
		return base
	}

	var sub = func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

func useClosureDemo2() {
	var add, sub = closureDemo2(100)
	var newAdd = add(100)
	fmt.Printf("%v\n", newAdd)

	var newSub = sub(50)
	fmt.Printf("%v\n", newSub)

	var add1, sub1 = closureDemo2(100)
	var newAdd1 = add1(1)
	fmt.Printf("%v\n", newAdd1)
	var newSub1 = sub1(2)
	fmt.Printf("%v\n", newSub1)

}

// go匿名函数
// 介绍匿名函数的概念和语法格式
// 匿名函数可以省略函数名称，但不能嵌套函数
// 匿名函数可以自己调用自己，实现更复杂的功能
func funcDemo6() {
	var demo = func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	fmt.Printf("%v\n", demo(5, 9))
}

// 匿名函数自己调用自己
func funcDemo7() {
	var demo = func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(100, 89)
	fmt.Printf("%v\n", demo)
}

// 函数的类型
func typeAdd(a int, b int) int {
	return a + b
}

func typeSub(a int, b int) int {
	return a - b
}

// 把函数作为返回值
func calculateDemo(operator string) func(int, int) int {
	switch operator {
	case "+":
		return typeAdd
	case "-":
		return typeSub
	default:
		return nil
	}
}

func testType() {
	var demo1 = calculateDemo("+")
	var demo2 = demo1(34, 22)
	fmt.Printf("%v\n", demo2)
	demo1 = calculateDemo("-")
	var demo3 = demo1(100, 49)
	fmt.Printf("%v\n", demo3)
}

func funcType() {
	var myfunc Myf

	myfunc = sum
	var demo = myfunc(45, 22)
	fmt.Printf("%v\n", demo)

	myfunc = maxDemo
	demo = myfunc(203, 234)
	fmt.Printf("%v\n", demo)
}

func sayHello(name string) {
	fmt.Printf("%v hello world go\n", name)
}

// 把函数作为参数
func typeDemo(name string, myf func(string)) {
	myf(name)
}

func sum(a int, b int) int {
	return a + b
}

func maxDemo(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 函数
// 1.g0语言中有3种函数：普通函数、匿名函数（（没有名称的函数）、方法（定义在struct上的函数）。receiver
// 2.g0语言中不允许函数重载（overload），也就是说不允许函数同名。
// 3.g0语言中的函数不能嵌套函数，但可以嵌套匿名函数。
// 4.函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数。
// 5.函数可以作为参数传递给另一个函数。
// 6.函数的返回值可以是一个函数。
// 7，函数调用的时候，如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数。
// 8，函数参数可以没有名称
func funcDemo(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func funcDemo2() (n string, a int) {
	n = "chen"
	a = 34
	return n, a
}

func funcDemo3(s []int) {
	s[0] = 89
	//var demo = []int{3, 5, 3, 3}
	//funcDemo3(demo)
	//fmt.Printf("%v\n", demo)
}

// 可变参数
func funcDemo4(args ...int) {
	for _, arg := range args {
		fmt.Printf("%v\n", arg)
	}
}

func funcDemo5(name string, age int, args ...int) {
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", age)
	for _, arg := range args {
		fmt.Printf("%v\n", arg)
	}
}

// map
// map是一种key:value键值对的数据结构容器。map内部实现是哈希表（hash）
// map最重要的一点是通过key来快速检索数据，key类似于索引，指向数据
// map是引用类型
// map无序
func mapDemo5() {
	var m1 = map[string]string{"name": "chen", "age": "23", "email": "239@qq.com"}

	for k, v := range m1 {
		fmt.Printf("%v: %v\n", k, v)
	}

	for k := range m1 {
		fmt.Printf("%v\n", k)
	}
}

func mapDemo4() {
	var m1 = map[string]string{"name": "chen", "age": "23", "email": "239@qq.com"}
	var key1 = "email"
	var key2 = "nama"
	var v, ok = m1[key1]
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", ok)
	v, ok = m1[key2]
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", ok)
}

func mapDemo3() {
	var m1 = map[string]string{"name": "chen", "age": "23", "email": "239@qq.com"}
	var value = m1["email"]
	fmt.Printf("%v\n", value)
}

func mapDemo() {
	//类型的声明
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("%v\n", m1)
	fmt.Printf("%T\n", m1)
}

func mapDemo2() {
	var m1 = map[string]string{"name": "chen", "age": "23", "email": "289@qq.com"}
	fmt.Printf("%v\n", m1)

	var m2 = make(map[string]string)
	m2["name"] = "liu"
	m2["age"] = "23"
	m2["eamil"] = "233@qq.com"
	fmt.Printf("%v\n", m2)
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
		A4 int = iota
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
