package model_of_person

type PersonRequest struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}
type PersonResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}
