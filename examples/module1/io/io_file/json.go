// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}
type Status struct {
	Text string
}

type User struct {
	Status Status
}
type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

/*
JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map [string] T（T 是 json 包中支持的任何类型）
Channel，复杂类型和函数类型不能被编码
不支持循环数据结构；它将引起序列化进入一个无限循环
指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
*/
func mainjson() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	/*数据结构可以是任何类型，只要其实现了某种接口，
	目标或源数据要能够被编码就必须实现 io.Writer
	或 io.Reader 接口。由于 Go 语言中到处都实现了 Reader 和 Writer，
	因此 Encoder 和 Decoder 可被应用的场景非常广泛，例如读取或写入 HTTP 连接、websockets 或文件。
	*/
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
	jsonStr := `{"Status":{"Text":"zhowuude"}}`
	user := User{Status{""}}
	json.Unmarshal([]byte(jsonStr), &user)
	fmt.Println(user)

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	//
	err = json.Unmarshal(b, &f)
	// f 指向的值是一个 map，key 是一个字符串，value 是自身存储作为空接口类型的值：
	if err == nil {
		fmt.Println(f)                        //map[Age:6 Name:Wednesday Parents:[Gomez Morticia]]
		fmt.Println(reflect.TypeOf(f).Kind()) //map
	}
	// 接口断言
	// fmt.Println(f["Name"]) //编译报错
	if m, ok := f.(map[string]interface{}); ok {
		fmt.Println(m["Name"]) //Wednesday
		// v 还是 interface{}接口类型
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)

			case []interface{}:
				fmt.Println(k, "is an array:")
				// u还是 interace{}
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don’t know how to handle")
			}
		}
	}
	var m FamilyMember
	// 程序实际上是分配了一个新的切片。这是一个典型的反序列化引用类型（指针、切片和 map）的例子。
	err = json.Unmarshal(b, &m)

}
