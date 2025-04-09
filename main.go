package main

import (
	"fmt"

	"github.com/demouth/learn-mcp-server/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Calculator and UUID Generator Demo",
		"1.1.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Add the calculator tool and handler
	s.AddTool(tools.CalculatorTool, tools.CalculatorHandler)

	// Add the UUID generator tool and handler
	s.AddTool(tools.UUIDTool, tools.UUIDHandler)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
