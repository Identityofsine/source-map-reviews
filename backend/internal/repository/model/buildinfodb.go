package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type BuildInfoDb struct {
	Version     string    `json:"version"`
	CommitHash  string    `json:"commit_hash"`
	BuildDate   time.Time `json:"build_date"`
	Environment string    `json:"environment"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetBuildInfo() ([]BuildInfoDb, db.DatabaseError) {
	query := "SELECT * FROM buildinfo"
	rows, err := db.Query[BuildInfoDb](query)

	if err != nil {
		return nil, err
	}

	return *rows, err
}

func GetBuildInfoByVersionAndCommitHash(version string, commitHash string) (BuildInfoDb, db.DatabaseError) {
	query := "SELECT * FROM buildinfo WHERE version = $1 AND commit_hash = $2"
	rows, err := db.Query[BuildInfoDb](query, version, commitHash)

	if err != nil {
		return BuildInfoDb{}, err
	}

	if len(*rows) == 0 {
		return BuildInfoDb{}, nil
	}

	return (*rows)[0], nil
}

func DoesVersionExist(version string, commitHash string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM buildinfo WHERE version = $1 AND commit_hash = $2)"
	rows, err := db.Query[bool](query, version, commitHash)

	if err != nil {
		return false, err
	}

	if len(*rows) == 0 {
		return false, nil
	}

	return (*rows)[0], nil
}

func InsertBuildInfo(buildInfo BuildInfoDb) db.DatabaseError {
	if buildInfo.Version == "" || buildInfo.CommitHash == "" {
		return db.NewDatabaseError("InsertBuildInfo", "Version and CommitHash cannot be empty", "", 400)
	}
	if exists, err := DoesVersionExist(buildInfo.Version, buildInfo.CommitHash); err != nil || exists {
		return db.NewDatabaseError("InsertBuildInfo", "Version already exists", "", 409)
	}

	query := "INSERT INTO buildinfo (version, commit_hash, build_date, environment, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Query[BuildInfoDb](query, buildInfo.Version, buildInfo.CommitHash, buildInfo.BuildDate, buildInfo.Environment, buildInfo.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
