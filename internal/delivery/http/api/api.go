// internal/delivery/http/api/api.go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huuvinhnguyen/binngo/internal/usecase/device"
)

type API interface {
	RegisterRoutes()
}

func NewAPI(deviceUsecase device.DeviceUsecase) API {
	return &api{
		deviceUsecase: deviceUsecase,
	}
}

type api struct {
	deviceUsecase device.DeviceUsecase
}

func (a *api) RegisterRoutes() {
	r := gin.Default()

	deviceHandler := NewDeviceHandler(a.deviceUsecase)

	r.GET("/api/devices", deviceHandler.GetDevices)

	r.Run(":8080")
}
