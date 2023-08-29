package main

import (
	//"./helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets int = 50
var bookings = make([]UserDate, 0)

type UserDate struct {
	firstName     string
	lastName      string
	email         string
	numberTickets int
}

// weight group 有三个用法 wg.Add(1)，wg.Wait(), wg.Done()
// 总的来说 wg.Add() 添加一个计数器， 应用程序应该等待并完成的协程会减少该计数器，当计数器为零，就意味着主进程没有协程再等待，
// 然后 wg.Wait() 它就可以退出应用程序了
var wg = sync.WaitGroup{}

func main() {
	// 调用问候用户的函数, 并将变量的赋值传递进去
	greetUsers()

	//注释代码用于完成协程测试 for {
	// for {
	//
	firstName, lastName, email, userTickets := getUserInput()
	//
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	// 如果用户购买门票 <= 剩余门票 那么则执行订票逻辑
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)                                              // 创建一个新线程 (如果有多个 go 那么改成相应数量即可)
		go sendTicket(userTickets, firstName, lastName, email) //开启协程并发处理 (将 for 注释，进程会只执行一次就退出，观察协程输出。此时没有任何输出结果，因为主进程不会等待协程完成后在退出。)

		// 调用打印用户名称函数， 并将变量的赋值传递进去
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// 门票售空了，跳出循环
			fmt.Println("Our conference is booked out. Come back next year")

			//注释代码用于完成协程测试 break
			//break
		}
		// 如果用户购买门票 > 剩余门票那么提示用户我们没有那么多门票，并让用户重新进行订票
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
		//fmt.Println("Your input data is invalid, try again!")
		//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
	}
	wg.Wait() //等待的所有线程都放在这里，以遍在应用程序退出之前完成其工作
	//注释代码用于完成协程测试 for {
	//}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// 定义结构体来创建 map 对象
	var userData = UserDate{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		numberTickets: userTickets,
	}

	// create a map for user
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData) //包含所有用户信息的 KV 列表
	fmt.Printf("List of bookings is %v\n", bookings)
	//bookings = append(bookings, firstName+" "+lastName) //包含所有用户的名称切片

	/* 用于打印订票的一些数据类型
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The slice type: %T\n", bookings)
	fmt.Printf("The slice length: %v\n", len(bookings))
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", userTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking) 以空格为换行读取booking中的列表
		//var firstName = names[0] //初始化firstName变量值
		//firstNames = append(firstNames, booking["firstName"]) booking[]里面是追加的 key值
		firstNames = append(firstNames, booking.firstName) //使用结构体类型时，点出那个类即可，里面是追加的 key值
	}
	return firstNames
	//fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

// 验证用户输入
// GO 可以在一个函数中返回多个值
// 第二个括号用于定义变量的 return 的类型
func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

// 获取用户输入函数
func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	// 为了保持函数签名的一致性，分支中的错误处理，也必须返回四个值
	fmt.Println("Enter your first name: ")
	if _, err := fmt.Scan(&firstName); err != nil {
		return "", "", "", 0
	}
	fmt.Println("Enter your last name: ")
	if _, err := fmt.Scan(&lastName); err != nil {
		return "", "", "", 0
	}
	fmt.Println("Enter your email address: ")
	if _, err := fmt.Scan(&email); err != nil {
		return "", "", "", 0
	}
	fmt.Println("Enter number of tickets: ")
	if _, err := fmt.Scan(&userTickets); err != nil {
		return "", "", "", 0
	}
	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // 使用 fmt.Sprintf()才能将两个不同类型的值存储在变量中
	time.Sleep(10 * time.Second)
	fmt.Println("################")
	fmt.Printf("Sending ticket\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done() //执行完就删除一个线程
}
