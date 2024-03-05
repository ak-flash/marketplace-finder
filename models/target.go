package models

import (
	"time"
)

type Target struct {
	Id     int  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Active bool `json:"active" gorm:"default:false"`
	UserID string
	//User      User   `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name      string `json:"name" gorm:"not null;type:varchar(100);default:null"`
	Price     int    `json:"target_price" gorm:"default:0"`
	Percent   int    `json:"target_percent" gorm:"default:0"`
	MaxPrice  int    `json:"max_price" gorm:"default:0"`
	Desc      string `json:"desc" gorm:"default:null"`
	Period    int    `json:"period" gorm:"default:10"`
	ErrCount  int    `json:"err_count" gorm:"default:0"`
	Error     string `json:"error" gorm:"default:null"`
	Status    string `json:"status"`
	CheckedAt time.Time
}

func GetTargets() []Target {

	var targets []Target
	var checkTargets []Target

	// Get all records
	DB.Table("targets").Where("active", true).Find(&targets)

	for _, each := range targets {
		if time.Now().UTC().Add(-time.Minute * time.Duration(each.Period-1)).After(each.CheckedAt) {
			checkTargets = append(checkTargets, each)
		}
	}

	return checkTargets
}

func GetAllTargets() []Target {

	var targets []Target
	// Get all records
	DB.Table("targets").Find(&targets)

	return targets
}

func GetActiveTargets() []Target {

	var targets []Target
	// Get all active records
	DB.Table("targets").Where("active", true).Find(&targets)

	return targets
}

func (target *Target) UpdateCheckTime() bool {

	DB.Table("targets").Where("id", target.Id).Update("checked_at", time.Now().UTC())

	return true
}

func (target *Target) SetError(text string) bool {

	if text == "not_found" {
		text = "Товары не найдены. Нyжно сменить поисковый запрос"
	}

	if target.ErrCount > 2 {
		DB.Table("targets").Where("id", target.Id).Update("error", text).Update("active", 0)
	} else {
		DB.Table("targets").Where("id", target.Id).Update("error", text).Update("err_count", target.ErrCount+1)
	}

	return true
}
