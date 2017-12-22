package pushwoosh

type Device struct {
	Application string `json:"application"`
	PushToken string `json:"push_token"`
	Language string `json:"language"`
	Hwid string `json:"hwid"`
	Timezone int `json:"timezone"`
	DeviceType DeviceType `json:"device_type"`
}
