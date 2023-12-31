// internal/delivery/http/api/device_handler.go
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huuvinhnguyen/binngo/internal/usecase"
  // "github.com/huuvinhnguyen/binngo/internal/repository/device"
)

type DeviceHandler struct {
	deviceUsecase usecase.DeviceUsecase
}

func NewDeviceHandler(deviceUsecase usecase.DeviceUsecase) *DeviceHandler {
	return &DeviceHandler{
		deviceUsecase: deviceUsecase,
	}
}

func (h *DeviceHandler) GetDevices(c *gin.Context) {
	devices, err := h.deviceUsecase.GetDevices()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get devices"})
		return
	}

	c.JSON(200, devices)
}

