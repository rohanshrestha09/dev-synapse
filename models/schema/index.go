package schema

import (
	"database/sql"
	"time"

	"github.com/rohanshrestha09/dev-synapse/enums"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	Model
	Email     string     `json:"email,omitempty" gorm:"not null;unique"`
	Name      string     `json:"name" gorm:"not null"`
	Bio       string     `json:"bio"`
	Password  string     `json:"-" gorm:"not null"`
	Image     string     `json:"image"`
	ImageName string     `json:"imageName"`
	Provider  string     `json:"provider" gorm:"type:enum('GOOGLE','EMAIL');default:EMAIL;not null" binding:"required"`
	Projects  []*Project `json:"projects,omitempty"`
}

type Project struct {
	Model
	Name              string              `json:"name" gorm:"not null"`
	Description       string              `json:"description" gorm:"not null"`
	Published         bool                `json:"published" gorm:"default:true;not null"`
	Image             string              `json:"image"`
	ImageName         string              `json:"imageName"`
	EstimatedDuration int                 `json:"estimatedDuration" gorm:"not null"`
	StartDate         sql.NullTime        `json:"startDate"`
	EndDate           sql.NullTime        `json:"endDate"`
	Status            enums.ProjectStatus `json:"status" gorm:"type:enum('OPEN','IN_PROGRESS','COMPLETED');default:OPEN;not null"`
	UserID            uint                `json:"userId" gorm:"not null"`
	User              *User               `json:"user"`
	Developers        *[]Developer        `json:"developers"`
}

type Request struct {
	Model
	ProjectID uint                `json:"projectId" gorm:"not null"`
	Project   *Project            `json:"project"`
	UserID    uint                `json:"userId" gorm:"not null"`
	User      *User               `json:"user"`
	Status    enums.RequestStatus `json:"status" gorm:"type:enum('APPROVED','PENDING','REJECTED');default:PENDING;not null"`
}

type Notification struct {
	Model
	Description string `json:"description" gorm:"not null"`
	InitiatorID uint   `json:"initiatorId" gorm:"not null"`
	Initiator   *User  `json:"initiator"`
	ListenerID  uint   `json:"listenerId" gorm:"not null"`
	Listener    *User  `json:"listener"`
}

type Developer struct {
	Model
	ProjectID uint     `json:"projectId" gorm:"not null"`
	Project   *Project `json:"project"`
	UserID    uint     `json:"userId" gorm:"not null"`
	User      *User    `json:"user"`
}

type Chat struct {
	Model
	ProjectID uint     `json:"projectId" gorm:"not null"`
	Project   *Project `json:"project"`
	UserID    uint     `json:"userId" gorm:"not null"`
	User      *User    `json:"user"`
	Message   string   `json:"message" gorm:"not null"`
}
