package main

import (
	"log"
	// "web/controller"

	"github.com/gin-gonic/gin"

  "github.com/huuvinhnguyen/binngo/internal/delivery/http/api"
	"github.com/huuvinhnguyen/binngo/internal/repository/device"
	"github.com/huuvinhnguyen/binngo/internal/usecase/device"
)

import "github.com/huuvinhnguyen/binngo/web/controller"


func main() {

  deviceRepo := device.NewDeviceRepository()
	deviceUsecase := device.NewDeviceUsecase(deviceRepo)

	api := api.NewAPI(deviceUsecase)
	api.RegisterRoutes()


	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	r.LoadHTMLGlob("web/templates/**/*")
	controller.Router(r)

	log.Println("Server started")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
