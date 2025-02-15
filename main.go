package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Cheolhwi/gator/internal/cli"
	"github.com/Cheolhwi/gator/internal/config"
)

func main() {
	// 读取配置文件
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// 创建全局状态
	state := &config.State{Config: &cfg}

	// 初始化命令注册系统
	commands := cli.NewCommands()

	// 注册 login 命令
	commands.Register("login", cli.HandlerLogin)

	// 解析命令行参数
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
	}

	// 提取命令和参数
	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:], // 剩余部分作为参数
	}

	// 运行命令
	if err := commands.Run(state, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
