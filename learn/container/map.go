package main

import (
	"fmt"
	"sort"
)

func mapTest() {
	var map1 map[int]string
	map2 := make(map[string]float32, 100)
	// 不要使用 new，永远用 make 来构造 map
	fmt.Println(map1)
	fmt.Println(map2)
	// 一个 key 要对应多个值时
	// 例如：处理unix机器上的所有进程，以父进程（pid 为整型）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value
	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)
	fmt.Println(mp1)
	fmt.Println(mp2)
}

func test2() {
	var value int
	var ok bool

	map1 := make(map[string]int)
	map1["Xi`an"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, ok = map1["Beijing"] // 键值对是否存在
	if ok {
		fmt.Println(value) // 满足条件时有value值
	} else {
		fmt.Println(ok) // 不满足条件时为false
	}

	// 删除某元素
	delete(map1, "Xi`an")
	fmt.Println(map1) // map[Beijing:20 Washington:25]
}

func forRange() {
	map1 := make(map[string]int)
	map1["Xian"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25

	for key, value := range map1 {
		fmt.Printf("key 是%s，value是%d\n", key, value)
	}

	// map的本质是散列表，map的增长扩容会导致重新进行散列，使遍历结果在扩容前后不同
	map1["China"] = 100
	for key, value := range map1 { // 遍历顺序与之前不同
		fmt.Printf("key 是%s，value是%d\n", key, value)
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

func map_slice() {
	// 想获取一个 map 类型的切片，必须使用两次 make() 函数，1 分配切片，2 分配切片中每个 map 元素
	items := make([]map[string]int, 5) // 创建切片
	for i := range items {
		items[i] = make(map[string]int, 1) // 每一项创建map
		items[i]["h"] = 2
	}
	fmt.Println(items) // map[h:2] map[h:2] map[h:2] map[h:2] map[h:2]]
}

func map_sort() {
	// map默认无序，需要排序时：先将 key（或 value）拷贝到一个切片，再对切片排序
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)
	fmt.Println("排序前！！！")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}
	fmt.Println("排序后！！！")
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v \n ", k, barVal[k])
	}
}

func map_reverse() {
	// 将map的键值对调：即 key与value
	// 前提是value是唯一的
	myMap := map[string]int{"h": 1, "e": 2, "l": 3, "o": 4}
	fmt.Println("kv反转之前！！")
	for k, v := range myMap {
		fmt.Printf("%s : %d\n", k, v)
	}
	// 开始对调
	myMap1 := make(map[int]string, len(myMap))
	for k, v := range myMap {
		myMap1[v] = k
	}
	fmt.Println("\nkv反转之后！！")
	for k, v := range myMap1 {
		fmt.Printf("%v : %v\n", k, v)
	}
}
func main() {
	// mapTest()
	// test2()
	// forRange()
	// map_slice()
	// map_sort()
	map_reverse()
}
