### 关键词

* [select](golang入门\go基础编程\select.go)

* [interface](golang入门\go基础编程\interface.go)

* [map](golang入门\go基础编程\map.go)
      
        引用类型

* [goto](golang入门\go基础编程\goto.go)

* [fallthrough](golang入门\go基础编程\fallthrough.go)  

        在执行一个case之后，默认地，控制立即从 switch语句中转移出来，如果不想跳出，而是继续执行下一个case 可以使用 fallthrough语句。

* [range](golang入门\go基础编程\range.go)

* [type]()  
    `定义类型:比如 type Personer interface`



### 内置函数

* [close]():关闭channel

* [len]():获取string,array,slice的长度，map key的数量，以及缓冲channel的可用数据数量

* [cap]():获取array的长度，slice的容量，以及缓冲channel的最大缓冲容量

* [new]():通常用于值类型，为指定类型分配初始化过的内存空间，返回指针。

* [make]():仅用于slice、map、channel这些引用类型，除了初始化内存还负责设置相关属性

* [append]():向slice追加(在其尾部添加)一个或多个元素

* [copy]():在不同slice间复制数据

* [print/println]():不支持format,要格式化输出，要使用fmt包

* [complex/real/imag]():复数处理

* [panic/recover]():错误处理


### 重点使用的地方

* [slice]()

* [map]()

* [struct]()

* [method]():结构体值传递还是指针传递

* [interface]()

