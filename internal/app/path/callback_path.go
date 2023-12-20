package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	Domain       string
	Subdomain    string
	CallbackName string
	CallbackData string
}

var ErrUnknownCallback = errors.New("unknown callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "_", 3)
	if len(callbackParts) < 3 {
		return CallbackPath{}, ErrUnknownCallback
	}

	return CallbackPath{
		Domain:       callbackParts[0],
		Subdomain:    callbackParts[1],
		CallbackName: callbackParts[2],
	}, nil
}

func (p CallbackPath) String() string {
	return fmt.Sprintf("%s_%s_%s", p.Domain, p.Subdomain, p.CallbackName)
}
