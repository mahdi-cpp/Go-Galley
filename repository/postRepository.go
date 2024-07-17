package repository

import "github.com/mahdi-cpp/Go-Galley/model"

func PostStart() {
	err := DB.AutoMigrate(&model.Post{})
	if err != nil {
		return
	}

	//for i := 0; i < 10; i++ {
	//	CreatUPosts()
	//}
	UpdatePost(6)
}

func UpdatePost(ID uint) {
	//DB.Update("id", ID).
	DB.Model(&model.Post{}).Where("id = ?", ID).Update("location", "Malard")
}

func GetPosts() []model.Post {
	var posts []model.Post
	DB.Limit(50).Offset(10).Find(&posts)
	return posts
}

func CreatUPosts() {

	post1 := model.Post{
		Location: "Tehran",
		Caption:  "Mahdi",
		HasVideo: true,
		Likes:    32,
		ChatId:   456,
		User: model.User{
			Name:       "Mahdi",
			Email:      "mahdi.cpp@gmail.com",
			Phone:      "09355512326",
			IsVerified: true,
			Avatar:     "photos/1234_0000_avatar.jpg",
			Info: map[string]interface{}{
				"website":   "www.mahdi.abdolmaleki.ir",
				"biography": "nothing",
				"fullName":  "Mahdi abdolmaleki",
				"age":       30,
				"city":      "Tehran",
				"height":    170,
				"weight":    70,
			},
			Persians: map[string]interface{}{
				"type1":    "Verb",
				"type2":    "Noun",
				"persian1": []string{"mahdi", "ali", "reza"},
				"persian2": []string{"sara", "maryam", "leyla"},
			},
		},
		Medias: []model.Media{
			{Thumbnail: "photos/photo_125.jpg", Photo: "photos/thumbnail/photo_138.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217,
				Tags: []model.Tag{
					{Username: "@mahdi.45", Dx: 45, Dy: 100},
					{Username: "@sara.2008", Dx: 45, Dy: 100},
					{Username: "@maryam.2008", Dx: 345, Dy: 500},
				},
			},
			{Thumbnail: "photos/photo_226.jpg", Photo: "photos/thumbnail/photo_158.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217},
			{Thumbnail: "photos/photo_229.jpg", Photo: "photos/thumbnail/photo_178.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217},
			{Thumbnail: "photos/photo_328.jpg", Photo: "photos/thumbnail/photo_153.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217,
				Tags: []model.Tag{
					{Username: "@jac.45", Dx: 45, Dy: 100},
					{Username: "@messi.37", Dx: 45, Dy: 100},
					{Username: "@alvarez.55", Dx: 345, Dy: 500},
				},
			},
		},
	}

	DB.Create(&post1)
}

//func QueryPosts() []model.Post {

//var posts []model.Post
//DB.Find(&posts).Limit(10)
//return posts

//for _, u := range posts {
//var p = u.Persians
//fmt.Println(u.Name
//fmt.Println(p)

//persian1 := model.Persian{}
//err := PrettyJson(u.Persians["persian1"], &persian1)
//if err != nil {
//	fmt.Println(err)
//	c.JSON(400, err.Error())
//	return
//}

//}
