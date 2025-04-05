package greeting

import (
	"fmt"
	"time"
)

const (
	GREETING_TOOL_NAME        = "greeting"
	GREETING_TOOL_DESCRIPTION = "Greet a person with the current date"
)

type ToolServer interface {
	RegisterTool(name string, description string, handler interface{}) error
}

type GreetingArguments struct {
	Name string `json:"name" jsonschema:"required,description=The name of the person to greet"`
}

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome! Today's date is %v.", name, time.Now().Format("2006-01-02"))
	return message
}

func GreetingToolHandler(args GreetingArguments) (string, error) {
	message := Hello(args.Name)
	return message, nil
}
