package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "root", "graphqldb",
		"--password=pass", "-e", "source seed.sql")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("err:%v", err)
		panic("panic!!")
	}
	fmt.Println("Seedが実行されました！")
}
