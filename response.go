package playrlo

type CompileResponse struct {
	Success bool   `json:"success,omitempty"`
	Code    string `json:"code,omitempty"`
	Stdout  string `json:"stdout,omitempty"`
	Stderr  string `json:"stderr,omitempty"`
}

type rawCompileResponse struct {
	Success bool   `json:"success,omitempty"`
	Code    string `json:"code,omitempty"`
	Stdout  string `json:"stdout,omitempty"`
	Stderr  string `json:"stderr,omitempty"`
	Error   string `json:"error,omitempty"`
}

type FormatResponse struct {
	Success bool   `json:"success,omitempty"`
	Code    string `json:"code,omitempty"`
	Stdout  string `json:"stdout,omitempty"`
	Stderr  string `json:"stderr,omitempty"`
}

type rawFormatResponse struct {
	Success bool   `json:"success,omitempty"`
	Code    string `json:"code,omitempty"`
	Stdout  string `json:"stdout,omitempty"`
	Stderr  string `json:"stderr,omitempty"`
	Error   string `json:"error,omitempty"`
}
