package main

import (
	"context"
	"fmt"

	"github.com/hassaku63/mcp-start-go/internal/tool/greeting"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	done := make(chan struct{})

	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
		server.WithResourceCapabilities(false, false),
	)

	tool := mcp.NewTool("greeting",
		mcp.WithDescription("Greet a person with the current date"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	// Add tool handler
	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, ok := request.Params.Arguments["name"].(string)
		if !ok {
			return nil, fmt.Errorf("name is not a string")
		}

		message, err := greeting.GreetingToolHandler(greeting.GreetingArguments{Name: name})
		if err != nil {
			return nil, fmt.Errorf("failed to greet: %w", err)
		}
		return mcp.NewToolResultText(message), nil
	})

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	<-done
}
