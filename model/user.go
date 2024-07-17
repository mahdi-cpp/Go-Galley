package model

type User struct {
	ID         uint                   `json:"id"`
	Name       string                 `json:"name"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	Avatar     string                 `json:"avatar"`
	IsVerified bool                   `json:"IsVerified"`
	Info       map[string]interface{} `gorm:"type:json;serializer:json" json:"info"`
	Persians   map[string]interface{} `gorm:"type:json;serializer:json" json:"Persians"`
}
