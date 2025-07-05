package schema

type APIMessage struct {
	Message string `json:"message"`
}
type APIError struct {
	Message string `json:"message"`
}
type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Status  int         `json:"status"`
}
