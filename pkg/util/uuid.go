package util

import (
	uuid "github.com/satori/go.uuid"
)

func NewUUID() uuid.UUID {
	return uuid.NewV4()
}
