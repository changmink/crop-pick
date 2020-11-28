package model

type CropInfo struct {
	Name  string `json:"name"`
	Area  string `json:"area"`
	Count string `json:"count"`
}

type CropCount struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}
