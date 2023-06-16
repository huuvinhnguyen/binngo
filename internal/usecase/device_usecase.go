// internal/usecase/device/device_usecase.go
package device

import (
	"github.com/huuvinhnguyen/binngo/internal/entity"
	"github.com/huuvinhnguyen/binngo/internal/repository/device"
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

func (u *deviceUsecase) GetDevices() ([]entity.Device, error) {
	devices, err := u.deviceRepo.GetDevices()
	if err != nil {
		return nil, err
	}

	return devices, nil
}
