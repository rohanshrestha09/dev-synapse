package models

type Notification struct {
	Model
	Description string `json:"description" gorm:"not null"`
	InitiatorID uint   `json:"initiatorId" gorm:"not null"`
	Initiator   *User  `json:"initiator"`
	ListenerID  uint   `json:"listenerId" gorm:"not null"`
	Listener    *User  `json:"listener"`
}
