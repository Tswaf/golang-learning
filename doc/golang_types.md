# Golang基本数据类型和变量

## 数据类型
Golang支持以下基本数据类型

| 名称 | 含义 |
| :--- | :--- |
| bool | 布尔值 true/false |  
| int8 | 8位整型 |  
| int16 | 16位整型 |  
| int32 | 32位整型 |  
| int64 | 64位整型 |
| int | 32或64位整型，依赖系统 |
| uint8 | 8位无符号整型 |
| uint16 | 16位无符号整型 |
| uint32 | 32位无符号整型 |
| uint64 | 64位无符号整型|
| uint | 32或64位无符号整型，依赖系统 |
| float32 | 32位浮点数 |
| float64 | 64位浮点数 |
| complex64 | 64位复数 |
| complex128 | 128位复数 |
| byte | unit8的别名 |
| rune | int32的别名 |
| string | 字符串 |

除了基本类型，Golang还支持数组、切片、map、指针和struct（结构体）类型，同时基于结构体类型实现方法、接口等来支持面向对象能力。

## 变量
### 声明变量
在Golang中，声明变量的语法为:  
*var varname type*  
*var*为变量声明关键字，*varname*为变量名称，*type*为变量类型。  

```
var a int
var b string
var c uint32
```

声明变量后还未显式给变量赋值，则变量会被自动赋值为*零值*。如整型类型变量的零值为*0*，字符串类型的零值为空串*“”*。  
可以在声明变量的同时为变量赋初始值：

```
var a int = 5
var b string = "golang"
```

### 类型推断
Golang允许在声明变量时省略变量类型，而根据初始值自动推断变量类型。  
*var varname = initValue*  

```
var a = 9
var b = "golang"
```

### 同时声明多个变量
可以一次性声明多个变量  
*var varA,varB type* //声明多个变量
*var varA,varB type = initValueA,initValueB* //声明多个变量，并分别赋予初始值
*var varA,varB = initValueA,initValueB* //类型推断

```
var a,b int
var a,b int = 3,6
var a,b = 3,"golang"
var (
    name = "golang"
    age = 5
)
```

### 短操作（short hand）声明变量
类型推断使得可以在声明变量时去掉变量类型，短操作再次省去*var*关键字
*varname := initValue*

```
a := 8
b := "golang"
c,d := 9,"golang"
```

使用短操作声明变量时有一个限制：*:=* 左边的变量至少有一个是之前未声明的，否则报错。

```
var a int = 4
a,b := 5,7 //error
```

## 类型转换
Golang是一门强数据类型的语言，变量被声明未一种类型后，只能使用该类型的值。

```
var a int = 6
var b float32 = 7.8
b = a //error
c := a + b //error
```

不同数据类型的变量之间相互使用，必选使用类型转换，类型转换的语法为：  
*Type(variable)*   *variable*是变量名称，*Type*是目标类型。  

```
a := 9
b := 9.8
c := a + int(b)
fmt.Println(a,b,c)
```

## 常量
使用*const*关键字声明常量。  

```
const a = 8
const b = "golang"
const (
    c = 9
    e = "month"
)
