package pushwoosh

type Error string

const (
	// ErrHwidRequired occurs when the HWID is not provided.
	ErrHwidRequired = Error("hwid_required")
)

func (e Error) Error() string {
	return string(e)
}
