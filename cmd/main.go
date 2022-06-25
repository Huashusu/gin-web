package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("no executable file path")
		return
	}
	cmdPath := os.Args[1]
	params := make([]string, 0, 4)
	if len(os.Args) >= 2 {
		params = os.Args[2:]
	}
	for true {
		cmd := exec.Command(cmdPath, params...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			fmt.Printf("[%s] %s 参数:%+v", time.Now().Format("2006-01-02 15:04:05.000"), cmdPath, params)
			fmt.Printf(" 启动错误:%+v\n", err)
			return
		}
		fmt.Printf("[%s] 启动成功：PID:%d\n", time.Now().Format("2006-01-02 15:04:05.000"), cmd.Process.Pid)
		err = cmd.Wait()
		if err == nil {
			return
		}
		fmt.Printf("[%s] %s 参数:%+v", time.Now().Format("2006-01-02 15:04:05.000"), cmdPath, params)
		fmt.Printf(" 意外终止:%+v\n", err)
		interval := time.Second * 5
		fmt.Printf("[%s] %s后重启命令\n", time.Now().Format("2006-01-02 15:04:05.000"), interval)
		time.Sleep(interval)
	}
}
