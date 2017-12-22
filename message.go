package pushwoosh

type Message struct {
	Application string `json:"application"`
	Auth string `json:"auth"`
	Notifications []*Notification `json:"notifications"`
}
