package tools

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// CalculatorTool defines the calculator tool
var CalculatorTool = mcp.NewTool("calculate",
	mcp.WithDescription("Perform basic arithmetic operations"),
	mcp.WithString("operation",
		mcp.Required(),
		mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
		mcp.Enum("add", "subtract", "multiply", "divide"),
	),
	mcp.WithNumber("x",
		mcp.Required(),
		mcp.Description("First number"),
	),
	mcp.WithNumber("y",
		mcp.Required(),
		mcp.Description("Second number"),
	),
)

// CalculatorHandler handles the calculator tool requests
func CalculatorHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	op := request.Params.Arguments["operation"].(string)
	x := request.Params.Arguments["x"].(float64)
	y := request.Params.Arguments["y"].(float64)

	var result float64
	switch op {
	case "add":
		result = x + y
	case "subtract":
		result = x - y
	case "multiply":
		result = x * y
	case "divide":
		if y == 0 {
			return nil, errors.New("Cannot divide by zero")
		}
		result = x / y
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
}
