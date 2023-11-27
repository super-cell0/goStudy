package main

import (
	"bufio"
	"bytes"
	"fmt"
	"godaily/user"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Timer
func timerDemo1() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("timer: %v\n", time.Now())
	timer01 := <-timer.C //阻塞的 知道时间到了
	fmt.Printf("%v\n", timer01)
}

func timerDemo2() {
	fmt.Printf("timer: %v\n", time.Now())
	var timer = time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("timerNow: %v\n", time.Now())
}

func timerDemo3() {
	<-time.After(time.Second * 2)
	fmt.Printf("time now\n")
}

func timeDemo4() {
	var timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Printf("func...\n")
	}()
	var stopOne = timer.Stop()
	if stopOne {
		fmt.Printf("stopOne...\n")
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("main end...\n")
}

func timeDemo05() {
	fmt.Printf("before\n")
	var timer = time.NewTimer(time.Second * 5) //原来设置5s
	timer.Reset(time.Second * 1)               //从新设置时间， 即修改newTimer的时间
	<-timer.C
	fmt.Printf("after\n")
}

// 1．select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中
// channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。
// select中的case语句必须是一个channel操作
// select中的default子句总是可运行的。
// 2．如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
// 3.如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
// 4.如果没有可运行的case语句，目没有default语句，select将阻塞，直到某个case通信可以运行
var chanInt = make(chan int, 0)
var chanString = make(chan string)

func selectDemo() {
	go func() {
		chanInt <- 100
		chanString <- "hello go"
		defer close(chanInt)
		defer close(chanString)
	}()

	for {
		select {
		case v := <-chanInt:
			fmt.Printf("chanInt: %v\n", v)
		case v := <-chanString:
			fmt.Printf("chanString: %v\n", v)
		default:
			fmt.Printf("default...\n")
		}
		time.Sleep(time.Second)
	}
}

// chan 循环
var chanValue = make(chan int)

func chanDemo1() {
	go func() {
		for i := 0; i < 2; i++ {
			chanValue <- i
		}
		close(chanValue)
	}()

	//var newValue = <-chanValue
	//fmt.Printf("%v\n", newValue)
	//newValue = <-chanValue
	//fmt.Printf("%v\n", newValue)

	//只有两个值但是读到第三个值的时候会发送死锁
	//for i := 0; i < 3; i++ {
	//	var newValue = <-chanValue
	//	fmt.Printf("%v\n", newValue)
	//
	//}

	for i2 := range chanValue {
		fmt.Printf("%v\n", i2)
	}

	//for {
	//	value, ok := <-chanValue
	//	if ok {
	//		fmt.Printf("%v\n", value)
	//	} else {
	//		break
	//	}
	//}
}

// golang并发编程之Mutex互斥锁实现同步
var a = 100
var waitg1 sync.WaitGroup
var lock sync.Mutex

func add() {
	defer waitg1.Done()
	lock.Lock()
	a += 1
	fmt.Printf("%v\n", a)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	defer waitg1.Done()
	a -= 1
	fmt.Printf("%v\n", a)
	lock.Unlock()

}

func testAdd() {
	for i := 0; i < 100; i++ {
		waitg1.Add(1)
		add()
		waitg1.Add(1)
		sub()
	}

	waitg1.Wait()

	fmt.Printf("end...: %v\n", a)
}

// GOMAXPROCS
func sendA() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %v\n", i)
	}
}

func sendB() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %v\n", i)
	}
}

func testSendAB() {
	fmt.Printf("numCPU: %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go sendA()
	go sendB()
	time.Sleep(time.Second)
}

// runtime.Goexit()
func showGoexit() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

// runtime Gosched
func showMessage(message string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("%v\n", message)
	}
}

func testShowMessage() {
	go showMessage("golang") //子协程
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Printf("%v\n", "gosched")
	}
	fmt.Printf("main golang\n")
}

var wp sync.WaitGroup

func sendMessage(i int) {
	defer wp.Done()
	fmt.Printf("hello goroutine %v\n", i)
}

func testSendMsg() {
	for i := 0; i < 5; i++ {
		go sendMessage(i)
		wp.Add(1)
	}

	wp.Wait()
	fmt.Printf("end...\n")
}

