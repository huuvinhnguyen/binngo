// internal/delivery/http/api/api.go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huuvinhnguyen/binngo/internal/usecase"
)

type API interface {
	RegisterRoutes(r *gin.Engine)
}

func NewAPI(deviceUsecase usecase.DeviceUsecase) API {
	return &api{
		deviceUsecase: deviceUsecase,
	}
}

type api struct {
	deviceUsecase usecase.DeviceUsecase
}

func (a *api) RegisterRoutes(r *gin.Engine) {

	deviceHandler := NewDeviceHandler(a.deviceUsecase)

	r.GET("/api/devices", deviceHandler.GetDevices)

	// r.Run(":8080")
}
