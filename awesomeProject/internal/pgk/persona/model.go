package persona

type PersonRequest struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
	Work string `json:"work"`
}
type PersonResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
	Work string `json:"work"`
}