# 2.4 struct类型
```Go
type person struct {
	name string
	age int
}
```
- 一个string类型的字段name，用来保存用户名称这个属性
- 一个int类型的字段age,用来保存用户年龄这个属性

使用struct
```Go
type person struct {
	name string
	age int
}

var P person  // P现在就是person类型的变量了

P.name = "Astaxie"  // 赋值"Astaxie"给P的name属性.
P.age = 25  // 赋值"25"给变量P的age属性
fmt.Printf("The person's name is %s", P.name)  // 访问P的name属性.
```
除了上面这种P的声明使用之外，还有另外几种声明使用方式：

- 1.按照顺序提供初始化值
	`P := person{"Tom", 25}`
- 2.通过`field:value`的方式初始化，这样可以任意顺序
	`P := person{age:24, name:"Tom"}`
- 3.当然也可以通过`new`函数分配一个指针，此处P的类型为*person
	`P := new(person)`

### struct的匿名字段
"struct_01" 介绍了如何定义一个struct，定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是**匿名字段**，也称为**嵌入字段**("struct_02" "struct_03"给出了示例)。

当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
> 内置类型和自定义类型也可以作为匿名字段的