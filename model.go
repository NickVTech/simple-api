package main

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []User{
	{Id: 100, Username: "NickVTech"},
	{Id: 101, Username: "Test1"},
	{Id: 102, Username: "Test2"},
}
