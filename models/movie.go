// MODELS - GO STRUCT REPRESENTING DATABASE TABLES
package models

type Movie struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `json:"title" gorm:"type:varchar(100)"`
	Director string `json:"director" gorm:"type:varchar(100)"`
	Year     int    `json:"year" gorm:"type:int"`
}
