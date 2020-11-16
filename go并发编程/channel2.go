package main

import (
	"fmt"
	"time"
)

func main() {
	tick := make(chan int, 1)
	go func() {
		time.Sleep(2 * time.Second)
		count := 1
		for {
			time.Sleep(1 * time.Second)
			tick <- count
			count++
			if count == 5 {
				//当通道被关闭时，迭代器的执行会自动结束
				close(tick)
			}
		}
	}()

	//可以用迭代器从通道中取数据。tick是非缓冲通道，所以只有当通道
	//有数据时，这里才会继续执行
	for v := range tick {
		fmt.Printf("%d\n", v)
		// if v == 5 {
		// 	break
		// }
	}
}
