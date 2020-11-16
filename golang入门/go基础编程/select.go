package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan int,1)
	ch2:=make(chan int,1)

	ch2<--1 
	// go func(){
	// 	ch1<-getNum()
	// }()
		

	select{
	case e1:=<-ch1:
		//如果ch1通道成功读取到数据，则执行该case处理语句
		fmt.Println("el=",e1)
	case e2:=<-ch2:
		//如果ch2通道成功读取到数据，则执行该case处理语句
		fmt.Println("e2=",e2)
	default:
		fmt.Println("default")
	}

	fmt.Println("next")
}

func getNum() int{
	time.Sleep(1000)
	return 1
}