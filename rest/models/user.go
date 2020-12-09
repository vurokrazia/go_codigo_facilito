package models

import (
	"errors"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type Users []User

var users = make(map[int]User)

func SetDeafaultUser() {
	user := User{Id: 1, Email: "jesus", Password: "123456"}
	users[user.Id] = user
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("User dont exits")
}

func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}
