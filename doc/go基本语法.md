主要基本语法可以参考[菜鸟教程](https://m.runoob.com/go/)，对于Golang的Api可以参考[这里](https://studygolang.com/)，本文主要记录和C/C++，Java等主流语言的不同之处。

### 标识符，关键字，命名规则

Go语言是一门区分大小写的语言，如果用大写，表示需要对外暴露；小写表示不需要对外暴露。

下面是一些命名的规范（大家约定俗成的写法）：

* 包名称要用小写，不要用下划线连接
* 文件名用小写，使用下划线连接
* 结构体和变量用驼峰，首字母根据是否暴露来决定大小写
* 接口和结构体类似，只是名字要用 er 结尾
* 常量用大写，下划线连接
* 单元测试文件命名 example_test.go，测试函数为 TestExample

### 变量

go语言声明之后必须使用，否则会报错

在变量声明的时候可以用短变量声明的方式，

```go
intVal := 1 
```

相等于：

```go
var intVal int 
intVal = 1 
// 或者
val intVal int = 1
// 或者
val intVal = 1 // 类型推断
```

**但是短变量声明不能放在函数的外面声明**，

匿名变量，表示变量没有被用到或者在函数返回值声明的时候使用，如

```
_, b := 1, 2
```

### 常量

在编译阶段已经确定的值，在用 const 声明就可以

iota 可以被认为是一个可被修改的常量，默认初始值为 0， 每调用一次自动加 1，遇到 const 被重置为 0，使用 _ 跳过某些值。

```
const A1 = iota // 0
_  // 1
const A2 = iota  // 2
const A3 = 100 
const A2 = iota  // 4
```

### 数据类型

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **布尔型** ：一个简单的例子：var b bool = true。             |
| 2    | **数字类型** 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型string:** Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 |
| 4    | **派生类型:** 包括：(a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型 |

数字类型比较多，更多内容可以参考[数字类型](https://m.runoob.com/go/go-data-types.html)

字符串可以用 ""表示，也可以用反引号 `来表示，反引号一般用来表示多行的字符串。

### 运算符

和其他语言基本一致

### 流程控制

#### 条件语句

不需要加括号，否则在保存的时候gofmt会自动把大括号去掉

```go
if a > b {

} else {

}
```

声明变量可以放在判断语句里面，但是作用域直在条件语句里有效

```go
if a:= 20, b := 10; a > b {

} else {

}
```

不能用 0 或者非 0 来表示真假

#### 循环语句

go语言中只有for循环

```go
for i := 0; i < 10; i++ {

}

x := []int {1, 2, 2}
for _, v := range x {  // 用 k, v 的方式循环
    
}

for i <= 10 {  // 类似 while 循环
    
}

for { // 永真循环
    
}
```

### 数组

数组如果没有初始化，会给每一个元素赋该类型的默认值

```go
var a [3]int 
```

初始化

```go
var a = [3]int {1, 2, 3}  // 使用初始化列表
var b = [...]int {1, 2, 3} // 初始化列表的长度表示数组的长度
var c = [...]int {0:1, 3:2, 5:3} // 索引 2, 4的值为默认值 0
```

### 切片（Slice）

是长度可变的数组

```go
var a []int
var b = make([]int)
var c = make([]int, 2)  // 设定初始化容量
```

初始化

```go
var a = []int {1, 2, 3}
var b := a[0:3] // 左闭右开区间
var c:= a[:] // 取全部
```

遍历

```go
x := []int {1, 2, 2}
for i := 0; i < len(x); i++ { 

}

for _, v := range x {  // 用 k, v 的方式循环
    
}
```

添加和删除

```go
x := []int {1, 2, 2}
x = x.append(x, 3)
x = x.append(x, x1, x2, x3...)
x = x.append(x[0:1], x[2:]...) // 相当于把 1 处的值给去掉了
```

### map

初始化

```go
m1 := make(map[string]string)
m1["name"] = "tom"

var m2 = map[string]string {
    "name" : "tom",
    "age" : "21"
}
```

遍历

```go
var m2 = map[string]string {
    "name" : "tom",
    "age" : "21"
}
for key := range m2 {  // 只遍历 key

}
for key, value := range m2 {  // 遍历 key 和 value

}
```

### 函数

* go语言里函数不允许重载
* 函数不能嵌套，但是匿名函数可以嵌套
* 函数可以作为参数，也可以作为函数值返回
* 函数的传参也是一个拷贝的过程
* 函数返回值可以有多个

```go
func sumAndMul(a int, b int) (int, int) {
	return a + b, a * b
}

c, d := sumAndMul(a, b)
```

函数可以传变长参数，传入的是拷贝的值，如果数据是引用类型的，传入后同样会改变元素的值

```go
func f(args...int) {
	for _, v := range args {
	
	}
}
```

函数类型和变量，所有满足下面这种类型的函数都可以直接被使用

```go
type f func(a int, b int) int
var a = f
a = sum
r := a(1, 2)
```

#### golang高阶函数（作为参数和返回值）

```go
func hello(name string) {
	fmt.Printf("hello %s!", name)
}
func f(name string, f func(name string)) {
	f(name)
}
```

#### 匿名函数

函数不能嵌套，但是可以定义匿名函数来事先某些功能

```go
func get_action() func() { 
    return func() { //匿名函数
        fmt.Println("Hello!")
    }
}

func main() {
    action := get_action()
    action()
}
```

匿名函数也可以自己调用自己

```go
func main() {
    func() {
        fmt.Println("Hello!")
    } ()
}
```

#### 闭包

引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量（例如局部变量）的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量。

```go
//1.简单闭包
func lbd() func() { //返回一个函数
	name := "lbdgood"
	return func() {//对于闭包来说，匿名函数被return回去，虽然name是局部变量，但也随匿名函数一起回去，也就是说即使lbd()这个函数结束并返回一个函数后，name依旧存在于引用它的函数里，没有消失
		fmt.Println("woo", name) //name已经被初始化一起return回去
	}
}
 
//2.复杂闭包
 
func makeSuffix(suffix string) func(string) string {
 
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
 
	}
}
 
func main() {
	//1
	gzq := lbd() //gz此时是一个闭包
	gzq()
 
	//2
	lbd := makeSuffix("love")
	fmt.Println(lbd("gzq"))
 
}
```

### defer语句

延迟（defer）语句，可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照**逆序**执行，最后该函数返回。

```go
func ReadWrite() bool {
    file.Open("file")
    defer file.Close()
    if failureX {
        return false
    }
    if failureY {
        return false
    }
    return true
}
```

### init函数

先于main函数执行，进行一些包初始化的操作。一个文件里可以有多个init函数，会从前往后自动执行，不能被其他函数调用

```go
func init() { //没有参数和返回值
	fmt.Printf("init...")
}
func main() {

}
```

### 指针

和C/C++类似，但是指针不允许预算

```go
var p *int
var a int = 100
p = &a
```

指向数组的指针，指向了数组的每一个元素

```go
const MAX int = 3
var p [MAX]*int // 表示数组里的元素类型是int指针
for i := 0; i < MAX; i++ {
	p[i] = &i;
}
```

### 类型定义与类型别名

和C类似

```go
tpye MyInt int  // 类型定义
tpye MyInt =  int  // 类型别名
```

* 类型定义是定义了一个新的类型MyInt，只不过它也是int类型的，而别名只是让给int起了一个别名。
* 类型别名只在编译中存在，编译后就不存在了。

### 结构体

没有初始化的结构体内部都是默认值

```go
type Books struct {
   title string
   subject string
   book_id int
}

func main() {
    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "Go 语言教程", 6495407})
    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", subject: "Go 语言教程", book_id: 6495407})
    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言"})
}
```

### 方法

特殊的函数，定义在结构体之上，和结构体之间做一个关联，该结构体就被成为接收者

```go
type Person struct {
    name string
}

func (person Person) eat() {
    fmt.Printf("eating %s", name)
}

func main() {
    per:= Preson {
        name : "tom",
    }
    per.eat()
}
```

### 接口

把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

```go
type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

}
```

要实现接口必须实现接口中的所有方法，在本例中只有一个方法。

另外要注意的是在方法的调用中我们一般都会取修改结构体中的值，如果原结构体中的值是基本类型，方法中会创建一个原值得拷贝，所以无法修改，但是引用类型可以被成功修改。

go语言接口可以被嵌套。

一个结构体如果实现了接口的所有方法，这个结构体就是该接口的类型。

### 继承

golang可以用结构体嵌套来实现继承

```go
type Animal struct {
	name string
	age int
}
type Dog struct {
	Animal
}
```

