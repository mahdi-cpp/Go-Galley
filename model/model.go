package model

type Music struct {
	Name   string  `json:"name"`
	Artist string  `json:"artist"`
	Ali    string  `json:"ali"`
	Width  float32 `json:"width"`
}

type Persian struct {
	Persian1 []string `json:"persian1"`
	Persian2 []string `json:"persian2"`
	Type1    string   `json:"type1"`
	Type2    string   `json:"type2"`
}
