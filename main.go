package main

import (
	"fmt"

	"github.com/demouth/learn-mcp-server/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Calculator and UUID Generator and Confluence search",
		"1.2.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Add the calculator tool and handler
	s.AddTool(tools.CalculatorTool, tools.CalculatorHandler)

	// Add the UUID generator tool and handler
	s.AddTool(tools.UUIDTool, tools.UUIDHandler)

	// Add the search tool and handler
	//
	// Note: The search tool requires a running Chrome instance with remote debugging enabled
	// You can start Chrome with the following command:
	// $ open -a 'Google Chrome' --args --remote-debugging-port=9222 --incognito
	namespace := "mynamespace" // Replace with your confluence namespace
	s.AddTool(tools.SearchConfluenceTool, tools.MakeHandler(namespace))

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
