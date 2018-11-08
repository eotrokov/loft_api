package model

import (
	"github.com/google/uuid"
)

type Person struct {
	Id uuid.UUID `json:"id"`
	LastName string `json:"lastName"`
	FirstName string `json:"firstName"`
	TypeId uuid.UUID `json:"typeId"`
	Type Type `json:"type"`
}
