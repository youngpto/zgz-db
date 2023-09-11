package base

import "fmt"

var BaseCommand = &CommandManager{
	Index: make(map[string]Command),
	Order: make([]Command, 0),
}

type CommandManager struct {
	Index map[string]Command
	Order []Command
}

func (c *CommandManager) Register(command Command) {
	_, ok := c.Index[command.Name()]
	if ok {
		fmt.Println(command.Name(), "已被注册")
		return
	}
	c.Index[command.Name()] = command
	for _, alias := range command.Alias() {
		_, ok := c.Index[alias]
		if ok {
			fmt.Println(alias, "别名已被使用")
			continue
		}
		c.Index[alias] = command
	}
	c.Order = append(c.Order, command)
}

type Command interface {
	Name() string
	Alias() []string
	Description() string
	Param() []CommandParam
}

type CommandParam struct {
	Name    string
	Type    ParamTypeEnum
	Require bool
}

type ParamTypeEnum int

const (
	Int32 ParamTypeEnum = iota
	Int64
	String
)
