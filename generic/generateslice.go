package generic

import (
	"math/rand"
	"sort"
	"time"
)

func GenerateSlice[T any](size int, generator func(*rand.Rand) T) []T {
	s := make([]T, size)
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range s {
		s[i] = generator(r)
	}
	return s
}

// 获取分片长度
func LengthSlice[T any](s []T) int {
	return len(s)
}

// 查找元素在分片中的位置，如果元素不存在则返回 -1
// 提供一个比较函数，用于判断元素是否相等
func FindSlice[T any](s []T, elem T, equals func(T, T) bool) int {
	for i, v := range s {
		if equals(v, elem) {
			return i
		}
	}
	return -1
}

// 过滤分片，返回一个新的分片，新的分片中的元素满足 provided condition
func FilterSlice[T any](s []T, condition func(T) bool) []T {
	var result []T
	for _, v := range s {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// 对分片进行映射，返回一个新的分片，新的分片中的元素是原始分片中元素经过函数 f 处理后的结果
func MapSlice[T any, R any](s []T, f func(T) R) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// 对分片进行排序
func SortSlice[T any](s []T, less func(T, T) bool) {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
}

// 删除分片中指定位置的元素
func RemoveAt[T any](s []T, index int) []T {
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
