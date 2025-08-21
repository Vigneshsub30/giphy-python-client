package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/giphy-api/mcp-server/config"
	"github.com/giphy-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetgifbyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		gifIdVal, ok := args["gifId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: gifId"), nil
		}
		gifId, ok := gifIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: gifId"), nil
		}
		queryParams := make([]string, 0)
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/gifs/%s%s", cfg.BaseURL, gifId, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			// API key already added to query string
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetgifbyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_gifs_gifId",
		mcp.WithDescription("Get GIF by Id"),
		mcp.WithNumber("gifId", mcp.Required(), mcp.Description("Filters results by specified GIF ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetgifbyidHandler(cfg),
	}
}
