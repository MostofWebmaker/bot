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
	commandParts := strings.SplitN(commandText, "__", 3)

	//if len(commandParts) == 1 {
	//	return CommandPath{
	//		CommandName: commandParts[0],
	//		Domain:      'domain',
	//		Subdomain:   'subdomain',
	//	}, nil
	//}

	len := len(commandParts)

	if len == 0 {
		return CommandPath{}, ErrUnknownCommand
	}

	domain := "demo"
	subdomain := "subdomain"

	if len == 3 {
		domain = commandParts[1]
		subdomain = commandParts[2]
	}

	//domain := "domain"
	//subdomain := "subdomain"
	//
	//domain, ok := commandParts[1]
	//if !ok {
	//	domain = "domain"
	//}
	//
	//subdomain, ok2 := commandParts[2]
	//if !ok2 {
	//	subdomain = "subdomain"
	//}

	//if commandParts[1] {
	//	domain := commandParts[1]
	//} else {
	//	domain := "domain"
	//}

	return CommandPath{
		CommandName: commandParts[0],
		Domain:      domain,
		Subdomain:   subdomain,
	}, nil

	//return CommandPath{
	//	CommandName: commandParts[0],
	//	Domain:      commandParts[1],
	//	Subdomain:   commandParts[2],
	//}, nil
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
