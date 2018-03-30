package pushwoosh

type Notification struct {
	SendDate string `json:"send_date"`
	Content string `json:"content,omitempty"`
	Devices []string `json:"devices"`
	IosRootParams IosRootParams `json:"ios_root_params"`
	IosSound string `json:"ios_sound,omitempty"`
	IosSilent SilentSetting `json:"ios_silent,omitempty"`
	AndroidRootParams AndroidRootParams `json:"android_root_params,omitempty"`
	AndroidSilent SilentSetting `json:"android_slient",omitempty`
	Platforms []int `json:"platforms"`
}

type IosRootParams struct {
	Aps Aps `json:"aps,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type AndroidRootParams map[string]interface{}

type Aps struct {
	ContentAvailable ContentAvailability `json:"content-available,omitempty"`
}