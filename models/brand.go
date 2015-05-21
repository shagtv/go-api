package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/shagtv/go-api/library/mongo"
	"errors"
)

type Brand struct {
	Id                 bson.ObjectId    `json:"id"        bson:"_id,omitempty"`
	Aliases            []string
	Info               BrandInfo
	Is_bad             bool
	Is_manufacturer    bool
	Is_original        bool
	Is_quality         bool
	Is_short_crosses   bool
	Lc_aliases         []string
	Lc_aliases_cleared []string
	Lc_name            string
	Lc_name_cleared    string
	Name               string
	Number_format      string
	Outer_number_func  string
	Short_number_func  string
}

type BrandInfo struct {
	Annotation string
	Text       string
	Logo       string
}

func FindBrandByName(name string) (brand Brand, err error) {
	sess := mongo.Connection()
	defer sess.Close()

	if len(name) == 0 {
		err = errors.New("empty name")
		return
	}

	c := sess.DB("brand_api").C("brand")
	if (c != nil) {
		err = c.Find(bson.M{"name": name}).One(&brand)
		if err != nil {
			return
		}
	} else {
		err = errors.New("database error")
		return
	}
	return
}