package tools

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type SearchResult struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type ToolDefinition struct {
	Name        string
	Description string
}

var SearchConfluenceTool = mcp.NewTool(
	"search_confluence",
	mcp.WithDescription("Search Confluence for a keyword and return results with URLs."),
	mcp.WithString("keyword",
		mcp.Required(),
		mcp.Description("The search keyword"),
	),
)

func MakeHandler(namespace string) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		keyword := request.Params.Arguments["keyword"].(string)

		contents, url := getContents(namespace, keyword)

		type SearchResult struct {
			Contents string `json:"contents"`
			URL      string `json:"url"`
		}
		result := SearchResult{
			Contents: contents,
			URL:      url,
		}
		jsonResult, _ := json.Marshal(result)

		return mcp.NewToolResultText(string(jsonResult)), nil
	}
}

func getContents(namespace, keyword string) (body string, location string) {
	ws := "ws://localhost:9222"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	allocatorContext, cancel := chromedp.NewRemoteAllocator(ctx, ws)
	defer cancel()

	ctx, cancel = chromedp.NewContext(allocatorContext)
	defer cancel()

	keyword = url.QueryEscape(keyword)
	_ = chromedp.Run(ctx,
		chromedp.Navigate("https://"+namespace+".atlassian.net/wiki/search?text="+keyword),
		chromedp.WaitVisible(".searchResultLink"),
		chromedp.Click(".searchResultLink", chromedp.NodeVisible),
		chromedp.WaitVisible(".PageContent"),
		chromedp.Text(".PageContent", &body, chromedp.NodeVisible),
		chromedp.Location(&location),
	)
	return body, location
}
