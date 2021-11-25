package web

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	errUnmarshal := json.Unmarshal(b, &ns.String)

	ns.Valid = (errUnmarshal == nil)
	return errUnmarshal
}

type NullTime struct {
	sql.NullTime
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {
	errUnmarshal := json.Unmarshal(b, &nt.Time)
	nt.Valid = (errUnmarshal == nil)
	return errUnmarshal
}

type CreateUserRequest struct {
	Id         int        `validate:"required" json:"id"`
	FirstName  string     `validate:"required,max=100,min=1" json:"firstname"`
	LastName   string     `validate:"required,max=100,min=1" json:"lastname"`
	Email      NullString `json:"email"`
	Birth_Date NullTime   `json:"birth_date"`
}
