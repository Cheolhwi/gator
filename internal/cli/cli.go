package cli

import (
	"errors"
	"fmt"

	"github.com/Cheolhwi/gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	handlers map[string]func(*config.State, Command) error
}

var ErrInvalidCommand = errors.New("invalid command or missing arguments")

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	username := cmd.Args[0]
	err := s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User set to: %s\n", username)
	return nil
}

// 创建新的 Commands 结构体
func NewCommands() *Commands {
	return &Commands{handlers: make(map[string]func(*config.State, Command) error)}
}

// 注册一个新命令
func (c *Commands) Register(name string, handler func(*config.State, Command) error) {
	c.handlers[name] = handler
}

// 运行一个命令
func (c *Commands) Run(s *config.State, cmd Command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return ErrInvalidCommand
	}
	return handler(s, cmd)
}
