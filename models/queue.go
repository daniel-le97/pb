package models

type Queue struct {
	ID      string  `db:"id" json:"id"`
	Project string  `db:"project" json:"project"`
	Active  bool    `db:"active" json:"active"`
	BuildTime float64 `db:"buildTime" json:"buildTime"`
	Logs     string  `db:"logs" json:"logs"`
}
