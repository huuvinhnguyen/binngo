// internal/usecase/device/device.go
package usecase

import (
  "github.com/huuvinhnguyen/binngo/internal/repository/device"
  "github.com/huuvinhnguyen/binngo/internal/entity"
)
type DeviceUsecase interface {
	GetDevices() ([]entity.Device, error)
}

type deviceUsecase struct {
	deviceRepo device.DeviceRepository
}

func NewDeviceUsecase(deviceRepo device.DeviceRepository) DeviceUsecase {
	return &deviceUsecase{
		deviceRepo: deviceRepo,
	}
}

func (uc *deviceUsecase) GetDevices() ([]entity.Device, error) {
	return uc.deviceRepo.GetDevices()
}
