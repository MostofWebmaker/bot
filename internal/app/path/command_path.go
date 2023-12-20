package path

import (
	"errors"
	"fmt"
	"strings"
)

type CommandPath struct {
	CommandName string
	Domain      string
	Subdomain   string
}

var ErrUnknownCommand = errors.New("unknown command")

func ParseCommand(commandText string) (CommandPath, error) {
	commandParts := strings.SplitN(commandText, "_", 3)

	l := len(commandParts)
	if l == 0 {
		return CommandPath{}, ErrUnknownCommand
	}

	domain := "demo"
	subdomain := "subdomain"
	commandName := "start"

	if l >= 3 {
		domain = commandParts[0]
		subdomain = commandParts[1]
		commandName = commandParts[2]
	}

	return CommandPath{
		CommandName: commandName,
		Domain:      domain,
		Subdomain:   subdomain,
	}, nil
}

func (c CommandPath) IsSimpleCommandPath() bool {
	return c.CommandName != "" && c.Domain == "" && c.Subdomain == ""
}

func (c CommandPath) WithCommandName(commandName string) CommandPath {
	c.CommandName = commandName

	return c
}

func (c CommandPath) String() string {
	return fmt.Sprintf("/%s__%s__%s", c.CommandName, c.Domain, c.Subdomain)
}
