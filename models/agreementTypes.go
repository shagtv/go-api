package models

type AgreementTypes struct {
	Id uint `sql:"AUTO_INCREMENT"`
	Name string
	IsDelete int
}

func (c AgreementTypes) TableName() string {
	return "agreementTypes"
}