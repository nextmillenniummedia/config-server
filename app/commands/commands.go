package commands

import (
	"config-server/app/errors"
	"config-server/app/types"

	"github.com/samber/do/v2"
)

func ProvideCommands(i do.Injector) (routes *Commands, err error) {
	commands := []types.Command{
		NewConfigAddCommand(),
	}
	commandMap := make(map[string]types.Command, len(commands))
	for _, command := range commands {
		commandMap[command.GetName()] = command
	}
	return &Commands{
		commandMap: commandMap,
	}, nil
}

type Commands struct {
	commandMap map[string]types.Command
}

func (c *Commands) Execute(name string, params []byte) (result any, err error) {
	command, has := c.commandMap[name]
	if !has {
		return nil, errors.CommandNotFound(name)
	}
	result, err = command.Execute(params)
	return
}
