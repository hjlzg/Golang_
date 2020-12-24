package main

import (
	"fmt"
)

type ValNode struct{
	row int
	col int
	val int
}

func main(){
	//1.创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2]=1
	chessMap[2][3]=2

	//2. 输出原始数组
	for _,v:=range chessMap{
		for _,v2:=range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	//3. 转成稀疏数组
	//思路
	//(1)遍历 chessmap,如果我们发现有一个元素的值不为0，创建一个node结构体
	//(2) 将其放入到对应的切片即可

	var spareArr []ValNode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组的规模（行和列，默认值）
	valNode:=ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	spareArr=append(spareArr,valNode)

	for i,v:=range chessMap{
		for j,v2:=range v{
			if v2!=0{
				//创建一个节点
				valNode:=ValNode{
					row: i,
					col:j,
					val: v2,
				}
				spareArr=append(spareArr,valNode)
			}
		}
	}

	//输出稀疏数组
	for i,valueNode:=range spareArr{
		fmt.Printf("%d: %d %d %d ",i,valueNode.row,valueNode.col,valueNode.val)
		fmt.Println()
	}

	//将稀疏数组存盘

	//恢复原始数组

	var chessMap2 [11][11]int

	//遍历 spareArr
	for i,valeNode:=range spareArr{
		if i!=0{
			chessMap2[valeNode.row][valeNode.col]=valNode.val
		}
	}

	//输出原始数组
	for _,v:=range chessMap2{
		for _,v2:=range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
}

