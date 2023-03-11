package model

import "time"

type Group struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Created time.Time `json:"created" xorm:"created"`
}