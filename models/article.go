package models

type Article struct {
	Id int `sql:"AUTO_INCREMENT"`
	ArticleBrand Brand
	ArticleBrandId int
	Number string
}