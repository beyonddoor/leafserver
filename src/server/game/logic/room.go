package logic

import (
	"log"
	"errors"
)

import("github.com/name5566/leaf/gate")

type Server struct {
	Rooms []*Room
	Hall Hall
}

type Hall struct {
	Users []*User
}

type Room struct {
	RoomId string
	Users []*User
}

type User struct {
	UserId string
	Agent gate.Agent
}

var server = Server{[]*Room{}, Hall{}}

func init()  {
	log.Println("init server")
	CreateRoom("def")
}

func CreateRoom(rid string)  {
	server.Rooms = append(server.Rooms, &Room{RoomId:rid})
}

func DestroyRoom(rid string)  {
	panic("todo")
}

func AddToRoom(rid string, user *User)  {
	room, err := GetRoom(rid)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range(room.Users) {
		if v.Agent == user.Agent {
			return
		}
	}
	
	room.Users = append(room.Users, user)
}

func RemoveFromRoom(rid string, user *User)  {
	_, err := GetRoom(rid)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetRoom(rid string) (room *Room, err error) {
	for _, v := range(server.Rooms) {
		if v.RoomId == rid {
			return v, nil
		}
	}
	return nil, errors.New("room not found")
}

func (room *Room) SendAll(text interface{})  {
	for _, v := range(room.Users) {
		v.Agent.WriteMsg(text)
	}
}

func SendAll(rid string,text interface{})  {
	room, err := GetRoom(rid)
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range(room.Users) {
		v.Agent.WriteMsg(text)
	}
}

