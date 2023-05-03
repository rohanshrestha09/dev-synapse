package models

import "github.com/rohanshrestha09/dev-synapse/enums"

type Request struct {
	Model
	ProjectID uint                `json:"projectId" gorm:"not null"`
	Project   *Project            `json:"project"`
	UserID    uint                `json:"userId" gorm:"not null"`
	User      *User               `json:"user"`
	Status    enums.RequestStatus `json:"status" gorm:"type:enum('APPROVED','PENDING','REJECTED');default:PENDING;not null"`
}
