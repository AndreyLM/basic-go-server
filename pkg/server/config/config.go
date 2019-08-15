package config

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/server/modules"
	"github.com/andreylm/basic-go-server.git/pkg/server/modules/v1/chat"
	"github.com/andreylm/basic-go-server.git/pkg/server/modules/v1/users"
)

// ServerConfigurations - server configurations
type ServerConfigurations struct {
	DbConfig   db.ConnectionConfig
	Port       string
	Driver     string
	KeyPublic  string
	KeyPrivate string
}

// GetModules - gets modules
func GetModules() []modules.Module {
	return []modules.Module{
		&users.Module{},
		&chat.Module{},
	}
}
