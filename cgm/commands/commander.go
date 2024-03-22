package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

type Commander struct {
}

type Command struct {
	Command string      `json:"command"`
	Param   interface{} `json:"param"`
}

type ExecFunc func(param interface{}) error

type commandFunc struct {
	exec ExecFunc
	help string
}

var registry = make(map[string]*commandFunc)

func RegisterCommand(name string, exec ExecFunc, help string) {
	var cf = &commandFunc{exec: exec, help: help}
	registry[name] = cf
}

func (cmd *Commander) Run(param *Command) error {
	if fx, ok := registry[param.Command]; ok {
		return fx.exec(param.Param)
	}
	return fmt.Errorf("command %v not found", param.Command)
}

func (cmd *Commander) RunFile(cmfFile string) error {
	cfg := &Command{}
	jsonb, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonb, &cfg); err != nil {
		return err
	}
	return cmd.Run(cfg)

}
