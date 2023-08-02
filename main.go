package main

import (
	"fmt"
	"github.com/leventtsang/geekgoclass1/generic"
	"math/rand"
)

func removeAt[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s // 索引越界，返回原始切片
	}
	copy(s[index:], s[index+1:])
	s = s[:len(s)-1]
	if cap(s) > 2*len(s) { // 如果容量是长度的两倍以上
		newSlice := make([]T, len(s))
		copy(newSlice, s)
		s = newSlice
	}
	return s
}

func main() {
	var index int
	fmt.Print("输入要删除的元素的索引（从0开始）：")
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println("请提供一个有效的整数作为索引")
		return
	}

	s := generic.GenerateSlice(10, func(r *rand.Rand) int { return r.Intn(100) })
	fmt.Println(removeAt(s, index))

	s2 := make([]string, len(s))
	for i := range s2 {
		s2[i] = string(rune('a' + i)) // 将整数转换为字母
	}
	fmt.Println(removeAt(s2, index))

}
