package models

type Credential struct {
	ID        uint
	AppName   string
	ClientID  string `gorm:"size:191;uniqueIndex"`
	SecretKey string
}
