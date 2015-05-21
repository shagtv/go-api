package models
import "time"

type ClientNotes struct {
	Id uint `sql:"AUTO_INCREMENT"`
	clientId int
	NoteDate time.Time
	Comment string
	IsDelete int
	LastEditorId int
}

func (c ClientNotes) TableName() string {
	return "clientNotes"
}