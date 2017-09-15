package main

import (
	"LordOfUltima/database"
	"LordOfUltima/gameEntities"

	_ "time"
)

type mapServer struct {
	Cities          []gameEntities.City
	Dungeons        []gameEntities.Dungeon
	Resources       []Resource
	transports      []*gameEntities.Transport
	militaryActions []*gameEntities.MilitaryAction

	Requests chan Request
	Answers  chan Answer
}

func (m *mapServer) LoadAndStart() {
	m.Requests = make(chan Request)
	m.Answers = make(chan Answer)
	db := database.Session.DB("lordofultima")
	db.C("cities").Find(nil).All(&m.Cities)
	go m.RequestProcessor()
}

func (m *mapServer) getCities(posX uint, posY uint, rang uint) *[]gameEntities.City {
	var listReturn []gameEntities.City
	for i := range m.Cities {
		if m.Cities[i].InRange(posX, posY, rang) {
			listReturn = append(listReturn, m.Cities[i])
		}
	}
	return &listReturn

}

func AnswerGenerator() {

}

func (m *mapServer) RequestProcessor() {

	request := <-m.Requests
	switch request.Code {
	case 0:
		println("caso 0")
		break
	}
}
