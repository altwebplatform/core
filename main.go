package main

import (
	"github.com/altwebplatform/core/models"
	"fmt"
	"time"
)

func main() {
	db := models.SetupDB("postgresql://root@localhost:26257/altwebplatform?sslmode=disable")
	defer db.Close()

	db.Create(&models.Service{
		Name: "test22",
		CreatedAt: time.Now(),
	})

	var service models.Service
	db.Find(&service)

	fmt.Println(service.Name)

	//_, err := k8s.CreateService("minio", 9000)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Created service!")
	//}
}
