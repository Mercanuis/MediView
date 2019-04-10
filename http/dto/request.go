package dto

type PatientAddRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
