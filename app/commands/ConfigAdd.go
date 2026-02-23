package commands

import "config-server/app/types"

type ConfigAddParams struct{}
type ConfigAddResult struct{}

func NewConfigAddCommand() types.Command {
	return &ConfigAddCommand{}
}

type ConfigAddCommand struct{}

func (c *ConfigAddCommand) GetName() string {
	return "ConfigAdd"
}

func (c *ConfigAddCommand) Execute(params []byte) (result any, err error) {
	return ConfigAddResult{}, nil
}
