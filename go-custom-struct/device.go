package mkilic


import (
	"time"
)

type Device struct {
	ID       string
	NAME string
	DATE     time.Time
	PROPERTY string
	VALUE    string
}

func NewDevice(id , name , property , value string, date time.Time)  *Device{
	d := new(Device)
	d.ID = id
	d.NAME= name
	d.DATE = date
	d.PROPERTY = property
	d.VALUE = value

	return d
}

