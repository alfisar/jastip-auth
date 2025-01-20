package domain

type ErrorData struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	HTTPCode int         `json:"-"`
	Message  string      `json:"message"`
	Errors   interface{} `json:"errors,omitempty"`
}
