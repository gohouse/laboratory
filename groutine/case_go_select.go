package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
)

type IndexController struct {
	Input string
}

type Student struct {
	Name string
	Age  string
	Sex  string
}

type Queue struct {
	List  []Student
}

func NewIndex()  *IndexController{
	return new(IndexController)
}

var scanner  *bufio.Scanner = bufio.NewScanner(os.Stdin)
var student   Student
var queue     Queue
var inputChan chan bool
func (this *IndexController)DrawMenu()  {
	fmt.Println("***********************终极一班管理系统***************************");
	fmt.Println("                       1.添加学生                                ")
	fmt.Println("                       2.删除学生                                ")
	fmt.Println("                       3.学生详情                                ")
	fmt.Println("                       4.学生列表                                ")
	fmt.Println("                       5.退出系统                                ")
	fmt.Println("****************************************************************")
	fmt.Println("欢迎使用班级管理系统,请输入对应的菜单编号:")
	//scanner
	scanner.Scan()
	input   := strings.Trim(scanner.Text()," ")
	this.Input = input
	this.ChooseMenu()
}

func (this *IndexController)ChooseMenu()  {
	switch this.Input {
	case "1":
		this.InputStudent()
	case "2":
	case "3":
	case "4":

	}
}

func (this *IndexController)InputStudent()  {
	fmt.Println("输入学生姓名:");
	scanner.Scan();
	student.Name = strings.Trim(scanner.Text(),"")
	fmt.Println("输入学生年龄:")
	scanner.Scan();
	student.Age = strings.Trim(scanner.Text(),"")
	fmt.Println("输入学生性别:")
	scanner.Scan();
	student.Sex = strings.Trim(scanner.Text(),"")
	queue.List =  append(queue.List, student)
	fmt.Println("添加中,请稍等...")
	//timerTick := time.NewTicker(time.Second * 3)
	//fmt.Println("添加成功,2秒后将跳转到\"学生列表\"菜单,继续输入请按'c'")
	fmt.Println("添加成功,2秒后将跳转")
	time.Sleep(time.Second* 2)
	this.StudentList()
	//go func(inputChan chan  bool) {
	//	scanner.Scan()
	//	input := strings.Trim(scanner.Text(),"\n")
	//	if input == "c"{
	//		inputChan<-true
	//		return
	//	}
	//}(inputChan)
	//select {
	//case <-timerTick.C:
	//	timerTick.Stop()
	//	this.StudentList()
	//case input := <-inputChan:
	//	if input{
	//		timerTick.Stop()
	//		this.InputStudent()
	//	}
	//}
}

func (this *IndexController)StudentList()  {
	fmt.Println("--------------------"+time.Now().Format("15:04:05")+"---------------------------")
	fmt.Println("姓名                       年龄                    性别");
	for _,value := range queue.List{
		fmt.Println(value.Name+"                           "+value.Age+"                        "+value.Sex)
	}
	fmt.Println("-------------------------------------------------------")
	// 用户输入协程
	go func(inputChan chan bool) {
		fmt.Println("按'b'回到菜单列表,按'q'退出程序...")
		for{
			input := strings.Trim(this.GetInput(scanner),"\n")
			switch input {
			case "b":
				inputChan <- true
				//return
				break
			case "q":
				inputChan <- false
				//return
				break
			default:
				fmt.Println("输入错误,请重新输入...")
			}
		}
	}(inputChan)
	// 主协程阻塞channel
	for  {
		select {
		case r := <-inputChan:
			if r == true {
				//this.DrawMenu()
				//close(inputChan)
				fmt.Println(r)
				close(inputChan)
				os.Exit(0)
			}else if r==false {
				fmt.Println(r)
				close(inputChan)
				os.Exit(0)
			}
		}
	}
}

func (this *IndexController)GetInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := strings.Trim(scanner.Text(),"\n")
	return  input
}

func (this *IndexController)Start()  {
	// 输入的线程
	inputChan = make(chan  bool)
	this.DrawMenu()
}

func main()  {
	index := NewIndex()
	index.Start()
}