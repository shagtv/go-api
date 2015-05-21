package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
)

var Mongo *mgo.Session

func Connection() (*mgo.Session) {
	var err error
	if Mongo == nil {
		//Mongo, err = mgo.Dial("127.0.0.1:27017")
		Mongo, err = mgo.Dial("t01.abcp.ru:27117")
		if err != nil {
			log.Fatal(err)
		}
		Mongo.SetMode(mgo.Monotonic, true)
	}
	return Mongo.Copy()
}