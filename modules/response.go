package modules

// ErrorModel ...
type ErrorModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// FileUploadedModel ...
type FileUploadedModel struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

type Response struct {
	StatusCode int        `json:"status_code"`
	Data       string     `json:"data"`
	SessionID  string     `json:"session_id"`
	Error      ErrorModel `json:"error"`
}
