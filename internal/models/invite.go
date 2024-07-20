package models

import (
	"time"

	"gorm.io/gorm"
)

type Invitation struct {
	ID             string `gorm:"type:uuid;primary_key;"`
	UserID         string `gorm:"type:uuid;"`
	OrganisationID string `gorm:"type:uuid;"`
	CreatedAt      time.Time
	ExpiresAt      time.Time
	IsValid        bool
	Email          string `gorm:"type:varchar(100);"`
}

func (i *Invitation) CreateInvitation(db *gorm.DB, invitation interface{}) error {
	err := db.Create(invitation).Error
	if err != nil {
		return err
	}
	return nil
}
