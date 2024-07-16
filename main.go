package main

import "github.com/mahdi-cpp/Go-Galley/repository"

func main() {
	//m := model.Music{Name: "name", Width: 58}
	//fmt.Println(m)
	repository.DatabaseInit()

	repository.UserStart()

}
