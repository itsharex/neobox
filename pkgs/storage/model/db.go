package model

import (
	"path/filepath"
	"time"

	"github.com/moqsien/hackbrowser/utils/hsqlite"
	"github.com/moqsien/neobox/pkgs/conf"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBEngine *gorm.DB
)

func NewDBEngine(cnf *conf.NeoConf) (db *gorm.DB, err error) {
	dbPath := filepath.Join(cnf.WorkDir, conf.SQLiteDBFileName)
	db, err = gorm.Open(
		hsqlite.Open(dbPath),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
			DryRun: true,
		},
	)
	db.Callback().Create().Replace("gorm:create", beforeCreate)
	db.Callback().Update().Replace("gorm:update", beforeUpdate)
	return
}

func beforeCreate(db *gorm.DB) {
	if db.Statement.Schema != nil {
		nowTime := time.Now().Unix()
		if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
			if _, isZero := createTimeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue); isZero {
				createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
			}
		}
	}
}

func beforeUpdate(db *gorm.DB) {
	if db.Statement.Schema != nil {
		nowTime := time.Now().Unix()
		if modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn"); modifyTimeField != nil {
			if _, isZero := modifyTimeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue); isZero {
				modifyTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
			}
		}
	}
}

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
}