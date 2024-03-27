package models

import (
	"time"
)

type RequestEvents struct {
	ReqID        int        `json:"reqid" gorm:"primaryKey;autoIncrement:true;unique"`
	BookID       int        `json:"bookid"`
	ReaderId     int        `json:"readerid"`
	RequestDate  time.Time  `json:"requestdate"`
	ApprovelDate *time.Time `json:"approveldate"`
	ApproverID   int        `json:"approverid"`
	RequestType  string     `json:"requesttype"`
}
