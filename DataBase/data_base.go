package DataBase

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task_rest/middleware"
)

// History struct for used values from Data Base
type History struct {
	ID     int    `gorm:"column:id"`
	Type   string `gorm:"column:type"`
	Input  string `gorm:"column:input"`
	Output string `gorm:"column:output"`
}

// Init function: initialized Data Base
func Init(dbURL string) (*gorm.DB, error) {
	middleware.Logs.Debug().Msgf("[DataBase] Init started")
	if dbURL == "" {
		dbURL = "postgres://user:u000000@localhost:5432/postgres"
	}
	middleware.Logs.Debug().Str("dbURL=", dbURL).Msgf("[DataBase]")
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}

func AddRec(db *gorm.DB, _type string, input string, output string) error {
	middleware.Logs.Debug().Msgf("[DataBase] AddRec started")
	var newRec = History{Type: _type, Input: input, Output: output}
	middleware.Logs.Debug().Interface("new record=", newRec).Msgf("[DataBase]")
	res := db.Table("history").Create(&newRec)
	return res.Error
}

func Show(db *gorm.DB, limit, offset int) (result []History, err error) {
	middleware.Logs.Debug().Msgf("[DataBase] Show started")
	middleware.Logs.Debug().Int("limit=", limit).Int("offset=", offset).Msgf("[DataBase]")
	err = db.Table("history").Limit(limit).Offset(offset).Scan(&result).Error
	return
}
