package main

import (
	"github.com/giphy-api/mcp-server/config"
	"github.com/giphy-api/mcp-server/models"
	tools_gifs "github.com/giphy-api/mcp-server/tools/gifs"
	tools_stickers "github.com/giphy-api/mcp-server/tools/stickers"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_gifs.CreateTranslategifTool(cfg),
		tools_gifs.CreateGetgifbyidTool(cfg),
		tools_stickers.CreateTrendingstickersTool(cfg),
		tools_gifs.CreateRandomgifTool(cfg),
		tools_stickers.CreateTranslatestickerTool(cfg),
		tools_stickers.CreateRandomstickerTool(cfg),
		tools_gifs.CreateGetgifsbyidTool(cfg),
		tools_gifs.CreateSearchgifsTool(cfg),
		tools_gifs.CreateTrendinggifsTool(cfg),
		tools_stickers.CreateSearchstickersTool(cfg),
	}
}
