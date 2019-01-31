package main

import (
	_ "airparking/order-bk/pkg/config"
	"airparking/order-bk/pkg/g"
	"airparking/order-bk/pkg/models"
)

func main() {
	defer models.Close()
	defer g.Logger.Sync()
}
