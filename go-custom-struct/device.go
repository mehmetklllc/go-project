package mkilic


import (
	"time"
)

type Device1 struct {
	ID       string
	NAME string
	DATE     time.Time
	PROPERTY string
	VALUE    string
}

func newDevice1(id , name , property , value string, date time.Time)  *Device1{
	d := new(Device1)
	d.ID = id
	d.NAME= name
	d.DATE = date
	d.PROPERTY = property
	d.VALUE = value

	return d
}

