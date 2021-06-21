package models

import (
	"time"
)

type Gif struct {
	Id int64			`gorm:"primary_key" json:"id"`
	Name	int64		`json:"name"`
	Description int64	`json:"description"`
	Created time.Time	`json:"created"`
	UserId string		`sql:"size:255" json:"user_id"`		
}
