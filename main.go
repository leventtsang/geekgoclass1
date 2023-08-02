package main

import (
	"fmt"
	"github.com/leventtsang/geekgoclass1/generic"
	"math/rand"
)

func main() {
	var index int
	fmt.Print("输入要删除的元素的索引（从0开始）：")
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println("请提供一个有效的整数作为索引")
		return
	}

	s := generic.GenerateSlice(10, func(r *rand.Rand) int { return r.Intn(100) })
	fmt.Println("原始整数切片：", s)
	fmt.Println("删除后的整数切片：", generic.RemoveAt(s, index))

	s2 := make([]string, 10)
	for i := range s2 {
		s2[i] = string(rune('a' + i)) // 将整数转换为字母
	}
	fmt.Println("原始字符串切片：", s2)
	fmt.Println("删除后的字符串切片：", generic.RemoveAt(s2, index))
}
