package models

import "time"

type IssueRegistery struct {
	IssueID            int        `json:"issueid" gorm:"primaryKey;autoIncrement:true"`
	ISBN               int        `json:"isbn"`
	ReaderId           int        `json:"readerid"`
	IssueApproverID    int        `json:"issueapprovelid"`
	IssueStatus        string     `json:"issuestatus"`
	IssueDate          time.Time  `json:"issuedate"`
	ExpectedReturnDate time.Time  `json:"expectedreturndate"`
	ReturnDate         *time.Time `json:"returndate"`
	ReturnApproverID   int        `json:"returnapprovelid"`
}
