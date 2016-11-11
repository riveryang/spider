package result

type Result struct {
	Status uint `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`

}