package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"github.com/pkg/errors"
)

type MongoConnection struct {
	Session  *mgo.Session
	MongoUrl string
}

func (mc *MongoConnection) ConnectMongo() error {
	var err error
	mc.Session, err = mgo.Dial(mc.MongoUrl)
	if err != nil {
		log.Fatal("error in connecting to database with url : ", mc.MongoUrl)
		return errors.New("error in connecting to database")
	}
	return nil
}
