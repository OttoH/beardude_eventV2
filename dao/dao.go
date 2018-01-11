package dao

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"

	"beardude_eventV2/models"
)

/* data access object */
type DAO struct {
  Server   string
	Database string
}

var db *mgo.Database

const (
	EVENT = "event"
	RACER = "racer"
)

func (m *DAO) Connect() *mgo.Session {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)

	return session.Clone()
}

// Events
func (m *DAO) CreateEvent(events *models.Events) error {
	err := db.C(EVENT).Insert(&events)
	return err
}


// Racer
func (m *DAO) ValidateRacer(user string, pwd string) bool {
	log.Println(user, pwd)
	result := &models.Racer{}
	err := db.C(RACER).Find(bson.M{"username": user, "password": pwd, "activate": 1}).One(result)
	if err == nil && len(result.Username) > 0 {
		return true
	}

	return false
}

func (m *DAO) GetRacerById(id string) (*models.Racer, error) {
	result := &models.Racer{}
	err := db.C(RACER).FindId(bson.ObjectIdHex(id)).One(result)

	return result, err
}

func (m *DAO) GetRacerByName(user string) (*models.Racer, error) {
	result := &models.Racer{}
	err := db.C(RACER).Find(bson.M{"username": user}).One(result)

	return result, err
}

func (m *DAO) CreateRacer(racer *models.Racer) error {
	result := &models.Racer{}

	// check if user exists
	if db.C(RACER).Find(bson.M{"username": racer.Username}).One(result);len(result.Username) > 0 {
		return errors.New("user already exists")
	}

	error := db.C(RACER).Insert(&racer)
	return error
}

func (m *DAO) UpdateRacer(racer *models.Racer) error {
	selector := bson.M{"username": racer.Username}
	data := bson.M{"$set": bson.M{"username": racer.Username, "password": racer.Password, "nockname": racer.Nickname}}
	err := db.C(RACER).Update(selector, data)
	return err
}

func (m *DAO) RemoveRacer(racer *models.Racer) error {
	selector := bson.M{"username": racer.Username}
	data := bson.M{"$set": bson.M{"activate": 0}}
	err := db.C(RACER).Update(selector, data)
	return err
}
