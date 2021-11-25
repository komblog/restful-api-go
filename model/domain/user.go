package domain

import "database/sql"

type User struct {
	Id         int
	FirstName  string
	LastName   string
	Email      sql.NullString
	Birth_Date sql.NullTime
}
