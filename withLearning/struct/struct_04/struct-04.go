package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human类型拥有的字段
}

type Employee struct {
	Human      // 匿名字段Human
	speciality string
	phone      string // 雇员的phone字段
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}

/*
输出结果
Bob's work phone is: 333-222
Bob's personal phone is: 777-444-XXXX
---
当匿名字段有相同的属性时(如示例Employee和Human都有"phone" 这个属性)，最外层的会被优先访问。
示例中外层的字段是 Employee.phone
*/
