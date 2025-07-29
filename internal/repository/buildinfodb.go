package repository

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type BuildInfoDB struct {
	Version     string    `json:"version" db:"version"`
	CommitHash  string    `json:"commit_hash" db:"commit_hash"`
	BuildDate   time.Time `json:"build_date" db:"build_date"`
	Environment string    `json:"environment" db:"environment"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func GetBuildInfo() ([]BuildInfoDB, db.DatabaseError) {
	rows, err := dao.SelectFromDatabaseByStruct(BuildInfoDB{}, "")
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, exception.ResourceNotFoundDatabase
	}

	return rows, err
}

func GetBuildInfoByVersionAndCommitHash(version string, commitHash string) (*BuildInfoDB, db.DatabaseError) {

	rows, err := dao.SelectFromDatabaseByStruct(BuildInfoDB{}, "version = $1 AND commit_hash = $2", version, commitHash)
	if err != nil {
		return &BuildInfoDB{}, err
	}

	if len(rows) == 0 {
		return nil, exception.ResourceNotFoundDatabase
	}

	return &(rows)[0], nil
}

func DoesVersionExist(version string, commitHash string) (bool, db.DatabaseError) {

	_, err := GetBuildInfoByVersionAndCommitHash(version, commitHash)
	if err != nil {
		if err.Code == 404 {
			return false, nil // Version does not exist
		}
		return false, err // Some other error occurred
	}

	return true, nil // Version exists

}

func InsertBuildInfo(buildInfo BuildInfoDB) db.DatabaseError {
	if buildInfo.Version == "" || buildInfo.CommitHash == "" {
		return db.NewDatabaseError("InsertBuildInfo", "Version and CommitHash cannot be empty", "", 400)
	}
	if exists, err := DoesVersionExist(buildInfo.Version, buildInfo.CommitHash); err != nil || exists {
		return db.NewDatabaseError("InsertBuildInfo", "Version already exists", "", 409)
	}

	err := dao.InsertIntoDatabaseByStruct(buildInfo)

	return err

}
