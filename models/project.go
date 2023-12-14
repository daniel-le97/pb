package models
import "time"
type Project struct {
	ID             string     `db:"id" json:"id"`
	CollectionID   *string    `db:"collectionId" json:"collectionId"`
	CollectionName *string    `db:"collectionName" json:"collectionName"`
	Created        *time.Time `db:"created" json:"created"`
	Updated        *time.Time `db:"updated" json:"updated"`
	RepoURL        *string    `db:"repoURL" json:"repoURL"`
	Name           *string    `db:"name" json:"name"`
	Deployed       *bool      `db:"deployed" json:"deployed"`
	Buildpack      *string    `db:"buildpack" json:"buildpack"`
	Configured     *bool      `db:"configured" json:"configured"`
	BaseDir        *string    `db:"baseDir" json:"baseDir"`
	BuildDir       *string    `db:"buildDir" json:"buildDir"`
	HTTPS          *bool      `db:"https" json:"https"`
	WWW            *bool      `db:"www" json:"www"`
	Managed        *bool      `db:"managed" json:"managed"`
	InstallCommand *string    `db:"installCommand" json:"installCommand"`
	BuildCommand   *string    `db:"buildCommand" json:"buildCommand"`
	StartCommand   *string    `db:"startCommand" json:"startCommand"`
	Ports          *string    `db:"ports" json:"ports"`
	ExposedPorts   *string    `db:"exposedPorts" json:"exposedPorts"`
}