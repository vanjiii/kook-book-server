package httpx

type APIResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
