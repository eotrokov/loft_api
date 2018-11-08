package model

import (
	"github.com/google/uuid"
)

type Type struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}
