package model

import (
	"github.com/volatiletech/null"
)

type User struct {
	Id        uint32    `json:"id"   db:"id"`
	Name      string    `json:"name" db:"name"`
	Mail      string    `json:"mail"  db:"mail"`
	UpdatedAt null.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt null.Time `json:"createdAt" db:"created_at"`
}