// Go提供了一种称为通道的机制，用于在goroutine之间共享数据。当您作为goroutine执行并发活动时，需要
// 在goroutine之间共享资源或数据，通道充当goroutine之间的管道（管道）并提供一种机制来保证同步交换。
// 需要在声明通道时指定数据类型。我们可以共享内置、命名、结构和引用类型的值和指针。数据在通道上传递：在
// 任何给定时间只有一个goroutine可以访问数据项：因此按照设计不会发生数据竞争。
// 根据数据交换的行为，有两种类型的通道：无缓冲通道和缓冲通道。无缓冲通道用于执行goroutine之间的同步通 信，而缓冲通道用于执行异步通信。
// 无缓冲通道保证在发送和接收发生的瞬间执行两个goroutine之间的交换。缓
// 冲通道没有这样的保证。
// 通道由make函数创建，该函数指定chan关键字和通道的元素类型。
// 这是将值发送到通道的代码块需要使用 <- 运算符：
// <- 运算符附加到通道变量goroutine的左侧，以接收来自通道的值
// 无缓冲通道
// 在无缓冲通道中，在接收到任何值之前没有能力保存它。在这种类型的通道中，发送和接收goroutine在任何发送
// 或接收操作完成之前的同一时刻都准备就绪。如果两个goroutine没有在同一时刻准备好，则通道会让执行其各自
// 发送或接收操作的goroutine首先等待。同步是通道上发送和接收之间交互的基础。没有另一个就不可能发生。
// Golang中的并发是函数相互独立运行的能力。Goroutines是并发运行的函数。Golang提供了Goroutines作为并发 处理操作的一种方式
// 缓冲通道
// 在缓冲通道中，有能力在接收到一个或多个值之前保存它们。在这种类型的通道中，不要强制goroutine在同一时 刻准备好执行发送和接收。
// 当发送或接收阳塞时也有不同的条件。只有当通道中没有要接收的值时，接收才会阻
// 塞。仅当没有可用缓冲区来放置正在发送的值时，发送才会阻塞。
// 通道的发送和接收特性
// 1．对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
// 2．发送操作和接收操作中对元素值的处理都是不可分割的。
// 3，发送操作在完全完成之前会被阻塞。接收操作也是如此
var values = make(chan int)

func send() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	var value = random.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value

}

func testSend() {
	defer close(values)
	go send()
	fmt.Printf("wait...\n")
	var value = <-values
	fmt.Printf("recieve: %v\n", value)
	fmt.Printf("end...\n")
}

// 协程
func show(message string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", message)
		time.Sleep(time.Millisecond * 100)
	}
}

func testShow() {
	////启动一个协程来执行
	go show("hello go")
	go show("golang...")
	time.Sleep(time.Millisecond * 2000)
	fmt.Printf("this is golang\n")
}

