package main

import "fmt"

func main() {
	//三要素指向底层数组的指针 长度 和容量
	s := make([]int, 20) //默认类型零值
	fmt.Println(s)
	s1 := s[:4]
	fmt.Println(s1) //截取 切片片段
	fmt.Println(cap(s), cap(s1))
	// 修改s1数据 s原切片的内容同样会修改
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s1[3] = 4
	fmt.Println(s1)
	fmt.Println(s)
	fmt.Printf("/n---%p---%p", s, s1)
	// 往 s1 里面加入数据
	/*[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
			[0 0 0 0]
			20 20
			[1 2 3 4]
		[1 2 3 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	/n---0xc000106000---0xc000106000 打印地址可以发现 s s1指向同一个地址
	*/
	fmt.Println()
	s2 := make([]int, len(s))
	copy(s2, s)
	fmt.Println(s2)
	fmt.Printf("/n--%p-%p \n", s2, s1)
	s2[1] = 100
	fmt.Println(s)
	fmt.Println(s2)
	// 复制了 s 切片之后地址不一样了 而且修改元素不影响 s切片 所以只有 copy 之后才会复制指针指向另外一块内存
	// ************切片重组是不会复制指针的-----
	/*[1 2 3 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		/n--0xc00012a0a0-0xc00012a000
		[1 2 3 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	[1 100 3 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]*/
	fmt.Println(make([]int, 2)) //[0 0]

	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		// season 只是 seasons 某个索引位置的值的一个拷贝，不能用来修改 seasons 该索引位置的值。
		fmt.Printf("%s\n", season)
	}
	// 修改切片得值只能这样
	for ix := range seasons {
		// seasons[ix] = newValue for ;;也一样
		fmt.Printf("%d", ix)
	}
}
