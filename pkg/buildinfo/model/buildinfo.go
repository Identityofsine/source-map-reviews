package model

type BuildInfo struct {
	// Version of the application
	Version string `json:"version" db:"version"`
	// CommitHash of the application
	CommitHash string `json:"commit_hash" db:"commit_hash"`
	// BuildDate of the application
	BuildDate string `json:"build_date" db:"build_date"`
	// Environment of the application
	Environment string `json:"environment" db:"environment"`
	// CreatedAt is the time when the build info was CreatedAt
	CreatedAt string `json:"created_at" db:"created_at"`
}
