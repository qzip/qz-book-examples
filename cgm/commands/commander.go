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

type Commands struct {
	name string
	exec ExecFunc
}

func (cmd *Commander) Run(param *Command) error {
	if fx, ok := Commands[param.Command]; ok {
		return fx(param.Param)
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
