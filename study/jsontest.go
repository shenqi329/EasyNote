package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float32
}

func main() {
	gobook := Book{Title: "Go语言编程",
		Authors:     []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		Publisher:   "ituring.com.cn",
		IsPublished: true,
	}

	b, err := json.Marshal(gobook)

	if err != nil {

	}

	str := string[b]
	fmt.Println(str)

	stb := &Book{}
	err = json.Unmarshal([]byte(b), &stb)

	fmt.Println(stb.Authors)
}

//
//Price:9.99
