package entity

import "time"

type Token struct {
	Tid  uint64 `json:"tid" gorm:"primarykey"`
	Sid  uint64 `json:"uid"`
	Uuid string
	Exp  time.Time
}
