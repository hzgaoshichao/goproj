package main

import (
	"fmt"
	"singleton/singleton"
)

func main() {
	instance := singleton.GetInstance()

	fmt.Println(instance)
}

// 参考链接
// https://www.liwenzhou.com/posts/Go/singleton/

// 关于原子操作
// https://blog.csdn.net/chai2010/article/details/120426568
// https://stackoverflow.com/questions/55840399/golang-what-is-atomic-read-used-for
