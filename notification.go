package pushwoosh

type Notification struct {
	SendDate string `json:"send_date"`
	Content string `json:"content,omitempty"`
	Devices []string `json:"devices"`
	IosRootParams `json:"ios_root_params"`
	IosSound string `json:"ios_sound,omitempty"`
	IosSilent SilentSetting `json:"ios_silent,omitempty"`
	Platforms []int `json:"platforms"`
}

type IosRootParams struct {
	Aps Aps `json:"aps,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type Aps struct {
	ContentAvailable ContentAvailability `json:"content-available,omitempty"`
}