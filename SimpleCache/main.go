package main

import (
	"MoreTask/SimpleCache/cache"
	_ "MoreTask/SimpleCache/cache"
	"fmt"
)

func main() {
	ca := cache.NewMemoryCache()
	ca.SetMaxMemory("100m")
	ca.Set("aaa", "aaa123123", 10)
	fmt.Println(ca.Get("aaa"))
	ca.Set("aaab", [3]int{1, 2, 3}, 100)
	fmt.Println(ca.Get("aaab"))

	// 添加定时器
	//timer := time.NewTimer(11 * time.Second)
	//select {
	//case <-timer.C:
	//	fmt.Println("aaa: ", ca.Get("aaa"))
	//}

	fmt.Println(ca.Get("aaa"))
	fmt.Println(ca.Exist("aaa"))
	fmt.Println(ca.Exist("aaab"))
}