func responseSize(url string) {
	fmt.Printf("step1: %v\n", url)
	var response, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("step2: %v\n", url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	fmt.Printf("step3: %v\n", url)
	var body, err01 = io.ReadAll(response.Body)
	if err01 != nil {
		log.Fatal(err01)
	}
	fmt.Printf("step4: %v\n", len(body))
}

func responseSize01(url string) {
	fmt.Printf("step1: %v\n", url)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	fmt.Printf("step2: %v\n", url)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("step3: %v\n", url)
	fmt.Printf("step4: %v\n", len(body))
}

func testResponseSize() {
	go responseSize("https://www.duoke360.com")
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	time.Sleep(10 * time.Second)
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
	fmt.Printf("pig eat...\n")
}

func (pig Pig) sleep() {
	fmt.Printf("pig sleep...\n")
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

type Fish struct {
}

func (f Fish) fly() {
	fmt.Printf("fly...\n")
}

func (f Fish) swim() {
	fmt.Printf("swim...\n")
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

// ticker
// Timer只执行一次，Ticker可以周期的执行。
func tickerDemo1() {
	var ticker = time.NewTicker(time.Second)
	var counter = 1
	for range ticker.C {
		fmt.Printf("ticker...\n")
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}

func tickerDemo2() {
	var ticker = time.NewTicker(time.Second)
	var chans = make(chan int)

	go func() {
		for range ticker.C {
			select {
			case chans <- 1:
			case chans <- 2:
			case chans <- 3:
			}
		}
	}()

	sum := 0
	for v := range chans {
		fmt.Printf("v: %v\n", v)
		sum += v
		if sum >= 10 {
			break
		}
	}
}

// 原子变量
var atomValue = 100

func add1() {
	lock.Lock()
	atomValue++
	lock.Unlock()
}

func sub1() {
	lock.Lock()
	atomValue--
	lock.Unlock()
}

func testAtom() {
	for i := 0; i < 100; i++ {
		go add1()
		go sub1()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("%v\n", atomValue)
}

var atom2 int32 = 100

func add2() {
	atomic.AddInt32(&atom2, 1)
}

func sub2() {
	atomic.AddInt32(&atom2, -1)
}

func testAtom2() {
	for i := 0; i < 100; i++ {
		go add2()
		go sub2()
	}

	time.Sleep(time.Second * 2)
	fmt.Printf("%v\n", atom2)
}

// atomic常见操作有：
// •增减
// 。载入read
// 。比较并交换cas
// •交换
// 。存储write
func atomDetail1() {
	var num int32 = 100
	atomic.AddInt32(&num, 20)
	fmt.Printf("%v\n", num)
	atomic.AddInt32(&num, -30)
	fmt.Printf("%v\n", num)
}

func atomDetail2() {
	var num int32 = 100
	atomic.LoadInt32(&num) //read
	fmt.Printf("%v\n", num)

	atomic.StoreInt32(&num, 200) //write
	fmt.Printf("%v\n", num)
}

func casDetail() {
	var num int32 = 100
	var newNum = atomic.CompareAndSwapInt32(&num, 100, 200)
	fmt.Printf("%v\n", newNum)
	fmt.Printf("%v\n", num)
}

// golang标准库os模块-文件目录
func createFile() {
	myFile, err := os.Create("my_golang.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", myFile.Name())
	}
}

// 创建目录
func createMake() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func createMake2() {
	err := os.MkdirAll("demo_dir/model/other", os.ModePerm)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

// 删除目录和文件
func removeFile() {
	err := os.Remove("my_golang2.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func removeDir() {
	err := os.RemoveAll("demo_dir")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

// 获得工作目录
func haveGetwd() {
	dir, _ := os.Getwd()
	fmt.Printf("%v\n", dir)
}

// 修改工作目录
func changeDir() {
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %vIn", err)
	}
	fmt.Printf(os.Getwd())
}

// 获取零时目录
func getTemp() {
	dir := os.TempDir()
	fmt.Printf("%v\n", dir)
}

// 重命名
func rename() {
	err := os.Rename("my_golang.txt", "my_golang.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

// 读文件
func readFile() {
	file, err := os.ReadFile("my_golang.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", string(file[:]))
	}
}

// 写文件
func writeFile() {
	var str = "hello golang"
	err := os.WriteFile("my_golang.txt", []byte(str), os.ModePerm)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func openCloseFile() {
	file, err := os.Open("my_golang.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", file.Name())
		err = file.Close()
		if err != nil {
			return
		}
	}
}

func openFile1() {
	file, err := os.OpenFile("my_golang2.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("fileName: %v\n", file.Name())
		err = file.Close()
		if err != nil {
			return
		}
	}
}

func readOpen2() {
	var file, err = os.Open("my_golang.txt")
	for {
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		var buffer = make([]byte, 10)
		var myLen, err1 = file.Read(buffer)
		if err1 == io.EOF {
			break
		}
		fmt.Printf("%v\n", myLen)
		fmt.Printf("%v\n", string(buffer))
	}
	err = file.Close()
	if err != nil {
		return
	}
}

func readAtOpen() {
	file, err := os.Open("my_golang.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	var buf = make([]byte, 3)
	n, _ := file.ReadAt(buf, 3)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", string(buf))
}

// 判断是文件还是一个目录
func demoFile() {
	dir, err := os.ReadDir("my_golang.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, d := range dir {
			fmt.Printf("d.isDir: %v\n", d.IsDir())
			fmt.Printf("d.Name: %v\n", d.Name())
		}
	}
}

func seekFile() {
	file, _ := os.Open("my_golang.txt")
	file.Seek(3, 0)
	var buf = make([]byte, 10)
	n, _ := file.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	file.Close()
}

func writeDemo() {
	file, err := os.OpenFile("my_golang.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("%n\n", err)
	}
	file.Write([]byte("hello golang... this is ... 这样做可以确保在函数返回之前关闭文件，即使在函数中间有其他操作。"))
	file.Close()
}

func writeString() {
	file, _ := os.OpenFile("my_golang.txt", os.O_RDWR|os.O_APPEND, 0755)
	file.WriteString("hello golang 8899")
	file.Close()
}

func writeAtDemo() {
	file, _ := os.OpenFile("my_golang.txt", os.O_RDWR, 0755)
	file.WriteAt([]byte("world"), 6)
	file.Close()
}

// os进程相关的 进程（Process）和线程（Thread）--window
func processDemoForWindow() {
	//获取当前正在运行的进程id
	fmt.Printf("os.Getpid: %v\n", os.Getpid())
	//父id
	fmt.Printf("os.Getppid: %v\n", os.Getppid())

	var attr = &os.ProcAttr{
		//files指定新进程继承的活动文件对象
		//前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},

		//新进程的环境变量
		Env: os.Environ(),
	}

	//开始一个进程
	process, err := os.StartProcess("/System/Applications/TextEdit.app", []string{"/Users/mac/Downloads/swift.txt"}, attr)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(process)
	fmt.Println("进程ID: ", process.Pid)

	//通过进程TD查找进程
	findProcess, _ := os.FindProcess(process.Pid)
	fmt.Println(findProcess)

	//等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		//向p进程发送退出信号
		process.Signal(os.Kill)
	})

	///等待进程p的退出，返回进程状态
	processWait, _ := process.Wait()
	fmt.Println(processWait.String())
}

// mac
func processDemoForMac() {
	fmt.Printf("os.Getpid: %v\n", os.Getpid())
	fmt.Printf("os.Getppid: %v\n", os.Getppid())

	cmd := exec.Command("open", "/Users/mac/Downloads/swift.txt")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("进程ID: ", cmd.Process.Pid)

	// 通过进程PID查找进程
	findProcess, _ := os.FindProcess(cmd.Process.Pid)
	fmt.Println("findProcess: ", findProcess)

	// 等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		// 向cmd进程发送退出信号
		cmd.Process.Signal(os.Kill)
	})

	// 等待进程cmd的退出，返回进程状态
	processWait, _ := cmd.Process.Wait()
	fmt.Println("processWait: ", processWait.String())
}

func processDemoForMac2() {
	fmt.Printf("os.Getpid: %v\n", os.Getpid())
	fmt.Printf("os.Getppid: %v\n", os.Getppid())

	cmd := exec.Command("open", "/Users/mac/Downloads/swift.txt")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("进程ID: ", cmd.Process.Pid)

	// 通过进程PID查找进程
	findProcess, _ := os.FindProcess(cmd.Process.Pid)
	fmt.Println(findProcess)

	// 等待应用程序完成
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error waiting for process: %v\n", err)
	}

	fmt.Println("应用程序完成")

	err = findProcess.Kill()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("应用程序被终止")
	}
}

// ENV
func envDemo() {
	//fmt.Printf("os.Environ%v\n", os.Environ())
	getenv := os.Getenv("GOPATH")
	fmt.Printf("%v\n", getenv)
}

func envDemo2() {
	env, b := os.LookupEnv("GOPATH")
	if b {
		fmt.Printf("%v\n", env)
	}
}

// Go语言中，为了方便开发者使用，将I0操作封装在了如下几个包中：
// •io为10原语（I/0primitives）提供基本的接口osFileReaderWriter
// •io/ioutil封装一些实用的1/0函数
// •fAht实现格式化1/0，类似C语言中的printf和scanf
// •bufio实现带缓冲1/0
func testCopy() {
	reader := strings.NewReader("hello golang\n")
	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO io包源码
// ReadAll
// 读取数据，返回读到的字节slice
// ReadDir
// 读取一个目录，返回目录入口数组［los.Filelnfo
// ReadFile
// 读一个文件，返回文件内容（字持Slice）
// WriteFile
// 根据文件路径，写入字节slice
// TempDir
// 在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径
// TempFile
// 在一个目录中创建指定前缀名的临时文件，返回os.File
func readAllDemo() {
	open, err := os.Open("my_golang.txt")
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(open)
	all, err := io.ReadAll(open)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", string(all))
}

// TODO
// bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时
// 还提供了缓冲和一些文本I/O的帮助函数的对象。
func bufioDemo() {
	//reader := strings.NewReader("hello go")
	open, _ := os.Open("my_golang.txt")
	newReader := bufio.NewReader(open)
	readString, _ := newReader.ReadString('\n')
	fmt.Printf("%v\n", readString)

}
func bufioDemo2() {

}

func bufioDemo3() {
	reader := strings.NewReader("ABCDEFGHIJK")
	newReader := bufio.NewReader(reader)

	readByte, _ := newReader.ReadByte()
	fmt.Printf("%v\n", string(readByte))

	b, _ := newReader.ReadByte()
	fmt.Printf("%v\n", string(b))

	err := newReader.UnreadByte()
	if err != nil {
		return
	}
	b2, _ := newReader.ReadByte()
	fmt.Printf("%v\n", string(b2))
}

func bufioDemo4() {
	reader := strings.NewReader("你好 世界！")
	newReader := bufio.NewReader(reader)

	r, size, _ := newReader.ReadRune()
	fmt.Printf("%v %v\n", string(r), size)

	readRune, s, _ := newReader.ReadRune()
	fmt.Printf("%v, %v\n", string(readRune), s)

	newReader.UnreadRune()
	r2, i, _ := newReader.ReadRune()
	fmt.Printf("%v, %v\n", string(r2), i)
}

func bufioDemo5() {
	reader := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	newReader := bufio.NewReader(reader)

	line, prefix, _ := newReader.ReadLine()
	fmt.Printf("%q, %v\n", line, prefix)

	readLine, isPrefix, _ := newReader.ReadLine()
	fmt.Printf("%q, %v\n", readLine, isPrefix)

	l, p, _ := newReader.ReadLine()
	fmt.Printf("%q, %v\n", l, p)

	line2, b, _ := newReader.ReadLine()
	fmt.Printf("%q, %v\n", line2, b)
}

func bufioDemo6() {
	reader := strings.NewReader("ABC,DFG,DIJ,UYT")
	newReader := bufio.NewReader(reader)

	line1, _ := newReader.ReadSlice(',')
	fmt.Printf("%v\n", string(line1))

	line2, _ := newReader.ReadSlice(',')
	fmt.Printf("%v\n", string(line2))

	line3, _ := newReader.ReadSlice(',')
	fmt.Printf("%v\n", string(line3))

	line4, _ := newReader.ReadSlice(',')
	fmt.Printf("%v\n", string(line4))
}

func bufioDemo7() {
	reader := strings.NewReader("ABC HDGSU SA KDS AS")
	newReader := bufio.NewReader(reader)

	readBytes1, _ := newReader.ReadBytes(' ')
	fmt.Printf("%q\n", readBytes1)

	readBytes2, _ := newReader.ReadBytes(' ')
	fmt.Printf("%q\n", readBytes2)

	readBytes3, _ := newReader.ReadBytes(' ')
	fmt.Printf("%q\n", readBytes3)

	readBytes4, _ := newReader.ReadBytes(' ')
	fmt.Printf("%q\n", readBytes4)

	readBytes5, _ := newReader.ReadBytes(' ')
	fmt.Printf("%q\n", readBytes5)
}

func bufioDemo8() {
	reader := strings.NewReader("\nAJSJABSDJSAS\n")
	newReader := bufio.NewReader(reader)

	//buffer := bytes.NewBuffer(make([]byte, 0))
	//写入文件
	file, _ := os.OpenFile("my_golang.txt", os.O_RDWR|os.O_APPEND, 0777)
	defer file.Close()
	newReader.WriteTo(file)
	//fmt.Printf("%v\n", buffer)

}

func bufioDemo9() {
	file, _ := os.OpenFile("my_golang.txt", os.O_RDWR|os.O_APPEND, 0777)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("concerning: 关于、涉及\n")
	writer.Flush()
}

func bufioDemo10() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}

// 11-27 TODO
func scannerDemo() {
	reader := strings.NewReader("hello golang fighting!")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func scanDemo() {
	reader := strings.NewReader("你 好 世界 你在干嘛哦 哎哟")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		fmt.Printf("%s\n", scanner.Text())
	}
}

func logDemo1() {
	log.Print("hello golang\n")

	log.Printf("%d", 239)

	name := "Tom"
	age := 23
	log.Println(name, " ", age)
}

func panicDemo1() {
	defer fmt.Printf("panic结束之后的执行。。。\n")
	log.Print("hello golang")
	log.Panic("my panic")
	fmt.Printf("end...\n")
}

func main() {
	panicDemo1()
}
