package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Brand struct {
	Id                 bson.ObjectId `json:"id"        bson:"_id,omitempty"`
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
