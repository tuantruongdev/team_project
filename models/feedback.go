package models

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	MAX_CONTENT = 4096
)

type Feedback struct {
	PackageName string `json:"package" gorm:"column:package"`
	Version     string `json:"version" gorm:"version""`
	Content     string `json:"content" gorm:"content"`
	Ip          string `json:"-" gorm:"ip"`
}

func (*Feedback) TableName() string {
	return "feedback"
}
func InsertFeedback(db *gorm.DB, feedback *Feedback) bool {
	if err := db.Create(&feedback).Error; err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
