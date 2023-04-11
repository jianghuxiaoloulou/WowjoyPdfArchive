package global

import "database/sql"

var (
	ReadDBEngine  *sql.DB
	WriteDBEngine *sql.DB
	ExistStatus   sql.NullInt16
)
