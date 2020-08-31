package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"unsafe"
)

type Home struct {
	beego.Controller
}

func (c *Home) Index() {
	//基本操作()
	//数据类型()
	//变量()
	//常量()
	//运算符()
	//条件语句()
	//循环语句()
	//函数()
	//变量的作用域()
	//数组()
	//指针()
	//结构体()
	动态数组()
}

func 基本操作() {
	//int [指定类型]
	var i int = 100

	//float [key := value] 根据值进行判断类型
	f := 123.05

	//定义字符串
	var str string = "World"

	fmt.Println("Console 输出语句")
	fmt.Printf("格式化:%d\n", i)
	fmt.Println("Hello" + " " + str)
	fmt.Printf("float:%.2f\n", f)
}

func 数据类型() {

	var b bool = true
	var i int = 0xFFFFFFFFFFFFFFFF / 2     //(int)根据当前系统32位 或者64位 8字节 -9223372036854775808 ~ 9223372036854775807 (18446744073709551615)
	var i8 int8 = 0xFF / 2                 //-128 ~ 127  (255)
	var i16 int16 = 0xFFFF / 2             //-32768 ~ 32767 (65535)
	var i32 int32 = 0xFFFFFFFF / 2         //-2147483648 ~ 2147483647 (4294967295)
	var i64 int64 = 0xFFFFFFFFFFFFFFFF / 2 //-9223372036854775808 ~ 9223372036854775807 (18446744073709551615)
	var ui uint = 1
	var ui8 uint8 = 1
	var ui16 uint16 = 1
	var ui32 uint32 = 1
	var ui64 uint64 = 1
	var f32 float32 = 0.5
	var f64 float64 = 0.55
	var str string = "string"
	var bt byte = 0b000000101
	var re rune = 0xFFFFFFFF / 2
	var ip *uint = &ui

	//fmt.Println(b,i,i8,i16,i32,i64,ui,ui8,ui16,ui32,ui64,f32,f64,str);
	fmt.Printf("b\t=> %v \t%d byte\n", b, unsafe.Sizeof(b))
	fmt.Printf("i\t=> %d \t%d byte\n", i, unsafe.Sizeof(i))
	fmt.Printf("i8\t=> %d \t%d byte\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("i16\t=> %d \t%d byte\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("i32\t=> %d \t%d byte\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("i64\t=> %d \t%d byte\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("ui\t=> %d \t%d byte\n", ui, unsafe.Sizeof(ui))
	fmt.Printf("ui8\t=> %d \t%d byte\n", ui8, unsafe.Sizeof(ui8))
	fmt.Printf("ui16\t=> %d \t%d byte\n", ui16, unsafe.Sizeof(ui16))
	fmt.Printf("ui32\t=> %d \t%d byte\n", ui32, unsafe.Sizeof(ui32))
	fmt.Printf("ui64\t=> %d \t%d byte\n", ui64, unsafe.Sizeof(ui64))
	fmt.Printf("f32\t=> %f \t%d byte\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("f64\t=> %f \t%d byte\n", f64, unsafe.Sizeof(f64))
	fmt.Printf("str\t=> %s \t%d byte\n", str, unsafe.Sizeof(str))
	fmt.Printf("bt\t=> %d \t%d byte\n", bt, unsafe.Sizeof(bt))
	fmt.Printf("re\t=> %d \t%d byte\n", re, unsafe.Sizeof(re))
	fmt.Printf("ip\t=> %d \t%d byte\n", ip, unsafe.Sizeof(ip))

}

func 变量() {
	//声明单个
	var i int = 100
	fmt.Printf("i\t=> %d\n", i)

	//声明多个
	var i1, i2 int = 1000, 2000
	fmt.Printf("i1\t=> %d\n", i1)
	fmt.Printf("i2\t=> %d\n", i2)

	//根据值自行判定变量类型
	var str = "string"
	fmt.Printf("str\t=> %s\n", str)

	//省略 var
	vBool := true
	fmt.Printf("vBool\t=> %v\n", vBool)

	//定义多个不同类型的变量
	var (
		i3   int    = 1
		str1 string = "string1"
	)
	fmt.Printf("i3\t=> %v\n", i3)
	fmt.Printf("str1\t=> %v\n", str1)

	//不同类型值可多赋值
	var i4 int
	var str2 string
	i4, str2 = 100, "string2"
	fmt.Printf("i4\t=> %v\n", i4)
	fmt.Printf("str2\t=> %v\n", str2)

	//不同类型进行初始化声明
	i5, str3, f := 100, "string3", 0.5
	fmt.Printf("i5\t=> %v\n", i5)
	fmt.Printf("str3\t=> %v\n", str3)
	fmt.Printf("f\t=> %v\n", f)

	//相同类型相互交换值
	var i6, i7 int = 1, 2
	i6, i7 = i7, i6
	fmt.Printf("i6\t=> %v\n", i6)
	fmt.Printf("i7\t=> %v\n", i7)
}

func 常量() {

	//单个声明初始化
	const i int = 100
	//省略类型
	const str = "string"
	//多变量声明
	const i1, str1 = 100, "string1"

	//枚举 iota每次递增
	const (
		stand int = iota
		walk
		run
	)

	fmt.Printf("stand: %d\n", stand)
	fmt.Printf("walk: %d\n", walk)
	fmt.Printf("run: %d\n", run)
}

func 运算符() {

	//加法 +
	var num = 5.0
	var money = 12.5

	num = num + 4
	fmt.Printf("num\t=> %.2f\n", num)

	//减法 -
	num = num - 2
	fmt.Printf("num\t=> %.2f\n", num)

	//乘法 *
	var total = money * num
	fmt.Printf("total\t=> %.2f\n", total)

	//除法 /
	var avg = total / num
	fmt.Printf("avg\t=> %.2f\n", avg)

	/*
		num     => 9.00
		num     => 7.00
		total   => 87.50
		avg     => 12.50
	*/

	//运算符
	var a = 1
	var b = 1
	var c = 2
	if a == b {
		fmt.Println("a == b")
	}
	if c != b {
		fmt.Println("c != b")
	}
	if c != b {
		fmt.Println("c != b")
	}
	if c >= b {
		fmt.Println("c >= b")
	}
	if b <= c {
		fmt.Println("b <= c")
	}
	if a == b && c != b {
		fmt.Println("a == b && c != b")
	}

	//位运算
	var b1 byte = 0b1010
	var b2 byte = 0b1100
	fmt.Printf("b1 & b2\t=> %.8b\n", (b1 & b2))
	fmt.Printf("b1 | b2\t=> %.8b\n", (b1 | b2))
	fmt.Printf("b1 ^ b2\t=> %.8b\n", (b1 ^ b2))

	//左右右移
	var bb1 byte = 0b10001101
	fmt.Printf("bb1>>4\t=> %.8b\n", (bb1 >> 4))
	fmt.Printf("bb1<<4\t=> %.8b\n", (bb1 << 4))

	//取地址 解引用
	var i int = 4
	var ptr *int = &i
	fmt.Printf("i\t=> %d\n", i)
	fmt.Printf("&i\t=> %x\n", &i)
	fmt.Printf("ptr\t=> %x\n", ptr)
	fmt.Printf("*ptr\t=> %d\n", *ptr)

}

func 条件语句() {

	var b bool = true
	var b1 bool = false

	//if
	if b {
		fmt.Println("b = true")
	} else if b1 {
		fmt.Println("b1 = true")
	} else {
		fmt.Println("b 和 b1 都是 false")
	}

	//switch
	var s1 = 3
	switch s1 {
	case 1:
		fmt.Println("s1 = 1")
		break
	case 2:
		fmt.Println("s1 = 2")
		break
	default:
		fmt.Println("s1 = " + fmt.Sprintf("%d", s1))
	}

	//select
	var c1, c2 chan int
	var i1, i2 int
	select {
	case c1 <- i1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	default:
		fmt.Printf("no communication\n")
	}
}

func 循环语句() {

	//for循环
	//(1)
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//(2)
	var i1 = 0
	for ; i1 <= 5; i1++ {
		fmt.Println(i1)
	}

	//(3)
	sum := 5
	for sum >= 0 {
		fmt.Println(sum)
		sum--
	}

	//(4) 无限循环
	//for {
	//fmt.Println(sum);
	//sum--
	//}

	//(5)
	str := []string{"google", "baidu"}
	for index, value := range str {
		fmt.Println(index, value)
	}

	number := [6]int{1, 2, 3, 4, 5}
	for k, v := range number {
		fmt.Println(k, v)
	}
}

func 函数() {

	var i1, i2 = Sum(100, 200)
	fmt.Println("Sum()", i1, i2)
}
func Sum(num1, num2 int) (int, int) {

	return num1, num2
}

var g int

func 变量的作用域() {
	var a, b int = 1, 2
	g = a + b

	fmt.Println(a, b, g)
}

func 数组() {
	//初始化
	var arr [10]int
	var i int
	fmt.Println("\n#声明和初始化数据")
	for i = 0; i < 10; i++ {
		arr[i] = i
		fmt.Printf("%d ", arr[i])
	}

	//初始化
	fmt.Println("\n#声明和初始化数据")
	var balance = []float32{1000.0, 2.0, 3.0, 70.05, 55.25}
	for i = 0; i < len(balance); i++ {
		fmt.Printf("%.2f ", balance[i])
	}

	/**
	 * 多维数组
	 */
	//声明
	var two [2][2]int

	var n, j int
	fmt.Println("\n#声明；循环初始化2x2数组")
	for n = 0; n < 2; n++ {
		for j = 0; j < 2; j++ {
			two[n][j] = n * j
			fmt.Printf("a[%d][%d] = %d \n", n, j, two[n][j])
		}
	}

	//声明和初始化
	var three = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var arr2 = []int{1, 3, 4}
	//输出数组
	coutArr(three, arr2)

	fmt.Println("\n结束")
}
func coutArr(arr [2][3]int, arr2 []int) {
	fmt.Println("\n#从函数输出：声明并初始化2x3数组")
	var n, j int
	for n = 0; n < 2; n++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("arr[%d][%d] = %d \n", n, j, arr[n][j])
		}
	}

	for n = 0; n < len(arr2); n++ {
		fmt.Printf("arr2[%d] = %d \n", n, arr2[n])
	}
}

func 指针() {
	var a int = 1000
	fmt.Printf("变量a的地址：%x \n", &a)

	var pa *int = &a
	fmt.Printf("pa = %x \n", pa)

	var ppa **int = &pa

	fmt.Printf("ppa = %x \n", ppa)
	fmt.Printf("pa = %x \n", *ppa)
	fmt.Printf("a = %d \n", **ppa)

	//指针判断
	if pa != nil {
		getObj(&pa)
		fmt.Printf("getObj() => %d \n", *pa)
	}
}
func getObj(pa **int) {
	var arr int = 100
	*pa = &arr
}

type Person struct {
	name  string
	age   int
	money float32
	sex   byte
}

func 结构体() {
	fmt.Println(Person{"张三", 21, 1542.34, 1})
	fmt.Println(Person{name: "李莹", age: 22, money: 8542.21, sex: 0})
	fmt.Println(Person{name: "李珊莎"})

	//声明赋值
	var san Person
	san.name = "珊莎"
	san.age = 22
	san.money = 1583.36
	san.sex = 0
	fmt.Println(san)

	//传参数
	printPs(&san)
}
func printPs(ps *Person) {
	fmt.Println("san.name", ps.name)
}

func 动态数组() {
	//var arr []int = make([]int,5)
	ia := []int{1, 2, 3, 4}
	ib := ia[0:3]
	ic := ia[0:]
	id := ia[:]
	fmt.Println(ia)
	fmt.Println(ib)
	fmt.Println(ic)
	fmt.Println(id)
}
