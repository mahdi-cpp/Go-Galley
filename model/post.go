package model

type Post struct {
	Location string  `json:"location"`
	Caption  string  `json:"caption"`
	HasVideo bool    `json:"hasVideo"`
	Likes    uint    `json:"likes"`
	ChatId   uint    `json:"ChatId"`
	User     User    `gorm:"type:json;serializer:json" json:"user"`
	Medias   []Media `gorm:"type:json;serializer:json" json:"medias"`
}

type Media struct {
	Thumbnail     string `json:"thumbnail"`
	Photo         string `json:"photo"`
	Video         string `json:"video"`
	VideoDuration int    `json:"videoDuration"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
}
