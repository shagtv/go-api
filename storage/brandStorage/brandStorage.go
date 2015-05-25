package brandStorage

import (
	"github.com/shagtv/go-api/httperrors"
	"github.com/shagtv/go-api/library/mongo"
	"github.com/shagtv/go-api/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func OneByName(name string) (brand models.Brand, err httperrors.IHttperror) {
	sess := mongo.Connection()
	defer sess.Close()

	if len(name) == 0 {
		err = httperrors.New("empty name", http.StatusBadRequest)
		return
	}

	c := sess.DB("brand_api").C("brand")
	if c != nil {
		mongoError := c.Find(bson.M{"name": name}).One(&brand)
		if mongoError != nil {
			err = httperrors.New(mongoError.Error(), http.StatusNotFound)
			return
		}
	} else {
		err = httperrors.New("database error", http.StatusInternalServerError)
		return
	}
	return
}

func List(skip int, limit int) (brands []models.Brand, err httperrors.IHttperror) {
	sess := mongo.Connection()
	defer sess.Close()

	c := sess.DB("brand_api").C("brand")
	if c != nil {
		mongoError := c.Find(bson.M{}).Skip(skip).Limit(limit).All(&brands)
		if mongoError != nil {
			err = httperrors.New("database error", http.StatusInternalServerError)
			return
		}
	} else {
		err = httperrors.New("database error", http.StatusInternalServerError)
		return
	}
	return
}

func Count() (count int, err httperrors.IHttperror) {
	sess := mongo.Connection()
	defer sess.Close()

	c := sess.DB("brand_api").C("brand")
	if c != nil {
		var mongoError error
		count, mongoError = c.Count()
		if mongoError != nil {
			err = httperrors.New("database error", http.StatusInternalServerError)
			return
		}
	} else {
		err = httperrors.New("database error", http.StatusInternalServerError)
		return
	}
	return
}
