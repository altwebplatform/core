package awp

import "github.com/altwebplatform/core/web"

func Start() {
	//db := models.SetupDB("postgresql://root@localhost:26257/altwebplatform?sslmode=disable")
	//defer db.Close()

	web.Start(":8080")

}