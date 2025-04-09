package tools

import (
	"context"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
)

// UUIDTool defines the UUID generator tool
var UUIDTool = mcp.NewTool("generate_uuid",
	mcp.WithDescription("Generate a UUID"),
	mcp.WithString("version",
		mcp.Description("UUID version to generate (v4 is random)"),
		mcp.Enum("v4"),
	),
)

// UUIDHandler handles the UUID generator tool requests
func UUIDHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Currently only supporting v4 (random) UUIDs
	newUUID := uuid.New().String()
	return mcp.NewToolResultText(newUUID), nil
}
