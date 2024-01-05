// filecopy.go
package main

import (
	"fmt"
	"io"
	"os"
)

func maincopy() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

/*

linux系统下文件的权限由三位二进制数表示，linux下文件权限分为所属用户权限user、
所属组权限group和其他用户权限other。
每一个又分为读r、写w、可执行x权限。0表示没有权限，1表示有权限。4对应2进制的100，
就是可读不可写不可执行。最高是7二进制111,可读可写可执行，777代表最高权限，所有用户都有所有权限！八进制表示是0777
一个八进制相当于三个二进制
0666 0b110110110
0644 0b110100100
*/

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	// 0644 -> 110 100 100有读写权限 0644八进制二的三次方 是8一个八进制可以换算成三个二进制 777 111111111
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
