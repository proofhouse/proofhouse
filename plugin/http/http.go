package http

import (
	"database/sql"
	"plugin"
)

type Plugin struct {
	plugin.Plugin
}

func init() {
	var plugin := Plugin{}

	sql.Register("http", plugin)
}