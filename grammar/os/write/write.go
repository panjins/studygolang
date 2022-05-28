package main

import (
	"fmt"
	"os"
)

func main() {
	WriteString()
	WriteByLine()

}

//字符串写入

func WriteString() {
	f, err := os.Create("grammar/os/write/string.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	mess, err := f.WriteString("Hello String")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(mess, "written successful")
	err = f.Close()
	if err != nil {
		panic(err)
	}

}

// 字符串一行一行写入

func WriteByLine() {
	f, err := os.Create("grammar/os/write/lines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	mess := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}

	for _, v := range mess {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successful")

}
