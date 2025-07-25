package id

import "github.com/google/uuid"

type UUID = uuid.UUID

func Parse(s string) (UUID, error) {
	return uuid.Parse(s)
}

func MustParse(s string) UUID {
	return uuid.MustParse(s)
}
