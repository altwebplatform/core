package awp

import (
	"flag"
	"github.com/altwebplatform/core/config"
	"github.com/altwebplatform/core/storage"
	"github.com/altwebplatform/core/web"
)

func Start() {
	config.DB_CONNECT = *flag.String("db", config.DB_CONNECT, "database connect string")
	storage.Init(config.DB_CONNECT)
	defer storage.Close()
	web.Start(*flag.String("listen", ":8080", "which address to start web dashboard"))
}
