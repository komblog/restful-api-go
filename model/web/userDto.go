package web

import "database/sql"

type UserDTO struct {
	Id         int            `json:"id"`
	FirstName  string         `json:"firstname"`
	LastName   string         `json:"lastname"`
	Email      sql.NullString `json:"email"`
	Birth_Date sql.NullTime   `json:"birth_date"`
}
