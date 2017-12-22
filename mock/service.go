package mock

import "github.com/swensonhe/pushwoosh-go"

type Service struct {
	RegisterDeviceFn func (*pushwoosh.Device) error
	RegisterDeviceInvoked bool

	UnregisterDeviceFn func (*pushwoosh.Device) error
	UnregisterDeviceInvoked bool

	CreateMessageFn func(*pushwoosh.Message) error
	CreateMessageInvoked bool
}

func (s *Service) RegisterDevice(device *pushwoosh.Device) error {
	s.RegisterDeviceInvoked = true
	return s.RegisterDeviceFn(device)
}

func (s *Service) UnregisterDevice(device *pushwoosh.Device) error {
	s.UnregisterDeviceInvoked = true
	return s.UnregisterDeviceFn(device)
}

func (s *Service) CreateMessage(message *pushwoosh.Message) error {
	s.CreateMessageInvoked = true
	return s.CreateMessageFn(message)
}
