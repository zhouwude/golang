package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "testing"
	fmt.Printf("%d\n", name)
	fmt.Printf("%s\n", name, name)
	// 3.4是索引
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(len(arrKeyValue[0]))

	fmt.Println([]int{1, 2, 3, 4, 5}[:])

	array := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := &array
	fmt.Println("slice", slice)
	fmt.Println("*slice", *slice) //*取内容符号 &取地址符号。
	fmt.Println(array)
	fmt.Printf("adreess %p \n", slice)
	slice1 := slice[:]
	fmt.Printf("adreess %p \n", slice1)
	// adreess 0xc00012c060
	// adreess 0xc00012c060
	fmt.Println(cap(slice))
	fmt.Println(cap(slice1))
	fmt.Println(len(slice[1:4])) //长度
	fmt.Println(cap(slice[1:4])) //容量
	fmt.Println(len(array[1:4])) //长度
	fmt.Println(cap(array[1:4])) //容量
	slice2 := array[2:5]
	fmt.Println(slice2)
	slice2 = slice2[0:6]
	// slice2 包含一个指向数组的指针容量 》=len
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	// 容量会少一位
	slice2 = slice2[1:]
	fmt.Println(cap(slice2))
	fmt.Printf("array %T \n", array)
	fmt.Printf("slice %T \n", slice)
	fmt.Printf("slice1 %T \n ", slice1)
	fmt.Printf("slice2 %T \n", slice2)
	fmt.Println(reflect.TypeOf(array).Kind())
	fmt.Println(reflect.TypeOf(slice).Kind()) //ptr 指针
	fmt.Println(reflect.TypeOf(slice2).Kind())
	// 带了个数的就是数组 不带的就是切片
	fmt.Println(reflect.TypeOf([1]float32{2.0}).Kind())
	fmt.Println(reflect.TypeOf([]float32{2.0}).Kind())
	fmt.Println(make([]int, 50, 100))
	fmt.Println("sjsjdj"[:3])
	fmt.Println(reflect.TypeOf("sjsjdj"[:3]).Kind())
	fmt.Println(reflect.TypeOf(new([10]int)).Kind()) //指针类型
	fmt.Println(new([10]int))
	items := []int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 10 //item 只是一个拷贝
	}
	fmt.Println(items)
	// 只需要下标参数来修改内部的值。
	for index := range items {
		// 吐过要改变切片或者 map 得值的话，不能直接操作元素 这里的方法是用下标取出来再复制。
		items[index] *= items[index]
	}
	fmt.Println(items)
	fmt.Println(make([]string, 20))

	dst := []byte{'s', 'b', 'c', 'd'} //字符串是一个切片 元素为 byte
	copy(dst[2:], []byte{'a', 'a'})
	fmt.Printf("-- %s", dst)
	str := "周武德"
	for i1 := 0; i1 < len(str); i1++ {
		fmt.Printf("%c \n", str[i1])
	}
	for _, value := range str {
		fmt.Printf("%c \n", value)
	}
	sliceStr := []byte(str)
	fmt.Printf("%s \n", sliceStr)
	fmt.Println(reflect.TypeOf(sliceStr).Kind(), len(sliceStr), cap(sliceStr))
	fmt.Println(len("周武德"))
	fmt.Println(len([]int32(str)))
	fmt.Println(sliceStr[0])
	// 获取数量
	fmt.Println(utf8.RuneCountInString(str))
	slice4 := append(sliceStr, "ddddd"...)
	fmt.Printf("%s \n", slice4)
	fmt.Printf("%d", len(slice4))
	// byte 字符必须使用单引号 ''
	fmt.Println(strings.Replace(str, "周", "梅", -1))
	slice6 := []int{1, 3, 6, 3, 5, 4, 9, 6}
	sort.Ints(slice6)
	fmt.Println(slice6)
	fmt.Println(sort.IntsAreSorted(slice6))
	fmt.Println(sort.SearchInts(slice6, 9)) //7
	// fmt.Println(sort.SearchStrings("周五得策空间范围", "周"))
	fmt.Println(strings.Index("小画册", "册"))
	fmt.Println(strings.IndexRune("小画册", '册'))

	//删除某个位置元素
	fmt.Printf("%v", append(slice6[:4], slice6[5:]...))
	fmt.Println(slice6[len(slice6)-1])
	fmt.Println(slice6[len(slice6)-1:])
	map1 := make(map[string]string)
	map1["djj"] = "s"
	map1["d"] = "v"
	fmt.Println(map1)
	value, have := map1["j"]
	if have {
		fmt.Println(value)
	} else {
		fmt.Println("not element")
	}
	// 合并语句
	if _, have := map1["j"]; have {
		fmt.Println(value)
	}
}
