package pushwoosh

type Service interface {
	RegisterDevice(device *Device) error
	UnregisterDevice(device *Device) error
	CreateMessage(message *Message) error
}
