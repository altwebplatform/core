package awp

import (
	"github.com/altwebplatform/core/web"
	"flag"
	"github.com/altwebplatform/core/db"
)

func Start() {
	db.Init(*flag.String("db", "postgresql://root@localhost:26257/altwebplatform?sslmode=disable", "a string"))
	web.Start(*flag.String("listen", ":8080", "a string"))
}