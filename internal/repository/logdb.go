package repository

import (
	"database/sql"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type LogDB struct {
	Id        int            `db:"id" dao:"omit"`
	Severity  string         `db:"severity"`
	Message   string         `db:"message"`
	CreatedAt time.Time      `db:"created_at" dao:"omit"`
	Version   sql.NullString `db:"version"`
	Commit    sql.NullString `db:"commit_hash"`
}

func GetLogs() ([]LogDB, db.DatabaseError) {
	return dao.SelectFromDatabaseByStruct[LogDB](
		LogDB{},
		"")
}

func SaveLogs(logDB LogDB) db.DatabaseError {

	err := dao.InsertIntoDatabaseByStruct(logDB)
	return err
}
