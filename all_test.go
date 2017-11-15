package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bastengao/bytesparser"
	"testing"
)

func TestPointer(t *testing.T) {
	//空指针，输出为nil
	var p *int
	fmt.Printf("p: %v\n", p)
	//指向局部变量，变量值初始为0
	var i int
	p = &i
	fmt.Printf("p: %v,%v\n", p, *p)
	//通过指针修改变量数值
	*p = 8
	fmt.Printf("p: %v,%v\n", p, *p)
	//数组的初始化及输出
	m := [3]int{3, 4, 5}
	fmt.Printf("m:%v--%v,%v,%v\n", m, m[0], m[1], m[2])
	//指针数组的初始化及输出
	//j, k, l := 3, 4, 5
	//x := [3]*int{&j, &k, &l}
	x := [3]*int{&m[0], &m[1], &m[2]}
	fmt.Printf("x:%v,%v,%v\n", x[0], x[1], x[2])
	fmt.Printf("*x:%v,%v,%v\n", *x[0], *x[1], *x[2])
	var n [3]*int
	n = x
	fmt.Printf("n:%v,%v,%v\n", n[0], n[1], n[2])
	//指向数组的指针，也即二级指针的使用
	y := []*[3]*int{&x}
	fmt.Printf("y:%v,%v\n", y, y[0])
	fmt.Printf("*y[0]:%v\n", *y[0])
	fmt.Printf("*y[][]:%v,%v,%v\n", *y[0][0], *y[0][1], *y[0][2])
}

type Student struct {
	name  string
	id    int
	score uint
}

func TestMemery(t *testing.T) {
	//new分配出来的数据是指针形式
	p := new(Student)
	p.name = "China"
	p.id = 63333
	p.score = 99
	fmt.Println(*p)
	//var定义的变量是数值形式
	var st Student
	st.name = "Chinese"
	st.id = 666333
	st.score = 100
	fmt.Println(st)
	//make分配slice、map和channel的空间，并且返回的不是指针
	var ptr *[]Student
	fmt.Println(ptr)     //ptr == nil
	ptr = new([]Student) //指向一个空的slice
	fmt.Println(ptr)
	*ptr = make([]Student, 3, 100)
	fmt.Println(ptr)
	stu := []Student{{"China", 3333, 66}, {"Chinese", 4444, 77}, {"Chince", 5555, 88}}
	fmt.Println(stu)
}

func TestParse(t *testing.T) {
	buff, _ := hex.DecodeString("6D7220306620303030303037363620303030303133653120303030303138303520303030303162663520323835322036352034303232343032322061303A300D0A")
	info := tofInfo{}
	offset, err := bytesparser.Parse(buff, &info)
	if err != nil {
		fmt.Println(err)
	}
	if offset > 0 {
		refine := refineTof(info)
		fmt.Printf("%v", refine)
	} else {
		t.Errorf("offset == %d", offset)
	}
}

func TestSender(t *testing.T) {
	startSender()
}

func TestExp(t *testing.T) {
	var fil filter
	fil.setR(3)
}
