package models

type Developer struct {
	Model
	ProjectID uint     `json:"projectId" gorm:"not null"`
	Project   *Project `json:"project"`
	UserID    uint     `json:"userId" gorm:"not null"`
	User      *User    `json:"user"`
}
