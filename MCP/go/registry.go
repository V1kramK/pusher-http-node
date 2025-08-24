package main

import (
	"github.com/pusher-api/mcp-server/config"
	"github.com/pusher-api/mcp-server/models"
	tools_path "github.com/pusher-api/mcp-server/tools/path"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_path.CreateGet_path_channelsTool(cfg),
	}
}
