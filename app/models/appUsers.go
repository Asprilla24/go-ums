package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type AppUsers struct {
	ID                     int64     `gorm:"primary_key" json:"id"`
	AccessFailedCount      int       `json:"accessFailedCount"`
	EmailAddress           string    `gorm:"type:varchar(100);unique_index" json:"emailAddress"`
	EmailConfirmationCode  string    `gorm:"size:20" json:"emailConfirmationCode"`
	IsActive               bool      `json:"isActive"`
	IsEmailConfirmed       bool      `json:"isEmailConfirmed"`
	IsPhoneNumberConfirmed bool      `json:"isPhoneNumberConfirmed"`
	LastLoginTime          time.Time `json:"lastLoginTime"`
	Password               string    `sql:"-" json:"-"`
	HashedPassword         []byte    `json:"-"`
	PasswordResetCode      string    `gorm:"size:20" json:"passwordResetCode"`
	PhoneNumber            string    `gorm:"size:20" json:"phoneNumber"`
	FirstName              string    `gorm:"size:50" json:"firstName"`
	LastName               string    `gorm:"size:50" json:"lastName"`
	TenantID               string    `gorm:"size:100" json:"tenantId"`
	Username               string    `gorm:"type:varchar(100);unique_index" json:"username"`
	Gender                 string    `gorm:"size:10" json:"gender"`
	LastUpdated            time.Time `json:"lastUpdated"`
}

// SetNewPassword set a new hashsed password to user
func (user *AppUsers) EncryptPassword() {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.HashedPassword = bcryptPassword
}

func (user *AppUsers) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.Must(uuid.NewV4()))
	scope.SetColumn("LastUpdated", time.Now())
	scope.SetColumn("LastLoginTime", time.Now())
	user.EncryptPassword()
	return nil
}

func (user *AppUsers) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("LastUpdated", time.Now())
	return nil
}
