package playrlo

type FormatRequest struct {
	Code    string  `json:"code,omitempty"`
	Edition Edition `json:"edition,omitempty"`
}
